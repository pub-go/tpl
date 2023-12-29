package html

import (
	"fmt"
	"html"
	"io"
	"reflect"
	"strings"

	"code.gopub.tech/tpl/internal/exp"
	"code.gopub.tech/tpl/internal/types"
)

var _ types.Template = (*htmlTemplate)(nil)

// htmlTemplate HTML 模板实例
type htmlTemplate struct {
	manager       *tplManager
	name          string
	node          *Node
	nodeCondition map[*Node]bool
}

// NewTemplate 构造一个模板实例
func NewTemplate(m *tplManager, name string, node *Node) *htmlTemplate {
	return &htmlTemplate{
		manager: m,
		name:    name,
		node:    node,
	}
}

// Execute implements types.Template. 执行模板
func (t *htmlTemplate) Execute(w io.Writer, data any) error {
	scope := exp.NewScope(data)
	scope = exp.Combine(scope, t.manager.globalScope)
	return t.execute(t.node, w, scope, nil)
}

// executeOption 执行模板的参数
type executeOption struct {
	noPrintToken bool // 不输出 tag 本身
	processChild func(w io.Writer) error
}

// nop 什么都不做
func nop(w io.Writer) error { return nil }

// execute 执行节点
func (t *htmlTemplate) execute(tree *Node, w io.Writer, data exp.Scope, opt *executeOption) error {
	if tree == nil {
		return nil
	}
	if opt == nil {
		opt = &executeOption{}
	}

	// 开始标签
	var (
		token    = tree.Token         // <p>child</p> ==> <p>是token, child 是子节点, </p> 是End
		tokenBuf = &strings.Builder{} // 开始标签缓冲区
	)
	if token != nil { // root doc is nil
		switch token.Kind {
		case TokenKindTag: // tag 标签
			if err := t.processTagStart(tree, tokenBuf, data, opt); err != nil {
				return fmt.Errorf("failed to process tag [%v](from postion %v to %v): %w",
					token.Value, token.Start, token.End, err)
			}
		case TokenKindComment: // 注释
			cmtValue := strings.TrimPrefix(token.Value, "<!--")
			cmtValue = strings.TrimSuffix(cmtValue, "-->")
			cmtValue = strings.TrimSpace(cmtValue)
			// <!-- /* */ -->
			if strings.HasPrefix(cmtValue, "/*") && strings.HasSuffix(cmtValue, "*/") {
				opt.noPrintToken = true
			}
			writeToBuf(opt, tokenBuf, token.Value)
		default: // text, cdata
			writeToBuf(opt, tokenBuf, token.Value)
		}
	}
	if _, err := w.Write([]byte(tokenBuf.String())); err != nil {
		return err // 从缓冲区真正输出开始标签
	}

	// 标签内的子节点内容
	if opt.processChild == nil {
		opt.processChild = func(w io.Writer) error {
			for _, child := range tree.Children {
				if err := t.execute(child, w, data, nil); err != nil {
					return err
				}
			}
			return nil
		}
	}
	if err := opt.processChild(w); err != nil {
		return err
	}

	// 结束标签
	if token = tree.End; token != nil && !opt.noPrintToken {
		_, err := w.Write([]byte(token.Value))
		if err != nil {
			return err
		}
	}
	return nil
}

// processTagStart 处理开始标签
func (t *htmlTemplate) processTagStart(node *Node, tokenBuf *strings.Builder, data exp.Scope, opt *executeOption) error {
	tag := node.Token.Tag
	if tag == nil {
		return fmt.Errorf("unexpected nil tag")
	}
	attrPrefix := t.manager.attrPrefix
	name := strings.ToLower(tag.Name)
	if name == t.manager.tagPrefix+tagNameBlock {
		opt.noPrintToken = true // <t:block> ... </t:block>
	}
	skipAttrNames := t.hasAnyAttr(tag,
		attrDefine,  // define 定义复用模板
		attrReplace, // replace 用指定模板替换本节点
		attrIf, attrElse_If, attrElseIf, attrElIf, attrElse,
		attrRange, // 循环输出本节点
	)
	if len(skipAttrNames) > 0 {
		// 有这些属性的, 本次不输出 tag 本身和子节点
		// <ul :if="xxx">xxx</ul>
		// <li :range="xxx">xxx<li>
		// 在下面处理具体属性时会重新调用 execute 再输出
		opt.noPrintToken = true
		opt.processChild = nop
	}
	if len(t.hasAnyAttr(tag, attrInsert)) > 0 {
		// insert 属性：本节点需要输出 子节点不输出
		// 会在下面执行 insert 属性时替换为指定模板
		opt.processChild = nop
	}

	var tagBuf = &strings.Builder{}
	writeToBuf(opt, tagBuf, "<"+tag.Name)

	attrMap := tag.AttrMap()
	for _, attr := range tag.SortedAttr(attrPrefix) {
		attr := attr
		an := attr.Name
		if strings.HasPrefix(an, attrPrefix) { // 指令属性
			cmd := strings.TrimPrefix(an, attrPrefix)
			switch cmd {
			case attrIf, attrElse_If, attrElseIf, attrElIf, attrElse: // 条件控制
				if err := t.processIfElse(node, attr, tokenBuf, data, opt); err != nil {
					return err
				}
			case attrRange: // 循环
				if err := t.processRange(node, attr, tokenBuf, data, opt); err != nil {
					return err
				}
			case attrRemove: // 移除
				t.processRemoveAttr(node, attr, data, opt)
			case attrText, attrRaw: // 替换内容
				if opt.processChild != nil {
					continue
				}
				opt.processChild = func(w io.Writer) error {
					av := attr.Value
					if av == nil {
						return fmt.Errorf("attribute `%s` should have value (at position %v)", an, attr.NameEnd)
					}
					result, err := attr.Evaluate(data)
					if err != nil {
						return err
					}
					if cmd == "text" {
						result = html.EscapeString(result)
					}
					_, err = w.Write([]byte(result))
					return err
				}
			case attrDefine: // 定义可复用组件 解析模板时已经处理 这里无需处理
			case attrInsert, attrReplace: // insert=插入组件作为内容 replce=用组件替换本身
				name, err := attr.Evaluate(data)
				if err != nil {
					return fmt.Errorf("failed to evaluate %v attribute [%v]: %w",
						attr.Name, *attr.Value, err)
				}
				tpl, ok := t.manager.files[name]
				if !ok {
					return fmt.Errorf("no such template `%v`:%w", name, ErrTplNotFound)
				}
				return NewTemplate(t.manager, name, tpl).execute(tpl, tokenBuf, data, nil)
			default: // 其他指令 替换普通属性
				result, err := attr.Evaluate(data)
				if err != nil {
					return err
				}
				result = html.EscapeString(result)
				result = fmt.Sprintf(" %v=%q", cmd, result)
				writeToBuf(opt, tagBuf, result)
			}
		} else { // 普通属性
			if _, ok := attrMap[attrPrefix+an]; !ok {
				// 如果有对应的指令属性则跳过 没有才输出
				if !opt.noPrintToken {
					attr.Print(tagBuf)
				}
			}
		}
	}

	writeToBuf(opt, tagBuf, ">")
	writeToBuf(opt, tokenBuf, tagBuf.String())
	return nil
}

func writeToBuf(opt *executeOption, buf *strings.Builder, data string) {
	if !opt.noPrintToken {
		buf.WriteString(data)
	}
}

// hasAnyAttr 是否有指定的属性且存在非空属性值
func (t *htmlTemplate) hasAnyAttr(tag *Tag, names ...string) (attrName map[string]struct{}) {
	attrMap := tag.AttrMap()
	attrName = make(map[string]struct{}, len(names))
	for _, n := range names {
		if attr, ok := attrMap[t.manager.attrPrefix+n]; ok {
			if attr.Value != nil && *attr.Value != "" {
				attrName[n] = struct{}{}
			}
		}
	}
	return
}

// processIfElse 处理 if-else 属性
func (t *htmlTemplate) processIfElse(node *Node, attr *Attr, tokenBuf *strings.Builder, data exp.Scope, opt *executeOption) error {
	an := attr.Name
	if attr.Value == nil {
		return fmt.Errorf("attribute `%s` should have value (at position %v)", an, attr.NameEnd)
	}
	if *attr.Value == "" {
		return nil
	}
	*attr.Value = ""       // 下面重新调用 execute 时 就会跳过本函数了
	opt.processChild = nop // 不要输出子内容 因为如果满足条件 在下面的 execute 会输出子内容
	cmd := strings.TrimPrefix(an, t.manager.attrPrefix)
	switch cmd {
	case "if":
		return t.evaluateCondition(node, attr, tokenBuf, data)
	case "else-if", "elseif", "elif", "else":
		p, ok := t.nodeCondition[node.GetPreviousSiblingTag()]
		if !ok {
			return fmt.Errorf("unexpected `%v` attribute at %v", attr.Name, attr.NameStart)
		}
		if !p {
			// 如果前一个节点是 false 才要计算本节点
			return t.evaluateCondition(node, attr, tokenBuf, data)
		}
	}
	return nil
}

// evaluateCondition 计算条件并执行节点
func (t *htmlTemplate) evaluateCondition(node *Node, attr *Attr, tokenBuf *strings.Builder, data exp.Scope) error {
	result, err := attr.Evaluate(data)
	if err != nil {
		return fmt.Errorf("failed to evaluate `%v` attribute [%v]: %w",
			attr.Name, *attr.Value, err)
	}
	t.addNodeCondition(node, false)
	if result == "true" {
		t.addNodeCondition(node, true)
		if err := t.execute(node, tokenBuf, data, nil); err != nil {
			return err
		}
	}
	return nil
}

// addNodeCondition 记录节点的条件结果
func (t *htmlTemplate) addNodeCondition(node *Node, cond bool) {
	if t.nodeCondition == nil {
		t.nodeCondition = map[*Node]bool{}
	}
	t.nodeCondition[node] = cond
}

// processRange 处理 range 属性
func (t *htmlTemplate) processRange(node *Node, attr *Attr, tokenBuf *strings.Builder, data exp.Scope, opt *executeOption) error {
	an := attr.Name
	av := attr.Value
	if av == nil {
		return fmt.Errorf("attribute `%s` should have value (at position %v)", an, attr.NameEnd)
	}
	attrValue := *av
	if attrValue == "" {
		return nil
	}
	attrValue = strings.TrimPrefix(attrValue, "'")
	attrValue = strings.TrimSuffix(attrValue, "'")
	attrValue = strings.TrimPrefix(attrValue, "\"")
	attrValue = strings.TrimSuffix(attrValue, "\"")
	indexName, itemName, objName, err := extractRange(attrValue)
	if err != nil {
		return err
	}
	scope := exp.WithDefaultScope(data)
	obj, err := scope.Get(objName)
	if err != nil {
		return fmt.Errorf("failed to process attribute `%v` at %v: %w", an, attr.ValueStart, err)
	}

	var (
		getRange = func() (any, any, bool) {
			return nil, nil, false
		}
		rv = reflect.ValueOf(obj)
		rk = rv.Kind()
	)
	switch rk {
	case reflect.Array, reflect.Slice, reflect.String:
		i := 0
		count := rv.Len()
		getRange = func() (any, any, bool) {
			var v any
			if i < count {
				v = rv.Index(i).Interface()
				i++
				return i, v, true
			}
			return 0, nil, false
		}
	case reflect.Map:
		iter := rv.MapRange()
		getRange = func() (any, any, bool) {
			if iter.Next() {
				return iter.Key().Interface(),
					iter.Value().Interface(), true
			}
			return nil, nil, false
		}
	default:
		return fmt.Errorf("`%v` only supports array, slice, string or map, got %v: %v",
			an, rk, objName)
	}

	// <ul>
	//   <li :range>当前 node 是这个</li>
	//   <li>
	// </ul>
	// 处理 range 的 li 项目时，找到并输出 </li> 后的空白字符
	var nextSiblingBlankTextNode = node.GetNextSibling()
	if !nextSiblingBlankTextNode.IsBlankText() {
		nextSiblingBlankTextNode = nil // 如果下一个不是空白文本就算了
	}

	*av = ""   // 标记 range 指令已经执行 下面 for 循环执行 node 时就会忽略 range 指令了
	count := 0 // 多个项目之间才输出空白
	for {
		i, n, ok := getRange()
		if !ok {
			break
		}
		childScope := exp.Combine(exp.NewScope(map[string]any{
			indexName: i,
			itemName:  n,
		}), scope)
		if count > 0 {
			// <ul :remove='all-but-first'>
			//   <li :range>repat</li>
			//   <li>xxx</li>
			// </ul>
			// 多个 li 【之间】才需要空白。最后一个项目之后的空白，会在下一个 node 处理
			if nextSiblingBlankTextNode != nil {
				if err := t.execute(nextSiblingBlankTextNode, tokenBuf, childScope, nil); err != nil {
					return err
				}
			}
		}
		if err := t.execute(node, tokenBuf, childScope, nil); err != nil {
			return err
		}
		count++

	}
	return nil
}

// extractRange 解析循环表达式语法 index, item : items
func extractRange(s string) (
	idxName string, itemName string, objName string, err error,
) {
	s = strings.TrimSpace(s)
	i := strings.Index(s, ":")
	if i < 0 { // 无 idxName, itemName :range="items"
		objName = s
		return
	}
	objName = strings.TrimSpace(s[i+1:])
	s = s[:i]
	i = strings.Index(s, ",")
	if i < 0 { // 无 itemName :range="key : items"
		idxName = strings.TrimSpace(s)
		return
	}
	// :range="idxName, itemName : items"
	idxName = strings.TrimSpace(s[:i])
	itemName = strings.TrimSpace(s[i+1:])
	return
}

// processRemoveAttr 处理 remove 属性
func (t *htmlTemplate) processRemoveAttr(node *Node, attr *Attr, data exp.Scope, opt *executeOption) {
	var av string
	if attr.Value != nil {
		av = *attr.Value
	}
	switch av {
	case removeAll1, removeAll2:
		opt.noPrintToken = true
		opt.processChild = nop
	case removeBody1, removeBody2:
		opt.processChild = nop
	case removeTag1, removeTag2:
		opt.noPrintToken = true
	case removeAllButFirst1, removeAllButFirst2:
		if opt.processChild != nil {
			return
		}
		// 只输出第一个子tag，如果本节点开头和最后一个节点是空白文本，则也输出
		opt.processChild = func(w io.Writer) error {
			var tagIndex int
			var blankTextBefore, tagNode, blankTextAfter *Node
			for i, child := range node.Children {
				if token := child.Token; token != nil && token.Kind == TokenKindTag {
					tagIndex = i
					tagNode = child
					break
				}
			}
			if tagIndex > 0 {
				child := node.Children[0] // 第一个子节点看是否是空白文本
				if child.IsBlankText() {
					blankTextBefore = child
				}
			}
			if count := len(node.Children); count > 0 {
				child := node.Children[count-1] // 最后一个子节点看是否空白文本
				if child.IsBlankText() {
					blankTextAfter = child
				}
			}
			if blankTextBefore != nil {
				if err := t.execute(blankTextBefore, w, data, nil); err != nil {
					return err
				}
			}
			if err := t.execute(tagNode, w, data, nil); err != nil {
				return err
			}
			if blankTextAfter != nil {
				if err := t.execute(blankTextAfter, w, data, nil); err != nil {
					return err
				}
			}
			return nil
		}
	default:
	}
}

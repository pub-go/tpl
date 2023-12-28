package html

import (
	"fmt"
	"html"
	"io"
	"reflect"
	"strings"

	"code.gopub.tech/tpl/internal/exp"
)

// htmlTemplate HTML 模板
type htmlTemplate struct {
	textTags     []string
	voidElements []string
	tagPrefix    string
	attrPrefix   string
	files        map[string]*Node
	current      string
}

// NewHtmlTemplate 新建一个 HTML 模板实例
func NewHtmlTemplate() *htmlTemplate {
	return (&htmlTemplate{
		files: map[string]*Node{},
	}).
		SetTextTags(GetDefaultTextTags()).
		SetVoidElements(GetDefaultVoidElements()).
		SetTagPrefix(DefaultTagPrefix).
		SetAttrPrefix(DefaultAttrPrefix)
}

// SetTextTags 设置只包含文本的标签
// @param textTags 只包含文本的标签, 如 script, title, style, textarea
func (h *htmlTemplate) SetTextTags(textTags []string) *htmlTemplate {
	if textTags != nil {
		h.textTags = textTags
	}
	return h
}

// SetVoidElements 设置空标签
// @param voidElements 空标签不包含闭合斜线也不包含内容, 如 meta, br, img
func (h *htmlTemplate) SetVoidElements(voidElements []string) *htmlTemplate {
	if voidElements != nil {
		h.voidElements = voidElements
	}
	return h
}

// SetTagPrefix 设置标签前缀
// @param prefix 要设置的前缀
func (h *htmlTemplate) SetTagPrefix(prefix string) *htmlTemplate {
	if prefix != "" {
		h.tagPrefix = prefix
	}
	return h
}

// SetAttrPrefix 设置属性前缀
// @param prefix 要设置的前缀
func (h *htmlTemplate) SetAttrPrefix(prefix string) *htmlTemplate {
	if prefix != "" {
		h.attrPrefix = prefix
	}
	return h
}

// Add 添加模板文件
// @param name 文件名
// @param reader 文件内容
func (h *htmlTemplate) Add(name string, reader io.Reader) error {
	tz := NewHtmlScanner(reader).
		SetTextTags(h.textTags).
		SetAttrPrefix(h.attrPrefix)
	tokens, err := tz.GetAllTokens()
	if err != nil {
		return fmt.Errorf("failed to read html tokens: %w", err)
	}
	p := NewParser().SetVoidElements(h.voidElements)
	tree, err := p.ParseTokens(tokens)
	if err != nil {
		return fmt.Errorf("failed to parse html tokens: %w", err)
	}
	h.files[name] = tree
	h.current = name
	return nil
}

// Execute 执行模板
//
// 功能
//   - t:text 文本替换。
//     示例 t:text="Hello, ${name}" t:text='Hello, ${name}'
//     说明 引号内默认是字面文本 ${} 内才是代码
//     t:class="${idx%2==0 ? 'even' : 'odd'}"
//     t:text="${user.name ?: t.T('No username')}"
//   - t:html html 替换。 t:html="<span>xxxx</span>" t:html='${item.content}'
//   - t:if t:else 条件执行。 t:if="true" t:else="${cond}" t:if="${ (len(list) gt 0) and (not isAdmin) }"
//   - t:range 循环。t:range="${ idx, item : array }"
//   - t:remove 移除。 t:remove="all-but-first"
//   - t:define t:insert t:replace <span t:define="name(attr)"></span> <div th:insert="file#name(attr)"></div>
//
// [Thymeleaf](https://www.thymeleaf.org/doc/tutorials/3.1/usingthymeleaf.html)
//
// 翻译文本
//   - ❌ <p th:text="#{home.welcome}">Welcome</p>
//   - ✅ <p th:text="${t.X('home page', 'Welcome')}">Welcome</p>
//
// HTML 文本
//   - ❌ <p th:utext="#{home.welcome}">Welcome</p>
//   - ✅ <p th:html="${t.X('home page', '<strong>Welcome</strong>')}">Welcome</p>
//
// 变量
//   - ✅ <span th:text="${today}">2023-10-31</span>
//   - ${...} 变量表达式 从全局作用域查找变量
//   - *{...} 选定变量表达式 搭配 th:object 使用
//   - #{...} 消息变量表达式
//   - @{...} URL 链接表达式 绝对 url、相对 url
//   - ~{...} 片段表达式
//
// 字面量
//   - 文本 'one text', ...
//   - 数字 0, 3.14, ...
//   - 布尔 true, false
//   - null
//   - 字面 token?: one, sometext, main, ...
//
// 文本操作符
//   - 加号拼接 +
//   - 字面替换 |The name is ${name}|
//
// 数学操作符
//   - 二元： +, -, *, /, %
//   - 一元： -
//
// 布尔操作符
//   - 二元 and, or
//   - 一元 not, !
//
// 比较操作
//   - gt, >, lt, <, ge, >=, le, <=
//   - eq, ==, ne, !=
//
// 条件操作符
//   - ?    <if> ? <then>
//   - ? :  <if> ? <then> : <else>
//   - ?:   <if> ?: <default>
//
// 特殊标记
//   - Nop 标记： _
//
// 示例
//   - 'User is of type ' + (${user.isAdmin()} ? 'Administrator' : (${user.type} ?: 'Unknown'))
func (h *htmlTemplate) Execute(w io.Writer, data any) error {
	tree := h.files[h.current]
	if tree == nil {
		return fmt.Errorf("template not found: %q", h.current)
	}
	return h.ExecuteTree(tree, w, exp.NewScope(data), nil)
}

// executeOption 执行模板的参数
type executeOption struct {
	noPrintToken bool // 不输出 tag 本身
	processChild func(w io.Writer) error
}

var notPrintChild = func(w io.Writer) error { return nil }

// ExecuteTree 执行一棵文档树
func (h *htmlTemplate) ExecuteTree(tree *Node, w io.Writer, data exp.Scope, opt *executeOption) error {
	if tree == nil {
		return nil
	}
	if opt == nil {
		opt = &executeOption{}
	}
	// 子节点处理函数
	if opt.processChild == nil {
		opt.processChild = func(w io.Writer) error {
			for _, child := range tree.Children {
				if err := h.ExecuteTree(child, w, data, nil); err != nil {
					return err
				}
			}
			return nil
		}
	}

	// 开始标签
	var (
		token    = tree.Token         // <p>child</p> ==> <p>是token, child 是子节点, </p> 是End
		tokenBuf = &strings.Builder{} // 开始标签缓冲区
	)
	if token != nil { // root doc is nil
		switch token.Kind {
		case TokenKindTag:
			if err := h.processTagStart(tree, tokenBuf, data, opt); err != nil {
				return fmt.Errorf("failed to process tag [%v](from postion %v to %v): %w",
					token.Value, token.Start, token.End, err)
			}
		case TokenKindComment:
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
func (h *htmlTemplate) processTagStart(node *Node, tokenBuf *strings.Builder,
	data exp.Scope, opt *executeOption) error {
	tag := node.Token.Tag
	if tag == nil {
		return fmt.Errorf("unexpected nil tag")
	}
	name := strings.ToLower(tag.Name)
	if name == h.tagPrefix+"block" {
		opt.noPrintToken = true // <t:block> ... </t:block>
	}
	hasRangeAttr := h.hasRangeAttr(tag)
	if hasRangeAttr {
		opt.noPrintToken = true
		// <li :range="">xxx<li> 不输出 tag 在处理 range 时再输出
	}
	var tagBuf = &strings.Builder{}
	writeToBuf(opt, tagBuf, "<"+tag.Name)

	m := tag.AttrMap()
	for _, attr := range tag.SortedAttr(h.attrPrefix) {
		attr := attr
		an := attr.Name
		if strings.HasPrefix(an, h.attrPrefix) { // 指令属性
			cmd := strings.TrimPrefix(an, h.attrPrefix)
			switch cmd {
			case "if", "else", "else-if": // 条件控制
			case "range": // 循环
				if err := h.processRange(node, attr, tokenBuf, data, opt); err != nil {
					return err
				}
			case "remove": // 移除
				h.processRemoveAttr(node, attr, data, opt)
			case "text", "raw": // 替换内容
				if hasRangeAttr {
					opt.processChild = notPrintChild
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
			case "define": // 定义可复用组件
			case "insert": // 插入组件作为内容
			case "replace": // 用组件替换本身
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
			if _, ok := m[h.attrPrefix+an]; !ok {
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

func (h *htmlTemplate) hasRangeAttr(tag *Tag) bool {
	m := tag.AttrMap()
	if attr, ok := m[h.attrPrefix+"range"]; ok {
		return attr.Value != nil && *attr.Value != ""
	}
	return false
}

// processRemoveAttr 处理 remove 属性
func (h *htmlTemplate) processRemoveAttr(node *Node, attr *Attr, data exp.Scope, opt *executeOption) {
	var av string
	if attr.Value != nil {
		av = *attr.Value
	}
	switch av {
	case `"all"`, `'all'`:
		opt.noPrintToken = true
		opt.processChild = notPrintChild
	case `"body"`, `'body'`:
		opt.processChild = notPrintChild
	case `"tag"`, `'tag'`:
		opt.noPrintToken = true
	case `"all-but-first"`, `'all-but-first'`:
		// 只输出第一个tag，如果之后有空白文本也输出
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
				if err := h.ExecuteTree(blankTextBefore, w, data, nil); err != nil {
					return err
				}
			}
			if err := h.ExecuteTree(tagNode, w, data, nil); err != nil {
				return err
			}
			if blankTextAfter != nil {
				if err := h.ExecuteTree(blankTextAfter, w, data, nil); err != nil {
					return err
				}
			}
			return nil
		}
	default:
	}
}

func writeToBuf(opt *executeOption, buf *strings.Builder, data string) {
	if !opt.noPrintToken {
		buf.WriteString(data)
	}
}

func (h *htmlTemplate) processRange(node *Node, attr *Attr, tokenBuf *strings.Builder, data exp.Scope, opt *executeOption) error {
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
		return err
	}
	rv := reflect.ValueOf(obj)
	rk := rv.Kind()
	var getRange = func() (any, any, bool) {
		return nil, nil, false
	}
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
		return fmt.Errorf("`%vrange` only supports array, slice, string or map, got %v: %v",
			h.attrPrefix, rk, objName)
	}

	// <ul>
	//   <li :range>当前 node 是这个</li>
	//   <li>
	// </ul>
	// 处理 range 的 li 项目时，找到并输出 </li> 后的空白字符
	var nextSiblingBlankTextNode *Node
	var nodeIndex int // 当前 node 在父节点的位置
	if node.Parent != nil {
		for i, child := range node.Parent.Children {
			if child == node {
				nodeIndex = i
				break
			}
		}

		if nodeIndex < len(node.Parent.Children)-1 {
			next := node.Parent.Children[nodeIndex+1]
			if next.IsBlankText() {
				nextSiblingBlankTextNode = next
			}
		}
	}

	*av = "" // 标记 range 指令已经执行 下面 for 循环执行 node 时就会忽略 range 指令了
	count := 0
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
				if err := h.ExecuteTree(nextSiblingBlankTextNode, tokenBuf, childScope, nil); err != nil {
					return err
				}
			}
		}
		if err := h.ExecuteTree(node, tokenBuf, childScope, nil); err != nil {
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

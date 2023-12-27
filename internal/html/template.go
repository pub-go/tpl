package html

import (
	"fmt"
	"html"
	"io"
	"strings"
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
	return h.ExecuteTree(tree, w, true, true, data)
}

// ExecuteTree 执行一棵文档树
func (h *htmlTemplate) ExecuteTree(tree *Node, w io.Writer, printToken bool, printChild bool, data any) error {
	if tree == nil {
		return nil
	}
	// text
	// cdata
	// comment
	// tag

	// :text
	// :raw
	// :if
	// :else-if
	// :else
	// :range
	// :remove
	// :define
	// :insert
	// :replace

	var buf = &strings.Builder{}
	var nop = func() error { return nil }
	var defaultProcessChild = func() error {
		for _, child := range tree.Children {
			if err := h.ExecuteTree(child, w, true, true, data); err != nil {
				return err
			}
		}
		return nil
	}
	var processChid = defaultProcessChild
	// 开始
	token := tree.Token
	if token != nil { // root doc is nil
		switch token.Kind {
		case TokenKindTag:
			tag := token.Tag
			if tag == nil {
				return fmt.Errorf("unexpected nil tag (from position %v to %v): %q",
					token.Start, token.End, token.Value)
			}
			var tokenBuf = &strings.Builder{}
			name := strings.ToLower(tag.Name)
			if name == h.tagPrefix+"block" {
				printToken = false // <t:block> ... </t:block>
			}
			m := tag.AttrMap()
			writeToBufIf(printToken, tokenBuf, "<"+tag.Name)
			for _, attr := range tag.Attrs {
				attr := attr
				an := attr.Name
				if strings.HasPrefix(an, h.attrPrefix) { // 指令属性
					cmd := strings.TrimPrefix(an, h.attrPrefix)
					switch cmd {
					case "text", "raw": // 替换内容
						processChid = func() error {
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
					case "if", "else", "else-if": // 条件控制
					case "range": // 循环
					case "remove": // 移除
						var av string
						if attr.Value != nil {
							av = *attr.Value
						}
						switch av {
						case `"all"`, `'all'`:
							printToken = false
							processChid = nop
						case `"body"`, `'body'`:
							processChid = nop
						case `"tag"`, `'tag'`:
							printToken = false
						case `"all-but-first"`, `'all-but-first'`:
						default:
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
						writeToBufIf(printToken, tokenBuf, result)
					}
				} else { // 普通属性
					if _, ok := m[h.attrPrefix+an]; !ok {
						// 如果有对应的指令属性则跳过 没有才输出
						if printToken {
							attr.Print(tokenBuf)
						}
					}
				}
			}
			writeToBufIf(printToken, tokenBuf, ">")
			writeToBufIf(printToken, buf, tokenBuf.String())
		case TokenKindComment:
			// <!-- /* */ -->
			value := strings.TrimPrefix(token.Value, "<!--")
			value = strings.TrimSuffix(value, "-->")
			value = strings.TrimSpace(value)
			if strings.HasPrefix(value, "/*") && strings.HasSuffix(value, "*/") {
				printToken = false
			}
			writeToBufIf(printToken, buf, token.Value)
		default: // text, cdata
			writeToBufIf(printToken, buf, token.Value)
		}
	}
	if _, err := w.Write([]byte(buf.String())); err != nil {
		return err
	}

	// 内容
	if err := processChid(); err != nil {
		return err
	}

	// 结束
	token = tree.End
	if token != nil && printToken {
		_, err := w.Write([]byte(token.Value))
		if err != nil {
			return err
		}
	}
	return nil
}

func writeToBufIf(cond bool, buf *strings.Builder, data string) {
	if cond {
		buf.WriteString(data)
	}
}

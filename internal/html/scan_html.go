package html

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"
)

// ErrUnexpectedEOF 非预期的 EOF. 读取标签过程中遇到 EOF 时会返回
var ErrUnexpectedEOF = errors.New("unecpected EOF")

// NewHtmlScanner 新建一个词法解析器实例
// 默认设置的文本 tag 有：["script", "style", "textarea", "title"]
func NewHtmlScanner(r io.Reader) *HtmlScanner {
	return (&HtmlScanner{
		BaseScanner: NewBaseScanner(r),
	}).
		SetTextTags(GetDefaultTextTags()).
		SetAttrPrefix(DefaultAttrPrefix)
}

// HtmlScanner HTML 词法解析器
type HtmlScanner struct {
	*BaseScanner
	textTags   []string // 只能包含文本 不能包含子元素的 tag
	attrPrefix string   // 属性前缀
	state      state    // 当前状态
	nextToken  *Token   // 下一个返回的 Token
	tokens     []*Token // 所有的 Tokens
}

// state 状态机
type state int

const (
	stateInit         state = iota // 开始状态
	stateText                      // 读取文本
	stateTagStart                  // 读取 Tag 下一个字符是 <
	stateTagName                   // 读取 Tag Name
	stateCDATA                     // 读取 CDATA 数据 <![CDATA[ 开头
	stateComment                   // 读取注释 <!-- 开头
	stateTagSpace                  // 读取 Tag 中的空白
	stateTagAttrName               // 读取 Tag 属性名
	stateTagAttrValue              // 读取 Tag 属性值
)

// SetTextTags 设置只能包含文本不能包含子标签的标签
// 如 ["script", "style", "textarea", "title"]
// 读取这些标签的文本内容时，允许包含 "</tagName" 但一旦遇到 "</tagName>" 会认为结束
// 如 <script>/*</script.*/</script>
func (s *HtmlScanner) SetTextTags(tagNames []string) *HtmlScanner {
	if tagNames != nil {
		s.textTags = tagNames
	}
	return s
}

func (s *HtmlScanner) SetAttrPrefix(prefix string) *HtmlScanner {
	s.attrPrefix = prefix
	return s
}

// GetAllTokens 解析所有词法单元
func (t *HtmlScanner) GetAllTokens() ([]*Token, error) {
	for t.err == nil && !t.done {
		t.NextToken()
	}
	if errors.Is(t.err, io.EOF) {
		t.err = nil
	}
	return t.tokens, t.err
}

// NextToken 解析一个词法单元
func (t *HtmlScanner) NextToken() (*Token, error) {
	// 如果有已经解析的下一个 token 直接消耗掉
	if t.nextToken != nil {
		// 示例：在解析 `<script>/*</script.*/</script>` 的文本内容时，遇到 `</script>` 才结束
		// 此时返回的是 Text token `/*</script.*/` 并设置 nextToken 为 `</script>`
		token := t.nextToken
		t.nextToken = nil
		return t.addToken(token), nil
	}
	// 如果是初始未知状态 读取一个字符看是不是 <
	if err := t.initState(); err != nil {
		return nil, err
	}
	switch t.state {
	case stateText: // 读取文本
		return t.readText()
	case stateTagStart: // 读取 tag
		return t.readTag()
	default:
		panic("unexpected state")
	}
}

func (s *HtmlScanner) addToken(tok *Token) *Token {
	s.tokens = append(s.tokens, tok)
	return tok
}

func (t *HtmlScanner) initState() error {
	if t.state == stateInit { // 初始状态 先读一个字符看是不是 <
		if err := t.NextRune(); err != nil {
			return err
		}
		if t.ch == '<' { // 如果是 < 说明是 tag
			t.state = stateTagStart
		} else {
			t.state = stateText
		}
		t.UnRead() // 还回去
	}
	return nil
}

func (t *HtmlScanner) readText() (tok *Token, err error) {
	var (
		textBuf bytes.Buffer // 文本缓冲区
		start   = t.pos      // 起始位置
		// 如果是正在读取文本标签内的文本 那么会顺带解析出闭合标签
		tagName, // tag 名。如 script
		inTextTag = t.isInRawTextTag() // 是否正在读取文本标签内的文本
		closeTag = "</" + tagName + ">" // 文本标签的闭合标签
		end      Pos                    // 文本的结束位置
		tagBuf   bytes.Buffer           // 闭合标签缓冲区
		nameBuf  bytes.Buffer           // 闭合标签名 用于和 closeTag 匹配
	)
	for {
		if err := t.NextRune(); err != nil {
			if errors.Is(err, io.EOF) {
				return t.addToken(&Token{
					Kind:  TokenKindText,
					Value: textBuf.String(),
					Start: start,
					End:   t.pos,
				}), nil
			}
			return nil, fmt.Errorf("read text failed from position %v to %v: %w",
				start, t.pos, err)
		}
		ch := t.ch
		if inTextTag {
			// 文本元素包含的文本可以有 < 字符。
			// 遇到 </tagName > 算结束。
			var closing = tagBuf.Len() > 0
			if ch == '<' {
				// 重置结束标签缓冲区
				tagBuf.Reset()
				nameBuf.Reset()
				closing = true
			}
			if !closing {
				// 记录结束标签前一个位置
				end = t.pos
			} else { // 已经有 < 了
				tagBuf.WriteRune(ch)
				if !unicode.IsSpace(ch) {
					// 用于比较 </tagName> 所以空白符号不写入
					nameBuf.WriteRune(ch)
				}

				name := strings.ToLower(nameBuf.String())
				closing = strings.HasPrefix(closeTag, name) // 如果不匹配开始标签 则说明不是结束标签

				if closing {
					if ch == '>' { // closed
						// +1: '>' 符号还没写入 textBuf, 但已写入 tagBuf
						// e.g. <script>a</script>
						// textBuf=a</script
						// tagBuf =</script>
						// after Truncate: textBuf=a
						textBuf.Truncate(textBuf.Len() + 1 - tagBuf.Len())
						textToken := t.addToken(&Token{
							Kind:  TokenKindText,
							Value: textBuf.String(),
							Start: start,
							End:   end,
						})
						tagToken := &Token{
							Kind:  TokenKindTag,
							Value: tagBuf.String(),
							Start: end,
							End:   t.pos,
							Tag: &Tag{
								Name: "/" + tagName,
								// 结束标签无属性
							},
						}
						t.nextToken = tagToken
						t.state = stateInit
						return textToken, nil
					}
				} else { // 不是 tag 结束，比如只是 `x<y` 的 "<y"
					tagBuf.Reset()
					nameBuf.Reset()
				}
			}
		} else if ch == '<' { // 其他元素内遇到 < 就当前一个新的 tag 开始
			t.state = stateTagStart // 状态扭转
			t.UnRead()              // 回退 < 符号
			return t.addToken(&Token{
				Kind:  TokenKindText,
				Value: textBuf.String(),
				Start: start,
				End:   t.pos,
			}), nil
		}

		// 写入文本
		textBuf.WriteRune(ch)
	}
}

// isInRawTextTag 正在读取文本 tag 里的文本
// 例如：正在读取 <script> 标签中的内容
func (t *HtmlScanner) isInRawTextTag() (string, bool) {
	size := len(t.tokens)
	if size == 0 {
		return "", false
	}
	last := t.tokens[size-1]
	if last.Kind != TokenKindTag || last.Tag == nil {
		return "", false
	}
	tagName := strings.ToLower(last.Tag.Name)
	for _, name := range t.textTags {
		if tagName == strings.ToLower(name) {
			return tagName, true
		}
	}
	return "", false
}

func (t *HtmlScanner) readTag() (tok *Token, err error) {
	var (
		buf            bytes.Buffer
		start          = t.pos
		tag            Tag
		tagName        bytes.Buffer
		commentText    bytes.Buffer
		cdata          bytes.Buffer
		attrName       bytes.Buffer
		attrNameStart  Pos
		attrNameEnd    Pos
		attrValue      bytes.Buffer
		attrValueStart Pos
		attrValueEnd   Pos
	)
	for {
		if err := t.NextRune(); err != nil {
			var tok *Token
			if errors.Is(err, io.EOF) {
				value := buf.String()
				// read tag 期间不应该提前结束 返回错误
				err = fmt.Errorf("read tag %q failed (from position %v to %v): %w",
					value, start, t.pos, ErrUnexpectedEOF)
			}
			t.err = err // 记录错误
			return tok, err
		}
		ch := t.ch
		buf.WriteRune(ch)

		switch t.state {
		case stateTagStart:
			if ch != '<' {
				panic(fmt.Sprintf("unexpected state, want '<', but got '%c'", ch))
			}
			t.state = stateTagName
		case stateTagName: // 读取 tagName
			if ch == '>' {
				break // tag 结束
			}
			if unicode.IsSpace(ch) {
				t.state = stateTagSpace
			} else { // 非空白字符都当做 tagName
				tagName.WriteRune(ch)
				name := tagName.String()
				if name == "!--" {
					t.state = stateComment
				}
				if name == "![CDATA[" {
					t.state = stateCDATA
				}
			}
		case stateComment:
			// https://html.spec.whatwg.org/multipage/syntax.html#syntax-comments
			// 1. "<!--" 开头
			// 2. 可选的 text, 要求：
			//    不能以 ">", "->" 开头；
			//    不能包含 "<!--", "-->", "--!>"；
			//    不能以 "<!-" 结尾
			// 3. “-->” 结尾
			commentText.WriteRune(ch)
			text := commentText.String()
			isEnd := strings.HasSuffix(text, "-->")
			text = strings.TrimSuffix(text, "-->")
			if strings.HasPrefix(text, ">") || strings.HasPrefix(text, "->") {
				err = fmt.Errorf(`read comment failed (from position %v to %v): comment text must not start with ">" or "->": %s`,
					start, t.pos, text)
				t.err = err
				return nil, err
			}
			if isEnd {
				if strings.Contains(text, "<!--") || strings.Contains(text, "-->") || strings.Contains(text, "--!>") {
					err = fmt.Errorf(`read comment failed (from position %v to %v): comment text must not contains "<!--", "-->" or "--!>": %s`,
						start, t.pos, text)
					t.err = err
					return nil, err
				}
				if strings.HasSuffix(text, "<!-") {
					err = fmt.Errorf(`read comment failed (from position %v to %v): comment text must not end with "<!-": %s`,
						start, t.pos, text)
					t.err = err
					return nil, err
				}
				t.state = stateInit // 状态扭转
				return t.addToken(&Token{
					Kind:  TokenKindComment,
					Value: "<!--" + text + "-->",
					Start: start,
					End:   t.pos,
				}), nil
			}
			continue
		case stateCDATA:
			cdata.WriteRune(ch)
			text := cdata.String()
			if strings.HasSuffix(text, "]]>") {
				t.state = stateInit // 状态扭转
				return t.addToken(&Token{
					Kind:  TokenKindCDATA,
					Value: "<![CDATA[" + text,
					Start: start,
					End:   t.pos,
				}), nil
			}
		case stateTagSpace: // 读取 tag 中的空白
			if ch == '>' {
				break // tag 结束
			}
			if !unicode.IsSpace(ch) { // 直到非空白字符
				t.UnRead()                         // 回退字符
				buf.Truncate(buf.Len() - t.chSize) // buf 中也回退

				t.state = stateTagAttrName // 状态扭转到读取属性名
				attrName.Reset()
				attrNameStart = t.pos
			}
		case stateTagAttrName: // 读取 attrName 属性名
			// <input disabled value = yes>
			//                ^     ^
			//                |     |-读取到‘=’结束
			//                -读取到非空白字符结束
			// <input disabled> -- 读取到 ‘>’ 结束
			// <input value=yes> -- 读取到 '=' 结束

			if unicode.IsSpace(ch) {
				attrName.WriteRune(' ')
			} else if ch == '>' {
				name := strings.TrimSuffix(attrName.String(), " ")
				attr := &Attr{
					Name:      name,
					NameStart: attrNameStart,
					NameEnd:   attrNameEnd,
				} // <p name>
				t.compileAttr(attr)
				if err := tag.AddAttr(attr); err != nil {
					return nil, t.Err("scan attr failed:%v. cause: %w", attr, err)
				}
				// 转到 state switch 后面的 > 字符判断去结束 tag 并转换 state 为 init
			} else if ch == '=' {
				t.state = stateTagAttrValue
				attrValue.Reset()
				attrValueStart = t.pos
			} else {
				name := attrName.String()
				if strings.HasSuffix(name, " ") {
					attr := &Attr{
						Name:      strings.TrimSuffix(name, " "),
						NameStart: attrNameStart,
						NameEnd:   attrNameEnd,
					} // <p foo bar> 读取到 b
					t.compileAttr(attr)
					if err := tag.AddAttr(attr); err != nil {
						return nil, t.Err("scan attr failed: %v. cause: %w", attr, err)
					}
					attrName.Reset()
					t.UnRead()
					attrNameStart = t.pos
					t.NextRune()
				}
				attrName.WriteRune(ch)
				attrNameEnd = t.pos
			}
		case stateTagAttrValue: // 读取 attrValue 属性值
			// <input value= 'a b c'>
			// <input value= 'a "b" c'>
			if attrValue.Len() == 0 { // 刚开始读取 value
				if unicode.IsSpace(ch) {
					attrValueStart = t.pos
					continue // 前导空白都忽略掉
				}
				attrValue.WriteRune(ch)
				attrValueEnd = t.pos
			} else { // 已经有 value 了
				firstCh := []rune(attrValue.String())[0]
				var finish bool
				if firstCh == '"' || firstCh == '\'' {
					if firstCh == ch { // 引号开头 现在又读取到引号 说明结束了
						attrValue.WriteRune(ch)
						attrValueEnd = t.pos
						finish = true
					}
				} else if unicode.IsSpace(ch) || ch == '>' {
					// 非引号开头 现在读取到空格或'>' 也说明结束了
					finish = true
				}
				if finish {
					value := attrValue.String()
					attr := &Attr{
						Name:       strings.TrimSuffix(attrName.String(), " "),
						NameStart:  attrNameStart,
						NameEnd:    attrNameEnd,
						Value:      &value,
						ValueStart: attrValueStart,
						ValueEnd:   attrValueEnd,
					}
					if err := t.compileAttr(attr); err != nil {
						return nil, t.Err("compile attr failed: %v. cause: %w", attr, err)
					}
					if err := tag.AddAttr(attr); err != nil {
						return nil, t.Err("scan attr failed:%v. cause: %w", attr, err)
					}
					// 去消耗掉连续的空白
					t.state = stateTagSpace
					// 如果是 '>' 执行到下方重新转 init
				} else {
					attrValue.WriteRune(ch)
					attrValueEnd = t.pos
					if firstCh == '"' || firstCh == '\'' {
						continue // 有引号时 不需要下方 > 判断
					}
				}
			}
		}

		if ch == '>' { // Tag 结束
			t.state = stateInit // 状态扭转
			tag.Name = tagName.String()
			return t.addToken(&Token{
				Kind:  TokenKindTag,
				Value: buf.String(),
				Start: start,
				End:   t.pos,
				Tag:   &tag,
			}), nil
		}
	}
}

func (t *HtmlScanner) compileAttr(attr *Attr) error {
	if attr.Value == nil && attr.Name == t.attrPrefix+"else" { // <p :else> ==> <p :else="true">
		t := `"true"`
		attr.Value = &t
	}
	if attr.Value == nil || !strings.HasPrefix(attr.Name, t.attrPrefix) {
		return nil
	}
	s := NewCodeScanner(attr.ValueStart, *attr.Value)
	toks, err := s.GetAllTokens()
	if err != nil {
		return err
	}
	attr.ValueTokens = toks
	return nil
}

package html

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"code.gopub.tech/tpl/internal/exp"
)

// ErrUnexpectedEOF 非预期的 EOF. 读取标签过程中遇到 EOF 时会返回
var ErrUnexpectedEOF = errors.New("unecpected EOF")

type Pos = exp.Pos

// Token 词法单元
type Token struct {
	Kind  TokenKind // 类型
	Value string    // 内容
	Start Pos       // 开始位置
	End   Pos       // 结束位置
	Tag   *Tag      // 如果是标签
}

func (t *Token) String() string {
	return fmt.Sprintf("{Kind=%s, Value=%q, Start=%v, End=%v, Tag=%v}",
		t.Kind, t.Value, t.Start, t.End, t.Tag)
}

// TokenKind 类型
type TokenKind int

const (
	TokenKindError   TokenKind = iota // 错误
	TokenKindTag                      // Tag 类型
	TokenKindText                     // Text 文本类型
	TokenKindComment                  // 注释
	TokenKindCDATA                    // CDATA
)

func (p TokenKind) String() string {
	switch p {
	case TokenKindTag:
		return "Tag"
	case TokenKindText:
		return "Text"
	case TokenKindComment:
		return "Comment"
	case TokenKindCDATA:
		return "CDATA"
	default:
		return "Error"
	}
}

// Tag <标签> 如果是结束标签则 Name 以 / 开头
type Tag struct {
	Name  string
	Attrs []Attr
	attrs map[string]Attr
}

func (t Tag) AttrMap() map[string]Attr {
	if t.attrs == nil {
		m := make(map[string]Attr, len(t.Attrs))
		for _, attr := range t.Attrs {
			m[attr.Name] = attr
		}
		t.attrs = m
	}
	return t.attrs
}

// IsClose 是否是闭合标签
// 包含两种：正常的闭合标签 </name>, 自闭合标签 <name />
func (t Tag) IsClose() bool {
	// </name>
	if strings.HasPrefix(t.Name, "/") {
		return true
	}
	return t.IsSelfClose()
}

// IsSelfClose 是否是自闭合标签
// 注意 <meta> 这种 void 标签不属于自闭合标签 只有 <name/> 这种以 / 结尾的才算
// 包含三种：无属性的 <name/>; 无属性值的 <name />; 有属性值的 <name k=v/>
func (t Tag) IsSelfClose() bool {
	as := len(t.Attrs)
	if as == 0 {
		// <xx/>
		return strings.HasSuffix(t.Name, "/")
	}
	a := t.Attrs[as-1]
	if a.Value == nil {
		// <xx />
		// <xx n/>
		return strings.HasSuffix(a.Name, "/")
	}
	// <xx k=v/>
	return strings.HasSuffix(*a.Value, "/")
}

func (t Tag) String() string {
	var sb strings.Builder
	sb.WriteString("<")
	sb.WriteString(t.Name)
	for _, attr := range t.Attrs {
		sb.WriteString(" ")
		sb.WriteString(attr.String())
	}
	sb.WriteString(">")
	return sb.String()
}

// Attr 属性
type Attr struct {
	Name        string       // 属性名
	NameStart   Pos          // 开始位置
	NameEnd     Pos          // 结束位置
	Value       *string      // 属性值 如果有
	ValueStart  Pos          // 开始位置
	ValueEnd    Pos          // 结束位置
	ValueTokens []*CodeToken // 解析后的属性值
}

func (a Attr) Print(w io.Writer) error {
	_, err := w.Write([]byte(" " + a.Name))
	if err != nil {
		return err
	}
	if a.Value != nil {
		_, err = w.Write([]byte("=" + *a.Value))
		if err != nil {
			return err
		}
	}
	return nil
}

func (a Attr) String() string {
	var sb strings.Builder
	sb.WriteString(a.NameStart.String())
	sb.WriteString("|")
	sb.WriteString(a.Name)
	sb.WriteString("|")
	sb.WriteString(a.NameEnd.String())
	if a.Value != nil {
		sb.WriteString("=")
		sb.WriteString(a.ValueStart.String())
		sb.WriteString("|")
		sb.WriteString(*a.Value)
		sb.WriteString("|")
		sb.WriteString(a.ValueEnd.String())
	}
	return sb.String()
}

func (a Attr) Evaluate(input any) (string, error) {
	if a.Value == nil {
		return "", fmt.Errorf("no value")
	}
	if len(a.ValueTokens) == 0 {
		return *a.Value, nil
	}
	var buf strings.Builder
	for _, tok := range a.ValueTokens {
		switch tok.Kind {
		case BegEnd:
		case Literal:
			buf.WriteString(tok.Value)
		case CodeStart: // ${
		case CodeValue:
			result, err := exp.Evaluate(tok.Start, tok.Tree, input)
			if err != nil {
				return "", err
			}
			buf.WriteString(fmt.Sprintf("%v", result))
		case CodeEnd: // }
		}
	}
	return buf.String(), nil
}

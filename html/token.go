package html

import (
	"fmt"
	"strings"

	"code.gopub.tech/tpl/exp"
)

type Pos = exp.Pos

// Token 词法单元
type Token struct {
	Kind  TokenKind // 类型
	Value string    // 内容
	Start Pos       // 开始位置
	End   Pos       // 结束位置
	Tag   *Tag      // 如果是标签
}

func (t *Token) IsBlankText() bool {
	if t.Kind == TokenKindText {
		return strings.TrimSpace(t.Value) == ""
	}
	return false
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

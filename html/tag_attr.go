package html

import (
	"fmt"
	"strings"

	"code.gopub.tech/tpl/exp"
)

// ErrAttrValueExpected 没有属性值
var ErrAttrValueExpected = fmt.Errorf("attribute value expected")

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

// Evaluate 在给定的作用域上计算属性值
func (a *Attr) Evaluate(input exp.Scope) (string, error) {
	if a.Value == nil {
		return "", fmt.Errorf(attrShouldHaveValue+": %w",
			a.Name, a.NameEnd, ErrAttrValueExpected)
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
				return "", fmt.Errorf("failed to evaluate %v attribute [%v]: %w",
					a.Name, *a.Value, err)
			}
			buf.WriteString(fmt.Sprintf("%v", result))
		case CodeEnd: // }
		}
	}
	return buf.String(), nil
}

// String 打印输出 不保证格式 debug only
func (a *Attr) String() string {
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

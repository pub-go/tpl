package html

import (
	"fmt"
	"strings"

	"code.gopub.tech/errors"
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
		return "", errors.Errorf(attrShouldHaveValue+": %w",
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
				return "", errors.Errorf("failed to evaluate %v attribute [%v]: %w",
					a.Name, *a.Value, err)
			}
			buf.WriteString(fmt.Sprintf("%v", result))
		case CodeEnd: // }
		}
	}
	return buf.String(), nil
}

// WithAssign 解析 :with="name := ${value}" 赋值属性 返回解析后的变量名，变量值
func (a *Attr) WithAssign(input exp.Scope) (string, any, error) {
	if a.Value == nil {
		return "", nil, errors.Errorf(attrShouldHaveValue+": %w",
			a.Name, a.NameEnd, ErrAttrValueExpected)
	}
	var name string
	var codes []*CodeToken
	for _, tok := range a.ValueTokens {
		switch tok.Kind {
		case Literal:
			if name != "" { // 多个变量暂不支持
				return "", nil, errors.Errorf("only support one variable, already found `%v`, got `%v`", name, tok.Value)
			}
			if len(codes) > 0 { // 先出现了代码 :with="${xxx}abc" 是异常情况
				return "", nil, errors.Errorf("variable name should not be after the code block", name)
			}
			name = strings.TrimSpace(tok.Value)
			if !strings.HasSuffix(name, ":=") {
				return "", nil, errors.Errorf("an assignment symbol(:=) should be after the name(`%v`)", name)
			}
			name = strings.TrimSpace(strings.TrimSuffix(name, ":="))

		case CodeValue:
			if name == "" {
				return "", nil, errors.Errorf("no variable name found before the code block: %v", tok)
			}
			codes = append(codes, tok)
		}
	}

	if len(codes) == 0 {
		return "", nil, errors.Errorf("code block(`${}`) not found: %v", a)
	}
	if len(codes) != 1 {
		return "", nil, errors.Errorf("too much code block(`${}`) found in `with` attr: %v", a)
	}
	tok := codes[0]
	result, err := exp.Evaluate(tok.Start, tok.Tree, input)
	return name, result, err
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

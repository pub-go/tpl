package exp

import (
	"code.gopub.tech/tpl/internal/exp/parser"
)

// Evaluate 对语法树求值
func Evaluate(startPos Pos, tree parser.IExpressionContext, data any) (any, error) {
	return NewVisitor(startPos, NewScope(data)).Evaluate(tree)
}

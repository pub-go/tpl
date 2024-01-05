package exp

import (
	"code.gopub.tech/tpl/exp/parser"
)

// Evaluate 对语法树求值
func Evaluate(startPos Pos, tree parser.IExpressionContext, data Scope) (any, error) {
	return NewVisitor(startPos, data).Evaluate(tree)
}

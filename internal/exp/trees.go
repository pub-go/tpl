package exp

import (
	"strings"

	"code.gopub.tech/tpl/internal/exp/parser"
	"github.com/antlr4-go/antlr/v4"
)

// ToString 打印语法树的工具函数
func ToString(tree parser.IExpressionContext) string {
	return toString(0, tree)
}

func toString(level int, tree antlr.Tree) string {
	var buf strings.Builder
	buf.WriteString("\n|-")
	buf.WriteString(strings.Repeat("-", level))
	buf.WriteString(" ")

	s := antlr.TreesGetNodeText(tree, parser.GoExpressionParserStaticData.RuleNames, nil)
	buf.WriteString(s)

	for _, child := range tree.GetChildren() {
		buf.WriteString(toString(level+1, child))
	}
	return buf.String()
}

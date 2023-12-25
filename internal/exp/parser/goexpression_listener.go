// Code generated from GoExpression.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // GoExpression
import "github.com/antlr4-go/antlr/v4"

// GoExpressionListener is a complete listener for a parse tree produced by GoExpression.
type GoExpressionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPrimaryExpr is called when entering the primaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterLiteralNil is called when entering the literalNil production.
	EnterLiteralNil(c *LiteralNilContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterLiteralFloat is called when entering the literalFloat production.
	EnterLiteralFloat(c *LiteralFloatContext)

	// EnterLiteralImag is called when entering the literalImag production.
	EnterLiteralImag(c *LiteralImagContext)

	// EnterOperandName is called when entering the operandName production.
	EnterOperandName(c *OperandNameContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterSlice is called when entering the slice production.
	EnterSlice(c *SliceContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPrimaryExpr is called when exiting the primaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitLiteralNil is called when exiting the literalNil production.
	ExitLiteralNil(c *LiteralNilContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitLiteralFloat is called when exiting the literalFloat production.
	ExitLiteralFloat(c *LiteralFloatContext)

	// ExitLiteralImag is called when exiting the literalImag production.
	ExitLiteralImag(c *LiteralImagContext)

	// ExitOperandName is called when exiting the operandName production.
	ExitOperandName(c *OperandNameContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitSlice is called when exiting the slice production.
	ExitSlice(c *SliceContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)
}

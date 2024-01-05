// Code generated from GoExpression.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // GoExpression
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GoExpression.
type GoExpressionVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GoExpression#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by GoExpression#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by GoExpression#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by GoExpression#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by GoExpression#literalNil.
	VisitLiteralNil(ctx *LiteralNilContext) interface{}

	// Visit a parse tree produced by GoExpression#integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by GoExpression#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by GoExpression#literalFloat.
	VisitLiteralFloat(ctx *LiteralFloatContext) interface{}

	// Visit a parse tree produced by GoExpression#literalImag.
	VisitLiteralImag(ctx *LiteralImagContext) interface{}

	// Visit a parse tree produced by GoExpression#operandName.
	VisitOperandName(ctx *OperandNameContext) interface{}

	// Visit a parse tree produced by GoExpression#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by GoExpression#index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by GoExpression#slice.
	VisitSlice(ctx *SliceContext) interface{}

	// Visit a parse tree produced by GoExpression#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by GoExpression#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}
}

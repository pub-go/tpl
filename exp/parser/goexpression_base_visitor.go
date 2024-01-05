// Code generated from GoExpression.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // GoExpression
import "github.com/antlr4-go/antlr/v4"

type BaseGoExpressionVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGoExpressionVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoExpressionVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoExpressionVisitor) VisitOperand(ctx *OperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoExpressionVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoExpressionVisitor) VisitLiteralNil(ctx *LiteralNilContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoExpressionVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoExpressionVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoExpressionVisitor) VisitLiteralFloat(ctx *LiteralFloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoExpressionVisitor) VisitLiteralImag(ctx *LiteralImagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoExpressionVisitor) VisitOperandName(ctx *OperandNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoExpressionVisitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoExpressionVisitor) VisitIndex(ctx *IndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoExpressionVisitor) VisitSlice(ctx *SliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoExpressionVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoExpressionVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

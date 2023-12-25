// Code generated from GoExpression.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // GoExpression
import "github.com/antlr4-go/antlr/v4"

// BaseGoExpressionListener is a complete listener for a parse tree produced by GoExpression.
type BaseGoExpressionListener struct{}

var _ GoExpressionListener = &BaseGoExpressionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGoExpressionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGoExpressionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGoExpressionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGoExpressionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGoExpressionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGoExpressionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BaseGoExpressionListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BaseGoExpressionListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseGoExpressionListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseGoExpressionListener) ExitOperand(ctx *OperandContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseGoExpressionListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseGoExpressionListener) ExitLiteral(ctx *LiteralContext) {}

// EnterLiteralNil is called when production literalNil is entered.
func (s *BaseGoExpressionListener) EnterLiteralNil(ctx *LiteralNilContext) {}

// ExitLiteralNil is called when production literalNil is exited.
func (s *BaseGoExpressionListener) ExitLiteralNil(ctx *LiteralNilContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaseGoExpressionListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaseGoExpressionListener) ExitInteger(ctx *IntegerContext) {}

// EnterString is called when production string is entered.
func (s *BaseGoExpressionListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseGoExpressionListener) ExitString(ctx *StringContext) {}

// EnterLiteralFloat is called when production literalFloat is entered.
func (s *BaseGoExpressionListener) EnterLiteralFloat(ctx *LiteralFloatContext) {}

// ExitLiteralFloat is called when production literalFloat is exited.
func (s *BaseGoExpressionListener) ExitLiteralFloat(ctx *LiteralFloatContext) {}

// EnterLiteralImag is called when production literalImag is entered.
func (s *BaseGoExpressionListener) EnterLiteralImag(ctx *LiteralImagContext) {}

// ExitLiteralImag is called when production literalImag is exited.
func (s *BaseGoExpressionListener) ExitLiteralImag(ctx *LiteralImagContext) {}

// EnterOperandName is called when production operandName is entered.
func (s *BaseGoExpressionListener) EnterOperandName(ctx *OperandNameContext) {}

// ExitOperandName is called when production operandName is exited.
func (s *BaseGoExpressionListener) ExitOperandName(ctx *OperandNameContext) {}

// EnterField is called when production field is entered.
func (s *BaseGoExpressionListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseGoExpressionListener) ExitField(ctx *FieldContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseGoExpressionListener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseGoExpressionListener) ExitIndex(ctx *IndexContext) {}

// EnterSlice is called when production slice is entered.
func (s *BaseGoExpressionListener) EnterSlice(ctx *SliceContext) {}

// ExitSlice is called when production slice is exited.
func (s *BaseGoExpressionListener) ExitSlice(ctx *SliceContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseGoExpressionListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseGoExpressionListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseGoExpressionListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseGoExpressionListener) ExitExpressionList(ctx *ExpressionListContext) {}

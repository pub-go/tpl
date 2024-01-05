// Generated from /data00/home/youthlin.chen/go/src/code.gopub.tech/tpl/exp/parser/GoExpression.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link GoExpression}.
 */
public interface GoExpressionListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link GoExpression#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(GoExpression.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoExpression#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(GoExpression.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoExpression#primaryExpr}.
	 * @param ctx the parse tree
	 */
	void enterPrimaryExpr(GoExpression.PrimaryExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoExpression#primaryExpr}.
	 * @param ctx the parse tree
	 */
	void exitPrimaryExpr(GoExpression.PrimaryExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoExpression#operand}.
	 * @param ctx the parse tree
	 */
	void enterOperand(GoExpression.OperandContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoExpression#operand}.
	 * @param ctx the parse tree
	 */
	void exitOperand(GoExpression.OperandContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoExpression#literal}.
	 * @param ctx the parse tree
	 */
	void enterLiteral(GoExpression.LiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoExpression#literal}.
	 * @param ctx the parse tree
	 */
	void exitLiteral(GoExpression.LiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoExpression#literalNil}.
	 * @param ctx the parse tree
	 */
	void enterLiteralNil(GoExpression.LiteralNilContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoExpression#literalNil}.
	 * @param ctx the parse tree
	 */
	void exitLiteralNil(GoExpression.LiteralNilContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoExpression#integer}.
	 * @param ctx the parse tree
	 */
	void enterInteger(GoExpression.IntegerContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoExpression#integer}.
	 * @param ctx the parse tree
	 */
	void exitInteger(GoExpression.IntegerContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoExpression#string}.
	 * @param ctx the parse tree
	 */
	void enterString(GoExpression.StringContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoExpression#string}.
	 * @param ctx the parse tree
	 */
	void exitString(GoExpression.StringContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoExpression#literalFloat}.
	 * @param ctx the parse tree
	 */
	void enterLiteralFloat(GoExpression.LiteralFloatContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoExpression#literalFloat}.
	 * @param ctx the parse tree
	 */
	void exitLiteralFloat(GoExpression.LiteralFloatContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoExpression#literalImag}.
	 * @param ctx the parse tree
	 */
	void enterLiteralImag(GoExpression.LiteralImagContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoExpression#literalImag}.
	 * @param ctx the parse tree
	 */
	void exitLiteralImag(GoExpression.LiteralImagContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoExpression#operandName}.
	 * @param ctx the parse tree
	 */
	void enterOperandName(GoExpression.OperandNameContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoExpression#operandName}.
	 * @param ctx the parse tree
	 */
	void exitOperandName(GoExpression.OperandNameContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoExpression#field}.
	 * @param ctx the parse tree
	 */
	void enterField(GoExpression.FieldContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoExpression#field}.
	 * @param ctx the parse tree
	 */
	void exitField(GoExpression.FieldContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoExpression#index}.
	 * @param ctx the parse tree
	 */
	void enterIndex(GoExpression.IndexContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoExpression#index}.
	 * @param ctx the parse tree
	 */
	void exitIndex(GoExpression.IndexContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoExpression#slice}.
	 * @param ctx the parse tree
	 */
	void enterSlice(GoExpression.SliceContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoExpression#slice}.
	 * @param ctx the parse tree
	 */
	void exitSlice(GoExpression.SliceContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoExpression#arguments}.
	 * @param ctx the parse tree
	 */
	void enterArguments(GoExpression.ArgumentsContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoExpression#arguments}.
	 * @param ctx the parse tree
	 */
	void exitArguments(GoExpression.ArgumentsContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoExpression#expressionList}.
	 * @param ctx the parse tree
	 */
	void enterExpressionList(GoExpression.ExpressionListContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoExpression#expressionList}.
	 * @param ctx the parse tree
	 */
	void exitExpressionList(GoExpression.ExpressionListContext ctx);
}
parser grammar GoExpression;

options {
	tokenVocab = GoLexer;
}

// 表达式
expression:
	primaryExpr                          // 基本表达式
	| unary_op = (
		PLUS
		| MINUS
		| EXCLAMATION
		| CARET
		| STAR
		| AMPERSAND
		| RECEIVE
	) expression                         // 一元表达式
	| expression mul_op = (
		STAR
		| DIV
		| MOD
		| LSHIFT
		| RSHIFT
		| AMPERSAND
		| BIT_CLEAR
	) expression                         // 二元表达式 乘除高优
	| expression add_op = (
		PLUS | MINUS | OR | CARET
	) expression                         // 二元表达式 加减中优
	| expression rel_op = (
		EQUALS
		| NOT_EQUALS
		| LESS
		| LESS_OR_EQUALS
		| GREATER
		| GREATER_OR_EQUALS
	) expression                         // 二元表达式 关系低优
	| expression LOGICAL_AND expression
	| expression LOGICAL_OR expression
    | expression Question expression COLON expression;

// 基本表达式
primaryExpr:
	operand            // 操作数
	| primaryExpr (
		field       // 取字段 struct.Field
		| index     // 取下标 array[index]
		| slice     // 取切片 array[:]
		| arguments // 取调用 function(args...)
	);                 // 基本表达式+后缀

// 操作数
operand:
    literal                        // 字面量
    | operandName                  // 操作数名称(变量)
    | L_PAREN expression R_PAREN;  // 括号表达式

// 字面量
literal:
    literalNil
	| integer
	| string
	| literalFloat
	| literalImag
	;
literalNil: NIL_LIT;
integer:
	DECIMAL_LIT
	| BINARY_LIT
	| OCTAL_LIT
	| HEX_LIT
	// | RUNE_LIT
	;
string: RAW_STRING_LIT | INTERPRETED_STRING_LIT | SINGER_QUOT_STRING_LIT;
literalFloat: FLOAT_LIT;
literalImag: IMAGINARY_LIT;
// 操作数名称
operandName: IDENTIFIER;

// 取字段
field:
	SafeIndex IDENTIFIER
	| DOT IDENTIFIER;
// 取下标
index: L_BRACKET expression R_BRACKET;
// 取切片
slice:
	L_BRACKET (
		lo=expression? COLON hi=expression?
		| lo=expression? COLON hi=expression COLON cap=expression
	) R_BRACKET;
// 取调用
arguments:
	L_PAREN (
		(expressionList) ELLIPSIS? COMMA?
	)? R_PAREN;
// 表达式列表
expressionList: expression (COMMA expression)*;

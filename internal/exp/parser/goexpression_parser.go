// Code generated from GoExpression.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // GoExpression
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GoExpression struct {
	*antlr.BaseParser
}

var GoExpressionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func goexpressionParserInit() {
	staticData := &GoExpressionParserStaticData
	staticData.LiteralNames = []string{
		"", "'break'", "'default'", "'func'", "'interface'", "'select'", "'case'",
		"'defer'", "'go'", "'map'", "'struct'", "'chan'", "'else'", "'goto'",
		"'package'", "'switch'", "'const'", "'fallthrough'", "'if'", "'range'",
		"'type'", "'continue'", "'for'", "'import'", "'return'", "'var'", "'nil'",
		"", "'('", "')'", "'{'", "'}'", "'['", "']'", "'='", "','", "';'", "':'",
		"'.'", "'++'", "'--'", "':='", "'...'", "'?'", "'?.'", "'||'", "'&&'",
		"'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'|'", "'/'", "'%'", "'<<'",
		"'>>'", "'&^'", "'~'", "'!'", "'+'", "'-'", "'^'", "'*'", "'&'", "'<-'",
	}
	staticData.SymbolicNames = []string{
		"", "BREAK", "DEFAULT", "FUNC", "INTERFACE", "SELECT", "CASE", "DEFER",
		"GO", "MAP", "STRUCT", "CHAN", "ELSE", "GOTO", "PACKAGE", "SWITCH",
		"CONST", "FALLTHROUGH", "IF", "RANGE", "TYPE", "CONTINUE", "FOR", "IMPORT",
		"RETURN", "VAR", "NIL_LIT", "IDENTIFIER", "L_PAREN", "R_PAREN", "L_CURLY",
		"R_CURLY", "L_BRACKET", "R_BRACKET", "ASSIGN", "COMMA", "SEMI", "COLON",
		"DOT", "PLUS_PLUS", "MINUS_MINUS", "DECLARE_ASSIGN", "ELLIPSIS", "Question",
		"SafeIndex", "LOGICAL_OR", "LOGICAL_AND", "EQUALS", "NOT_EQUALS", "LESS",
		"LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS", "OR", "DIV", "MOD",
		"LSHIFT", "RSHIFT", "BIT_CLEAR", "UNDERLYING", "EXCLAMATION", "PLUS",
		"MINUS", "CARET", "STAR", "AMPERSAND", "RECEIVE", "DECIMAL_LIT", "BINARY_LIT",
		"OCTAL_LIT", "HEX_LIT", "FLOAT_LIT", "DECIMAL_FLOAT_LIT", "HEX_FLOAT_LIT",
		"IMAGINARY_LIT", "BYTE_VALUE", "OCTAL_BYTE_VALUE", "HEX_BYTE_VALUE",
		"LITTLE_U_VALUE", "BIG_U_VALUE", "RAW_STRING_LIT", "INTERPRETED_STRING_LIT",
		"SINGER_QUOT_STRING_LIT", "WS", "COMMENT", "TERMINATOR", "LINE_COMMENT",
		"WS_NLSEMI", "COMMENT_NLSEMI", "LINE_COMMENT_NLSEMI", "EOS", "OTHER",
	}
	staticData.RuleNames = []string{
		"expression", "primaryExpr", "operand", "literal", "literalNil", "integer",
		"string", "literalFloat", "literalImag", "operandName", "field", "index",
		"slice", "arguments", "expressionList",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 91, 155, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 1, 0,
		1, 0, 1, 0, 3, 0, 35, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 5, 0, 58, 8, 0, 10, 0, 12, 0, 61, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 71, 8, 1, 5, 1, 73, 8, 1, 10, 1, 12, 1, 76,
		9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 84, 8, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 3, 3, 91, 8, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 109, 8,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 3, 12, 117, 8, 12, 1, 12,
		1, 12, 3, 12, 121, 8, 12, 1, 12, 3, 12, 124, 8, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 3, 12, 131, 8, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 3,
		13, 138, 8, 13, 1, 13, 3, 13, 141, 8, 13, 3, 13, 143, 8, 13, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 5, 14, 150, 8, 14, 10, 14, 12, 14, 153, 9, 14,
		1, 14, 0, 2, 0, 2, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 0, 6, 1, 0, 60, 66, 2, 0, 54, 58, 64, 65, 2, 0, 53, 53, 61, 63, 1,
		0, 47, 52, 1, 0, 67, 70, 1, 0, 80, 82, 165, 0, 34, 1, 0, 0, 0, 2, 62, 1,
		0, 0, 0, 4, 83, 1, 0, 0, 0, 6, 90, 1, 0, 0, 0, 8, 92, 1, 0, 0, 0, 10, 94,
		1, 0, 0, 0, 12, 96, 1, 0, 0, 0, 14, 98, 1, 0, 0, 0, 16, 100, 1, 0, 0, 0,
		18, 102, 1, 0, 0, 0, 20, 108, 1, 0, 0, 0, 22, 110, 1, 0, 0, 0, 24, 114,
		1, 0, 0, 0, 26, 134, 1, 0, 0, 0, 28, 146, 1, 0, 0, 0, 30, 31, 6, 0, -1,
		0, 31, 35, 3, 2, 1, 0, 32, 33, 7, 0, 0, 0, 33, 35, 3, 0, 0, 7, 34, 30,
		1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 35, 59, 1, 0, 0, 0, 36, 37, 10, 6, 0, 0,
		37, 38, 7, 1, 0, 0, 38, 58, 3, 0, 0, 7, 39, 40, 10, 5, 0, 0, 40, 41, 7,
		2, 0, 0, 41, 58, 3, 0, 0, 6, 42, 43, 10, 4, 0, 0, 43, 44, 7, 3, 0, 0, 44,
		58, 3, 0, 0, 5, 45, 46, 10, 3, 0, 0, 46, 47, 5, 46, 0, 0, 47, 58, 3, 0,
		0, 4, 48, 49, 10, 2, 0, 0, 49, 50, 5, 45, 0, 0, 50, 58, 3, 0, 0, 3, 51,
		52, 10, 1, 0, 0, 52, 53, 5, 43, 0, 0, 53, 54, 3, 0, 0, 0, 54, 55, 5, 37,
		0, 0, 55, 56, 3, 0, 0, 2, 56, 58, 1, 0, 0, 0, 57, 36, 1, 0, 0, 0, 57, 39,
		1, 0, 0, 0, 57, 42, 1, 0, 0, 0, 57, 45, 1, 0, 0, 0, 57, 48, 1, 0, 0, 0,
		57, 51, 1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1,
		0, 0, 0, 60, 1, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 62, 63, 6, 1, -1, 0, 63,
		64, 3, 4, 2, 0, 64, 74, 1, 0, 0, 0, 65, 70, 10, 1, 0, 0, 66, 71, 3, 20,
		10, 0, 67, 71, 3, 22, 11, 0, 68, 71, 3, 24, 12, 0, 69, 71, 3, 26, 13, 0,
		70, 66, 1, 0, 0, 0, 70, 67, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 69, 1,
		0, 0, 0, 71, 73, 1, 0, 0, 0, 72, 65, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74,
		72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 3, 1, 0, 0, 0, 76, 74, 1, 0, 0,
		0, 77, 84, 3, 6, 3, 0, 78, 84, 3, 18, 9, 0, 79, 80, 5, 28, 0, 0, 80, 81,
		3, 0, 0, 0, 81, 82, 5, 29, 0, 0, 82, 84, 1, 0, 0, 0, 83, 77, 1, 0, 0, 0,
		83, 78, 1, 0, 0, 0, 83, 79, 1, 0, 0, 0, 84, 5, 1, 0, 0, 0, 85, 91, 3, 8,
		4, 0, 86, 91, 3, 10, 5, 0, 87, 91, 3, 12, 6, 0, 88, 91, 3, 14, 7, 0, 89,
		91, 3, 16, 8, 0, 90, 85, 1, 0, 0, 0, 90, 86, 1, 0, 0, 0, 90, 87, 1, 0,
		0, 0, 90, 88, 1, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 7, 1, 0, 0, 0, 92, 93,
		5, 26, 0, 0, 93, 9, 1, 0, 0, 0, 94, 95, 7, 4, 0, 0, 95, 11, 1, 0, 0, 0,
		96, 97, 7, 5, 0, 0, 97, 13, 1, 0, 0, 0, 98, 99, 5, 71, 0, 0, 99, 15, 1,
		0, 0, 0, 100, 101, 5, 74, 0, 0, 101, 17, 1, 0, 0, 0, 102, 103, 5, 27, 0,
		0, 103, 19, 1, 0, 0, 0, 104, 105, 5, 44, 0, 0, 105, 109, 5, 27, 0, 0, 106,
		107, 5, 38, 0, 0, 107, 109, 5, 27, 0, 0, 108, 104, 1, 0, 0, 0, 108, 106,
		1, 0, 0, 0, 109, 21, 1, 0, 0, 0, 110, 111, 5, 32, 0, 0, 111, 112, 3, 0,
		0, 0, 112, 113, 5, 33, 0, 0, 113, 23, 1, 0, 0, 0, 114, 130, 5, 32, 0, 0,
		115, 117, 3, 0, 0, 0, 116, 115, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117,
		118, 1, 0, 0, 0, 118, 120, 5, 37, 0, 0, 119, 121, 3, 0, 0, 0, 120, 119,
		1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 131, 1, 0, 0, 0, 122, 124, 3, 0,
		0, 0, 123, 122, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0,
		125, 126, 5, 37, 0, 0, 126, 127, 3, 0, 0, 0, 127, 128, 5, 37, 0, 0, 128,
		129, 3, 0, 0, 0, 129, 131, 1, 0, 0, 0, 130, 116, 1, 0, 0, 0, 130, 123,
		1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 133, 5, 33, 0, 0, 133, 25, 1, 0,
		0, 0, 134, 142, 5, 28, 0, 0, 135, 137, 3, 28, 14, 0, 136, 138, 5, 42, 0,
		0, 137, 136, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 140, 1, 0, 0, 0, 139,
		141, 5, 35, 0, 0, 140, 139, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 143,
		1, 0, 0, 0, 142, 135, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 144, 1, 0,
		0, 0, 144, 145, 5, 29, 0, 0, 145, 27, 1, 0, 0, 0, 146, 151, 3, 0, 0, 0,
		147, 148, 5, 35, 0, 0, 148, 150, 3, 0, 0, 0, 149, 147, 1, 0, 0, 0, 150,
		153, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 29, 1,
		0, 0, 0, 153, 151, 1, 0, 0, 0, 16, 34, 57, 59, 70, 74, 83, 90, 108, 116,
		120, 123, 130, 137, 140, 142, 151,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GoExpressionInit initializes any static state used to implement GoExpression. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGoExpression(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoExpressionInit() {
	staticData := &GoExpressionParserStaticData
	staticData.once.Do(goexpressionParserInit)
}

// NewGoExpression produces a new parser instance for the optional input antlr.TokenStream.
func NewGoExpression(input antlr.TokenStream) *GoExpression {
	GoExpressionInit()
	this := new(GoExpression)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GoExpressionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "GoExpression.g4"

	return this
}

// GoExpression tokens.
const (
	GoExpressionEOF                    = antlr.TokenEOF
	GoExpressionBREAK                  = 1
	GoExpressionDEFAULT                = 2
	GoExpressionFUNC                   = 3
	GoExpressionINTERFACE              = 4
	GoExpressionSELECT                 = 5
	GoExpressionCASE                   = 6
	GoExpressionDEFER                  = 7
	GoExpressionGO                     = 8
	GoExpressionMAP                    = 9
	GoExpressionSTRUCT                 = 10
	GoExpressionCHAN                   = 11
	GoExpressionELSE                   = 12
	GoExpressionGOTO                   = 13
	GoExpressionPACKAGE                = 14
	GoExpressionSWITCH                 = 15
	GoExpressionCONST                  = 16
	GoExpressionFALLTHROUGH            = 17
	GoExpressionIF                     = 18
	GoExpressionRANGE                  = 19
	GoExpressionTYPE                   = 20
	GoExpressionCONTINUE               = 21
	GoExpressionFOR                    = 22
	GoExpressionIMPORT                 = 23
	GoExpressionRETURN                 = 24
	GoExpressionVAR                    = 25
	GoExpressionNIL_LIT                = 26
	GoExpressionIDENTIFIER             = 27
	GoExpressionL_PAREN                = 28
	GoExpressionR_PAREN                = 29
	GoExpressionL_CURLY                = 30
	GoExpressionR_CURLY                = 31
	GoExpressionL_BRACKET              = 32
	GoExpressionR_BRACKET              = 33
	GoExpressionASSIGN                 = 34
	GoExpressionCOMMA                  = 35
	GoExpressionSEMI                   = 36
	GoExpressionCOLON                  = 37
	GoExpressionDOT                    = 38
	GoExpressionPLUS_PLUS              = 39
	GoExpressionMINUS_MINUS            = 40
	GoExpressionDECLARE_ASSIGN         = 41
	GoExpressionELLIPSIS               = 42
	GoExpressionQuestion               = 43
	GoExpressionSafeIndex              = 44
	GoExpressionLOGICAL_OR             = 45
	GoExpressionLOGICAL_AND            = 46
	GoExpressionEQUALS                 = 47
	GoExpressionNOT_EQUALS             = 48
	GoExpressionLESS                   = 49
	GoExpressionLESS_OR_EQUALS         = 50
	GoExpressionGREATER                = 51
	GoExpressionGREATER_OR_EQUALS      = 52
	GoExpressionOR                     = 53
	GoExpressionDIV                    = 54
	GoExpressionMOD                    = 55
	GoExpressionLSHIFT                 = 56
	GoExpressionRSHIFT                 = 57
	GoExpressionBIT_CLEAR              = 58
	GoExpressionUNDERLYING             = 59
	GoExpressionEXCLAMATION            = 60
	GoExpressionPLUS                   = 61
	GoExpressionMINUS                  = 62
	GoExpressionCARET                  = 63
	GoExpressionSTAR                   = 64
	GoExpressionAMPERSAND              = 65
	GoExpressionRECEIVE                = 66
	GoExpressionDECIMAL_LIT            = 67
	GoExpressionBINARY_LIT             = 68
	GoExpressionOCTAL_LIT              = 69
	GoExpressionHEX_LIT                = 70
	GoExpressionFLOAT_LIT              = 71
	GoExpressionDECIMAL_FLOAT_LIT      = 72
	GoExpressionHEX_FLOAT_LIT          = 73
	GoExpressionIMAGINARY_LIT          = 74
	GoExpressionBYTE_VALUE             = 75
	GoExpressionOCTAL_BYTE_VALUE       = 76
	GoExpressionHEX_BYTE_VALUE         = 77
	GoExpressionLITTLE_U_VALUE         = 78
	GoExpressionBIG_U_VALUE            = 79
	GoExpressionRAW_STRING_LIT         = 80
	GoExpressionINTERPRETED_STRING_LIT = 81
	GoExpressionSINGER_QUOT_STRING_LIT = 82
	GoExpressionWS                     = 83
	GoExpressionCOMMENT                = 84
	GoExpressionTERMINATOR             = 85
	GoExpressionLINE_COMMENT           = 86
	GoExpressionWS_NLSEMI              = 87
	GoExpressionCOMMENT_NLSEMI         = 88
	GoExpressionLINE_COMMENT_NLSEMI    = 89
	GoExpressionEOS                    = 90
	GoExpressionOTHER                  = 91
)

// GoExpression rules.
const (
	GoExpressionRULE_expression     = 0
	GoExpressionRULE_primaryExpr    = 1
	GoExpressionRULE_operand        = 2
	GoExpressionRULE_literal        = 3
	GoExpressionRULE_literalNil     = 4
	GoExpressionRULE_integer        = 5
	GoExpressionRULE_string         = 6
	GoExpressionRULE_literalFloat   = 7
	GoExpressionRULE_literalImag    = 8
	GoExpressionRULE_operandName    = 9
	GoExpressionRULE_field          = 10
	GoExpressionRULE_index          = 11
	GoExpressionRULE_slice          = 12
	GoExpressionRULE_arguments      = 13
	GoExpressionRULE_expressionList = 14
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetUnary_op returns the unary_op token.
	GetUnary_op() antlr.Token

	// GetMul_op returns the mul_op token.
	GetMul_op() antlr.Token

	// GetAdd_op returns the add_op token.
	GetAdd_op() antlr.Token

	// GetRel_op returns the rel_op token.
	GetRel_op() antlr.Token

	// SetUnary_op sets the unary_op token.
	SetUnary_op(antlr.Token)

	// SetMul_op sets the mul_op token.
	SetMul_op(antlr.Token)

	// SetAdd_op sets the add_op token.
	SetAdd_op(antlr.Token)

	// SetRel_op sets the rel_op token.
	SetRel_op(antlr.Token)

	// Getter signatures
	PrimaryExpr() IPrimaryExprContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	EXCLAMATION() antlr.TerminalNode
	CARET() antlr.TerminalNode
	STAR() antlr.TerminalNode
	AMPERSAND() antlr.TerminalNode
	RECEIVE() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode
	LSHIFT() antlr.TerminalNode
	RSHIFT() antlr.TerminalNode
	BIT_CLEAR() antlr.TerminalNode
	OR() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	NOT_EQUALS() antlr.TerminalNode
	LESS() antlr.TerminalNode
	LESS_OR_EQUALS() antlr.TerminalNode
	GREATER() antlr.TerminalNode
	GREATER_OR_EQUALS() antlr.TerminalNode
	LOGICAL_AND() antlr.TerminalNode
	LOGICAL_OR() antlr.TerminalNode
	Question() antlr.TerminalNode
	COLON() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	unary_op antlr.Token
	mul_op   antlr.Token
	add_op   antlr.Token
	rel_op   antlr.Token
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoExpressionRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetUnary_op() antlr.Token { return s.unary_op }

func (s *ExpressionContext) GetMul_op() antlr.Token { return s.mul_op }

func (s *ExpressionContext) GetAdd_op() antlr.Token { return s.add_op }

func (s *ExpressionContext) GetRel_op() antlr.Token { return s.rel_op }

func (s *ExpressionContext) SetUnary_op(v antlr.Token) { s.unary_op = v }

func (s *ExpressionContext) SetMul_op(v antlr.Token) { s.mul_op = v }

func (s *ExpressionContext) SetAdd_op(v antlr.Token) { s.add_op = v }

func (s *ExpressionContext) SetRel_op(v antlr.Token) { s.rel_op = v }

func (s *ExpressionContext) PrimaryExpr() IPrimaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(GoExpressionPLUS, 0)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoExpressionMINUS, 0)
}

func (s *ExpressionContext) EXCLAMATION() antlr.TerminalNode {
	return s.GetToken(GoExpressionEXCLAMATION, 0)
}

func (s *ExpressionContext) CARET() antlr.TerminalNode {
	return s.GetToken(GoExpressionCARET, 0)
}

func (s *ExpressionContext) STAR() antlr.TerminalNode {
	return s.GetToken(GoExpressionSTAR, 0)
}

func (s *ExpressionContext) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(GoExpressionAMPERSAND, 0)
}

func (s *ExpressionContext) RECEIVE() antlr.TerminalNode {
	return s.GetToken(GoExpressionRECEIVE, 0)
}

func (s *ExpressionContext) DIV() antlr.TerminalNode {
	return s.GetToken(GoExpressionDIV, 0)
}

func (s *ExpressionContext) MOD() antlr.TerminalNode {
	return s.GetToken(GoExpressionMOD, 0)
}

func (s *ExpressionContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(GoExpressionLSHIFT, 0)
}

func (s *ExpressionContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(GoExpressionRSHIFT, 0)
}

func (s *ExpressionContext) BIT_CLEAR() antlr.TerminalNode {
	return s.GetToken(GoExpressionBIT_CLEAR, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(GoExpressionOR, 0)
}

func (s *ExpressionContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(GoExpressionEQUALS, 0)
}

func (s *ExpressionContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(GoExpressionNOT_EQUALS, 0)
}

func (s *ExpressionContext) LESS() antlr.TerminalNode {
	return s.GetToken(GoExpressionLESS, 0)
}

func (s *ExpressionContext) LESS_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(GoExpressionLESS_OR_EQUALS, 0)
}

func (s *ExpressionContext) GREATER() antlr.TerminalNode {
	return s.GetToken(GoExpressionGREATER, 0)
}

func (s *ExpressionContext) GREATER_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(GoExpressionGREATER_OR_EQUALS, 0)
}

func (s *ExpressionContext) LOGICAL_AND() antlr.TerminalNode {
	return s.GetToken(GoExpressionLOGICAL_AND, 0)
}

func (s *ExpressionContext) LOGICAL_OR() antlr.TerminalNode {
	return s.GetToken(GoExpressionLOGICAL_OR, 0)
}

func (s *ExpressionContext) Question() antlr.TerminalNode {
	return s.GetToken(GoExpressionQuestion, 0)
}

func (s *ExpressionContext) COLON() antlr.TerminalNode {
	return s.GetToken(GoExpressionCOLON, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoExpressionVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoExpression) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *GoExpression) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, GoExpressionRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoExpressionNIL_LIT, GoExpressionIDENTIFIER, GoExpressionL_PAREN, GoExpressionDECIMAL_LIT, GoExpressionBINARY_LIT, GoExpressionOCTAL_LIT, GoExpressionHEX_LIT, GoExpressionFLOAT_LIT, GoExpressionIMAGINARY_LIT, GoExpressionRAW_STRING_LIT, GoExpressionINTERPRETED_STRING_LIT, GoExpressionSINGER_QUOT_STRING_LIT:
		{
			p.SetState(31)
			p.primaryExpr(0)
		}

	case GoExpressionEXCLAMATION, GoExpressionPLUS, GoExpressionMINUS, GoExpressionCARET, GoExpressionSTAR, GoExpressionAMPERSAND, GoExpressionRECEIVE:
		{
			p.SetState(32)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).unary_op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-60)) & ^0x3f) == 0 && ((int64(1)<<(_la-60))&127) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).unary_op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(33)
			p.expression(7)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GoExpressionRULE_expression)
				p.SetState(36)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(37)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).mul_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-54)) & ^0x3f) == 0 && ((int64(1)<<(_la-54))&3103) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).mul_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(38)
					p.expression(7)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GoExpressionRULE_expression)
				p.SetState(39)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(40)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).add_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2296835809958952960) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).add_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(41)
					p.expression(6)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GoExpressionRULE_expression)
				p.SetState(42)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(43)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).rel_op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8866461766385664) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).rel_op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(44)
					p.expression(5)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GoExpressionRULE_expression)
				p.SetState(45)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(46)
					p.Match(GoExpressionLOGICAL_AND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(47)
					p.expression(4)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GoExpressionRULE_expression)
				p.SetState(48)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(49)
					p.Match(GoExpressionLOGICAL_OR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(50)
					p.expression(3)
				}

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GoExpressionRULE_expression)
				p.SetState(51)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(52)
					p.Match(GoExpressionQuestion)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(53)
					p.expression(0)
				}
				{
					p.SetState(54)
					p.Match(GoExpressionCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(55)
					p.expression(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Operand() IOperandContext
	PrimaryExpr() IPrimaryExprContext
	Field() IFieldContext
	Index() IIndexContext
	Slice() ISliceContext
	Arguments() IArgumentsContext

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_primaryExpr
	return p
}

func InitEmptyPrimaryExprContext(p *PrimaryExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_primaryExpr
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoExpressionRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *PrimaryExprContext) PrimaryExpr() IPrimaryExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *PrimaryExprContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *PrimaryExprContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *PrimaryExprContext) Slice() ISliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceContext)
}

func (s *PrimaryExprContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoExpressionVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoExpression) PrimaryExpr() (localctx IPrimaryExprContext) {
	return p.primaryExpr(0)
}

func (p *GoExpression) primaryExpr(_p int) (localctx IPrimaryExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, GoExpressionRULE_primaryExpr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Operand()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPrimaryExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, GoExpressionRULE_primaryExpr)
			p.SetState(65)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			p.SetState(70)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(66)
					p.Field()
				}

			case 2:
				{
					p.SetState(67)
					p.Index()
				}

			case 3:
				{
					p.SetState(68)
					p.Slice()
				}

			case 4:
				{
					p.SetState(69)
					p.Arguments()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	OperandName() IOperandNameContext
	L_PAREN() antlr.TerminalNode
	Expression() IExpressionContext
	R_PAREN() antlr.TerminalNode

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_operand
	return p
}

func InitEmptyOperandContext(p *OperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_operand
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoExpressionRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *OperandContext) OperandName() IOperandNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandNameContext)
}

func (s *OperandContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GoExpressionL_PAREN, 0)
}

func (s *OperandContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OperandContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GoExpressionR_PAREN, 0)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (s *OperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoExpressionVisitor:
		return t.VisitOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoExpression) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GoExpressionRULE_operand)
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoExpressionNIL_LIT, GoExpressionDECIMAL_LIT, GoExpressionBINARY_LIT, GoExpressionOCTAL_LIT, GoExpressionHEX_LIT, GoExpressionFLOAT_LIT, GoExpressionIMAGINARY_LIT, GoExpressionRAW_STRING_LIT, GoExpressionINTERPRETED_STRING_LIT, GoExpressionSINGER_QUOT_STRING_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.Literal()
		}

	case GoExpressionIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(78)
			p.OperandName()
		}

	case GoExpressionL_PAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(79)
			p.Match(GoExpressionL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.expression(0)
		}
		{
			p.SetState(81)
			p.Match(GoExpressionR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LiteralNil() ILiteralNilContext
	Integer() IIntegerContext
	String_() IStringContext
	LiteralFloat() ILiteralFloatContext
	LiteralImag() ILiteralImagContext

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoExpressionRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) LiteralNil() ILiteralNilContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralNilContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralNilContext)
}

func (s *LiteralContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *LiteralContext) String_() IStringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringContext)
}

func (s *LiteralContext) LiteralFloat() ILiteralFloatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralFloatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralFloatContext)
}

func (s *LiteralContext) LiteralImag() ILiteralImagContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralImagContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralImagContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoExpressionVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoExpression) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GoExpressionRULE_literal)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoExpressionNIL_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)
			p.LiteralNil()
		}

	case GoExpressionDECIMAL_LIT, GoExpressionBINARY_LIT, GoExpressionOCTAL_LIT, GoExpressionHEX_LIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)
			p.Integer()
		}

	case GoExpressionRAW_STRING_LIT, GoExpressionINTERPRETED_STRING_LIT, GoExpressionSINGER_QUOT_STRING_LIT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(87)
			p.String_()
		}

	case GoExpressionFLOAT_LIT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(88)
			p.LiteralFloat()
		}

	case GoExpressionIMAGINARY_LIT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(89)
			p.LiteralImag()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralNilContext is an interface to support dynamic dispatch.
type ILiteralNilContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NIL_LIT() antlr.TerminalNode

	// IsLiteralNilContext differentiates from other interfaces.
	IsLiteralNilContext()
}

type LiteralNilContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralNilContext() *LiteralNilContext {
	var p = new(LiteralNilContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_literalNil
	return p
}

func InitEmptyLiteralNilContext(p *LiteralNilContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_literalNil
}

func (*LiteralNilContext) IsLiteralNilContext() {}

func NewLiteralNilContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralNilContext {
	var p = new(LiteralNilContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoExpressionRULE_literalNil

	return p
}

func (s *LiteralNilContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralNilContext) NIL_LIT() antlr.TerminalNode {
	return s.GetToken(GoExpressionNIL_LIT, 0)
}

func (s *LiteralNilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralNilContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralNilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.EnterLiteralNil(s)
	}
}

func (s *LiteralNilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.ExitLiteralNil(s)
	}
}

func (s *LiteralNilContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoExpressionVisitor:
		return t.VisitLiteralNil(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoExpression) LiteralNil() (localctx ILiteralNilContext) {
	localctx = NewLiteralNilContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GoExpressionRULE_literalNil)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(GoExpressionNIL_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DECIMAL_LIT() antlr.TerminalNode
	BINARY_LIT() antlr.TerminalNode
	OCTAL_LIT() antlr.TerminalNode
	HEX_LIT() antlr.TerminalNode

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_integer
	return p
}

func InitEmptyIntegerContext(p *IntegerContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_integer
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoExpressionRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) DECIMAL_LIT() antlr.TerminalNode {
	return s.GetToken(GoExpressionDECIMAL_LIT, 0)
}

func (s *IntegerContext) BINARY_LIT() antlr.TerminalNode {
	return s.GetToken(GoExpressionBINARY_LIT, 0)
}

func (s *IntegerContext) OCTAL_LIT() antlr.TerminalNode {
	return s.GetToken(GoExpressionOCTAL_LIT, 0)
}

func (s *IntegerContext) HEX_LIT() antlr.TerminalNode {
	return s.GetToken(GoExpressionHEX_LIT, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoExpressionVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoExpression) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GoExpressionRULE_integer)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-67)) & ^0x3f) == 0 && ((int64(1)<<(_la-67))&15) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStringContext is an interface to support dynamic dispatch.
type IStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RAW_STRING_LIT() antlr.TerminalNode
	INTERPRETED_STRING_LIT() antlr.TerminalNode
	SINGER_QUOT_STRING_LIT() antlr.TerminalNode

	// IsStringContext differentiates from other interfaces.
	IsStringContext()
}

type StringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringContext() *StringContext {
	var p = new(StringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_string
	return p
}

func InitEmptyStringContext(p *StringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_string
}

func (*StringContext) IsStringContext() {}

func NewStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringContext {
	var p = new(StringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoExpressionRULE_string

	return p
}

func (s *StringContext) GetParser() antlr.Parser { return s.parser }

func (s *StringContext) RAW_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(GoExpressionRAW_STRING_LIT, 0)
}

func (s *StringContext) INTERPRETED_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(GoExpressionINTERPRETED_STRING_LIT, 0)
}

func (s *StringContext) SINGER_QUOT_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(GoExpressionSINGER_QUOT_STRING_LIT, 0)
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.ExitString(s)
	}
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoExpressionVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoExpression) String_() (localctx IStringContext) {
	localctx = NewStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GoExpressionRULE_string)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-80)) & ^0x3f) == 0 && ((int64(1)<<(_la-80))&7) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralFloatContext is an interface to support dynamic dispatch.
type ILiteralFloatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOAT_LIT() antlr.TerminalNode

	// IsLiteralFloatContext differentiates from other interfaces.
	IsLiteralFloatContext()
}

type LiteralFloatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralFloatContext() *LiteralFloatContext {
	var p = new(LiteralFloatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_literalFloat
	return p
}

func InitEmptyLiteralFloatContext(p *LiteralFloatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_literalFloat
}

func (*LiteralFloatContext) IsLiteralFloatContext() {}

func NewLiteralFloatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralFloatContext {
	var p = new(LiteralFloatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoExpressionRULE_literalFloat

	return p
}

func (s *LiteralFloatContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralFloatContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(GoExpressionFLOAT_LIT, 0)
}

func (s *LiteralFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralFloatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.EnterLiteralFloat(s)
	}
}

func (s *LiteralFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.ExitLiteralFloat(s)
	}
}

func (s *LiteralFloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoExpressionVisitor:
		return t.VisitLiteralFloat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoExpression) LiteralFloat() (localctx ILiteralFloatContext) {
	localctx = NewLiteralFloatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GoExpressionRULE_literalFloat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(GoExpressionFLOAT_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralImagContext is an interface to support dynamic dispatch.
type ILiteralImagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMAGINARY_LIT() antlr.TerminalNode

	// IsLiteralImagContext differentiates from other interfaces.
	IsLiteralImagContext()
}

type LiteralImagContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralImagContext() *LiteralImagContext {
	var p = new(LiteralImagContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_literalImag
	return p
}

func InitEmptyLiteralImagContext(p *LiteralImagContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_literalImag
}

func (*LiteralImagContext) IsLiteralImagContext() {}

func NewLiteralImagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralImagContext {
	var p = new(LiteralImagContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoExpressionRULE_literalImag

	return p
}

func (s *LiteralImagContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralImagContext) IMAGINARY_LIT() antlr.TerminalNode {
	return s.GetToken(GoExpressionIMAGINARY_LIT, 0)
}

func (s *LiteralImagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralImagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralImagContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.EnterLiteralImag(s)
	}
}

func (s *LiteralImagContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.ExitLiteralImag(s)
	}
}

func (s *LiteralImagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoExpressionVisitor:
		return t.VisitLiteralImag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoExpression) LiteralImag() (localctx ILiteralImagContext) {
	localctx = NewLiteralImagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GoExpressionRULE_literalImag)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(GoExpressionIMAGINARY_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperandNameContext is an interface to support dynamic dispatch.
type IOperandNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsOperandNameContext differentiates from other interfaces.
	IsOperandNameContext()
}

type OperandNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandNameContext() *OperandNameContext {
	var p = new(OperandNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_operandName
	return p
}

func InitEmptyOperandNameContext(p *OperandNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_operandName
}

func (*OperandNameContext) IsOperandNameContext() {}

func NewOperandNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandNameContext {
	var p = new(OperandNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoExpressionRULE_operandName

	return p
}

func (s *OperandNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GoExpressionIDENTIFIER, 0)
}

func (s *OperandNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.EnterOperandName(s)
	}
}

func (s *OperandNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.ExitOperandName(s)
	}
}

func (s *OperandNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoExpressionVisitor:
		return t.VisitOperandName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoExpression) OperandName() (localctx IOperandNameContext) {
	localctx = NewOperandNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GoExpressionRULE_operandName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Match(GoExpressionIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SafeIndex() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	DOT() antlr.TerminalNode

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoExpressionRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) SafeIndex() antlr.TerminalNode {
	return s.GetToken(GoExpressionSafeIndex, 0)
}

func (s *FieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GoExpressionIDENTIFIER, 0)
}

func (s *FieldContext) DOT() antlr.TerminalNode {
	return s.GetToken(GoExpressionDOT, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoExpressionVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoExpression) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GoExpressionRULE_field)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoExpressionSafeIndex:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.Match(GoExpressionSafeIndex)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Match(GoExpressionIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoExpressionDOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.Match(GoExpressionDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Match(GoExpressionIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_BRACKET() antlr.TerminalNode
	Expression() IExpressionContext
	R_BRACKET() antlr.TerminalNode

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_index
	return p
}

func InitEmptyIndexContext(p *IndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_index
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoExpressionRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(GoExpressionL_BRACKET, 0)
}

func (s *IndexContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(GoExpressionR_BRACKET, 0)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.EnterIndex(s)
	}
}

func (s *IndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.ExitIndex(s)
	}
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoExpressionVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoExpression) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GoExpressionRULE_index)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(GoExpressionL_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.expression(0)
	}
	{
		p.SetState(112)
		p.Match(GoExpressionR_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISliceContext is an interface to support dynamic dispatch.
type ISliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLo returns the lo rule contexts.
	GetLo() IExpressionContext

	// GetHi returns the hi rule contexts.
	GetHi() IExpressionContext

	// GetCap_ returns the cap_ rule contexts.
	GetCap_() IExpressionContext

	// SetLo sets the lo rule contexts.
	SetLo(IExpressionContext)

	// SetHi sets the hi rule contexts.
	SetHi(IExpressionContext)

	// SetCap_ sets the cap_ rule contexts.
	SetCap_(IExpressionContext)

	// Getter signatures
	L_BRACKET() antlr.TerminalNode
	R_BRACKET() antlr.TerminalNode
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsSliceContext differentiates from other interfaces.
	IsSliceContext()
}

type SliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	lo     IExpressionContext
	hi     IExpressionContext
	cap_   IExpressionContext
}

func NewEmptySliceContext() *SliceContext {
	var p = new(SliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_slice
	return p
}

func InitEmptySliceContext(p *SliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_slice
}

func (*SliceContext) IsSliceContext() {}

func NewSliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceContext {
	var p = new(SliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoExpressionRULE_slice

	return p
}

func (s *SliceContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceContext) GetLo() IExpressionContext { return s.lo }

func (s *SliceContext) GetHi() IExpressionContext { return s.hi }

func (s *SliceContext) GetCap_() IExpressionContext { return s.cap_ }

func (s *SliceContext) SetLo(v IExpressionContext) { s.lo = v }

func (s *SliceContext) SetHi(v IExpressionContext) { s.hi = v }

func (s *SliceContext) SetCap_(v IExpressionContext) { s.cap_ = v }

func (s *SliceContext) L_BRACKET() antlr.TerminalNode {
	return s.GetToken(GoExpressionL_BRACKET, 0)
}

func (s *SliceContext) R_BRACKET() antlr.TerminalNode {
	return s.GetToken(GoExpressionR_BRACKET, 0)
}

func (s *SliceContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(GoExpressionCOLON)
}

func (s *SliceContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(GoExpressionCOLON, i)
}

func (s *SliceContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *SliceContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.EnterSlice(s)
	}
}

func (s *SliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.ExitSlice(s)
	}
}

func (s *SliceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoExpressionVisitor:
		return t.VisitSlice(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoExpression) Slice() (localctx ISliceContext) {
	localctx = NewSliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GoExpressionRULE_slice)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(GoExpressionL_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-26)) & ^0x3f) == 0 && ((int64(1)<<(_la-26))&126452616107393031) != 0 {
			{
				p.SetState(115)

				var _x = p.expression(0)

				localctx.(*SliceContext).lo = _x
			}

		}
		{
			p.SetState(118)
			p.Match(GoExpressionCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-26)) & ^0x3f) == 0 && ((int64(1)<<(_la-26))&126452616107393031) != 0 {
			{
				p.SetState(119)

				var _x = p.expression(0)

				localctx.(*SliceContext).hi = _x
			}

		}

	case 2:
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-26)) & ^0x3f) == 0 && ((int64(1)<<(_la-26))&126452616107393031) != 0 {
			{
				p.SetState(122)

				var _x = p.expression(0)

				localctx.(*SliceContext).lo = _x
			}

		}
		{
			p.SetState(125)
			p.Match(GoExpressionCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)

			var _x = p.expression(0)

			localctx.(*SliceContext).hi = _x
		}
		{
			p.SetState(127)
			p.Match(GoExpressionCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)

			var _x = p.expression(0)

			localctx.(*SliceContext).cap_ = _x
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(132)
		p.Match(GoExpressionR_BRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	ExpressionList() IExpressionListContext
	ELLIPSIS() antlr.TerminalNode
	COMMA() antlr.TerminalNode

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoExpressionRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GoExpressionL_PAREN, 0)
}

func (s *ArgumentsContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GoExpressionR_PAREN, 0)
}

func (s *ArgumentsContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArgumentsContext) ELLIPSIS() antlr.TerminalNode {
	return s.GetToken(GoExpressionELLIPSIS, 0)
}

func (s *ArgumentsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GoExpressionCOMMA, 0)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoExpressionVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoExpression) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GoExpressionRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(GoExpressionL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-26)) & ^0x3f) == 0 && ((int64(1)<<(_la-26))&126452616107393031) != 0 {
		{
			p.SetState(135)
			p.ExpressionList()
		}

		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GoExpressionELLIPSIS {
			{
				p.SetState(136)
				p.Match(GoExpressionELLIPSIS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GoExpressionCOMMA {
			{
				p.SetState(139)
				p.Match(GoExpressionCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(144)
		p.Match(GoExpressionR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_expressionList
	return p
}

func InitEmptyExpressionListContext(p *ExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoExpressionRULE_expressionList
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoExpressionRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GoExpressionCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GoExpressionCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoExpressionListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoExpressionVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoExpression) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GoExpressionRULE_expressionList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.expression(0)
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(147)
				p.Match(GoExpressionCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(148)
				p.expression(0)
			}

		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *GoExpression) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 1:
		var t *PrimaryExprContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExprContext)
		}
		return p.PrimaryExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GoExpression) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GoExpression) PrimaryExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

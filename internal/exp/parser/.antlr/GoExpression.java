// Generated from /data00/home/youthlin.chen/go/src/code.gopub.tech/tpl/internal/exp/parser/GoExpression.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class GoExpression extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		BREAK=1, DEFAULT=2, FUNC=3, INTERFACE=4, SELECT=5, CASE=6, DEFER=7, GO=8, 
		MAP=9, STRUCT=10, CHAN=11, ELSE=12, GOTO=13, PACKAGE=14, SWITCH=15, CONST=16, 
		FALLTHROUGH=17, IF=18, RANGE=19, TYPE=20, CONTINUE=21, FOR=22, IMPORT=23, 
		RETURN=24, VAR=25, NIL_LIT=26, IDENTIFIER=27, L_PAREN=28, R_PAREN=29, 
		L_CURLY=30, R_CURLY=31, L_BRACKET=32, R_BRACKET=33, ASSIGN=34, COMMA=35, 
		SEMI=36, COLON=37, DOT=38, PLUS_PLUS=39, MINUS_MINUS=40, DECLARE_ASSIGN=41, 
		ELLIPSIS=42, Question=43, SafeIndex=44, LOGICAL_OR=45, LOGICAL_AND=46, 
		EQUALS=47, NOT_EQUALS=48, LESS=49, LESS_OR_EQUALS=50, GREATER=51, GREATER_OR_EQUALS=52, 
		OR=53, DIV=54, MOD=55, LSHIFT=56, RSHIFT=57, BIT_CLEAR=58, UNDERLYING=59, 
		EXCLAMATION=60, PLUS=61, MINUS=62, CARET=63, STAR=64, AMPERSAND=65, RECEIVE=66, 
		DECIMAL_LIT=67, BINARY_LIT=68, OCTAL_LIT=69, HEX_LIT=70, FLOAT_LIT=71, 
		DECIMAL_FLOAT_LIT=72, HEX_FLOAT_LIT=73, IMAGINARY_LIT=74, BYTE_VALUE=75, 
		OCTAL_BYTE_VALUE=76, HEX_BYTE_VALUE=77, LITTLE_U_VALUE=78, BIG_U_VALUE=79, 
		RAW_STRING_LIT=80, INTERPRETED_STRING_LIT=81, SINGER_QUOT_STRING_LIT=82, 
		WS=83, COMMENT=84, TERMINATOR=85, LINE_COMMENT=86, WS_NLSEMI=87, COMMENT_NLSEMI=88, 
		LINE_COMMENT_NLSEMI=89, EOS=90, OTHER=91;
	public static final int
		RULE_expression = 0, RULE_primaryExpr = 1, RULE_operand = 2, RULE_literal = 3, 
		RULE_literalNil = 4, RULE_integer = 5, RULE_string = 6, RULE_literalFloat = 7, 
		RULE_literalImag = 8, RULE_operandName = 9, RULE_field = 10, RULE_index = 11, 
		RULE_slice = 12, RULE_arguments = 13, RULE_expressionList = 14;
	private static String[] makeRuleNames() {
		return new String[] {
			"expression", "primaryExpr", "operand", "literal", "literalNil", "integer", 
			"string", "literalFloat", "literalImag", "operandName", "field", "index", 
			"slice", "arguments", "expressionList"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'break'", "'default'", "'func'", "'interface'", "'select'", "'case'", 
			"'defer'", "'go'", "'map'", "'struct'", "'chan'", "'else'", "'goto'", 
			"'package'", "'switch'", "'const'", "'fallthrough'", "'if'", "'range'", 
			"'type'", "'continue'", "'for'", "'import'", "'return'", "'var'", "'nil'", 
			null, "'('", "')'", "'{'", "'}'", "'['", "']'", "'='", "','", "';'", 
			"':'", "'.'", "'++'", "'--'", "':='", "'...'", "'?'", "'?.'", "'||'", 
			"'&&'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'|'", "'/'", "'%'", 
			"'<<'", "'>>'", "'&^'", "'~'", "'!'", "'+'", "'-'", "'^'", "'*'", "'&'", 
			"'<-'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "BREAK", "DEFAULT", "FUNC", "INTERFACE", "SELECT", "CASE", "DEFER", 
			"GO", "MAP", "STRUCT", "CHAN", "ELSE", "GOTO", "PACKAGE", "SWITCH", "CONST", 
			"FALLTHROUGH", "IF", "RANGE", "TYPE", "CONTINUE", "FOR", "IMPORT", "RETURN", 
			"VAR", "NIL_LIT", "IDENTIFIER", "L_PAREN", "R_PAREN", "L_CURLY", "R_CURLY", 
			"L_BRACKET", "R_BRACKET", "ASSIGN", "COMMA", "SEMI", "COLON", "DOT", 
			"PLUS_PLUS", "MINUS_MINUS", "DECLARE_ASSIGN", "ELLIPSIS", "Question", 
			"SafeIndex", "LOGICAL_OR", "LOGICAL_AND", "EQUALS", "NOT_EQUALS", "LESS", 
			"LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS", "OR", "DIV", "MOD", 
			"LSHIFT", "RSHIFT", "BIT_CLEAR", "UNDERLYING", "EXCLAMATION", "PLUS", 
			"MINUS", "CARET", "STAR", "AMPERSAND", "RECEIVE", "DECIMAL_LIT", "BINARY_LIT", 
			"OCTAL_LIT", "HEX_LIT", "FLOAT_LIT", "DECIMAL_FLOAT_LIT", "HEX_FLOAT_LIT", 
			"IMAGINARY_LIT", "BYTE_VALUE", "OCTAL_BYTE_VALUE", "HEX_BYTE_VALUE", 
			"LITTLE_U_VALUE", "BIG_U_VALUE", "RAW_STRING_LIT", "INTERPRETED_STRING_LIT", 
			"SINGER_QUOT_STRING_LIT", "WS", "COMMENT", "TERMINATOR", "LINE_COMMENT", 
			"WS_NLSEMI", "COMMENT_NLSEMI", "LINE_COMMENT_NLSEMI", "EOS", "OTHER"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "GoExpression.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public GoExpression(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public Token unary_op;
		public Token mul_op;
		public Token add_op;
		public Token rel_op;
		public PrimaryExprContext primaryExpr() {
			return getRuleContext(PrimaryExprContext.class,0);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode PLUS() { return getToken(GoExpression.PLUS, 0); }
		public TerminalNode MINUS() { return getToken(GoExpression.MINUS, 0); }
		public TerminalNode EXCLAMATION() { return getToken(GoExpression.EXCLAMATION, 0); }
		public TerminalNode CARET() { return getToken(GoExpression.CARET, 0); }
		public TerminalNode STAR() { return getToken(GoExpression.STAR, 0); }
		public TerminalNode AMPERSAND() { return getToken(GoExpression.AMPERSAND, 0); }
		public TerminalNode RECEIVE() { return getToken(GoExpression.RECEIVE, 0); }
		public TerminalNode DIV() { return getToken(GoExpression.DIV, 0); }
		public TerminalNode MOD() { return getToken(GoExpression.MOD, 0); }
		public TerminalNode LSHIFT() { return getToken(GoExpression.LSHIFT, 0); }
		public TerminalNode RSHIFT() { return getToken(GoExpression.RSHIFT, 0); }
		public TerminalNode BIT_CLEAR() { return getToken(GoExpression.BIT_CLEAR, 0); }
		public TerminalNode OR() { return getToken(GoExpression.OR, 0); }
		public TerminalNode EQUALS() { return getToken(GoExpression.EQUALS, 0); }
		public TerminalNode NOT_EQUALS() { return getToken(GoExpression.NOT_EQUALS, 0); }
		public TerminalNode LESS() { return getToken(GoExpression.LESS, 0); }
		public TerminalNode LESS_OR_EQUALS() { return getToken(GoExpression.LESS_OR_EQUALS, 0); }
		public TerminalNode GREATER() { return getToken(GoExpression.GREATER, 0); }
		public TerminalNode GREATER_OR_EQUALS() { return getToken(GoExpression.GREATER_OR_EQUALS, 0); }
		public TerminalNode LOGICAL_AND() { return getToken(GoExpression.LOGICAL_AND, 0); }
		public TerminalNode LOGICAL_OR() { return getToken(GoExpression.LOGICAL_OR, 0); }
		public TerminalNode Question() { return getToken(GoExpression.Question, 0); }
		public TerminalNode COLON() { return getToken(GoExpression.COLON, 0); }
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 0;
		enterRecursionRule(_localctx, 0, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(34);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NIL_LIT:
			case IDENTIFIER:
			case L_PAREN:
			case DECIMAL_LIT:
			case BINARY_LIT:
			case OCTAL_LIT:
			case HEX_LIT:
			case FLOAT_LIT:
			case IMAGINARY_LIT:
			case RAW_STRING_LIT:
			case INTERPRETED_STRING_LIT:
			case SINGER_QUOT_STRING_LIT:
				{
				setState(31);
				primaryExpr(0);
				}
				break;
			case EXCLAMATION:
			case PLUS:
			case MINUS:
			case CARET:
			case STAR:
			case AMPERSAND:
			case RECEIVE:
				{
				setState(32);
				((ExpressionContext)_localctx).unary_op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(((((_la - 60)) & ~0x3f) == 0 && ((1L << (_la - 60)) & 127L) != 0)) ) {
					((ExpressionContext)_localctx).unary_op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(33);
				expression(7);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(59);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(57);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
					case 1:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(36);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(37);
						((ExpressionContext)_localctx).mul_op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(((((_la - 54)) & ~0x3f) == 0 && ((1L << (_la - 54)) & 3103L) != 0)) ) {
							((ExpressionContext)_localctx).mul_op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(38);
						expression(7);
						}
						break;
					case 2:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(39);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(40);
						((ExpressionContext)_localctx).add_op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & -2296835809958952960L) != 0)) ) {
							((ExpressionContext)_localctx).add_op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(41);
						expression(6);
						}
						break;
					case 3:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(42);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(43);
						((ExpressionContext)_localctx).rel_op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 8866461766385664L) != 0)) ) {
							((ExpressionContext)_localctx).rel_op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(44);
						expression(5);
						}
						break;
					case 4:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(45);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(46);
						match(LOGICAL_AND);
						setState(47);
						expression(4);
						}
						break;
					case 5:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(48);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(49);
						match(LOGICAL_OR);
						setState(50);
						expression(3);
						}
						break;
					case 6:
						{
						_localctx = new ExpressionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(51);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(52);
						match(Question);
						setState(53);
						expression(0);
						setState(54);
						match(COLON);
						setState(55);
						expression(2);
						}
						break;
					}
					} 
				}
				setState(61);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrimaryExprContext extends ParserRuleContext {
		public OperandContext operand() {
			return getRuleContext(OperandContext.class,0);
		}
		public PrimaryExprContext primaryExpr() {
			return getRuleContext(PrimaryExprContext.class,0);
		}
		public FieldContext field() {
			return getRuleContext(FieldContext.class,0);
		}
		public IndexContext index() {
			return getRuleContext(IndexContext.class,0);
		}
		public SliceContext slice() {
			return getRuleContext(SliceContext.class,0);
		}
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public PrimaryExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primaryExpr; }
	}

	public final PrimaryExprContext primaryExpr() throws RecognitionException {
		return primaryExpr(0);
	}

	private PrimaryExprContext primaryExpr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		PrimaryExprContext _localctx = new PrimaryExprContext(_ctx, _parentState);
		PrimaryExprContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_primaryExpr, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(63);
			operand();
			}
			_ctx.stop = _input.LT(-1);
			setState(74);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new PrimaryExprContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_primaryExpr);
					setState(65);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(70);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
					case 1:
						{
						setState(66);
						field();
						}
						break;
					case 2:
						{
						setState(67);
						index();
						}
						break;
					case 3:
						{
						setState(68);
						slice();
						}
						break;
					case 4:
						{
						setState(69);
						arguments();
						}
						break;
					}
					}
					} 
				}
				setState(76);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OperandContext extends ParserRuleContext {
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public OperandNameContext operandName() {
			return getRuleContext(OperandNameContext.class,0);
		}
		public TerminalNode L_PAREN() { return getToken(GoExpression.L_PAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode R_PAREN() { return getToken(GoExpression.R_PAREN, 0); }
		public OperandContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operand; }
	}

	public final OperandContext operand() throws RecognitionException {
		OperandContext _localctx = new OperandContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_operand);
		try {
			setState(83);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NIL_LIT:
			case DECIMAL_LIT:
			case BINARY_LIT:
			case OCTAL_LIT:
			case HEX_LIT:
			case FLOAT_LIT:
			case IMAGINARY_LIT:
			case RAW_STRING_LIT:
			case INTERPRETED_STRING_LIT:
			case SINGER_QUOT_STRING_LIT:
				enterOuterAlt(_localctx, 1);
				{
				setState(77);
				literal();
				}
				break;
			case IDENTIFIER:
				enterOuterAlt(_localctx, 2);
				{
				setState(78);
				operandName();
				}
				break;
			case L_PAREN:
				enterOuterAlt(_localctx, 3);
				{
				setState(79);
				match(L_PAREN);
				setState(80);
				expression(0);
				setState(81);
				match(R_PAREN);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralContext extends ParserRuleContext {
		public LiteralNilContext literalNil() {
			return getRuleContext(LiteralNilContext.class,0);
		}
		public IntegerContext integer() {
			return getRuleContext(IntegerContext.class,0);
		}
		public StringContext string() {
			return getRuleContext(StringContext.class,0);
		}
		public LiteralFloatContext literalFloat() {
			return getRuleContext(LiteralFloatContext.class,0);
		}
		public LiteralImagContext literalImag() {
			return getRuleContext(LiteralImagContext.class,0);
		}
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_literal);
		try {
			setState(90);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NIL_LIT:
				enterOuterAlt(_localctx, 1);
				{
				setState(85);
				literalNil();
				}
				break;
			case DECIMAL_LIT:
			case BINARY_LIT:
			case OCTAL_LIT:
			case HEX_LIT:
				enterOuterAlt(_localctx, 2);
				{
				setState(86);
				integer();
				}
				break;
			case RAW_STRING_LIT:
			case INTERPRETED_STRING_LIT:
			case SINGER_QUOT_STRING_LIT:
				enterOuterAlt(_localctx, 3);
				{
				setState(87);
				string();
				}
				break;
			case FLOAT_LIT:
				enterOuterAlt(_localctx, 4);
				{
				setState(88);
				literalFloat();
				}
				break;
			case IMAGINARY_LIT:
				enterOuterAlt(_localctx, 5);
				{
				setState(89);
				literalImag();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralNilContext extends ParserRuleContext {
		public TerminalNode NIL_LIT() { return getToken(GoExpression.NIL_LIT, 0); }
		public LiteralNilContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literalNil; }
	}

	public final LiteralNilContext literalNil() throws RecognitionException {
		LiteralNilContext _localctx = new LiteralNilContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_literalNil);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(92);
			match(NIL_LIT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IntegerContext extends ParserRuleContext {
		public TerminalNode DECIMAL_LIT() { return getToken(GoExpression.DECIMAL_LIT, 0); }
		public TerminalNode BINARY_LIT() { return getToken(GoExpression.BINARY_LIT, 0); }
		public TerminalNode OCTAL_LIT() { return getToken(GoExpression.OCTAL_LIT, 0); }
		public TerminalNode HEX_LIT() { return getToken(GoExpression.HEX_LIT, 0); }
		public IntegerContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_integer; }
	}

	public final IntegerContext integer() throws RecognitionException {
		IntegerContext _localctx = new IntegerContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_integer);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(94);
			_la = _input.LA(1);
			if ( !(((((_la - 67)) & ~0x3f) == 0 && ((1L << (_la - 67)) & 15L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StringContext extends ParserRuleContext {
		public TerminalNode RAW_STRING_LIT() { return getToken(GoExpression.RAW_STRING_LIT, 0); }
		public TerminalNode INTERPRETED_STRING_LIT() { return getToken(GoExpression.INTERPRETED_STRING_LIT, 0); }
		public TerminalNode SINGER_QUOT_STRING_LIT() { return getToken(GoExpression.SINGER_QUOT_STRING_LIT, 0); }
		public StringContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_string; }
	}

	public final StringContext string() throws RecognitionException {
		StringContext _localctx = new StringContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_string);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(96);
			_la = _input.LA(1);
			if ( !(((((_la - 80)) & ~0x3f) == 0 && ((1L << (_la - 80)) & 7L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralFloatContext extends ParserRuleContext {
		public TerminalNode FLOAT_LIT() { return getToken(GoExpression.FLOAT_LIT, 0); }
		public LiteralFloatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literalFloat; }
	}

	public final LiteralFloatContext literalFloat() throws RecognitionException {
		LiteralFloatContext _localctx = new LiteralFloatContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_literalFloat);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(98);
			match(FLOAT_LIT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralImagContext extends ParserRuleContext {
		public TerminalNode IMAGINARY_LIT() { return getToken(GoExpression.IMAGINARY_LIT, 0); }
		public LiteralImagContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literalImag; }
	}

	public final LiteralImagContext literalImag() throws RecognitionException {
		LiteralImagContext _localctx = new LiteralImagContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_literalImag);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(100);
			match(IMAGINARY_LIT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OperandNameContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(GoExpression.IDENTIFIER, 0); }
		public OperandNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operandName; }
	}

	public final OperandNameContext operandName() throws RecognitionException {
		OperandNameContext _localctx = new OperandNameContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_operandName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(102);
			match(IDENTIFIER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FieldContext extends ParserRuleContext {
		public TerminalNode SafeIndex() { return getToken(GoExpression.SafeIndex, 0); }
		public TerminalNode IDENTIFIER() { return getToken(GoExpression.IDENTIFIER, 0); }
		public TerminalNode DOT() { return getToken(GoExpression.DOT, 0); }
		public FieldContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_field; }
	}

	public final FieldContext field() throws RecognitionException {
		FieldContext _localctx = new FieldContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_field);
		try {
			setState(108);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SafeIndex:
				enterOuterAlt(_localctx, 1);
				{
				setState(104);
				match(SafeIndex);
				setState(105);
				match(IDENTIFIER);
				}
				break;
			case DOT:
				enterOuterAlt(_localctx, 2);
				{
				setState(106);
				match(DOT);
				setState(107);
				match(IDENTIFIER);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IndexContext extends ParserRuleContext {
		public TerminalNode L_BRACKET() { return getToken(GoExpression.L_BRACKET, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode R_BRACKET() { return getToken(GoExpression.R_BRACKET, 0); }
		public IndexContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_index; }
	}

	public final IndexContext index() throws RecognitionException {
		IndexContext _localctx = new IndexContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_index);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(110);
			match(L_BRACKET);
			setState(111);
			expression(0);
			setState(112);
			match(R_BRACKET);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SliceContext extends ParserRuleContext {
		public ExpressionContext lo;
		public ExpressionContext hi;
		public ExpressionContext cap;
		public TerminalNode L_BRACKET() { return getToken(GoExpression.L_BRACKET, 0); }
		public TerminalNode R_BRACKET() { return getToken(GoExpression.R_BRACKET, 0); }
		public List<TerminalNode> COLON() { return getTokens(GoExpression.COLON); }
		public TerminalNode COLON(int i) {
			return getToken(GoExpression.COLON, i);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public SliceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_slice; }
	}

	public final SliceContext slice() throws RecognitionException {
		SliceContext _localctx = new SliceContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_slice);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(114);
			match(L_BRACKET);
			setState(130);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				{
				setState(116);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 26)) & ~0x3f) == 0 && ((1L << (_la - 26)) & 126452616107393031L) != 0)) {
					{
					setState(115);
					((SliceContext)_localctx).lo = expression(0);
					}
				}

				setState(118);
				match(COLON);
				setState(120);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 26)) & ~0x3f) == 0 && ((1L << (_la - 26)) & 126452616107393031L) != 0)) {
					{
					setState(119);
					((SliceContext)_localctx).hi = expression(0);
					}
				}

				}
				break;
			case 2:
				{
				setState(123);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 26)) & ~0x3f) == 0 && ((1L << (_la - 26)) & 126452616107393031L) != 0)) {
					{
					setState(122);
					((SliceContext)_localctx).lo = expression(0);
					}
				}

				setState(125);
				match(COLON);
				setState(126);
				((SliceContext)_localctx).hi = expression(0);
				setState(127);
				match(COLON);
				setState(128);
				((SliceContext)_localctx).cap = expression(0);
				}
				break;
			}
			setState(132);
			match(R_BRACKET);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ArgumentsContext extends ParserRuleContext {
		public TerminalNode L_PAREN() { return getToken(GoExpression.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(GoExpression.R_PAREN, 0); }
		public ExpressionListContext expressionList() {
			return getRuleContext(ExpressionListContext.class,0);
		}
		public TerminalNode ELLIPSIS() { return getToken(GoExpression.ELLIPSIS, 0); }
		public TerminalNode COMMA() { return getToken(GoExpression.COMMA, 0); }
		public ArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arguments; }
	}

	public final ArgumentsContext arguments() throws RecognitionException {
		ArgumentsContext _localctx = new ArgumentsContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_arguments);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(134);
			match(L_PAREN);
			setState(142);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 26)) & ~0x3f) == 0 && ((1L << (_la - 26)) & 126452616107393031L) != 0)) {
				{
				{
				setState(135);
				expressionList();
				}
				setState(137);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ELLIPSIS) {
					{
					setState(136);
					match(ELLIPSIS);
					}
				}

				setState(140);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COMMA) {
					{
					setState(139);
					match(COMMA);
					}
				}

				}
			}

			setState(144);
			match(R_PAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionListContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(GoExpression.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(GoExpression.COMMA, i);
		}
		public ExpressionListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionList; }
	}

	public final ExpressionListContext expressionList() throws RecognitionException {
		ExpressionListContext _localctx = new ExpressionListContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_expressionList);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(146);
			expression(0);
			setState(151);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(147);
					match(COMMA);
					setState(148);
					expression(0);
					}
					} 
				}
				setState(153);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 0:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		case 1:
			return primaryExpr_sempred((PrimaryExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 6);
		case 1:
			return precpred(_ctx, 5);
		case 2:
			return precpred(_ctx, 4);
		case 3:
			return precpred(_ctx, 3);
		case 4:
			return precpred(_ctx, 2);
		case 5:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean primaryExpr_sempred(PrimaryExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 6:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001[\u009b\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0003\u0000#\b\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0005\u0000:\b\u0000\n\u0000\f\u0000=\t\u0000\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0003\u0001G\b\u0001\u0005\u0001I\b\u0001\n\u0001\f\u0001"+
		"L\t\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0003\u0002T\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0003\u0003[\b\u0003\u0001\u0004\u0001\u0004"+
		"\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007"+
		"\u0001\b\u0001\b\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\n\u0003"+
		"\nm\b\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\f\u0001"+
		"\f\u0003\fu\b\f\u0001\f\u0001\f\u0003\fy\b\f\u0001\f\u0003\f|\b\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0003\f\u0083\b\f\u0001\f\u0001\f\u0001"+
		"\r\u0001\r\u0001\r\u0003\r\u008a\b\r\u0001\r\u0003\r\u008d\b\r\u0003\r"+
		"\u008f\b\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0005\u000e"+
		"\u0096\b\u000e\n\u000e\f\u000e\u0099\t\u000e\u0001\u000e\u0000\u0002\u0000"+
		"\u0002\u000f\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016"+
		"\u0018\u001a\u001c\u0000\u0006\u0001\u0000<B\u0002\u00006:@A\u0002\u0000"+
		"55=?\u0001\u0000/4\u0001\u0000CF\u0001\u0000PR\u00a5\u0000\"\u0001\u0000"+
		"\u0000\u0000\u0002>\u0001\u0000\u0000\u0000\u0004S\u0001\u0000\u0000\u0000"+
		"\u0006Z\u0001\u0000\u0000\u0000\b\\\u0001\u0000\u0000\u0000\n^\u0001\u0000"+
		"\u0000\u0000\f`\u0001\u0000\u0000\u0000\u000eb\u0001\u0000\u0000\u0000"+
		"\u0010d\u0001\u0000\u0000\u0000\u0012f\u0001\u0000\u0000\u0000\u0014l"+
		"\u0001\u0000\u0000\u0000\u0016n\u0001\u0000\u0000\u0000\u0018r\u0001\u0000"+
		"\u0000\u0000\u001a\u0086\u0001\u0000\u0000\u0000\u001c\u0092\u0001\u0000"+
		"\u0000\u0000\u001e\u001f\u0006\u0000\uffff\uffff\u0000\u001f#\u0003\u0002"+
		"\u0001\u0000 !\u0007\u0000\u0000\u0000!#\u0003\u0000\u0000\u0007\"\u001e"+
		"\u0001\u0000\u0000\u0000\" \u0001\u0000\u0000\u0000#;\u0001\u0000\u0000"+
		"\u0000$%\n\u0006\u0000\u0000%&\u0007\u0001\u0000\u0000&:\u0003\u0000\u0000"+
		"\u0007\'(\n\u0005\u0000\u0000()\u0007\u0002\u0000\u0000):\u0003\u0000"+
		"\u0000\u0006*+\n\u0004\u0000\u0000+,\u0007\u0003\u0000\u0000,:\u0003\u0000"+
		"\u0000\u0005-.\n\u0003\u0000\u0000./\u0005.\u0000\u0000/:\u0003\u0000"+
		"\u0000\u000401\n\u0002\u0000\u000012\u0005-\u0000\u00002:\u0003\u0000"+
		"\u0000\u000334\n\u0001\u0000\u000045\u0005+\u0000\u000056\u0003\u0000"+
		"\u0000\u000067\u0005%\u0000\u000078\u0003\u0000\u0000\u00028:\u0001\u0000"+
		"\u0000\u00009$\u0001\u0000\u0000\u00009\'\u0001\u0000\u0000\u00009*\u0001"+
		"\u0000\u0000\u00009-\u0001\u0000\u0000\u000090\u0001\u0000\u0000\u0000"+
		"93\u0001\u0000\u0000\u0000:=\u0001\u0000\u0000\u0000;9\u0001\u0000\u0000"+
		"\u0000;<\u0001\u0000\u0000\u0000<\u0001\u0001\u0000\u0000\u0000=;\u0001"+
		"\u0000\u0000\u0000>?\u0006\u0001\uffff\uffff\u0000?@\u0003\u0004\u0002"+
		"\u0000@J\u0001\u0000\u0000\u0000AF\n\u0001\u0000\u0000BG\u0003\u0014\n"+
		"\u0000CG\u0003\u0016\u000b\u0000DG\u0003\u0018\f\u0000EG\u0003\u001a\r"+
		"\u0000FB\u0001\u0000\u0000\u0000FC\u0001\u0000\u0000\u0000FD\u0001\u0000"+
		"\u0000\u0000FE\u0001\u0000\u0000\u0000GI\u0001\u0000\u0000\u0000HA\u0001"+
		"\u0000\u0000\u0000IL\u0001\u0000\u0000\u0000JH\u0001\u0000\u0000\u0000"+
		"JK\u0001\u0000\u0000\u0000K\u0003\u0001\u0000\u0000\u0000LJ\u0001\u0000"+
		"\u0000\u0000MT\u0003\u0006\u0003\u0000NT\u0003\u0012\t\u0000OP\u0005\u001c"+
		"\u0000\u0000PQ\u0003\u0000\u0000\u0000QR\u0005\u001d\u0000\u0000RT\u0001"+
		"\u0000\u0000\u0000SM\u0001\u0000\u0000\u0000SN\u0001\u0000\u0000\u0000"+
		"SO\u0001\u0000\u0000\u0000T\u0005\u0001\u0000\u0000\u0000U[\u0003\b\u0004"+
		"\u0000V[\u0003\n\u0005\u0000W[\u0003\f\u0006\u0000X[\u0003\u000e\u0007"+
		"\u0000Y[\u0003\u0010\b\u0000ZU\u0001\u0000\u0000\u0000ZV\u0001\u0000\u0000"+
		"\u0000ZW\u0001\u0000\u0000\u0000ZX\u0001\u0000\u0000\u0000ZY\u0001\u0000"+
		"\u0000\u0000[\u0007\u0001\u0000\u0000\u0000\\]\u0005\u001a\u0000\u0000"+
		"]\t\u0001\u0000\u0000\u0000^_\u0007\u0004\u0000\u0000_\u000b\u0001\u0000"+
		"\u0000\u0000`a\u0007\u0005\u0000\u0000a\r\u0001\u0000\u0000\u0000bc\u0005"+
		"G\u0000\u0000c\u000f\u0001\u0000\u0000\u0000de\u0005J\u0000\u0000e\u0011"+
		"\u0001\u0000\u0000\u0000fg\u0005\u001b\u0000\u0000g\u0013\u0001\u0000"+
		"\u0000\u0000hi\u0005,\u0000\u0000im\u0005\u001b\u0000\u0000jk\u0005&\u0000"+
		"\u0000km\u0005\u001b\u0000\u0000lh\u0001\u0000\u0000\u0000lj\u0001\u0000"+
		"\u0000\u0000m\u0015\u0001\u0000\u0000\u0000no\u0005 \u0000\u0000op\u0003"+
		"\u0000\u0000\u0000pq\u0005!\u0000\u0000q\u0017\u0001\u0000\u0000\u0000"+
		"r\u0082\u0005 \u0000\u0000su\u0003\u0000\u0000\u0000ts\u0001\u0000\u0000"+
		"\u0000tu\u0001\u0000\u0000\u0000uv\u0001\u0000\u0000\u0000vx\u0005%\u0000"+
		"\u0000wy\u0003\u0000\u0000\u0000xw\u0001\u0000\u0000\u0000xy\u0001\u0000"+
		"\u0000\u0000y\u0083\u0001\u0000\u0000\u0000z|\u0003\u0000\u0000\u0000"+
		"{z\u0001\u0000\u0000\u0000{|\u0001\u0000\u0000\u0000|}\u0001\u0000\u0000"+
		"\u0000}~\u0005%\u0000\u0000~\u007f\u0003\u0000\u0000\u0000\u007f\u0080"+
		"\u0005%\u0000\u0000\u0080\u0081\u0003\u0000\u0000\u0000\u0081\u0083\u0001"+
		"\u0000\u0000\u0000\u0082t\u0001\u0000\u0000\u0000\u0082{\u0001\u0000\u0000"+
		"\u0000\u0083\u0084\u0001\u0000\u0000\u0000\u0084\u0085\u0005!\u0000\u0000"+
		"\u0085\u0019\u0001\u0000\u0000\u0000\u0086\u008e\u0005\u001c\u0000\u0000"+
		"\u0087\u0089\u0003\u001c\u000e\u0000\u0088\u008a\u0005*\u0000\u0000\u0089"+
		"\u0088\u0001\u0000\u0000\u0000\u0089\u008a\u0001\u0000\u0000\u0000\u008a"+
		"\u008c\u0001\u0000\u0000\u0000\u008b\u008d\u0005#\u0000\u0000\u008c\u008b"+
		"\u0001\u0000\u0000\u0000\u008c\u008d\u0001\u0000\u0000\u0000\u008d\u008f"+
		"\u0001\u0000\u0000\u0000\u008e\u0087\u0001\u0000\u0000\u0000\u008e\u008f"+
		"\u0001\u0000\u0000\u0000\u008f\u0090\u0001\u0000\u0000\u0000\u0090\u0091"+
		"\u0005\u001d\u0000\u0000\u0091\u001b\u0001\u0000\u0000\u0000\u0092\u0097"+
		"\u0003\u0000\u0000\u0000\u0093\u0094\u0005#\u0000\u0000\u0094\u0096\u0003"+
		"\u0000\u0000\u0000\u0095\u0093\u0001\u0000\u0000\u0000\u0096\u0099\u0001"+
		"\u0000\u0000\u0000\u0097\u0095\u0001\u0000\u0000\u0000\u0097\u0098\u0001"+
		"\u0000\u0000\u0000\u0098\u001d\u0001\u0000\u0000\u0000\u0099\u0097\u0001"+
		"\u0000\u0000\u0000\u0010\"9;FJSZltx{\u0082\u0089\u008c\u008e\u0097";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
// Generated from c:\Users\brian\OneDrive\Escritorio\Repositories\proyecto\Calc.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class CalcParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T_INT=3, T_DOUBLE=4, T_STRING=5, T_CHAR=6, T_BOOLEAN=7, 
		VAR=8, LET=9, IF=10, ELSE=11, FOR=12, WHILE=13, REPEAT=14, BREAK=15, CONTINUE=16, 
		SWITCH=17, CASE=18, DEFAULT=19, ADD=20, SUB=21, MUL=22, DIV=23, MOD=24, 
		EQU=25, DIF=26, GREATER=27, GOETO=28, LESS=29, LOETO=30, AND=31, OR=32, 
		NOT=33, CLOSE_RANGE=34, SOPEN_RANGE=35, INCREMENT=36, DECREMENT=37, ID=38, 
		DOUBLE=39, INT=40, WHITESPACE=41;
	public static final int
		RULE_s = 0, RULE_e = 1, RULE_ep = 2, RULE_t = 3, RULE_tp = 4, RULE_f = 5, 
		RULE_value = 6;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "e", "ep", "t", "tp", "f", "value"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'('", "')'", "'int'", "'double'", "'string'", "'char'", "'bool'", 
			"'var'", "'let'", "'if'", "'else'", "'for'", "'while'", "'repear'", "'break'", 
			"'continue'", "'switch'", "'case'", "'default'", "'+'", "'-'", "'*'", 
			"'/'", "'%'", "'=='", "'!='", "'>'", "'>='", "'<'", "'<='", "'&&'", "'||'", 
			"'!'", "'...'", "'..<'", "'++'", "'--'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, "T_INT", "T_DOUBLE", "T_STRING", "T_CHAR", "T_BOOLEAN", 
			"VAR", "LET", "IF", "ELSE", "FOR", "WHILE", "REPEAT", "BREAK", "CONTINUE", 
			"SWITCH", "CASE", "DEFAULT", "ADD", "SUB", "MUL", "DIV", "MOD", "EQU", 
			"DIF", "GREATER", "GOETO", "LESS", "LOETO", "AND", "OR", "NOT", "CLOSE_RANGE", 
			"SOPEN_RANGE", "INCREMENT", "DECREMENT", "ID", "DOUBLE", "INT", "WHITESPACE"
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
	public String getGrammarFileName() { return "Calc.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public CalcParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class SContext extends ParserRuleContext {
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode EOF() { return getToken(CalcParser.EOF, 0); }
		public SContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_s; }
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(14);
			e();
			setState(15);
			match(EOF);
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

	public static class EContext extends ParserRuleContext {
		public EContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_e; }
	 
		public EContext() { }
		public void copyFrom(EContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class ETContext extends EContext {
		public TContext t() {
			return getRuleContext(TContext.class,0);
		}
		public EpContext ep() {
			return getRuleContext(EpContext.class,0);
		}
		public ETContext(EContext ctx) { copyFrom(ctx); }
	}

	public final EContext e() throws RecognitionException {
		EContext _localctx = new EContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_e);
		try {
			_localctx = new ETContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(17);
			t();
			setState(18);
			ep();
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

	public static class EpContext extends ParserRuleContext {
		public EpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ep; }
	 
		public EpContext() { }
		public void copyFrom(EpContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class AddContext extends EpContext {
		public TerminalNode ADD() { return getToken(CalcParser.ADD, 0); }
		public TContext t() {
			return getRuleContext(TContext.class,0);
		}
		public EpContext ep() {
			return getRuleContext(EpContext.class,0);
		}
		public AddContext(EpContext ctx) { copyFrom(ctx); }
	}
	public static class SubContext extends EpContext {
		public TerminalNode SUB() { return getToken(CalcParser.SUB, 0); }
		public TContext t() {
			return getRuleContext(TContext.class,0);
		}
		public EpContext ep() {
			return getRuleContext(EpContext.class,0);
		}
		public SubContext(EpContext ctx) { copyFrom(ctx); }
	}
	public static class EpsAddSubContext extends EpContext {
		public EpsAddSubContext(EpContext ctx) { copyFrom(ctx); }
	}

	public final EpContext ep() throws RecognitionException {
		EpContext _localctx = new EpContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_ep);
		try {
			setState(29);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ADD:
				_localctx = new AddContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(20);
				match(ADD);
				setState(21);
				t();
				setState(22);
				ep();
				}
				break;
			case SUB:
				_localctx = new SubContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(24);
				match(SUB);
				setState(25);
				t();
				setState(26);
				ep();
				}
				break;
			case EOF:
			case T__1:
				_localctx = new EpsAddSubContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
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

	public static class TContext extends ParserRuleContext {
		public TContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_t; }
	 
		public TContext() { }
		public void copyFrom(TContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class TFContext extends TContext {
		public FContext f() {
			return getRuleContext(FContext.class,0);
		}
		public TpContext tp() {
			return getRuleContext(TpContext.class,0);
		}
		public TFContext(TContext ctx) { copyFrom(ctx); }
	}

	public final TContext t() throws RecognitionException {
		TContext _localctx = new TContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_t);
		try {
			_localctx = new TFContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(31);
			f();
			setState(32);
			tp();
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

	public static class TpContext extends ParserRuleContext {
		public TpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tp; }
	 
		public TpContext() { }
		public void copyFrom(TpContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class DivContext extends TpContext {
		public TerminalNode DIV() { return getToken(CalcParser.DIV, 0); }
		public FContext f() {
			return getRuleContext(FContext.class,0);
		}
		public TpContext tp() {
			return getRuleContext(TpContext.class,0);
		}
		public DivContext(TpContext ctx) { copyFrom(ctx); }
	}
	public static class MulContext extends TpContext {
		public TerminalNode MUL() { return getToken(CalcParser.MUL, 0); }
		public FContext f() {
			return getRuleContext(FContext.class,0);
		}
		public TpContext tp() {
			return getRuleContext(TpContext.class,0);
		}
		public MulContext(TpContext ctx) { copyFrom(ctx); }
	}
	public static class EpsMulDivContext extends TpContext {
		public EpsMulDivContext(TpContext ctx) { copyFrom(ctx); }
	}

	public final TpContext tp() throws RecognitionException {
		TpContext _localctx = new TpContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_tp);
		try {
			setState(43);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MUL:
				_localctx = new MulContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(34);
				match(MUL);
				setState(35);
				f();
				setState(36);
				tp();
				}
				break;
			case DIV:
				_localctx = new DivContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(38);
				match(DIV);
				setState(39);
				f();
				setState(40);
				tp();
				}
				break;
			case EOF:
			case T__1:
			case ADD:
			case SUB:
				_localctx = new EpsMulDivContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
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

	public static class FContext extends ParserRuleContext {
		public FContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_f; }
	 
		public FContext() { }
		public void copyFrom(FContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class ValContext extends FContext {
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public ValContext(FContext ctx) { copyFrom(ctx); }
	}
	public static class BraceContext extends FContext {
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public BraceContext(FContext ctx) { copyFrom(ctx); }
	}

	public final FContext f() throws RecognitionException {
		FContext _localctx = new FContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_f);
		try {
			setState(50);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				_localctx = new BraceContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(45);
				match(T__0);
				setState(46);
				e();
				setState(47);
				match(T__1);
				}
				break;
			case ID:
			case DOUBLE:
			case INT:
				_localctx = new ValContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(49);
				value();
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

	public static class ValueContext extends ParserRuleContext {
		public ValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_value; }
	 
		public ValueContext() { }
		public void copyFrom(ValueContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class IdContext extends ValueContext {
		public TerminalNode ID() { return getToken(CalcParser.ID, 0); }
		public IdContext(ValueContext ctx) { copyFrom(ctx); }
	}
	public static class DoubleContext extends ValueContext {
		public TerminalNode DOUBLE() { return getToken(CalcParser.DOUBLE, 0); }
		public DoubleContext(ValueContext ctx) { copyFrom(ctx); }
	}
	public static class IntContext extends ValueContext {
		public TerminalNode INT() { return getToken(CalcParser.INT, 0); }
		public IntContext(ValueContext ctx) { copyFrom(ctx); }
	}

	public final ValueContext value() throws RecognitionException {
		ValueContext _localctx = new ValueContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_value);
		try {
			setState(55);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				_localctx = new IntContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(52);
				match(INT);
				}
				break;
			case DOUBLE:
				_localctx = new DoubleContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(53);
				match(DOUBLE);
				}
				break;
			case ID:
				_localctx = new IdContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(54);
				match(ID);
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

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3+<\4\2\t\2\4\3\t\3"+
		"\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\3\2\3\2\3\2\3\3\3\3\3\3\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4 \n\4\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\5\6.\n\6\3\7\3\7\3\7\3\7\3\7\5\7\65\n\7\3\b\3\b\3\b"+
		"\5\b:\n\b\3\b\2\2\t\2\4\6\b\n\f\16\2\2\2;\2\20\3\2\2\2\4\23\3\2\2\2\6"+
		"\37\3\2\2\2\b!\3\2\2\2\n-\3\2\2\2\f\64\3\2\2\2\169\3\2\2\2\20\21\5\4\3"+
		"\2\21\22\7\2\2\3\22\3\3\2\2\2\23\24\5\b\5\2\24\25\5\6\4\2\25\5\3\2\2\2"+
		"\26\27\7\26\2\2\27\30\5\b\5\2\30\31\5\6\4\2\31 \3\2\2\2\32\33\7\27\2\2"+
		"\33\34\5\b\5\2\34\35\5\6\4\2\35 \3\2\2\2\36 \3\2\2\2\37\26\3\2\2\2\37"+
		"\32\3\2\2\2\37\36\3\2\2\2 \7\3\2\2\2!\"\5\f\7\2\"#\5\n\6\2#\t\3\2\2\2"+
		"$%\7\30\2\2%&\5\f\7\2&\'\5\n\6\2\'.\3\2\2\2()\7\31\2\2)*\5\f\7\2*+\5\n"+
		"\6\2+.\3\2\2\2,.\3\2\2\2-$\3\2\2\2-(\3\2\2\2-,\3\2\2\2.\13\3\2\2\2/\60"+
		"\7\3\2\2\60\61\5\4\3\2\61\62\7\4\2\2\62\65\3\2\2\2\63\65\5\16\b\2\64/"+
		"\3\2\2\2\64\63\3\2\2\2\65\r\3\2\2\2\66:\7*\2\2\67:\7)\2\28:\7(\2\29\66"+
		"\3\2\2\29\67\3\2\2\298\3\2\2\2:\17\3\2\2\2\6\37-\649";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
// Generated from c:\Users\brian\OneDrive\Escritorio\Repositories\proyecto\Grammar.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class GrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T_INT=23, T_FLOAT=24, 
		T_CHAR=25, T_STRING=26, T_BOOL=27, VAR=28, LET=29, DOUBLE=30, INT=31, 
		ID=32, CHAR=33, STRING=34, MULTCOMMENT=35, COMMENT=36, WHITESPACE=37;
	public static final int
		RULE_init = 0, RULE_block = 1, RULE_stmt = 2, RULE_declstmt = 3, RULE_asignstmt = 4, 
		RULE_vartype = 5, RULE_expr = 6;
	private static String[] makeRuleNames() {
		return new String[] {
			"init", "block", "stmt", "declstmt", "asignstmt", "vartype", "expr"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "':'", "'='", "'?'", "'!'", "'-'", "'*'", "'/'", "'%'", "'+'", 
			"'>='", "'>'", "'<='", "'<'", "'=='", "'!='", "'&&'", "'||'", "'('", 
			"')'", "'true'", "'false'", "'nil'", "'Int'", "'Float'", "'Character'", 
			"'String'", "'Bool'", "'var'", "'let'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, "T_INT", 
			"T_FLOAT", "T_CHAR", "T_STRING", "T_BOOL", "VAR", "LET", "DOUBLE", "INT", 
			"ID", "CHAR", "STRING", "MULTCOMMENT", "COMMENT", "WHITESPACE"
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
	public String getGrammarFileName() { return "Grammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public GrammarParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class InitContext extends ParserRuleContext {
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode EOF() { return getToken(GrammarParser.EOF, 0); }
		public InitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_init; }
	}

	public final InitContext init() throws RecognitionException {
		InitContext _localctx = new InitContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_init);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(14);
			block();
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

	public static class BlockContext extends ParserRuleContext {
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(20);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << VAR) | (1L << LET) | (1L << ID))) != 0)) {
				{
				{
				setState(17);
				stmt();
				}
				}
				setState(22);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	public static class StmtContext extends ParserRuleContext {
		public DeclstmtContext declstmt() {
			return getRuleContext(DeclstmtContext.class,0);
		}
		public AsignstmtContext asignstmt() {
			return getRuleContext(AsignstmtContext.class,0);
		}
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_stmt);
		try {
			setState(25);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case VAR:
			case LET:
				enterOuterAlt(_localctx, 1);
				{
				setState(23);
				declstmt();
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 2);
				{
				setState(24);
				asignstmt();
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

	public static class DeclstmtContext extends ParserRuleContext {
		public DeclstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declstmt; }
	 
		public DeclstmtContext() { }
		public void copyFrom(DeclstmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class DeclstmtWithTypeAndExprContext extends DeclstmtContext {
		public Token vtype;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public VartypeContext vartype() {
			return getRuleContext(VartypeContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode VAR() { return getToken(GrammarParser.VAR, 0); }
		public TerminalNode LET() { return getToken(GrammarParser.LET, 0); }
		public DeclstmtWithTypeAndExprContext(DeclstmtContext ctx) { copyFrom(ctx); }
	}
	public static class DeclstmtWithExprContext extends DeclstmtContext {
		public Token vtype;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode VAR() { return getToken(GrammarParser.VAR, 0); }
		public TerminalNode LET() { return getToken(GrammarParser.LET, 0); }
		public DeclstmtWithExprContext(DeclstmtContext ctx) { copyFrom(ctx); }
	}
	public static class DeclstmtWithTypeContext extends DeclstmtContext {
		public Token vtype;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public VartypeContext vartype() {
			return getRuleContext(VartypeContext.class,0);
		}
		public TerminalNode VAR() { return getToken(GrammarParser.VAR, 0); }
		public TerminalNode LET() { return getToken(GrammarParser.LET, 0); }
		public DeclstmtWithTypeContext(DeclstmtContext ctx) { copyFrom(ctx); }
	}

	public final DeclstmtContext declstmt() throws RecognitionException {
		DeclstmtContext _localctx = new DeclstmtContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_declstmt);
		int _la;
		try {
			setState(44);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				_localctx = new DeclstmtWithTypeAndExprContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(27);
				((DeclstmtWithTypeAndExprContext)_localctx).vtype = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==VAR || _la==LET) ) {
					((DeclstmtWithTypeAndExprContext)_localctx).vtype = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(28);
				match(ID);
				setState(29);
				match(T__0);
				setState(30);
				vartype();
				setState(31);
				match(T__1);
				setState(32);
				expr(0);
				}
				break;
			case 2:
				_localctx = new DeclstmtWithExprContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(34);
				((DeclstmtWithExprContext)_localctx).vtype = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==VAR || _la==LET) ) {
					((DeclstmtWithExprContext)_localctx).vtype = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(35);
				match(ID);
				setState(36);
				match(T__1);
				setState(37);
				expr(0);
				}
				break;
			case 3:
				_localctx = new DeclstmtWithTypeContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(38);
				((DeclstmtWithTypeContext)_localctx).vtype = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==VAR || _la==LET) ) {
					((DeclstmtWithTypeContext)_localctx).vtype = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(39);
				match(ID);
				setState(40);
				match(T__0);
				setState(41);
				vartype();
				setState(42);
				match(T__2);
				}
				break;
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

	public static class AsignstmtContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AsignstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignstmt; }
	}

	public final AsignstmtContext asignstmt() throws RecognitionException {
		AsignstmtContext _localctx = new AsignstmtContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_asignstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(46);
			match(ID);
			setState(47);
			match(T__1);
			setState(48);
			expr(0);
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

	public static class VartypeContext extends ParserRuleContext {
		public VartypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vartype; }
	 
		public VartypeContext() { }
		public void copyFrom(VartypeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class PrimitiveTypeContext extends VartypeContext {
		public Token tipo;
		public TerminalNode T_INT() { return getToken(GrammarParser.T_INT, 0); }
		public TerminalNode T_FLOAT() { return getToken(GrammarParser.T_FLOAT, 0); }
		public TerminalNode T_CHAR() { return getToken(GrammarParser.T_CHAR, 0); }
		public TerminalNode T_STRING() { return getToken(GrammarParser.T_STRING, 0); }
		public TerminalNode T_BOOL() { return getToken(GrammarParser.T_BOOL, 0); }
		public PrimitiveTypeContext(VartypeContext ctx) { copyFrom(ctx); }
	}

	public final VartypeContext vartype() throws RecognitionException {
		VartypeContext _localctx = new VartypeContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_vartype);
		try {
			setState(55);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T_INT:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(50);
				((PrimitiveTypeContext)_localctx).tipo = match(T_INT);
				}
				break;
			case T_FLOAT:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(51);
				((PrimitiveTypeContext)_localctx).tipo = match(T_FLOAT);
				}
				break;
			case T_CHAR:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(52);
				((PrimitiveTypeContext)_localctx).tipo = match(T_CHAR);
				}
				break;
			case T_STRING:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(53);
				((PrimitiveTypeContext)_localctx).tipo = match(T_STRING);
				}
				break;
			case T_BOOL:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(54);
				((PrimitiveTypeContext)_localctx).tipo = match(T_BOOL);
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

	public static class ExprContext extends ParserRuleContext {
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	 
		public ExprContext() { }
		public void copyFrom(ExprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class DoubleExprContext extends ExprContext {
		public TerminalNode DOUBLE() { return getToken(GrammarParser.DOUBLE, 0); }
		public DoubleExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class BoolExprContext extends ExprContext {
		public BoolExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class IdExprContext extends ExprContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public IdExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class ParExprContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ParExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class StrExprContext extends ExprContext {
		public TerminalNode STRING() { return getToken(GrammarParser.STRING, 0); }
		public StrExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class NotExprContext extends ExprContext {
		public ExprContext right;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public NotExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class IntExprContext extends ExprContext {
		public TerminalNode INT() { return getToken(GrammarParser.INT, 0); }
		public IntExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class NilExprContext extends ExprContext {
		public NilExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class OpExprContext extends ExprContext {
		public ExprContext left;
		public Token op;
		public ExprContext right;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public OpExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class UnaryExprContext extends ExprContext {
		public ExprContext right;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public UnaryExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class CharExprContext extends ExprContext {
		public TerminalNode CHAR() { return getToken(GrammarParser.CHAR, 0); }
		public CharExprContext(ExprContext ctx) { copyFrom(ctx); }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 12;
		enterRecursionRule(_localctx, 12, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(73);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__3:
				{
				_localctx = new NotExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(58);
				match(T__3);
				setState(59);
				((NotExprContext)_localctx).right = expr(17);
				}
				break;
			case T__4:
				{
				_localctx = new UnaryExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(60);
				match(T__4);
				setState(61);
				((UnaryExprContext)_localctx).right = expr(16);
				}
				break;
			case T__17:
				{
				_localctx = new ParExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(62);
				match(T__17);
				setState(63);
				expr(0);
				setState(64);
				match(T__18);
				}
				break;
			case INT:
				{
				_localctx = new IntExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(66);
				match(INT);
				}
				break;
			case DOUBLE:
				{
				_localctx = new DoubleExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(67);
				match(DOUBLE);
				}
				break;
			case ID:
				{
				_localctx = new IdExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(68);
				match(ID);
				}
				break;
			case CHAR:
				{
				_localctx = new CharExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(69);
				match(CHAR);
				}
				break;
			case STRING:
				{
				_localctx = new StrExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(70);
				match(STRING);
				}
				break;
			case T__19:
			case T__20:
				{
				_localctx = new BoolExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(71);
				_la = _input.LA(1);
				if ( !(_la==T__19 || _la==T__20) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case T__21:
				{
				_localctx = new NilExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(72);
				match(T__21);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(98);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(96);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
					case 1:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(75);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(76);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__5) | (1L << T__6) | (1L << T__7))) != 0)) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(77);
						((OpExprContext)_localctx).right = expr(16);
						}
						break;
					case 2:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(78);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(79);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__4 || _la==T__8) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(80);
						((OpExprContext)_localctx).right = expr(15);
						}
						break;
					case 3:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(81);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(82);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__9 || _la==T__10) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(83);
						((OpExprContext)_localctx).right = expr(14);
						}
						break;
					case 4:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(84);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(85);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__11 || _la==T__12) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(86);
						((OpExprContext)_localctx).right = expr(13);
						}
						break;
					case 5:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(87);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(88);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__13 || _la==T__14) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(89);
						((OpExprContext)_localctx).right = expr(12);
						}
						break;
					case 6:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(90);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(91);
						match(T__15);
						setState(92);
						((OpExprContext)_localctx).right = expr(11);
						}
						break;
					case 7:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(93);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(94);
						match(T__16);
						setState(95);
						((OpExprContext)_localctx).right = expr(10);
						}
						break;
					}
					} 
				}
				setState(100);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 6:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 15);
		case 1:
			return precpred(_ctx, 14);
		case 2:
			return precpred(_ctx, 13);
		case 3:
			return precpred(_ctx, 12);
		case 4:
			return precpred(_ctx, 11);
		case 5:
			return precpred(_ctx, 10);
		case 6:
			return precpred(_ctx, 9);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\'h\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\3\2\3\2\3\2\3\3\7\3\25\n\3"+
		"\f\3\16\3\30\13\3\3\4\3\4\5\4\34\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\5\5/\n\5\3\6\3\6\3\6\3\6\3\7\3\7\3"+
		"\7\3\7\3\7\5\7:\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\5\bL\n\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\7\bc\n\b\f\b\16\bf\13\b\3\b\2\3"+
		"\16\t\2\4\6\b\n\f\16\2\t\3\2\36\37\3\2\26\27\3\2\b\n\4\2\7\7\13\13\3\2"+
		"\f\r\3\2\16\17\3\2\20\21\2x\2\20\3\2\2\2\4\26\3\2\2\2\6\33\3\2\2\2\b."+
		"\3\2\2\2\n\60\3\2\2\2\f9\3\2\2\2\16K\3\2\2\2\20\21\5\4\3\2\21\22\7\2\2"+
		"\3\22\3\3\2\2\2\23\25\5\6\4\2\24\23\3\2\2\2\25\30\3\2\2\2\26\24\3\2\2"+
		"\2\26\27\3\2\2\2\27\5\3\2\2\2\30\26\3\2\2\2\31\34\5\b\5\2\32\34\5\n\6"+
		"\2\33\31\3\2\2\2\33\32\3\2\2\2\34\7\3\2\2\2\35\36\t\2\2\2\36\37\7\"\2"+
		"\2\37 \7\3\2\2 !\5\f\7\2!\"\7\4\2\2\"#\5\16\b\2#/\3\2\2\2$%\t\2\2\2%&"+
		"\7\"\2\2&\'\7\4\2\2\'/\5\16\b\2()\t\2\2\2)*\7\"\2\2*+\7\3\2\2+,\5\f\7"+
		"\2,-\7\5\2\2-/\3\2\2\2.\35\3\2\2\2.$\3\2\2\2.(\3\2\2\2/\t\3\2\2\2\60\61"+
		"\7\"\2\2\61\62\7\4\2\2\62\63\5\16\b\2\63\13\3\2\2\2\64:\7\31\2\2\65:\7"+
		"\32\2\2\66:\7\33\2\2\67:\7\34\2\28:\7\35\2\29\64\3\2\2\29\65\3\2\2\29"+
		"\66\3\2\2\29\67\3\2\2\298\3\2\2\2:\r\3\2\2\2;<\b\b\1\2<=\7\6\2\2=L\5\16"+
		"\b\23>?\7\7\2\2?L\5\16\b\22@A\7\24\2\2AB\5\16\b\2BC\7\25\2\2CL\3\2\2\2"+
		"DL\7!\2\2EL\7 \2\2FL\7\"\2\2GL\7#\2\2HL\7$\2\2IL\t\3\2\2JL\7\30\2\2K;"+
		"\3\2\2\2K>\3\2\2\2K@\3\2\2\2KD\3\2\2\2KE\3\2\2\2KF\3\2\2\2KG\3\2\2\2K"+
		"H\3\2\2\2KI\3\2\2\2KJ\3\2\2\2Ld\3\2\2\2MN\f\21\2\2NO\t\4\2\2Oc\5\16\b"+
		"\22PQ\f\20\2\2QR\t\5\2\2Rc\5\16\b\21ST\f\17\2\2TU\t\6\2\2Uc\5\16\b\20"+
		"VW\f\16\2\2WX\t\7\2\2Xc\5\16\b\17YZ\f\r\2\2Z[\t\b\2\2[c\5\16\b\16\\]\f"+
		"\f\2\2]^\7\22\2\2^c\5\16\b\r_`\f\13\2\2`a\7\23\2\2ac\5\16\b\fbM\3\2\2"+
		"\2bP\3\2\2\2bS\3\2\2\2bV\3\2\2\2bY\3\2\2\2b\\\3\2\2\2b_\3\2\2\2cf\3\2"+
		"\2\2db\3\2\2\2de\3\2\2\2e\17\3\2\2\2fd\3\2\2\2\t\26\33.9Kbd";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
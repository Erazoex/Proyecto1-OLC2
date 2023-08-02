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
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, DOUBLE=29, INT=30, ID=31, STRING=32, 
		WHITESPACE=33;
	public static final int
		RULE_init = 0, RULE_block = 1, RULE_stmt = 2, RULE_declstmt = 3, RULE_printstmt = 4, 
		RULE_ifstmt = 5, RULE_forstmt = 6, RULE_expr = 7;
	private static String[] makeRuleNames() {
		return new String[] {
			"init", "block", "stmt", "declstmt", "printstmt", "ifstmt", "forstmt", 
			"expr"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'let'", "'='", "'var'", "'print'", "'('", "')'", "'if'", "'{'", 
			"'}'", "'for'", "'in'", "'!'", "'-'", "'*'", "'/'", "'%'", "'+'", "'>='", 
			"'>'", "'<='", "'<'", "'=='", "'!='", "'&&'", "'||'", "'true'", "'false'", 
			"'nil'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, "DOUBLE", "INT", "ID", "STRING", "WHITESPACE"
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
			setState(16);
			block();
			setState(17);
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
			setState(22);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__2) | (1L << T__3) | (1L << T__6) | (1L << T__9))) != 0)) {
				{
				{
				setState(19);
				stmt();
				}
				}
				setState(24);
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
		public PrintstmtContext printstmt() {
			return getRuleContext(PrintstmtContext.class,0);
		}
		public IfstmtContext ifstmt() {
			return getRuleContext(IfstmtContext.class,0);
		}
		public ForstmtContext forstmt() {
			return getRuleContext(ForstmtContext.class,0);
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
			setState(29);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
			case T__2:
				enterOuterAlt(_localctx, 1);
				{
				setState(25);
				declstmt();
				}
				break;
			case T__3:
				enterOuterAlt(_localctx, 2);
				{
				setState(26);
				printstmt();
				}
				break;
			case T__6:
				enterOuterAlt(_localctx, 3);
				{
				setState(27);
				ifstmt();
				}
				break;
			case T__9:
				enterOuterAlt(_localctx, 4);
				{
				setState(28);
				forstmt();
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
		public Token typeVar;
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DeclstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declstmt; }
	}

	public final DeclstmtContext declstmt() throws RecognitionException {
		DeclstmtContext _localctx = new DeclstmtContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_declstmt);
		try {
			setState(39);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				enterOuterAlt(_localctx, 1);
				{
				setState(31);
				((DeclstmtContext)_localctx).typeVar = match(T__0);
				setState(32);
				match(ID);
				setState(33);
				match(T__1);
				setState(34);
				expr(0);
				}
				break;
			case T__2:
				enterOuterAlt(_localctx, 2);
				{
				setState(35);
				((DeclstmtContext)_localctx).typeVar = match(T__2);
				setState(36);
				match(ID);
				setState(37);
				match(T__1);
				setState(38);
				expr(0);
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

	public static class PrintstmtContext extends ParserRuleContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public PrintstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printstmt; }
	}

	public final PrintstmtContext printstmt() throws RecognitionException {
		PrintstmtContext _localctx = new PrintstmtContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_printstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(41);
			match(T__3);
			setState(42);
			match(T__4);
			setState(43);
			expr(0);
			setState(44);
			match(T__5);
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

	public static class IfstmtContext extends ParserRuleContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public IfstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifstmt; }
	}

	public final IfstmtContext ifstmt() throws RecognitionException {
		IfstmtContext _localctx = new IfstmtContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_ifstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(46);
			match(T__6);
			setState(47);
			expr(0);
			setState(48);
			match(T__7);
			setState(49);
			block();
			setState(50);
			match(T__8);
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

	public static class ForstmtContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(GrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(GrammarParser.ID, i);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ForstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forstmt; }
	}

	public final ForstmtContext forstmt() throws RecognitionException {
		ForstmtContext _localctx = new ForstmtContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_forstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(52);
			match(T__9);
			setState(53);
			match(ID);
			setState(54);
			match(T__10);
			setState(55);
			match(ID);
			setState(56);
			match(T__7);
			setState(57);
			block();
			setState(58);
			match(T__8);
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

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(75);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__11:
				{
				_localctx = new NotExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(61);
				match(T__11);
				setState(62);
				((NotExprContext)_localctx).right = expr(16);
				}
				break;
			case T__12:
				{
				_localctx = new UnaryExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(63);
				match(T__12);
				setState(64);
				((UnaryExprContext)_localctx).right = expr(15);
				}
				break;
			case T__4:
				{
				_localctx = new ParExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(65);
				match(T__4);
				setState(66);
				expr(0);
				setState(67);
				match(T__5);
				}
				break;
			case INT:
				{
				_localctx = new IntExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(69);
				match(INT);
				}
				break;
			case DOUBLE:
				{
				_localctx = new DoubleExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(70);
				match(DOUBLE);
				}
				break;
			case ID:
				{
				_localctx = new IdExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(71);
				match(ID);
				}
				break;
			case STRING:
				{
				_localctx = new StrExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(72);
				match(STRING);
				}
				break;
			case T__25:
			case T__26:
				{
				_localctx = new BoolExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(73);
				_la = _input.LA(1);
				if ( !(_la==T__25 || _la==T__26) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case T__27:
				{
				_localctx = new NilExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(74);
				match(T__27);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(100);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(98);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
					case 1:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(77);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(78);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__13) | (1L << T__14) | (1L << T__15))) != 0)) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(79);
						((OpExprContext)_localctx).right = expr(15);
						}
						break;
					case 2:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(80);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(81);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__12 || _la==T__16) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(82);
						((OpExprContext)_localctx).right = expr(14);
						}
						break;
					case 3:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(83);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(84);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__17 || _la==T__18) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(85);
						((OpExprContext)_localctx).right = expr(13);
						}
						break;
					case 4:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(86);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(87);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__19 || _la==T__20) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(88);
						((OpExprContext)_localctx).right = expr(12);
						}
						break;
					case 5:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(89);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(90);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__21 || _la==T__22) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(91);
						((OpExprContext)_localctx).right = expr(11);
						}
						break;
					case 6:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(92);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(93);
						match(T__23);
						setState(94);
						((OpExprContext)_localctx).right = expr(10);
						}
						break;
					case 7:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(95);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(96);
						match(T__24);
						setState(97);
						((OpExprContext)_localctx).right = expr(9);
						}
						break;
					}
					} 
				}
				setState(102);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
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
		case 7:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 14);
		case 1:
			return precpred(_ctx, 13);
		case 2:
			return precpred(_ctx, 12);
		case 3:
			return precpred(_ctx, 11);
		case 4:
			return precpred(_ctx, 10);
		case 5:
			return precpred(_ctx, 9);
		case 6:
			return precpred(_ctx, 8);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3#j\4\2\t\2\4\3\t\3"+
		"\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\3\2\3\2\3\2\3\3\7\3\27"+
		"\n\3\f\3\16\3\32\13\3\3\4\3\4\3\4\3\4\5\4 \n\4\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\3\5\3\5\5\5*\n\5\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\5\tN\n\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\7\te\n\t\f\t\16\th\13\t\3\t\2\3"+
		"\20\n\2\4\6\b\n\f\16\20\2\b\3\2\34\35\3\2\20\22\4\2\17\17\23\23\3\2\24"+
		"\25\3\2\26\27\3\2\30\31\2u\2\22\3\2\2\2\4\30\3\2\2\2\6\37\3\2\2\2\b)\3"+
		"\2\2\2\n+\3\2\2\2\f\60\3\2\2\2\16\66\3\2\2\2\20M\3\2\2\2\22\23\5\4\3\2"+
		"\23\24\7\2\2\3\24\3\3\2\2\2\25\27\5\6\4\2\26\25\3\2\2\2\27\32\3\2\2\2"+
		"\30\26\3\2\2\2\30\31\3\2\2\2\31\5\3\2\2\2\32\30\3\2\2\2\33 \5\b\5\2\34"+
		" \5\n\6\2\35 \5\f\7\2\36 \5\16\b\2\37\33\3\2\2\2\37\34\3\2\2\2\37\35\3"+
		"\2\2\2\37\36\3\2\2\2 \7\3\2\2\2!\"\7\3\2\2\"#\7!\2\2#$\7\4\2\2$*\5\20"+
		"\t\2%&\7\5\2\2&\'\7!\2\2\'(\7\4\2\2(*\5\20\t\2)!\3\2\2\2)%\3\2\2\2*\t"+
		"\3\2\2\2+,\7\6\2\2,-\7\7\2\2-.\5\20\t\2./\7\b\2\2/\13\3\2\2\2\60\61\7"+
		"\t\2\2\61\62\5\20\t\2\62\63\7\n\2\2\63\64\5\4\3\2\64\65\7\13\2\2\65\r"+
		"\3\2\2\2\66\67\7\f\2\2\678\7!\2\289\7\r\2\29:\7!\2\2:;\7\n\2\2;<\5\4\3"+
		"\2<=\7\13\2\2=\17\3\2\2\2>?\b\t\1\2?@\7\16\2\2@N\5\20\t\22AB\7\17\2\2"+
		"BN\5\20\t\21CD\7\7\2\2DE\5\20\t\2EF\7\b\2\2FN\3\2\2\2GN\7 \2\2HN\7\37"+
		"\2\2IN\7!\2\2JN\7\"\2\2KN\t\2\2\2LN\7\36\2\2M>\3\2\2\2MA\3\2\2\2MC\3\2"+
		"\2\2MG\3\2\2\2MH\3\2\2\2MI\3\2\2\2MJ\3\2\2\2MK\3\2\2\2ML\3\2\2\2Nf\3\2"+
		"\2\2OP\f\20\2\2PQ\t\3\2\2Qe\5\20\t\21RS\f\17\2\2ST\t\4\2\2Te\5\20\t\20"+
		"UV\f\16\2\2VW\t\5\2\2We\5\20\t\17XY\f\r\2\2YZ\t\6\2\2Ze\5\20\t\16[\\\f"+
		"\f\2\2\\]\t\7\2\2]e\5\20\t\r^_\f\13\2\2_`\7\32\2\2`e\5\20\t\fab\f\n\2"+
		"\2bc\7\33\2\2ce\5\20\t\13dO\3\2\2\2dR\3\2\2\2dU\3\2\2\2dX\3\2\2\2d[\3"+
		"\2\2\2d^\3\2\2\2da\3\2\2\2eh\3\2\2\2fd\3\2\2\2fg\3\2\2\2g\21\3\2\2\2h"+
		"f\3\2\2\2\b\30\37)Mdf";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
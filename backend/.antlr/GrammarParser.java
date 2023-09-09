// Generated from c:\Users\brian\OneDrive\Escritorio\Repositories\proyecto\backend\Grammar.g4 by ANTLR 4.9.2
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
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		T_INT=32, T_FLOAT=33, T_CHAR=34, T_STRING=35, T_BOOL=36, VAR=37, LET=38, 
		IF=39, ELSE=40, SWITCH=41, CASE=42, DEFAULT=43, WHILE=44, FOR=45, VECTOR=46, 
		FUNC=47, GUARD=48, BREAK=49, RETURN=50, CONTINUE=51, DOUBLE=52, INT=53, 
		ID=54, CHAR=55, STRING=56, MULTCOMMENT=57, COMMENT=58, WHITESPACE=59;
	public static final int
		RULE_init = 0, RULE_block = 1, RULE_stmt = 2, RULE_breakstmt = 3, RULE_continuestmt = 4, 
		RULE_returnstmt = 5, RULE_declstmt = 6, RULE_asignstmt = 7, RULE_incstmt = 8, 
		RULE_decstmt = 9, RULE_ifstmt = 10, RULE_switchstmt = 11, RULE_switchcase = 12, 
		RULE_printlnstmt = 13, RULE_intstmt = 14, RULE_floatstmt = 15, RULE_stringstmt = 16, 
		RULE_exprparams = 17, RULE_whilestmt = 18, RULE_forstmt = 19, RULE_forrange = 20, 
		RULE_guardstmt = 21, RULE_array = 22, RULE_array_def = 23, RULE_funcstmt = 24, 
		RULE_listaparametros = 25, RULE_parametro = 26, RULE_callstmt = 27, RULE_vartype = 28, 
		RULE_expr = 29;
	private static String[] makeRuleNames() {
		return new String[] {
			"init", "block", "stmt", "breakstmt", "continuestmt", "returnstmt", "declstmt", 
			"asignstmt", "incstmt", "decstmt", "ifstmt", "switchstmt", "switchcase", 
			"printlnstmt", "intstmt", "floatstmt", "stringstmt", "exprparams", "whilestmt", 
			"forstmt", "forrange", "guardstmt", "array", "array_def", "funcstmt", 
			"listaparametros", "parametro", "callstmt", "vartype", "expr"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "':'", "'='", "'?'", "'+='", "'-='", "'{'", "'}'", "'print'", "'('", 
			"')'", "','", "'in'", "'...'", "'<'", "'>'", "'->'", "'!'", "'-'", "'*'", 
			"'/'", "'%'", "'+'", "'>='", "'<='", "'=='", "'!='", "'&&'", "'||'", 
			"'true'", "'false'", "'nil'", "'Int'", "'Float'", "'Character'", "'String'", 
			"'Bool'", "'var'", "'let'", "'if'", "'else'", "'switch'", "'case'", "'default'", 
			"'while'", "'for'", "'vector'", "'func'", "'guard'", "'break'", "'return'", 
			"'continue'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, "T_INT", "T_FLOAT", "T_CHAR", 
			"T_STRING", "T_BOOL", "VAR", "LET", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", 
			"WHILE", "FOR", "VECTOR", "FUNC", "GUARD", "BREAK", "RETURN", "CONTINUE", 
			"DOUBLE", "INT", "ID", "CHAR", "STRING", "MULTCOMMENT", "COMMENT", "WHITESPACE"
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
			setState(60);
			block();
			setState(61);
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
			setState(66);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__7) | (1L << VAR) | (1L << LET) | (1L << IF) | (1L << SWITCH) | (1L << WHILE) | (1L << FOR) | (1L << FUNC) | (1L << GUARD) | (1L << BREAK) | (1L << RETURN) | (1L << CONTINUE) | (1L << ID))) != 0)) {
				{
				{
				setState(63);
				stmt();
				}
				}
				setState(68);
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
		public IncstmtContext incstmt() {
			return getRuleContext(IncstmtContext.class,0);
		}
		public DecstmtContext decstmt() {
			return getRuleContext(DecstmtContext.class,0);
		}
		public IfstmtContext ifstmt() {
			return getRuleContext(IfstmtContext.class,0);
		}
		public SwitchstmtContext switchstmt() {
			return getRuleContext(SwitchstmtContext.class,0);
		}
		public PrintlnstmtContext printlnstmt() {
			return getRuleContext(PrintlnstmtContext.class,0);
		}
		public WhilestmtContext whilestmt() {
			return getRuleContext(WhilestmtContext.class,0);
		}
		public ForstmtContext forstmt() {
			return getRuleContext(ForstmtContext.class,0);
		}
		public GuardstmtContext guardstmt() {
			return getRuleContext(GuardstmtContext.class,0);
		}
		public BreakstmtContext breakstmt() {
			return getRuleContext(BreakstmtContext.class,0);
		}
		public ContinuestmtContext continuestmt() {
			return getRuleContext(ContinuestmtContext.class,0);
		}
		public ReturnstmtContext returnstmt() {
			return getRuleContext(ReturnstmtContext.class,0);
		}
		public FuncstmtContext funcstmt() {
			return getRuleContext(FuncstmtContext.class,0);
		}
		public CallstmtContext callstmt() {
			return getRuleContext(CallstmtContext.class,0);
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
			setState(84);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(69);
				declstmt();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(70);
				asignstmt();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(71);
				incstmt();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(72);
				decstmt();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(73);
				ifstmt();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(74);
				switchstmt();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(75);
				printlnstmt();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(76);
				whilestmt();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(77);
				forstmt();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(78);
				guardstmt();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(79);
				breakstmt();
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(80);
				continuestmt();
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(81);
				returnstmt();
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(82);
				funcstmt();
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(83);
				callstmt();
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

	public static class BreakstmtContext extends ParserRuleContext {
		public TerminalNode BREAK() { return getToken(GrammarParser.BREAK, 0); }
		public BreakstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_breakstmt; }
	}

	public final BreakstmtContext breakstmt() throws RecognitionException {
		BreakstmtContext _localctx = new BreakstmtContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_breakstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(86);
			match(BREAK);
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

	public static class ContinuestmtContext extends ParserRuleContext {
		public TerminalNode CONTINUE() { return getToken(GrammarParser.CONTINUE, 0); }
		public ContinuestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continuestmt; }
	}

	public final ContinuestmtContext continuestmt() throws RecognitionException {
		ContinuestmtContext _localctx = new ContinuestmtContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_continuestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(88);
			match(CONTINUE);
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

	public static class ReturnstmtContext extends ParserRuleContext {
		public TerminalNode RETURN() { return getToken(GrammarParser.RETURN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ReturnstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnstmt; }
	}

	public final ReturnstmtContext returnstmt() throws RecognitionException {
		ReturnstmtContext _localctx = new ReturnstmtContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_returnstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(90);
			match(RETURN);
			setState(92);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				{
				setState(91);
				expr(0);
				}
				break;
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
		enterRule(_localctx, 12, RULE_declstmt);
		int _la;
		try {
			setState(111);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				_localctx = new DeclstmtWithTypeAndExprContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(94);
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
				setState(95);
				match(ID);
				setState(96);
				match(T__0);
				setState(97);
				vartype();
				setState(98);
				match(T__1);
				setState(99);
				expr(0);
				}
				break;
			case 2:
				_localctx = new DeclstmtWithExprContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(101);
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
				setState(102);
				match(ID);
				setState(103);
				match(T__1);
				setState(104);
				expr(0);
				}
				break;
			case 3:
				_localctx = new DeclstmtWithTypeContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(105);
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
				setState(106);
				match(ID);
				setState(107);
				match(T__0);
				setState(108);
				vartype();
				setState(109);
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
		enterRule(_localctx, 14, RULE_asignstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(113);
			match(ID);
			setState(114);
			match(T__1);
			setState(115);
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

	public static class IncstmtContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public IncstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_incstmt; }
	}

	public final IncstmtContext incstmt() throws RecognitionException {
		IncstmtContext _localctx = new IncstmtContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_incstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(117);
			match(ID);
			setState(118);
			match(T__3);
			setState(119);
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

	public static class DecstmtContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DecstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decstmt; }
	}

	public final DecstmtContext decstmt() throws RecognitionException {
		DecstmtContext _localctx = new DecstmtContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_decstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(121);
			match(ID);
			setState(122);
			match(T__4);
			setState(123);
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

	public static class IfstmtContext extends ParserRuleContext {
		public IfstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifstmt; }
	 
		public IfstmtContext() { }
		public void copyFrom(IfstmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class IfWithElseIfContext extends IfstmtContext {
		public TerminalNode IF() { return getToken(GrammarParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(GrammarParser.ELSE, 0); }
		public IfstmtContext ifstmt() {
			return getRuleContext(IfstmtContext.class,0);
		}
		public IfWithElseIfContext(IfstmtContext ctx) { copyFrom(ctx); }
	}
	public static class IfWithElseContext extends IfstmtContext {
		public BlockContext trueCondition;
		public BlockContext falseCondition;
		public TerminalNode IF() { return getToken(GrammarParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(GrammarParser.ELSE, 0); }
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public IfWithElseContext(IfstmtContext ctx) { copyFrom(ctx); }
	}
	public static class IfSimpleContext extends IfstmtContext {
		public TerminalNode IF() { return getToken(GrammarParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public IfSimpleContext(IfstmtContext ctx) { copyFrom(ctx); }
	}

	public final IfstmtContext ifstmt() throws RecognitionException {
		IfstmtContext _localctx = new IfstmtContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_ifstmt);
		try {
			setState(149);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				_localctx = new IfSimpleContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(125);
				match(IF);
				setState(126);
				expr(0);
				setState(127);
				match(T__5);
				setState(128);
				block();
				setState(129);
				match(T__6);
				}
				break;
			case 2:
				_localctx = new IfWithElseContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(131);
				match(IF);
				setState(132);
				expr(0);
				setState(133);
				match(T__5);
				setState(134);
				((IfWithElseContext)_localctx).trueCondition = block();
				setState(135);
				match(T__6);
				setState(136);
				match(ELSE);
				setState(137);
				match(T__5);
				setState(138);
				((IfWithElseContext)_localctx).falseCondition = block();
				setState(139);
				match(T__6);
				}
				break;
			case 3:
				_localctx = new IfWithElseIfContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(141);
				match(IF);
				setState(142);
				expr(0);
				setState(143);
				match(T__5);
				setState(144);
				block();
				setState(145);
				match(T__6);
				setState(146);
				match(ELSE);
				setState(147);
				ifstmt();
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

	public static class SwitchstmtContext extends ParserRuleContext {
		public TerminalNode SWITCH() { return getToken(GrammarParser.SWITCH, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<SwitchcaseContext> switchcase() {
			return getRuleContexts(SwitchcaseContext.class);
		}
		public SwitchcaseContext switchcase(int i) {
			return getRuleContext(SwitchcaseContext.class,i);
		}
		public SwitchstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchstmt; }
	}

	public final SwitchstmtContext switchstmt() throws RecognitionException {
		SwitchstmtContext _localctx = new SwitchstmtContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_switchstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(151);
			match(SWITCH);
			setState(152);
			expr(0);
			setState(153);
			match(T__5);
			setState(155); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(154);
				switchcase();
				}
				}
				setState(157); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE || _la==DEFAULT );
			setState(159);
			match(T__6);
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

	public static class SwitchcaseContext extends ParserRuleContext {
		public Token casetype;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode CASE() { return getToken(GrammarParser.CASE, 0); }
		public TerminalNode DEFAULT() { return getToken(GrammarParser.DEFAULT, 0); }
		public SwitchcaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchcase; }
	}

	public final SwitchcaseContext switchcase() throws RecognitionException {
		SwitchcaseContext _localctx = new SwitchcaseContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_switchcase);
		try {
			setState(169);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CASE:
				enterOuterAlt(_localctx, 1);
				{
				setState(161);
				((SwitchcaseContext)_localctx).casetype = match(CASE);
				setState(162);
				expr(0);
				setState(163);
				match(T__0);
				setState(164);
				block();
				}
				break;
			case DEFAULT:
				enterOuterAlt(_localctx, 2);
				{
				setState(166);
				((SwitchcaseContext)_localctx).casetype = match(DEFAULT);
				setState(167);
				match(T__0);
				setState(168);
				block();
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

	public static class PrintlnstmtContext extends ParserRuleContext {
		public ExprparamsContext exprparams() {
			return getRuleContext(ExprparamsContext.class,0);
		}
		public PrintlnstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printlnstmt; }
	}

	public final PrintlnstmtContext printlnstmt() throws RecognitionException {
		PrintlnstmtContext _localctx = new PrintlnstmtContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_printlnstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(171);
			match(T__7);
			setState(172);
			match(T__8);
			setState(174);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__16) | (1L << T__17) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T_INT) | (1L << T_FLOAT) | (1L << T_STRING) | (1L << DOUBLE) | (1L << INT) | (1L << ID) | (1L << CHAR) | (1L << STRING))) != 0)) {
				{
				setState(173);
				exprparams();
				}
			}

			setState(176);
			match(T__9);
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

	public static class IntstmtContext extends ParserRuleContext {
		public TerminalNode T_INT() { return getToken(GrammarParser.T_INT, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public IntstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_intstmt; }
	}

	public final IntstmtContext intstmt() throws RecognitionException {
		IntstmtContext _localctx = new IntstmtContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_intstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(178);
			match(T_INT);
			setState(179);
			match(T__8);
			setState(180);
			expr(0);
			setState(181);
			match(T__9);
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

	public static class FloatstmtContext extends ParserRuleContext {
		public TerminalNode T_FLOAT() { return getToken(GrammarParser.T_FLOAT, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public FloatstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_floatstmt; }
	}

	public final FloatstmtContext floatstmt() throws RecognitionException {
		FloatstmtContext _localctx = new FloatstmtContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_floatstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(183);
			match(T_FLOAT);
			setState(184);
			match(T__8);
			setState(185);
			expr(0);
			setState(186);
			match(T__9);
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

	public static class StringstmtContext extends ParserRuleContext {
		public TerminalNode T_STRING() { return getToken(GrammarParser.T_STRING, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public StringstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stringstmt; }
	}

	public final StringstmtContext stringstmt() throws RecognitionException {
		StringstmtContext _localctx = new StringstmtContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_stringstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(188);
			match(T_STRING);
			setState(189);
			match(T__8);
			setState(190);
			expr(0);
			setState(191);
			match(T__9);
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

	public static class ExprparamsContext extends ParserRuleContext {
		public ExprparamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exprparams; }
	 
		public ExprparamsContext() { }
		public void copyFrom(ExprparamsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class ExprWithParamsContext extends ExprparamsContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ExprparamsContext exprparams() {
			return getRuleContext(ExprparamsContext.class,0);
		}
		public ExprWithParamsContext(ExprparamsContext ctx) { copyFrom(ctx); }
	}
	public static class ExprWithParamContext extends ExprparamsContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ExprWithParamContext(ExprparamsContext ctx) { copyFrom(ctx); }
	}

	public final ExprparamsContext exprparams() throws RecognitionException {
		ExprparamsContext _localctx = new ExprparamsContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_exprparams);
		try {
			setState(198);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				_localctx = new ExprWithParamsContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(193);
				expr(0);
				setState(194);
				match(T__10);
				setState(195);
				exprparams();
				}
				break;
			case 2:
				_localctx = new ExprWithParamContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(197);
				expr(0);
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

	public static class WhilestmtContext extends ParserRuleContext {
		public TerminalNode WHILE() { return getToken(GrammarParser.WHILE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public WhilestmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whilestmt; }
	}

	public final WhilestmtContext whilestmt() throws RecognitionException {
		WhilestmtContext _localctx = new WhilestmtContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_whilestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(200);
			match(WHILE);
			setState(201);
			expr(0);
			setState(202);
			match(T__5);
			setState(203);
			block();
			setState(204);
			match(T__6);
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
		public ForstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forstmt; }
	 
		public ForstmtContext() { }
		public void copyFrom(ForstmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class ForWithRangeContext extends ForstmtContext {
		public TerminalNode FOR() { return getToken(GrammarParser.FOR, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public ForrangeContext forrange() {
			return getRuleContext(ForrangeContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ForWithRangeContext(ForstmtContext ctx) { copyFrom(ctx); }
	}
	public static class ForWithExprContext extends ForstmtContext {
		public TerminalNode FOR() { return getToken(GrammarParser.FOR, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ForWithExprContext(ForstmtContext ctx) { copyFrom(ctx); }
	}

	public final ForstmtContext forstmt() throws RecognitionException {
		ForstmtContext _localctx = new ForstmtContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_forstmt);
		try {
			setState(222);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				_localctx = new ForWithExprContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(206);
				match(FOR);
				setState(207);
				match(ID);
				setState(208);
				match(T__11);
				setState(209);
				expr(0);
				setState(210);
				match(T__5);
				setState(211);
				block();
				setState(212);
				match(T__6);
				}
				break;
			case 2:
				_localctx = new ForWithRangeContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(214);
				match(FOR);
				setState(215);
				match(ID);
				setState(216);
				match(T__11);
				setState(217);
				forrange();
				setState(218);
				match(T__5);
				setState(219);
				block();
				setState(220);
				match(T__6);
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

	public static class ForrangeContext extends ParserRuleContext {
		public ExprContext beginsWith;
		public ExprContext endsWith;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ForrangeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forrange; }
	}

	public final ForrangeContext forrange() throws RecognitionException {
		ForrangeContext _localctx = new ForrangeContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_forrange);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(224);
			((ForrangeContext)_localctx).beginsWith = expr(0);
			setState(225);
			match(T__12);
			setState(226);
			((ForrangeContext)_localctx).endsWith = expr(0);
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

	public static class GuardstmtContext extends ParserRuleContext {
		public TerminalNode GUARD() { return getToken(GrammarParser.GUARD, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(GrammarParser.ELSE, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public GuardstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_guardstmt; }
	}

	public final GuardstmtContext guardstmt() throws RecognitionException {
		GuardstmtContext _localctx = new GuardstmtContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_guardstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(228);
			match(GUARD);
			setState(229);
			expr(0);
			setState(230);
			match(ELSE);
			setState(231);
			match(T__5);
			setState(232);
			block();
			setState(233);
			match(T__6);
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

	public static class ArrayContext extends ParserRuleContext {
		public TerminalNode VECTOR() { return getToken(GrammarParser.VECTOR, 0); }
		public VartypeContext vartype() {
			return getRuleContext(VartypeContext.class,0);
		}
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public Array_defContext array_def() {
			return getRuleContext(Array_defContext.class,0);
		}
		public ArrayContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_array; }
	}

	public final ArrayContext array() throws RecognitionException {
		ArrayContext _localctx = new ArrayContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_array);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(235);
			match(VECTOR);
			setState(236);
			match(T__13);
			setState(237);
			vartype();
			setState(238);
			match(T__14);
			setState(239);
			match(ID);
			setState(240);
			array_def();
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

	public static class Array_defContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public Array_defContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_array_def; }
	}

	public final Array_defContext array_def() throws RecognitionException {
		Array_defContext _localctx = new Array_defContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_array_def);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(242);
			match(T__1);
			setState(246);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__16) | (1L << T__17) | (1L << T__28) | (1L << T__29) | (1L << T__30) | (1L << T_INT) | (1L << T_FLOAT) | (1L << T_STRING) | (1L << DOUBLE) | (1L << INT) | (1L << ID) | (1L << CHAR) | (1L << STRING))) != 0)) {
				{
				{
				setState(243);
				expr(0);
				}
				}
				setState(248);
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

	public static class FuncstmtContext extends ParserRuleContext {
		public TerminalNode FUNC() { return getToken(GrammarParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ListaparametrosContext listaparametros() {
			return getRuleContext(ListaparametrosContext.class,0);
		}
		public VartypeContext vartype() {
			return getRuleContext(VartypeContext.class,0);
		}
		public FuncstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcstmt; }
	}

	public final FuncstmtContext funcstmt() throws RecognitionException {
		FuncstmtContext _localctx = new FuncstmtContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_funcstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(249);
			match(FUNC);
			setState(250);
			match(ID);
			setState(251);
			match(T__8);
			setState(253);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(252);
				listaparametros();
				}
			}

			setState(255);
			match(T__9);
			setState(258);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__15) {
				{
				setState(256);
				match(T__15);
				setState(257);
				vartype();
				}
			}

			setState(260);
			match(T__5);
			setState(261);
			block();
			setState(262);
			match(T__6);
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

	public static class ListaparametrosContext extends ParserRuleContext {
		public ListaparametrosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaparametros; }
	 
		public ListaparametrosContext() { }
		public void copyFrom(ListaparametrosContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class FuncParameterContext extends ListaparametrosContext {
		public ParametroContext parametro() {
			return getRuleContext(ParametroContext.class,0);
		}
		public FuncParameterContext(ListaparametrosContext ctx) { copyFrom(ctx); }
	}
	public static class FuncParametersContext extends ListaparametrosContext {
		public ParametroContext parametro() {
			return getRuleContext(ParametroContext.class,0);
		}
		public ListaparametrosContext listaparametros() {
			return getRuleContext(ListaparametrosContext.class,0);
		}
		public FuncParametersContext(ListaparametrosContext ctx) { copyFrom(ctx); }
	}

	public final ListaparametrosContext listaparametros() throws RecognitionException {
		ListaparametrosContext _localctx = new ListaparametrosContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_listaparametros);
		try {
			setState(269);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				_localctx = new FuncParametersContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(264);
				parametro();
				setState(265);
				match(T__10);
				setState(266);
				listaparametros();
				}
				break;
			case 2:
				_localctx = new FuncParameterContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(268);
				parametro();
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

	public static class ParametroContext extends ParserRuleContext {
		public Token extVarId;
		public Token varId;
		public VartypeContext vartype() {
			return getRuleContext(VartypeContext.class,0);
		}
		public List<TerminalNode> ID() { return getTokens(GrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(GrammarParser.ID, i);
		}
		public ParametroContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametro; }
	}

	public final ParametroContext parametro() throws RecognitionException {
		ParametroContext _localctx = new ParametroContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_parametro);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(272);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				{
				setState(271);
				((ParametroContext)_localctx).extVarId = match(ID);
				}
				break;
			}
			setState(274);
			((ParametroContext)_localctx).varId = match(ID);
			setState(275);
			match(T__0);
			setState(276);
			vartype();
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

	public static class CallstmtContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public ExprparamsContext exprparams() {
			return getRuleContext(ExprparamsContext.class,0);
		}
		public CallstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_callstmt; }
	}

	public final CallstmtContext callstmt() throws RecognitionException {
		CallstmtContext _localctx = new CallstmtContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_callstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(278);
			match(ID);
			setState(279);
			match(T__8);
			setState(280);
			exprparams();
			setState(281);
			match(T__9);
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
		enterRule(_localctx, 56, RULE_vartype);
		try {
			setState(288);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T_INT:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(283);
				((PrimitiveTypeContext)_localctx).tipo = match(T_INT);
				}
				break;
			case T_FLOAT:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(284);
				((PrimitiveTypeContext)_localctx).tipo = match(T_FLOAT);
				}
				break;
			case T_CHAR:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(285);
				((PrimitiveTypeContext)_localctx).tipo = match(T_CHAR);
				}
				break;
			case T_STRING:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(286);
				((PrimitiveTypeContext)_localctx).tipo = match(T_STRING);
				}
				break;
			case T_BOOL:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(287);
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
	public static class BoolExprContext extends ExprContext {
		public BoolExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class IdExprContext extends ExprContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public IdExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class StringConvExprContext extends ExprContext {
		public StringstmtContext stringstmt() {
			return getRuleContext(StringstmtContext.class,0);
		}
		public StringConvExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	public static class IntConvExprContext extends ExprContext {
		public IntstmtContext intstmt() {
			return getRuleContext(IntstmtContext.class,0);
		}
		public IntConvExprContext(ExprContext ctx) { copyFrom(ctx); }
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
	public static class DoubleExprContext extends ExprContext {
		public TerminalNode DOUBLE() { return getToken(GrammarParser.DOUBLE, 0); }
		public DoubleExprContext(ExprContext ctx) { copyFrom(ctx); }
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
	public static class FloatConvExprContext extends ExprContext {
		public FloatstmtContext floatstmt() {
			return getRuleContext(FloatstmtContext.class,0);
		}
		public FloatConvExprContext(ExprContext ctx) { copyFrom(ctx); }
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
	public static class CallExprContext extends ExprContext {
		public CallstmtContext callstmt() {
			return getRuleContext(CallstmtContext.class,0);
		}
		public CallExprContext(ExprContext ctx) { copyFrom(ctx); }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 58;
		enterRecursionRule(_localctx, 58, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(310);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				{
				_localctx = new NotExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(291);
				match(T__16);
				setState(292);
				((NotExprContext)_localctx).right = expr(21);
				}
				break;
			case 2:
				{
				_localctx = new UnaryExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(293);
				match(T__17);
				setState(294);
				((UnaryExprContext)_localctx).right = expr(20);
				}
				break;
			case 3:
				{
				_localctx = new ParExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(295);
				match(T__8);
				setState(296);
				expr(0);
				setState(297);
				match(T__9);
				}
				break;
			case 4:
				{
				_localctx = new IntExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(299);
				match(INT);
				}
				break;
			case 5:
				{
				_localctx = new DoubleExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(300);
				match(DOUBLE);
				}
				break;
			case 6:
				{
				_localctx = new IdExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(301);
				match(ID);
				}
				break;
			case 7:
				{
				_localctx = new CharExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(302);
				match(CHAR);
				}
				break;
			case 8:
				{
				_localctx = new StrExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(303);
				match(STRING);
				}
				break;
			case 9:
				{
				_localctx = new BoolExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(304);
				_la = _input.LA(1);
				if ( !(_la==T__28 || _la==T__29) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case 10:
				{
				_localctx = new NilExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(305);
				match(T__30);
				}
				break;
			case 11:
				{
				_localctx = new IntConvExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(306);
				intstmt();
				}
				break;
			case 12:
				{
				_localctx = new FloatConvExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(307);
				floatstmt();
				}
				break;
			case 13:
				{
				_localctx = new StringConvExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(308);
				stringstmt();
				}
				break;
			case 14:
				{
				_localctx = new CallExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(309);
				callstmt();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(335);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(333);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
					case 1:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(312);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(313);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__18) | (1L << T__19) | (1L << T__20))) != 0)) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(314);
						((OpExprContext)_localctx).right = expr(20);
						}
						break;
					case 2:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(315);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(316);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__17 || _la==T__21) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(317);
						((OpExprContext)_localctx).right = expr(19);
						}
						break;
					case 3:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(318);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(319);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__14 || _la==T__22) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(320);
						((OpExprContext)_localctx).right = expr(18);
						}
						break;
					case 4:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(321);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(322);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__13 || _la==T__23) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(323);
						((OpExprContext)_localctx).right = expr(17);
						}
						break;
					case 5:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(324);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(325);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__24 || _la==T__25) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(326);
						((OpExprContext)_localctx).right = expr(16);
						}
						break;
					case 6:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(327);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(328);
						((OpExprContext)_localctx).op = match(T__26);
						setState(329);
						((OpExprContext)_localctx).right = expr(15);
						}
						break;
					case 7:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(330);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(331);
						((OpExprContext)_localctx).op = match(T__27);
						setState(332);
						((OpExprContext)_localctx).right = expr(14);
						}
						break;
					}
					} 
				}
				setState(337);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
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
		case 29:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 19);
		case 1:
			return precpred(_ctx, 18);
		case 2:
			return precpred(_ctx, 17);
		case 3:
			return precpred(_ctx, 16);
		case 4:
			return precpred(_ctx, 15);
		case 5:
			return precpred(_ctx, 14);
		case 6:
			return precpred(_ctx, 13);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3=\u0155\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\3\2\3\2\3"+
		"\2\3\3\7\3C\n\3\f\3\16\3F\13\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\5\4W\n\4\3\5\3\5\3\6\3\6\3\7\3\7\5\7_\n\7\3\b\3\b"+
		"\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\5\br\n\b"+
		"\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\5\f\u0098\n\f\3\r\3\r\3\r\3\r\6\r\u009e\n\r\r\r\16\r\u009f\3"+
		"\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\5\16\u00ac\n\16\3\17\3"+
		"\17\3\17\5\17\u00b1\n\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\21\3\21"+
		"\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\5\23"+
		"\u00c9\n\23\3\24\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25"+
		"\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\5\25\u00e1\n\25\3\26"+
		"\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30"+
		"\3\30\3\30\3\30\3\31\3\31\7\31\u00f7\n\31\f\31\16\31\u00fa\13\31\3\32"+
		"\3\32\3\32\3\32\5\32\u0100\n\32\3\32\3\32\3\32\5\32\u0105\n\32\3\32\3"+
		"\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\5\33\u0110\n\33\3\34\5\34\u0113"+
		"\n\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36"+
		"\3\36\5\36\u0123\n\36\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37"+
		"\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\5\37\u0139\n\37\3\37"+
		"\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37"+
		"\3\37\3\37\3\37\3\37\3\37\3\37\7\37\u0150\n\37\f\37\16\37\u0153\13\37"+
		"\3\37\2\3< \2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\66"+
		"8:<\2\t\3\2\'(\3\2\37 \3\2\25\27\4\2\24\24\30\30\4\2\21\21\31\31\4\2\20"+
		"\20\32\32\3\2\33\34\2\u016c\2>\3\2\2\2\4D\3\2\2\2\6V\3\2\2\2\bX\3\2\2"+
		"\2\nZ\3\2\2\2\f\\\3\2\2\2\16q\3\2\2\2\20s\3\2\2\2\22w\3\2\2\2\24{\3\2"+
		"\2\2\26\u0097\3\2\2\2\30\u0099\3\2\2\2\32\u00ab\3\2\2\2\34\u00ad\3\2\2"+
		"\2\36\u00b4\3\2\2\2 \u00b9\3\2\2\2\"\u00be\3\2\2\2$\u00c8\3\2\2\2&\u00ca"+
		"\3\2\2\2(\u00e0\3\2\2\2*\u00e2\3\2\2\2,\u00e6\3\2\2\2.\u00ed\3\2\2\2\60"+
		"\u00f4\3\2\2\2\62\u00fb\3\2\2\2\64\u010f\3\2\2\2\66\u0112\3\2\2\28\u0118"+
		"\3\2\2\2:\u0122\3\2\2\2<\u0138\3\2\2\2>?\5\4\3\2?@\7\2\2\3@\3\3\2\2\2"+
		"AC\5\6\4\2BA\3\2\2\2CF\3\2\2\2DB\3\2\2\2DE\3\2\2\2E\5\3\2\2\2FD\3\2\2"+
		"\2GW\5\16\b\2HW\5\20\t\2IW\5\22\n\2JW\5\24\13\2KW\5\26\f\2LW\5\30\r\2"+
		"MW\5\34\17\2NW\5&\24\2OW\5(\25\2PW\5,\27\2QW\5\b\5\2RW\5\n\6\2SW\5\f\7"+
		"\2TW\5\62\32\2UW\58\35\2VG\3\2\2\2VH\3\2\2\2VI\3\2\2\2VJ\3\2\2\2VK\3\2"+
		"\2\2VL\3\2\2\2VM\3\2\2\2VN\3\2\2\2VO\3\2\2\2VP\3\2\2\2VQ\3\2\2\2VR\3\2"+
		"\2\2VS\3\2\2\2VT\3\2\2\2VU\3\2\2\2W\7\3\2\2\2XY\7\63\2\2Y\t\3\2\2\2Z["+
		"\7\65\2\2[\13\3\2\2\2\\^\7\64\2\2]_\5<\37\2^]\3\2\2\2^_\3\2\2\2_\r\3\2"+
		"\2\2`a\t\2\2\2ab\78\2\2bc\7\3\2\2cd\5:\36\2de\7\4\2\2ef\5<\37\2fr\3\2"+
		"\2\2gh\t\2\2\2hi\78\2\2ij\7\4\2\2jr\5<\37\2kl\t\2\2\2lm\78\2\2mn\7\3\2"+
		"\2no\5:\36\2op\7\5\2\2pr\3\2\2\2q`\3\2\2\2qg\3\2\2\2qk\3\2\2\2r\17\3\2"+
		"\2\2st\78\2\2tu\7\4\2\2uv\5<\37\2v\21\3\2\2\2wx\78\2\2xy\7\6\2\2yz\5<"+
		"\37\2z\23\3\2\2\2{|\78\2\2|}\7\7\2\2}~\5<\37\2~\25\3\2\2\2\177\u0080\7"+
		")\2\2\u0080\u0081\5<\37\2\u0081\u0082\7\b\2\2\u0082\u0083\5\4\3\2\u0083"+
		"\u0084\7\t\2\2\u0084\u0098\3\2\2\2\u0085\u0086\7)\2\2\u0086\u0087\5<\37"+
		"\2\u0087\u0088\7\b\2\2\u0088\u0089\5\4\3\2\u0089\u008a\7\t\2\2\u008a\u008b"+
		"\7*\2\2\u008b\u008c\7\b\2\2\u008c\u008d\5\4\3\2\u008d\u008e\7\t\2\2\u008e"+
		"\u0098\3\2\2\2\u008f\u0090\7)\2\2\u0090\u0091\5<\37\2\u0091\u0092\7\b"+
		"\2\2\u0092\u0093\5\4\3\2\u0093\u0094\7\t\2\2\u0094\u0095\7*\2\2\u0095"+
		"\u0096\5\26\f\2\u0096\u0098\3\2\2\2\u0097\177\3\2\2\2\u0097\u0085\3\2"+
		"\2\2\u0097\u008f\3\2\2\2\u0098\27\3\2\2\2\u0099\u009a\7+\2\2\u009a\u009b"+
		"\5<\37\2\u009b\u009d\7\b\2\2\u009c\u009e\5\32\16\2\u009d\u009c\3\2\2\2"+
		"\u009e\u009f\3\2\2\2\u009f\u009d\3\2\2\2\u009f\u00a0\3\2\2\2\u00a0\u00a1"+
		"\3\2\2\2\u00a1\u00a2\7\t\2\2\u00a2\31\3\2\2\2\u00a3\u00a4\7,\2\2\u00a4"+
		"\u00a5\5<\37\2\u00a5\u00a6\7\3\2\2\u00a6\u00a7\5\4\3\2\u00a7\u00ac\3\2"+
		"\2\2\u00a8\u00a9\7-\2\2\u00a9\u00aa\7\3\2\2\u00aa\u00ac\5\4\3\2\u00ab"+
		"\u00a3\3\2\2\2\u00ab\u00a8\3\2\2\2\u00ac\33\3\2\2\2\u00ad\u00ae\7\n\2"+
		"\2\u00ae\u00b0\7\13\2\2\u00af\u00b1\5$\23\2\u00b0\u00af\3\2\2\2\u00b0"+
		"\u00b1\3\2\2\2\u00b1\u00b2\3\2\2\2\u00b2\u00b3\7\f\2\2\u00b3\35\3\2\2"+
		"\2\u00b4\u00b5\7\"\2\2\u00b5\u00b6\7\13\2\2\u00b6\u00b7\5<\37\2\u00b7"+
		"\u00b8\7\f\2\2\u00b8\37\3\2\2\2\u00b9\u00ba\7#\2\2\u00ba\u00bb\7\13\2"+
		"\2\u00bb\u00bc\5<\37\2\u00bc\u00bd\7\f\2\2\u00bd!\3\2\2\2\u00be\u00bf"+
		"\7%\2\2\u00bf\u00c0\7\13\2\2\u00c0\u00c1\5<\37\2\u00c1\u00c2\7\f\2\2\u00c2"+
		"#\3\2\2\2\u00c3\u00c4\5<\37\2\u00c4\u00c5\7\r\2\2\u00c5\u00c6\5$\23\2"+
		"\u00c6\u00c9\3\2\2\2\u00c7\u00c9\5<\37\2\u00c8\u00c3\3\2\2\2\u00c8\u00c7"+
		"\3\2\2\2\u00c9%\3\2\2\2\u00ca\u00cb\7.\2\2\u00cb\u00cc\5<\37\2\u00cc\u00cd"+
		"\7\b\2\2\u00cd\u00ce\5\4\3\2\u00ce\u00cf\7\t\2\2\u00cf\'\3\2\2\2\u00d0"+
		"\u00d1\7/\2\2\u00d1\u00d2\78\2\2\u00d2\u00d3\7\16\2\2\u00d3\u00d4\5<\37"+
		"\2\u00d4\u00d5\7\b\2\2\u00d5\u00d6\5\4\3\2\u00d6\u00d7\7\t\2\2\u00d7\u00e1"+
		"\3\2\2\2\u00d8\u00d9\7/\2\2\u00d9\u00da\78\2\2\u00da\u00db\7\16\2\2\u00db"+
		"\u00dc\5*\26\2\u00dc\u00dd\7\b\2\2\u00dd\u00de\5\4\3\2\u00de\u00df\7\t"+
		"\2\2\u00df\u00e1\3\2\2\2\u00e0\u00d0\3\2\2\2\u00e0\u00d8\3\2\2\2\u00e1"+
		")\3\2\2\2\u00e2\u00e3\5<\37\2\u00e3\u00e4\7\17\2\2\u00e4\u00e5\5<\37\2"+
		"\u00e5+\3\2\2\2\u00e6\u00e7\7\62\2\2\u00e7\u00e8\5<\37\2\u00e8\u00e9\7"+
		"*\2\2\u00e9\u00ea\7\b\2\2\u00ea\u00eb\5\4\3\2\u00eb\u00ec\7\t\2\2\u00ec"+
		"-\3\2\2\2\u00ed\u00ee\7\60\2\2\u00ee\u00ef\7\20\2\2\u00ef\u00f0\5:\36"+
		"\2\u00f0\u00f1\7\21\2\2\u00f1\u00f2\78\2\2\u00f2\u00f3\5\60\31\2\u00f3"+
		"/\3\2\2\2\u00f4\u00f8\7\4\2\2\u00f5\u00f7\5<\37\2\u00f6\u00f5\3\2\2\2"+
		"\u00f7\u00fa\3\2\2\2\u00f8\u00f6\3\2\2\2\u00f8\u00f9\3\2\2\2\u00f9\61"+
		"\3\2\2\2\u00fa\u00f8\3\2\2\2\u00fb\u00fc\7\61\2\2\u00fc\u00fd\78\2\2\u00fd"+
		"\u00ff\7\13\2\2\u00fe\u0100\5\64\33\2\u00ff\u00fe\3\2\2\2\u00ff\u0100"+
		"\3\2\2\2\u0100\u0101\3\2\2\2\u0101\u0104\7\f\2\2\u0102\u0103\7\22\2\2"+
		"\u0103\u0105\5:\36\2\u0104\u0102\3\2\2\2\u0104\u0105\3\2\2\2\u0105\u0106"+
		"\3\2\2\2\u0106\u0107\7\b\2\2\u0107\u0108\5\4\3\2\u0108\u0109\7\t\2\2\u0109"+
		"\63\3\2\2\2\u010a\u010b\5\66\34\2\u010b\u010c\7\r\2\2\u010c\u010d\5\64"+
		"\33\2\u010d\u0110\3\2\2\2\u010e\u0110\5\66\34\2\u010f\u010a\3\2\2\2\u010f"+
		"\u010e\3\2\2\2\u0110\65\3\2\2\2\u0111\u0113\78\2\2\u0112\u0111\3\2\2\2"+
		"\u0112\u0113\3\2\2\2\u0113\u0114\3\2\2\2\u0114\u0115\78\2\2\u0115\u0116"+
		"\7\3\2\2\u0116\u0117\5:\36\2\u0117\67\3\2\2\2\u0118\u0119\78\2\2\u0119"+
		"\u011a\7\13\2\2\u011a\u011b\5$\23\2\u011b\u011c\7\f\2\2\u011c9\3\2\2\2"+
		"\u011d\u0123\7\"\2\2\u011e\u0123\7#\2\2\u011f\u0123\7$\2\2\u0120\u0123"+
		"\7%\2\2\u0121\u0123\7&\2\2\u0122\u011d\3\2\2\2\u0122\u011e\3\2\2\2\u0122"+
		"\u011f\3\2\2\2\u0122\u0120\3\2\2\2\u0122\u0121\3\2\2\2\u0123;\3\2\2\2"+
		"\u0124\u0125\b\37\1\2\u0125\u0126\7\23\2\2\u0126\u0139\5<\37\27\u0127"+
		"\u0128\7\24\2\2\u0128\u0139\5<\37\26\u0129\u012a\7\13\2\2\u012a\u012b"+
		"\5<\37\2\u012b\u012c\7\f\2\2\u012c\u0139\3\2\2\2\u012d\u0139\7\67\2\2"+
		"\u012e\u0139\7\66\2\2\u012f\u0139\78\2\2\u0130\u0139\79\2\2\u0131\u0139"+
		"\7:\2\2\u0132\u0139\t\3\2\2\u0133\u0139\7!\2\2\u0134\u0139\5\36\20\2\u0135"+
		"\u0139\5 \21\2\u0136\u0139\5\"\22\2\u0137\u0139\58\35\2\u0138\u0124\3"+
		"\2\2\2\u0138\u0127\3\2\2\2\u0138\u0129\3\2\2\2\u0138\u012d\3\2\2\2\u0138"+
		"\u012e\3\2\2\2\u0138\u012f\3\2\2\2\u0138\u0130\3\2\2\2\u0138\u0131\3\2"+
		"\2\2\u0138\u0132\3\2\2\2\u0138\u0133\3\2\2\2\u0138\u0134\3\2\2\2\u0138"+
		"\u0135\3\2\2\2\u0138\u0136\3\2\2\2\u0138\u0137\3\2\2\2\u0139\u0151\3\2"+
		"\2\2\u013a\u013b\f\25\2\2\u013b\u013c\t\4\2\2\u013c\u0150\5<\37\26\u013d"+
		"\u013e\f\24\2\2\u013e\u013f\t\5\2\2\u013f\u0150\5<\37\25\u0140\u0141\f"+
		"\23\2\2\u0141\u0142\t\6\2\2\u0142\u0150\5<\37\24\u0143\u0144\f\22\2\2"+
		"\u0144\u0145\t\7\2\2\u0145\u0150\5<\37\23\u0146\u0147\f\21\2\2\u0147\u0148"+
		"\t\b\2\2\u0148\u0150\5<\37\22\u0149\u014a\f\20\2\2\u014a\u014b\7\35\2"+
		"\2\u014b\u0150\5<\37\21\u014c\u014d\f\17\2\2\u014d\u014e\7\36\2\2\u014e"+
		"\u0150\5<\37\20\u014f\u013a\3\2\2\2\u014f\u013d\3\2\2\2\u014f\u0140\3"+
		"\2\2\2\u014f\u0143\3\2\2\2\u014f\u0146\3\2\2\2\u014f\u0149\3\2\2\2\u014f"+
		"\u014c\3\2\2\2\u0150\u0153\3\2\2\2\u0151\u014f\3\2\2\2\u0151\u0152\3\2"+
		"\2\2\u0152=\3\2\2\2\u0153\u0151\3\2\2\2\25DV^q\u0097\u009f\u00ab\u00b0"+
		"\u00c8\u00e0\u00f8\u00ff\u0104\u010f\u0112\u0122\u0138\u014f\u0151";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
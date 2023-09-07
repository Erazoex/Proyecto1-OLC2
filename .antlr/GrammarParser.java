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
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T_INT=31, 
		T_FLOAT=32, T_CHAR=33, T_STRING=34, T_BOOL=35, VAR=36, LET=37, IF=38, 
		ELSE=39, SWITCH=40, CASE=41, DEFAULT=42, WHILE=43, FOR=44, VECTOR=45, 
		FUNC=46, DOUBLE=47, INT=48, ID=49, CHAR=50, STRING=51, MULTCOMMENT=52, 
		COMMENT=53, WHITESPACE=54;
	public static final int
		RULE_init = 0, RULE_block = 1, RULE_stmt = 2, RULE_declstmt = 3, RULE_asignstmt = 4, 
		RULE_incstmt = 5, RULE_decstmt = 6, RULE_ifstmt = 7, RULE_switchstmt = 8, 
		RULE_switchcase = 9, RULE_printlnstmt = 10, RULE_exprparams = 11, RULE_whilestmt = 12, 
		RULE_forstmt = 13, RULE_forrange = 14, RULE_array = 15, RULE_array_def = 16, 
		RULE_vartype = 17, RULE_expr = 18;
	private static String[] makeRuleNames() {
		return new String[] {
			"init", "block", "stmt", "declstmt", "asignstmt", "incstmt", "decstmt", 
			"ifstmt", "switchstmt", "switchcase", "printlnstmt", "exprparams", "whilestmt", 
			"forstmt", "forrange", "array", "array_def", "vartype", "expr"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "':'", "'='", "'?'", "'+='", "'-='", "'{'", "'}'", "'print'", "'('", 
			"')'", "','", "'in'", "'...'", "'<'", "'>'", "'!'", "'-'", "'*'", "'/'", 
			"'%'", "'+'", "'>='", "'<='", "'=='", "'!='", "'&&'", "'||'", "'true'", 
			"'false'", "'nil'", "'Int'", "'Float'", "'Character'", "'String'", "'Bool'", 
			"'var'", "'let'", "'if'", "'else'", "'switch'", "'case'", "'default'", 
			"'while'", "'for'", "'vector'", "'func'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, "T_INT", "T_FLOAT", "T_CHAR", 
			"T_STRING", "T_BOOL", "VAR", "LET", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", 
			"WHILE", "FOR", "VECTOR", "FUNC", "DOUBLE", "INT", "ID", "CHAR", "STRING", 
			"MULTCOMMENT", "COMMENT", "WHITESPACE"
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
			setState(38);
			block();
			setState(39);
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
			setState(44);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__7) | (1L << VAR) | (1L << LET) | (1L << IF) | (1L << SWITCH) | (1L << WHILE) | (1L << FOR) | (1L << ID))) != 0)) {
				{
				{
				setState(41);
				stmt();
				}
				}
				setState(46);
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
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_stmt);
		try {
			setState(56);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(47);
				declstmt();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(48);
				asignstmt();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(49);
				incstmt();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(50);
				decstmt();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(51);
				ifstmt();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(52);
				switchstmt();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(53);
				printlnstmt();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(54);
				whilestmt();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(55);
				forstmt();
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
			setState(75);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				_localctx = new DeclstmtWithTypeAndExprContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(58);
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
				setState(59);
				match(ID);
				setState(60);
				match(T__0);
				setState(61);
				vartype();
				setState(62);
				match(T__1);
				setState(63);
				expr(0);
				}
				break;
			case 2:
				_localctx = new DeclstmtWithExprContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(65);
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
				setState(66);
				match(ID);
				setState(67);
				match(T__1);
				setState(68);
				expr(0);
				}
				break;
			case 3:
				_localctx = new DeclstmtWithTypeContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(69);
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
				setState(70);
				match(ID);
				setState(71);
				match(T__0);
				setState(72);
				vartype();
				setState(73);
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
			setState(77);
			match(ID);
			setState(78);
			match(T__1);
			setState(79);
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
		enterRule(_localctx, 10, RULE_incstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(81);
			match(ID);
			setState(82);
			match(T__3);
			setState(83);
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
		enterRule(_localctx, 12, RULE_decstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(85);
			match(ID);
			setState(86);
			match(T__4);
			setState(87);
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
		enterRule(_localctx, 14, RULE_ifstmt);
		try {
			setState(113);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				_localctx = new IfSimpleContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(89);
				match(IF);
				setState(90);
				expr(0);
				setState(91);
				match(T__5);
				setState(92);
				block();
				setState(93);
				match(T__6);
				}
				break;
			case 2:
				_localctx = new IfWithElseContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(95);
				match(IF);
				setState(96);
				expr(0);
				setState(97);
				match(T__5);
				setState(98);
				((IfWithElseContext)_localctx).trueCondition = block();
				setState(99);
				match(T__6);
				setState(100);
				match(ELSE);
				setState(101);
				match(T__5);
				setState(102);
				((IfWithElseContext)_localctx).falseCondition = block();
				setState(103);
				match(T__6);
				}
				break;
			case 3:
				_localctx = new IfWithElseIfContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(105);
				match(IF);
				setState(106);
				expr(0);
				setState(107);
				match(T__5);
				setState(108);
				block();
				setState(109);
				match(T__6);
				setState(110);
				match(ELSE);
				setState(111);
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
		enterRule(_localctx, 16, RULE_switchstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(115);
			match(SWITCH);
			setState(116);
			expr(0);
			setState(117);
			match(T__5);
			setState(119); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(118);
				switchcase();
				}
				}
				setState(121); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE || _la==DEFAULT );
			setState(123);
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
		enterRule(_localctx, 18, RULE_switchcase);
		try {
			setState(133);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CASE:
				enterOuterAlt(_localctx, 1);
				{
				setState(125);
				((SwitchcaseContext)_localctx).casetype = match(CASE);
				setState(126);
				expr(0);
				setState(127);
				match(T__0);
				setState(128);
				block();
				}
				break;
			case DEFAULT:
				enterOuterAlt(_localctx, 2);
				{
				setState(130);
				((SwitchcaseContext)_localctx).casetype = match(DEFAULT);
				setState(131);
				match(T__0);
				setState(132);
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
		enterRule(_localctx, 20, RULE_printlnstmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(135);
			match(T__7);
			setState(136);
			match(T__8);
			setState(138);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__15) | (1L << T__16) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << DOUBLE) | (1L << INT) | (1L << ID) | (1L << CHAR) | (1L << STRING))) != 0)) {
				{
				setState(137);
				exprparams();
				}
			}

			setState(140);
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
		enterRule(_localctx, 22, RULE_exprparams);
		try {
			setState(147);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				_localctx = new ExprWithParamsContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(142);
				expr(0);
				setState(143);
				match(T__10);
				setState(144);
				exprparams();
				}
				break;
			case 2:
				_localctx = new ExprWithParamContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(146);
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
		enterRule(_localctx, 24, RULE_whilestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(149);
			match(WHILE);
			setState(150);
			expr(0);
			setState(151);
			match(T__5);
			setState(152);
			block();
			setState(153);
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
		enterRule(_localctx, 26, RULE_forstmt);
		try {
			setState(171);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				_localctx = new ForWithExprContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(155);
				match(FOR);
				setState(156);
				match(ID);
				setState(157);
				match(T__11);
				setState(158);
				expr(0);
				setState(159);
				match(T__5);
				setState(160);
				block();
				setState(161);
				match(T__6);
				}
				break;
			case 2:
				_localctx = new ForWithRangeContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(163);
				match(FOR);
				setState(164);
				match(ID);
				setState(165);
				match(T__11);
				setState(166);
				forrange();
				setState(167);
				match(T__5);
				setState(168);
				block();
				setState(169);
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
		enterRule(_localctx, 28, RULE_forrange);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(173);
			((ForrangeContext)_localctx).beginsWith = expr(0);
			setState(174);
			match(T__12);
			setState(175);
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
		enterRule(_localctx, 30, RULE_array);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(177);
			match(VECTOR);
			setState(178);
			match(T__13);
			setState(179);
			vartype();
			setState(180);
			match(T__14);
			setState(181);
			match(ID);
			setState(182);
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
		enterRule(_localctx, 32, RULE_array_def);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(184);
			match(T__1);
			setState(188);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__8) | (1L << T__15) | (1L << T__16) | (1L << T__27) | (1L << T__28) | (1L << T__29) | (1L << DOUBLE) | (1L << INT) | (1L << ID) | (1L << CHAR) | (1L << STRING))) != 0)) {
				{
				{
				setState(185);
				expr(0);
				}
				}
				setState(190);
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
		enterRule(_localctx, 34, RULE_vartype);
		try {
			setState(196);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T_INT:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(191);
				((PrimitiveTypeContext)_localctx).tipo = match(T_INT);
				}
				break;
			case T_FLOAT:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(192);
				((PrimitiveTypeContext)_localctx).tipo = match(T_FLOAT);
				}
				break;
			case T_CHAR:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(193);
				((PrimitiveTypeContext)_localctx).tipo = match(T_CHAR);
				}
				break;
			case T_STRING:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(194);
				((PrimitiveTypeContext)_localctx).tipo = match(T_STRING);
				}
				break;
			case T_BOOL:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(195);
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
		int _startState = 36;
		enterRecursionRule(_localctx, 36, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(214);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__15:
				{
				_localctx = new NotExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(199);
				match(T__15);
				setState(200);
				((NotExprContext)_localctx).right = expr(17);
				}
				break;
			case T__16:
				{
				_localctx = new UnaryExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(201);
				match(T__16);
				setState(202);
				((UnaryExprContext)_localctx).right = expr(16);
				}
				break;
			case T__8:
				{
				_localctx = new ParExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(203);
				match(T__8);
				setState(204);
				expr(0);
				setState(205);
				match(T__9);
				}
				break;
			case INT:
				{
				_localctx = new IntExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(207);
				match(INT);
				}
				break;
			case DOUBLE:
				{
				_localctx = new DoubleExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(208);
				match(DOUBLE);
				}
				break;
			case ID:
				{
				_localctx = new IdExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(209);
				match(ID);
				}
				break;
			case CHAR:
				{
				_localctx = new CharExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(210);
				match(CHAR);
				}
				break;
			case STRING:
				{
				_localctx = new StrExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(211);
				match(STRING);
				}
				break;
			case T__27:
			case T__28:
				{
				_localctx = new BoolExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(212);
				_la = _input.LA(1);
				if ( !(_la==T__27 || _la==T__28) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case T__29:
				{
				_localctx = new NilExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(213);
				match(T__29);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(239);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(237);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
					case 1:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(216);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(217);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__17) | (1L << T__18) | (1L << T__19))) != 0)) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(218);
						((OpExprContext)_localctx).right = expr(16);
						}
						break;
					case 2:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(219);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(220);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__16 || _la==T__20) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(221);
						((OpExprContext)_localctx).right = expr(15);
						}
						break;
					case 3:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(222);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(223);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__14 || _la==T__21) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(224);
						((OpExprContext)_localctx).right = expr(14);
						}
						break;
					case 4:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(225);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(226);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__13 || _la==T__22) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(227);
						((OpExprContext)_localctx).right = expr(13);
						}
						break;
					case 5:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(228);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(229);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__23 || _la==T__24) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(230);
						((OpExprContext)_localctx).right = expr(12);
						}
						break;
					case 6:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(231);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(232);
						((OpExprContext)_localctx).op = match(T__25);
						setState(233);
						((OpExprContext)_localctx).right = expr(11);
						}
						break;
					case 7:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(234);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(235);
						((OpExprContext)_localctx).op = match(T__26);
						setState(236);
						((OpExprContext)_localctx).right = expr(10);
						}
						break;
					}
					} 
				}
				setState(241);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
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
		case 18:
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\38\u00f5\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\3\2\3\2\3\2\3\3\7\3-\n\3\f\3\16\3\60\13\3\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4;\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\5\5N\n\5\3\6\3\6\3\6\3\6\3\7\3\7"+
		"\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\5\tt\n\t\3\n\3\n\3"+
		"\n\3\n\6\nz\n\n\r\n\16\n{\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3"+
		"\13\5\13\u0088\n\13\3\f\3\f\3\f\5\f\u008d\n\f\3\f\3\f\3\r\3\r\3\r\3\r"+
		"\3\r\5\r\u0096\n\r\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3"+
		"\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\5\17\u00ae"+
		"\n\17\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22"+
		"\7\22\u00bd\n\22\f\22\16\22\u00c0\13\22\3\23\3\23\3\23\3\23\3\23\5\23"+
		"\u00c7\n\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\5\24\u00d9\n\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\7\24\u00f0\n\24\f\24\16\24\u00f3\13\24\3\24\2\3&\25\2\4\6\b\n\f\16\20"+
		"\22\24\26\30\32\34\36 \"$&\2\t\3\2&\'\3\2\36\37\3\2\24\26\4\2\23\23\27"+
		"\27\4\2\21\21\30\30\4\2\20\20\31\31\3\2\32\33\2\u0108\2(\3\2\2\2\4.\3"+
		"\2\2\2\6:\3\2\2\2\bM\3\2\2\2\nO\3\2\2\2\fS\3\2\2\2\16W\3\2\2\2\20s\3\2"+
		"\2\2\22u\3\2\2\2\24\u0087\3\2\2\2\26\u0089\3\2\2\2\30\u0095\3\2\2\2\32"+
		"\u0097\3\2\2\2\34\u00ad\3\2\2\2\36\u00af\3\2\2\2 \u00b3\3\2\2\2\"\u00ba"+
		"\3\2\2\2$\u00c6\3\2\2\2&\u00d8\3\2\2\2()\5\4\3\2)*\7\2\2\3*\3\3\2\2\2"+
		"+-\5\6\4\2,+\3\2\2\2-\60\3\2\2\2.,\3\2\2\2./\3\2\2\2/\5\3\2\2\2\60.\3"+
		"\2\2\2\61;\5\b\5\2\62;\5\n\6\2\63;\5\f\7\2\64;\5\16\b\2\65;\5\20\t\2\66"+
		";\5\22\n\2\67;\5\26\f\28;\5\32\16\29;\5\34\17\2:\61\3\2\2\2:\62\3\2\2"+
		"\2:\63\3\2\2\2:\64\3\2\2\2:\65\3\2\2\2:\66\3\2\2\2:\67\3\2\2\2:8\3\2\2"+
		"\2:9\3\2\2\2;\7\3\2\2\2<=\t\2\2\2=>\7\63\2\2>?\7\3\2\2?@\5$\23\2@A\7\4"+
		"\2\2AB\5&\24\2BN\3\2\2\2CD\t\2\2\2DE\7\63\2\2EF\7\4\2\2FN\5&\24\2GH\t"+
		"\2\2\2HI\7\63\2\2IJ\7\3\2\2JK\5$\23\2KL\7\5\2\2LN\3\2\2\2M<\3\2\2\2MC"+
		"\3\2\2\2MG\3\2\2\2N\t\3\2\2\2OP\7\63\2\2PQ\7\4\2\2QR\5&\24\2R\13\3\2\2"+
		"\2ST\7\63\2\2TU\7\6\2\2UV\5&\24\2V\r\3\2\2\2WX\7\63\2\2XY\7\7\2\2YZ\5"+
		"&\24\2Z\17\3\2\2\2[\\\7(\2\2\\]\5&\24\2]^\7\b\2\2^_\5\4\3\2_`\7\t\2\2"+
		"`t\3\2\2\2ab\7(\2\2bc\5&\24\2cd\7\b\2\2de\5\4\3\2ef\7\t\2\2fg\7)\2\2g"+
		"h\7\b\2\2hi\5\4\3\2ij\7\t\2\2jt\3\2\2\2kl\7(\2\2lm\5&\24\2mn\7\b\2\2n"+
		"o\5\4\3\2op\7\t\2\2pq\7)\2\2qr\5\20\t\2rt\3\2\2\2s[\3\2\2\2sa\3\2\2\2"+
		"sk\3\2\2\2t\21\3\2\2\2uv\7*\2\2vw\5&\24\2wy\7\b\2\2xz\5\24\13\2yx\3\2"+
		"\2\2z{\3\2\2\2{y\3\2\2\2{|\3\2\2\2|}\3\2\2\2}~\7\t\2\2~\23\3\2\2\2\177"+
		"\u0080\7+\2\2\u0080\u0081\5&\24\2\u0081\u0082\7\3\2\2\u0082\u0083\5\4"+
		"\3\2\u0083\u0088\3\2\2\2\u0084\u0085\7,\2\2\u0085\u0086\7\3\2\2\u0086"+
		"\u0088\5\4\3\2\u0087\177\3\2\2\2\u0087\u0084\3\2\2\2\u0088\25\3\2\2\2"+
		"\u0089\u008a\7\n\2\2\u008a\u008c\7\13\2\2\u008b\u008d\5\30\r\2\u008c\u008b"+
		"\3\2\2\2\u008c\u008d\3\2\2\2\u008d\u008e\3\2\2\2\u008e\u008f\7\f\2\2\u008f"+
		"\27\3\2\2\2\u0090\u0091\5&\24\2\u0091\u0092\7\r\2\2\u0092\u0093\5\30\r"+
		"\2\u0093\u0096\3\2\2\2\u0094\u0096\5&\24\2\u0095\u0090\3\2\2\2\u0095\u0094"+
		"\3\2\2\2\u0096\31\3\2\2\2\u0097\u0098\7-\2\2\u0098\u0099\5&\24\2\u0099"+
		"\u009a\7\b\2\2\u009a\u009b\5\4\3\2\u009b\u009c\7\t\2\2\u009c\33\3\2\2"+
		"\2\u009d\u009e\7.\2\2\u009e\u009f\7\63\2\2\u009f\u00a0\7\16\2\2\u00a0"+
		"\u00a1\5&\24\2\u00a1\u00a2\7\b\2\2\u00a2\u00a3\5\4\3\2\u00a3\u00a4\7\t"+
		"\2\2\u00a4\u00ae\3\2\2\2\u00a5\u00a6\7.\2\2\u00a6\u00a7\7\63\2\2\u00a7"+
		"\u00a8\7\16\2\2\u00a8\u00a9\5\36\20\2\u00a9\u00aa\7\b\2\2\u00aa\u00ab"+
		"\5\4\3\2\u00ab\u00ac\7\t\2\2\u00ac\u00ae\3\2\2\2\u00ad\u009d\3\2\2\2\u00ad"+
		"\u00a5\3\2\2\2\u00ae\35\3\2\2\2\u00af\u00b0\5&\24\2\u00b0\u00b1\7\17\2"+
		"\2\u00b1\u00b2\5&\24\2\u00b2\37\3\2\2\2\u00b3\u00b4\7/\2\2\u00b4\u00b5"+
		"\7\20\2\2\u00b5\u00b6\5$\23\2\u00b6\u00b7\7\21\2\2\u00b7\u00b8\7\63\2"+
		"\2\u00b8\u00b9\5\"\22\2\u00b9!\3\2\2\2\u00ba\u00be\7\4\2\2\u00bb\u00bd"+
		"\5&\24\2\u00bc\u00bb\3\2\2\2\u00bd\u00c0\3\2\2\2\u00be\u00bc\3\2\2\2\u00be"+
		"\u00bf\3\2\2\2\u00bf#\3\2\2\2\u00c0\u00be\3\2\2\2\u00c1\u00c7\7!\2\2\u00c2"+
		"\u00c7\7\"\2\2\u00c3\u00c7\7#\2\2\u00c4\u00c7\7$\2\2\u00c5\u00c7\7%\2"+
		"\2\u00c6\u00c1\3\2\2\2\u00c6\u00c2\3\2\2\2\u00c6\u00c3\3\2\2\2\u00c6\u00c4"+
		"\3\2\2\2\u00c6\u00c5\3\2\2\2\u00c7%\3\2\2\2\u00c8\u00c9\b\24\1\2\u00c9"+
		"\u00ca\7\22\2\2\u00ca\u00d9\5&\24\23\u00cb\u00cc\7\23\2\2\u00cc\u00d9"+
		"\5&\24\22\u00cd\u00ce\7\13\2\2\u00ce\u00cf\5&\24\2\u00cf\u00d0\7\f\2\2"+
		"\u00d0\u00d9\3\2\2\2\u00d1\u00d9\7\62\2\2\u00d2\u00d9\7\61\2\2\u00d3\u00d9"+
		"\7\63\2\2\u00d4\u00d9\7\64\2\2\u00d5\u00d9\7\65\2\2\u00d6\u00d9\t\3\2"+
		"\2\u00d7\u00d9\7 \2\2\u00d8\u00c8\3\2\2\2\u00d8\u00cb\3\2\2\2\u00d8\u00cd"+
		"\3\2\2\2\u00d8\u00d1\3\2\2\2\u00d8\u00d2\3\2\2\2\u00d8\u00d3\3\2\2\2\u00d8"+
		"\u00d4\3\2\2\2\u00d8\u00d5\3\2\2\2\u00d8\u00d6\3\2\2\2\u00d8\u00d7\3\2"+
		"\2\2\u00d9\u00f1\3\2\2\2\u00da\u00db\f\21\2\2\u00db\u00dc\t\4\2\2\u00dc"+
		"\u00f0\5&\24\22\u00dd\u00de\f\20\2\2\u00de\u00df\t\5\2\2\u00df\u00f0\5"+
		"&\24\21\u00e0\u00e1\f\17\2\2\u00e1\u00e2\t\6\2\2\u00e2\u00f0\5&\24\20"+
		"\u00e3\u00e4\f\16\2\2\u00e4\u00e5\t\7\2\2\u00e5\u00f0\5&\24\17\u00e6\u00e7"+
		"\f\r\2\2\u00e7\u00e8\t\b\2\2\u00e8\u00f0\5&\24\16\u00e9\u00ea\f\f\2\2"+
		"\u00ea\u00eb\7\34\2\2\u00eb\u00f0\5&\24\r\u00ec\u00ed\f\13\2\2\u00ed\u00ee"+
		"\7\35\2\2\u00ee\u00f0\5&\24\f\u00ef\u00da\3\2\2\2\u00ef\u00dd\3\2\2\2"+
		"\u00ef\u00e0\3\2\2\2\u00ef\u00e3\3\2\2\2\u00ef\u00e6\3\2\2\2\u00ef\u00e9"+
		"\3\2\2\2\u00ef\u00ec\3\2\2\2\u00f0\u00f3\3\2\2\2\u00f1\u00ef\3\2\2\2\u00f1"+
		"\u00f2\3\2\2\2\u00f2\'\3\2\2\2\u00f3\u00f1\3\2\2\2\20.:Ms{\u0087\u008c"+
		"\u0095\u00ad\u00be\u00c6\u00d8\u00ef\u00f1";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
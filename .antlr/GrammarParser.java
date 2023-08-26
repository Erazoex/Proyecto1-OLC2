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
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T_INT=30, T_FLOAT=31, 
		T_CHAR=32, T_STRING=33, T_BOOL=34, VAR=35, LET=36, IF=37, ELSE=38, SWITCH=39, 
		CASE=40, DEFAULT=41, WHILE=42, FOR=43, DOUBLE=44, INT=45, ID=46, CHAR=47, 
		STRING=48, MULTCOMMENT=49, COMMENT=50, WHITESPACE=51;
	public static final int
		RULE_init = 0, RULE_block = 1, RULE_stmt = 2, RULE_declstmt = 3, RULE_asignstmt = 4, 
		RULE_incstmt = 5, RULE_decstmt = 6, RULE_ifstmt = 7, RULE_switchstmt = 8, 
		RULE_switchcase = 9, RULE_printlnstmt = 10, RULE_whilestmt = 11, RULE_forstmt = 12, 
		RULE_forrange = 13, RULE_vartype = 14, RULE_expr = 15;
	private static String[] makeRuleNames() {
		return new String[] {
			"init", "block", "stmt", "declstmt", "asignstmt", "incstmt", "decstmt", 
			"ifstmt", "switchstmt", "switchcase", "printlnstmt", "whilestmt", "forstmt", 
			"forrange", "vartype", "expr"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "':'", "'='", "'?'", "'+='", "'-='", "'{'", "'}'", "'println'", 
			"'('", "')'", "'in'", "'...'", "'!'", "'-'", "'*'", "'/'", "'%'", "'+'", 
			"'>='", "'>'", "'<='", "'<'", "'=='", "'!='", "'&&'", "'||'", "'true'", 
			"'false'", "'nil'", "'Int'", "'Float'", "'Character'", "'String'", "'Bool'", 
			"'var'", "'let'", "'if'", "'else'", "'switch'", "'case'", "'default'", 
			"'while'", "'for'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, "T_INT", "T_FLOAT", "T_CHAR", "T_STRING", 
			"T_BOOL", "VAR", "LET", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "WHILE", 
			"FOR", "DOUBLE", "INT", "ID", "CHAR", "STRING", "MULTCOMMENT", "COMMENT", 
			"WHITESPACE"
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
			setState(32);
			block();
			setState(33);
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
			setState(38);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__7) | (1L << VAR) | (1L << LET) | (1L << IF) | (1L << SWITCH) | (1L << WHILE) | (1L << FOR) | (1L << ID))) != 0)) {
				{
				{
				setState(35);
				stmt();
				}
				}
				setState(40);
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
			setState(50);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(41);
				declstmt();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(42);
				asignstmt();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(43);
				incstmt();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(44);
				decstmt();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(45);
				ifstmt();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(46);
				switchstmt();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(47);
				printlnstmt();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(48);
				whilestmt();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(49);
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
			setState(69);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				_localctx = new DeclstmtWithTypeAndExprContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(52);
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
				setState(53);
				match(ID);
				setState(54);
				match(T__0);
				setState(55);
				vartype();
				setState(56);
				match(T__1);
				setState(57);
				expr(0);
				}
				break;
			case 2:
				_localctx = new DeclstmtWithExprContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(59);
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
				setState(60);
				match(ID);
				setState(61);
				match(T__1);
				setState(62);
				expr(0);
				}
				break;
			case 3:
				_localctx = new DeclstmtWithTypeContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(63);
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
				setState(64);
				match(ID);
				setState(65);
				match(T__0);
				setState(66);
				vartype();
				setState(67);
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
			setState(71);
			match(ID);
			setState(72);
			match(T__1);
			setState(73);
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
			setState(75);
			match(ID);
			setState(76);
			match(T__3);
			setState(77);
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
			setState(79);
			match(ID);
			setState(80);
			match(T__4);
			setState(81);
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
			setState(107);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				_localctx = new IfSimpleContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(83);
				match(IF);
				setState(84);
				expr(0);
				setState(85);
				match(T__5);
				setState(86);
				block();
				setState(87);
				match(T__6);
				}
				break;
			case 2:
				_localctx = new IfWithElseContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(89);
				match(IF);
				setState(90);
				expr(0);
				setState(91);
				match(T__5);
				setState(92);
				((IfWithElseContext)_localctx).trueCondition = block();
				setState(93);
				match(T__6);
				setState(94);
				match(ELSE);
				setState(95);
				match(T__5);
				setState(96);
				((IfWithElseContext)_localctx).falseCondition = block();
				setState(97);
				match(T__6);
				}
				break;
			case 3:
				_localctx = new IfWithElseIfContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(99);
				match(IF);
				setState(100);
				expr(0);
				setState(101);
				match(T__5);
				setState(102);
				block();
				setState(103);
				match(T__6);
				setState(104);
				match(ELSE);
				setState(105);
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
			setState(109);
			match(SWITCH);
			setState(110);
			expr(0);
			setState(111);
			match(T__5);
			setState(113); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(112);
				switchcase();
				}
				}
				setState(115); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE || _la==DEFAULT );
			setState(117);
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
			setState(127);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CASE:
				enterOuterAlt(_localctx, 1);
				{
				setState(119);
				((SwitchcaseContext)_localctx).casetype = match(CASE);
				setState(120);
				expr(0);
				setState(121);
				match(T__0);
				setState(122);
				block();
				}
				break;
			case DEFAULT:
				enterOuterAlt(_localctx, 2);
				{
				setState(124);
				((SwitchcaseContext)_localctx).casetype = match(DEFAULT);
				setState(125);
				match(T__0);
				setState(126);
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
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public PrintlnstmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printlnstmt; }
	}

	public final PrintlnstmtContext printlnstmt() throws RecognitionException {
		PrintlnstmtContext _localctx = new PrintlnstmtContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_printlnstmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(129);
			match(T__7);
			setState(130);
			match(T__8);
			setState(131);
			expr(0);
			setState(132);
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
		enterRule(_localctx, 22, RULE_whilestmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(134);
			match(WHILE);
			setState(135);
			expr(0);
			setState(136);
			match(T__5);
			setState(137);
			block();
			setState(138);
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
		enterRule(_localctx, 24, RULE_forstmt);
		try {
			setState(156);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				_localctx = new ForWithExprContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(140);
				match(FOR);
				setState(141);
				match(ID);
				setState(142);
				match(T__10);
				setState(143);
				expr(0);
				setState(144);
				match(T__5);
				setState(145);
				block();
				setState(146);
				match(T__6);
				}
				break;
			case 2:
				_localctx = new ForWithRangeContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(148);
				match(FOR);
				setState(149);
				match(ID);
				setState(150);
				match(T__10);
				setState(151);
				forrange();
				setState(152);
				match(T__5);
				setState(153);
				block();
				setState(154);
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
		public Token beginsWith;
		public Token endsWith;
		public List<TerminalNode> INT() { return getTokens(GrammarParser.INT); }
		public TerminalNode INT(int i) {
			return getToken(GrammarParser.INT, i);
		}
		public ForrangeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forrange; }
	}

	public final ForrangeContext forrange() throws RecognitionException {
		ForrangeContext _localctx = new ForrangeContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_forrange);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(158);
			((ForrangeContext)_localctx).beginsWith = match(INT);
			setState(159);
			match(T__11);
			setState(160);
			((ForrangeContext)_localctx).endsWith = match(INT);
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
		enterRule(_localctx, 28, RULE_vartype);
		try {
			setState(167);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T_INT:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(162);
				((PrimitiveTypeContext)_localctx).tipo = match(T_INT);
				}
				break;
			case T_FLOAT:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(163);
				((PrimitiveTypeContext)_localctx).tipo = match(T_FLOAT);
				}
				break;
			case T_CHAR:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(164);
				((PrimitiveTypeContext)_localctx).tipo = match(T_CHAR);
				}
				break;
			case T_STRING:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(165);
				((PrimitiveTypeContext)_localctx).tipo = match(T_STRING);
				}
				break;
			case T_BOOL:
				_localctx = new PrimitiveTypeContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(166);
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
		int _startState = 30;
		enterRecursionRule(_localctx, 30, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(185);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__12:
				{
				_localctx = new NotExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(170);
				match(T__12);
				setState(171);
				((NotExprContext)_localctx).right = expr(17);
				}
				break;
			case T__13:
				{
				_localctx = new UnaryExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(172);
				match(T__13);
				setState(173);
				((UnaryExprContext)_localctx).right = expr(16);
				}
				break;
			case T__8:
				{
				_localctx = new ParExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(174);
				match(T__8);
				setState(175);
				expr(0);
				setState(176);
				match(T__9);
				}
				break;
			case INT:
				{
				_localctx = new IntExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(178);
				match(INT);
				}
				break;
			case DOUBLE:
				{
				_localctx = new DoubleExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(179);
				match(DOUBLE);
				}
				break;
			case ID:
				{
				_localctx = new IdExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(180);
				match(ID);
				}
				break;
			case CHAR:
				{
				_localctx = new CharExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(181);
				match(CHAR);
				}
				break;
			case STRING:
				{
				_localctx = new StrExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(182);
				match(STRING);
				}
				break;
			case T__26:
			case T__27:
				{
				_localctx = new BoolExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(183);
				_la = _input.LA(1);
				if ( !(_la==T__26 || _la==T__27) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case T__28:
				{
				_localctx = new NilExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(184);
				match(T__28);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(210);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(208);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
					case 1:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(187);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(188);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__14) | (1L << T__15) | (1L << T__16))) != 0)) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(189);
						((OpExprContext)_localctx).right = expr(16);
						}
						break;
					case 2:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(190);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(191);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__13 || _la==T__17) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(192);
						((OpExprContext)_localctx).right = expr(15);
						}
						break;
					case 3:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(193);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(194);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__18 || _la==T__19) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(195);
						((OpExprContext)_localctx).right = expr(14);
						}
						break;
					case 4:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(196);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(197);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__20 || _la==T__21) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(198);
						((OpExprContext)_localctx).right = expr(13);
						}
						break;
					case 5:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(199);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(200);
						((OpExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__22 || _la==T__23) ) {
							((OpExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(201);
						((OpExprContext)_localctx).right = expr(12);
						}
						break;
					case 6:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(202);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(203);
						match(T__24);
						setState(204);
						((OpExprContext)_localctx).right = expr(11);
						}
						break;
					case 7:
						{
						_localctx = new OpExprContext(new ExprContext(_parentctx, _parentState));
						((OpExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(205);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(206);
						match(T__25);
						setState(207);
						((OpExprContext)_localctx).right = expr(10);
						}
						break;
					}
					} 
				}
				setState(212);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
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
		case 15:
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\65\u00d8\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\3\2\3\2"+
		"\3\2\3\3\7\3\'\n\3\f\3\16\3*\13\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\5\4\65\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\3\5\5\5H\n\5\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\5\tn\n\t\3\n\3\n\3\n\3\n\6\nt\n\n\r\n\16\nu\3"+
		"\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13\u0082\n\13\3\f\3\f"+
		"\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\5\16\u009f\n\16\3\17\3\17"+
		"\3\17\3\17\3\20\3\20\3\20\3\20\3\20\5\20\u00aa\n\20\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\5\21\u00bc"+
		"\n\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\7\21\u00d3\n\21\f\21\16\21\u00d6"+
		"\13\21\3\21\2\3 \22\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \2\t\3\2%&"+
		"\3\2\35\36\3\2\21\23\4\2\20\20\24\24\3\2\25\26\3\2\27\30\3\2\31\32\2\u00eb"+
		"\2\"\3\2\2\2\4(\3\2\2\2\6\64\3\2\2\2\bG\3\2\2\2\nI\3\2\2\2\fM\3\2\2\2"+
		"\16Q\3\2\2\2\20m\3\2\2\2\22o\3\2\2\2\24\u0081\3\2\2\2\26\u0083\3\2\2\2"+
		"\30\u0088\3\2\2\2\32\u009e\3\2\2\2\34\u00a0\3\2\2\2\36\u00a9\3\2\2\2 "+
		"\u00bb\3\2\2\2\"#\5\4\3\2#$\7\2\2\3$\3\3\2\2\2%\'\5\6\4\2&%\3\2\2\2\'"+
		"*\3\2\2\2(&\3\2\2\2()\3\2\2\2)\5\3\2\2\2*(\3\2\2\2+\65\5\b\5\2,\65\5\n"+
		"\6\2-\65\5\f\7\2.\65\5\16\b\2/\65\5\20\t\2\60\65\5\22\n\2\61\65\5\26\f"+
		"\2\62\65\5\30\r\2\63\65\5\32\16\2\64+\3\2\2\2\64,\3\2\2\2\64-\3\2\2\2"+
		"\64.\3\2\2\2\64/\3\2\2\2\64\60\3\2\2\2\64\61\3\2\2\2\64\62\3\2\2\2\64"+
		"\63\3\2\2\2\65\7\3\2\2\2\66\67\t\2\2\2\678\7\60\2\289\7\3\2\29:\5\36\20"+
		"\2:;\7\4\2\2;<\5 \21\2<H\3\2\2\2=>\t\2\2\2>?\7\60\2\2?@\7\4\2\2@H\5 \21"+
		"\2AB\t\2\2\2BC\7\60\2\2CD\7\3\2\2DE\5\36\20\2EF\7\5\2\2FH\3\2\2\2G\66"+
		"\3\2\2\2G=\3\2\2\2GA\3\2\2\2H\t\3\2\2\2IJ\7\60\2\2JK\7\4\2\2KL\5 \21\2"+
		"L\13\3\2\2\2MN\7\60\2\2NO\7\6\2\2OP\5 \21\2P\r\3\2\2\2QR\7\60\2\2RS\7"+
		"\7\2\2ST\5 \21\2T\17\3\2\2\2UV\7\'\2\2VW\5 \21\2WX\7\b\2\2XY\5\4\3\2Y"+
		"Z\7\t\2\2Zn\3\2\2\2[\\\7\'\2\2\\]\5 \21\2]^\7\b\2\2^_\5\4\3\2_`\7\t\2"+
		"\2`a\7(\2\2ab\7\b\2\2bc\5\4\3\2cd\7\t\2\2dn\3\2\2\2ef\7\'\2\2fg\5 \21"+
		"\2gh\7\b\2\2hi\5\4\3\2ij\7\t\2\2jk\7(\2\2kl\5\20\t\2ln\3\2\2\2mU\3\2\2"+
		"\2m[\3\2\2\2me\3\2\2\2n\21\3\2\2\2op\7)\2\2pq\5 \21\2qs\7\b\2\2rt\5\24"+
		"\13\2sr\3\2\2\2tu\3\2\2\2us\3\2\2\2uv\3\2\2\2vw\3\2\2\2wx\7\t\2\2x\23"+
		"\3\2\2\2yz\7*\2\2z{\5 \21\2{|\7\3\2\2|}\5\4\3\2}\u0082\3\2\2\2~\177\7"+
		"+\2\2\177\u0080\7\3\2\2\u0080\u0082\5\4\3\2\u0081y\3\2\2\2\u0081~\3\2"+
		"\2\2\u0082\25\3\2\2\2\u0083\u0084\7\n\2\2\u0084\u0085\7\13\2\2\u0085\u0086"+
		"\5 \21\2\u0086\u0087\7\f\2\2\u0087\27\3\2\2\2\u0088\u0089\7,\2\2\u0089"+
		"\u008a\5 \21\2\u008a\u008b\7\b\2\2\u008b\u008c\5\4\3\2\u008c\u008d\7\t"+
		"\2\2\u008d\31\3\2\2\2\u008e\u008f\7-\2\2\u008f\u0090\7\60\2\2\u0090\u0091"+
		"\7\r\2\2\u0091\u0092\5 \21\2\u0092\u0093\7\b\2\2\u0093\u0094\5\4\3\2\u0094"+
		"\u0095\7\t\2\2\u0095\u009f\3\2\2\2\u0096\u0097\7-\2\2\u0097\u0098\7\60"+
		"\2\2\u0098\u0099\7\r\2\2\u0099\u009a\5\34\17\2\u009a\u009b\7\b\2\2\u009b"+
		"\u009c\5\4\3\2\u009c\u009d\7\t\2\2\u009d\u009f\3\2\2\2\u009e\u008e\3\2"+
		"\2\2\u009e\u0096\3\2\2\2\u009f\33\3\2\2\2\u00a0\u00a1\7/\2\2\u00a1\u00a2"+
		"\7\16\2\2\u00a2\u00a3\7/\2\2\u00a3\35\3\2\2\2\u00a4\u00aa\7 \2\2\u00a5"+
		"\u00aa\7!\2\2\u00a6\u00aa\7\"\2\2\u00a7\u00aa\7#\2\2\u00a8\u00aa\7$\2"+
		"\2\u00a9\u00a4\3\2\2\2\u00a9\u00a5\3\2\2\2\u00a9\u00a6\3\2\2\2\u00a9\u00a7"+
		"\3\2\2\2\u00a9\u00a8\3\2\2\2\u00aa\37\3\2\2\2\u00ab\u00ac\b\21\1\2\u00ac"+
		"\u00ad\7\17\2\2\u00ad\u00bc\5 \21\23\u00ae\u00af\7\20\2\2\u00af\u00bc"+
		"\5 \21\22\u00b0\u00b1\7\13\2\2\u00b1\u00b2\5 \21\2\u00b2\u00b3\7\f\2\2"+
		"\u00b3\u00bc\3\2\2\2\u00b4\u00bc\7/\2\2\u00b5\u00bc\7.\2\2\u00b6\u00bc"+
		"\7\60\2\2\u00b7\u00bc\7\61\2\2\u00b8\u00bc\7\62\2\2\u00b9\u00bc\t\3\2"+
		"\2\u00ba\u00bc\7\37\2\2\u00bb\u00ab\3\2\2\2\u00bb\u00ae\3\2\2\2\u00bb"+
		"\u00b0\3\2\2\2\u00bb\u00b4\3\2\2\2\u00bb\u00b5\3\2\2\2\u00bb\u00b6\3\2"+
		"\2\2\u00bb\u00b7\3\2\2\2\u00bb\u00b8\3\2\2\2\u00bb\u00b9\3\2\2\2\u00bb"+
		"\u00ba\3\2\2\2\u00bc\u00d4\3\2\2\2\u00bd\u00be\f\21\2\2\u00be\u00bf\t"+
		"\4\2\2\u00bf\u00d3\5 \21\22\u00c0\u00c1\f\20\2\2\u00c1\u00c2\t\5\2\2\u00c2"+
		"\u00d3\5 \21\21\u00c3\u00c4\f\17\2\2\u00c4\u00c5\t\6\2\2\u00c5\u00d3\5"+
		" \21\20\u00c6\u00c7\f\16\2\2\u00c7\u00c8\t\7\2\2\u00c8\u00d3\5 \21\17"+
		"\u00c9\u00ca\f\r\2\2\u00ca\u00cb\t\b\2\2\u00cb\u00d3\5 \21\16\u00cc\u00cd"+
		"\f\f\2\2\u00cd\u00ce\7\33\2\2\u00ce\u00d3\5 \21\r\u00cf\u00d0\f\13\2\2"+
		"\u00d0\u00d1\7\34\2\2\u00d1\u00d3\5 \21\f\u00d2\u00bd\3\2\2\2\u00d2\u00c0"+
		"\3\2\2\2\u00d2\u00c3\3\2\2\2\u00d2\u00c6\3\2\2\2\u00d2\u00c9\3\2\2\2\u00d2"+
		"\u00cc\3\2\2\2\u00d2\u00cf\3\2\2\2\u00d3\u00d6\3\2\2\2\u00d4\u00d2\3\2"+
		"\2\2\u00d4\u00d5\3\2\2\2\u00d5!\3\2\2\2\u00d6\u00d4\3\2\2\2\r(\64Gmu\u0081"+
		"\u009e\u00a9\u00bb\u00d2\u00d4";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
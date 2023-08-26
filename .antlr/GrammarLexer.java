// Generated from c:\Users\brian\OneDrive\Escritorio\Repositories\proyecto\Grammar.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class GrammarLexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
			"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24", 
			"T__25", "T__26", "T__27", "T__28", "T_INT", "T_FLOAT", "T_CHAR", "T_STRING", 
			"T_BOOL", "VAR", "LET", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "WHILE", 
			"FOR", "DOUBLE", "INT", "ID", "CHAR", "STRING", "MULTCOMMENT", "COMMENT", 
			"WHITESPACE"
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


	public GrammarLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Grammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\65\u015c\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64"+
		"\t\64\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\5\3\6\3\6\3\6\3\7\3\7\3\b\3\b"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\f\3\r\3\r"+
		"\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3"+
		"\24\3\24\3\24\3\25\3\25\3\26\3\26\3\26\3\27\3\27\3\30\3\30\3\30\3\31\3"+
		"\31\3\31\3\32\3\32\3\32\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\35\3"+
		"\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3 \3 "+
		"\3 \3 \3 \3 \3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3\"\3\""+
		"\3#\3#\3#\3#\3#\3$\3$\3$\3$\3%\3%\3%\3%\3&\3&\3&\3\'\3\'\3\'\3\'\3\'\3"+
		"(\3(\3(\3(\3(\3(\3(\3)\3)\3)\3)\3)\3*\3*\3*\3*\3*\3*\3*\3*\3+\3+\3+\3"+
		"+\3+\3+\3,\3,\3,\3,\3-\6-\u010d\n-\r-\16-\u010e\3-\3-\6-\u0113\n-\r-\16"+
		"-\u0114\3.\6.\u0118\n.\r.\16.\u0119\3/\3/\7/\u011e\n/\f/\16/\u0121\13"+
		"/\3\60\3\60\5\60\u0125\n\60\3\60\3\60\3\61\3\61\3\61\3\61\7\61\u012d\n"+
		"\61\f\61\16\61\u0130\13\61\3\61\3\61\3\62\3\62\3\62\3\62\3\62\3\62\7\62"+
		"\u013a\n\62\f\62\16\62\u013d\13\62\3\62\3\62\3\62\3\62\3\62\3\63\3\63"+
		"\3\63\3\63\7\63\u0148\n\63\f\63\16\63\u014b\13\63\3\63\5\63\u014e\n\63"+
		"\3\63\3\63\5\63\u0152\n\63\3\63\3\63\3\64\6\64\u0157\n\64\r\64\16\64\u0158"+
		"\3\64\3\64\2\2\65\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31"+
		"\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65"+
		"\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64"+
		"g\65\3\2\13\3\2\62;\5\2C\\aac|\6\2\62;C\\aac|\5\2\13\f\17\17$$\5\2\f\f"+
		"\17\17$$\4\2,,\61\61\3\2\61\61\4\2\f\f\17\17\5\2\13\f\17\17\"\"\2\u0168"+
		"\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2"+
		"\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2"+
		"\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2"+
		"\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2"+
		"\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3"+
		"\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2"+
		"\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2"+
		"U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3"+
		"\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\3i\3\2\2\2\5k\3\2\2\2\7m\3\2\2"+
		"\2\to\3\2\2\2\13r\3\2\2\2\ru\3\2\2\2\17w\3\2\2\2\21y\3\2\2\2\23\u0081"+
		"\3\2\2\2\25\u0083\3\2\2\2\27\u0085\3\2\2\2\31\u0088\3\2\2\2\33\u008c\3"+
		"\2\2\2\35\u008e\3\2\2\2\37\u0090\3\2\2\2!\u0092\3\2\2\2#\u0094\3\2\2\2"+
		"%\u0096\3\2\2\2\'\u0098\3\2\2\2)\u009b\3\2\2\2+\u009d\3\2\2\2-\u00a0\3"+
		"\2\2\2/\u00a2\3\2\2\2\61\u00a5\3\2\2\2\63\u00a8\3\2\2\2\65\u00ab\3\2\2"+
		"\2\67\u00ae\3\2\2\29\u00b3\3\2\2\2;\u00b9\3\2\2\2=\u00bd\3\2\2\2?\u00c1"+
		"\3\2\2\2A\u00c7\3\2\2\2C\u00d1\3\2\2\2E\u00d8\3\2\2\2G\u00dd\3\2\2\2I"+
		"\u00e1\3\2\2\2K\u00e5\3\2\2\2M\u00e8\3\2\2\2O\u00ed\3\2\2\2Q\u00f4\3\2"+
		"\2\2S\u00f9\3\2\2\2U\u0101\3\2\2\2W\u0107\3\2\2\2Y\u010c\3\2\2\2[\u0117"+
		"\3\2\2\2]\u011b\3\2\2\2_\u0122\3\2\2\2a\u0128\3\2\2\2c\u0133\3\2\2\2e"+
		"\u0143\3\2\2\2g\u0156\3\2\2\2ij\7<\2\2j\4\3\2\2\2kl\7?\2\2l\6\3\2\2\2"+
		"mn\7A\2\2n\b\3\2\2\2op\7-\2\2pq\7?\2\2q\n\3\2\2\2rs\7/\2\2st\7?\2\2t\f"+
		"\3\2\2\2uv\7}\2\2v\16\3\2\2\2wx\7\177\2\2x\20\3\2\2\2yz\7r\2\2z{\7t\2"+
		"\2{|\7k\2\2|}\7p\2\2}~\7v\2\2~\177\7n\2\2\177\u0080\7p\2\2\u0080\22\3"+
		"\2\2\2\u0081\u0082\7*\2\2\u0082\24\3\2\2\2\u0083\u0084\7+\2\2\u0084\26"+
		"\3\2\2\2\u0085\u0086\7k\2\2\u0086\u0087\7p\2\2\u0087\30\3\2\2\2\u0088"+
		"\u0089\7\60\2\2\u0089\u008a\7\60\2\2\u008a\u008b\7\60\2\2\u008b\32\3\2"+
		"\2\2\u008c\u008d\7#\2\2\u008d\34\3\2\2\2\u008e\u008f\7/\2\2\u008f\36\3"+
		"\2\2\2\u0090\u0091\7,\2\2\u0091 \3\2\2\2\u0092\u0093\7\61\2\2\u0093\""+
		"\3\2\2\2\u0094\u0095\7\'\2\2\u0095$\3\2\2\2\u0096\u0097\7-\2\2\u0097&"+
		"\3\2\2\2\u0098\u0099\7@\2\2\u0099\u009a\7?\2\2\u009a(\3\2\2\2\u009b\u009c"+
		"\7@\2\2\u009c*\3\2\2\2\u009d\u009e\7>\2\2\u009e\u009f\7?\2\2\u009f,\3"+
		"\2\2\2\u00a0\u00a1\7>\2\2\u00a1.\3\2\2\2\u00a2\u00a3\7?\2\2\u00a3\u00a4"+
		"\7?\2\2\u00a4\60\3\2\2\2\u00a5\u00a6\7#\2\2\u00a6\u00a7\7?\2\2\u00a7\62"+
		"\3\2\2\2\u00a8\u00a9\7(\2\2\u00a9\u00aa\7(\2\2\u00aa\64\3\2\2\2\u00ab"+
		"\u00ac\7~\2\2\u00ac\u00ad\7~\2\2\u00ad\66\3\2\2\2\u00ae\u00af\7v\2\2\u00af"+
		"\u00b0\7t\2\2\u00b0\u00b1\7w\2\2\u00b1\u00b2\7g\2\2\u00b28\3\2\2\2\u00b3"+
		"\u00b4\7h\2\2\u00b4\u00b5\7c\2\2\u00b5\u00b6\7n\2\2\u00b6\u00b7\7u\2\2"+
		"\u00b7\u00b8\7g\2\2\u00b8:\3\2\2\2\u00b9\u00ba\7p\2\2\u00ba\u00bb\7k\2"+
		"\2\u00bb\u00bc\7n\2\2\u00bc<\3\2\2\2\u00bd\u00be\7K\2\2\u00be\u00bf\7"+
		"p\2\2\u00bf\u00c0\7v\2\2\u00c0>\3\2\2\2\u00c1\u00c2\7H\2\2\u00c2\u00c3"+
		"\7n\2\2\u00c3\u00c4\7q\2\2\u00c4\u00c5\7c\2\2\u00c5\u00c6\7v\2\2\u00c6"+
		"@\3\2\2\2\u00c7\u00c8\7E\2\2\u00c8\u00c9\7j\2\2\u00c9\u00ca\7c\2\2\u00ca"+
		"\u00cb\7t\2\2\u00cb\u00cc\7c\2\2\u00cc\u00cd\7e\2\2\u00cd\u00ce\7v\2\2"+
		"\u00ce\u00cf\7g\2\2\u00cf\u00d0\7t\2\2\u00d0B\3\2\2\2\u00d1\u00d2\7U\2"+
		"\2\u00d2\u00d3\7v\2\2\u00d3\u00d4\7t\2\2\u00d4\u00d5\7k\2\2\u00d5\u00d6"+
		"\7p\2\2\u00d6\u00d7\7i\2\2\u00d7D\3\2\2\2\u00d8\u00d9\7D\2\2\u00d9\u00da"+
		"\7q\2\2\u00da\u00db\7q\2\2\u00db\u00dc\7n\2\2\u00dcF\3\2\2\2\u00dd\u00de"+
		"\7x\2\2\u00de\u00df\7c\2\2\u00df\u00e0\7t\2\2\u00e0H\3\2\2\2\u00e1\u00e2"+
		"\7n\2\2\u00e2\u00e3\7g\2\2\u00e3\u00e4\7v\2\2\u00e4J\3\2\2\2\u00e5\u00e6"+
		"\7k\2\2\u00e6\u00e7\7h\2\2\u00e7L\3\2\2\2\u00e8\u00e9\7g\2\2\u00e9\u00ea"+
		"\7n\2\2\u00ea\u00eb\7u\2\2\u00eb\u00ec\7g\2\2\u00ecN\3\2\2\2\u00ed\u00ee"+
		"\7u\2\2\u00ee\u00ef\7y\2\2\u00ef\u00f0\7k\2\2\u00f0\u00f1\7v\2\2\u00f1"+
		"\u00f2\7e\2\2\u00f2\u00f3\7j\2\2\u00f3P\3\2\2\2\u00f4\u00f5\7e\2\2\u00f5"+
		"\u00f6\7c\2\2\u00f6\u00f7\7u\2\2\u00f7\u00f8\7g\2\2\u00f8R\3\2\2\2\u00f9"+
		"\u00fa\7f\2\2\u00fa\u00fb\7g\2\2\u00fb\u00fc\7h\2\2\u00fc\u00fd\7c\2\2"+
		"\u00fd\u00fe\7w\2\2\u00fe\u00ff\7n\2\2\u00ff\u0100\7v\2\2\u0100T\3\2\2"+
		"\2\u0101\u0102\7y\2\2\u0102\u0103\7j\2\2\u0103\u0104\7k\2\2\u0104\u0105"+
		"\7n\2\2\u0105\u0106\7g\2\2\u0106V\3\2\2\2\u0107\u0108\7h\2\2\u0108\u0109"+
		"\7q\2\2\u0109\u010a\7t\2\2\u010aX\3\2\2\2\u010b\u010d\t\2\2\2\u010c\u010b"+
		"\3\2\2\2\u010d\u010e\3\2\2\2\u010e\u010c\3\2\2\2\u010e\u010f\3\2\2\2\u010f"+
		"\u0110\3\2\2\2\u0110\u0112\7\60\2\2\u0111\u0113\t\2\2\2\u0112\u0111\3"+
		"\2\2\2\u0113\u0114\3\2\2\2\u0114\u0112\3\2\2\2\u0114\u0115\3\2\2\2\u0115"+
		"Z\3\2\2\2\u0116\u0118\t\2\2\2\u0117\u0116\3\2\2\2\u0118\u0119\3\2\2\2"+
		"\u0119\u0117\3\2\2\2\u0119\u011a\3\2\2\2\u011a\\\3\2\2\2\u011b\u011f\t"+
		"\3\2\2\u011c\u011e\t\4\2\2\u011d\u011c\3\2\2\2\u011e\u0121\3\2\2\2\u011f"+
		"\u011d\3\2\2\2\u011f\u0120\3\2\2\2\u0120^\3\2\2\2\u0121\u011f\3\2\2\2"+
		"\u0122\u0124\7$\2\2\u0123\u0125\13\2\2\2\u0124\u0123\3\2\2\2\u0124\u0125"+
		"\3\2\2\2\u0125\u0126\3\2\2\2\u0126\u0127\7$\2\2\u0127`\3\2\2\2\u0128\u012e"+
		"\7$\2\2\u0129\u012a\7^\2\2\u012a\u012d\t\5\2\2\u012b\u012d\n\6\2\2\u012c"+
		"\u0129\3\2\2\2\u012c\u012b\3\2\2\2\u012d\u0130\3\2\2\2\u012e\u012c\3\2"+
		"\2\2\u012e\u012f\3\2\2\2\u012f\u0131\3\2\2\2\u0130\u012e\3\2\2\2\u0131"+
		"\u0132\7$\2\2\u0132b\3\2\2\2\u0133\u0134\7\61\2\2\u0134\u0135\7,\2\2\u0135"+
		"\u013b\3\2\2\2\u0136\u013a\n\7\2\2\u0137\u0138\7,\2\2\u0138\u013a\n\b"+
		"\2\2\u0139\u0136\3\2\2\2\u0139\u0137\3\2\2\2\u013a\u013d\3\2\2\2\u013b"+
		"\u0139\3\2\2\2\u013b\u013c\3\2\2\2\u013c\u013e\3\2\2\2\u013d\u013b\3\2"+
		"\2\2\u013e\u013f\7,\2\2\u013f\u0140\7\61\2\2\u0140\u0141\3\2\2\2\u0141"+
		"\u0142\b\62\2\2\u0142d\3\2\2\2\u0143\u0144\7\61\2\2\u0144\u0145\7\61\2"+
		"\2\u0145\u0149\3\2\2\2\u0146\u0148\n\t\2\2\u0147\u0146\3\2\2\2\u0148\u014b"+
		"\3\2\2\2\u0149\u0147\3\2\2\2\u0149\u014a\3\2\2\2\u014a\u0151\3\2\2\2\u014b"+
		"\u0149\3\2\2\2\u014c\u014e\7\17\2\2\u014d\u014c\3\2\2\2\u014d\u014e\3"+
		"\2\2\2\u014e\u014f\3\2\2\2\u014f\u0152\7\f\2\2\u0150\u0152\7\2\2\3\u0151"+
		"\u014d\3\2\2\2\u0151\u0150\3\2\2\2\u0152\u0153\3\2\2\2\u0153\u0154\b\63"+
		"\2\2\u0154f\3\2\2\2\u0155\u0157\t\n\2\2\u0156\u0155\3\2\2\2\u0157\u0158"+
		"\3\2\2\2\u0158\u0156\3\2\2\2\u0158\u0159\3\2\2\2\u0159\u015a\3\2\2\2\u015a"+
		"\u015b\b\64\2\2\u015bh\3\2\2\2\20\2\u010e\u0114\u0119\u011f\u0124\u012c"+
		"\u012e\u0139\u013b\u0149\u014d\u0151\u0158\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
// Generated from c:\Users\brian\OneDrive\Escritorio\Repositories\proyecto\backend\Grammar.g4 by ANTLR 4.9.2
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
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		T_INT=32, T_FLOAT=33, T_CHAR=34, T_STRING=35, T_BOOL=36, VAR=37, LET=38, 
		IF=39, ELSE=40, SWITCH=41, CASE=42, DEFAULT=43, WHILE=44, FOR=45, VECTOR=46, 
		FUNC=47, GUARD=48, BREAK=49, RETURN=50, CONTINUE=51, DOUBLE=52, INT=53, 
		ID=54, CHAR=55, STRING=56, MULTCOMMENT=57, COMMENT=58, WHITESPACE=59;
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
			"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T_INT", "T_FLOAT", 
			"T_CHAR", "T_STRING", "T_BOOL", "VAR", "LET", "IF", "ELSE", "SWITCH", 
			"CASE", "DEFAULT", "WHILE", "FOR", "VECTOR", "FUNC", "GUARD", "BREAK", 
			"RETURN", "CONTINUE", "DOUBLE", "INT", "ID", "CHAR", "STRING", "MULTCOMMENT", 
			"COMMENT", "WHITESPACE"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2=\u0197\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\3\2\3"+
		"\2\3\3\3\3\3\4\3\4\3\5\3\5\3\5\3\6\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\r\3\16\3\16\3\16\3\16"+
		"\3\17\3\17\3\20\3\20\3\21\3\21\3\21\3\22\3\22\3\23\3\23\3\24\3\24\3\25"+
		"\3\25\3\26\3\26\3\27\3\27\3\30\3\30\3\30\3\31\3\31\3\31\3\32\3\32\3\32"+
		"\3\33\3\33\3\33\3\34\3\34\3\34\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36"+
		"\3\37\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3!\3!\3!\3!\3\"\3\"\3\"\3\""+
		"\3\"\3\"\3#\3#\3#\3#\3#\3#\3#\3#\3#\3#\3$\3$\3$\3$\3$\3$\3$\3%\3%\3%\3"+
		"%\3%\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3(\3(\3(\3)\3)\3)\3)\3)\3*\3*\3*\3*\3"+
		"*\3*\3*\3+\3+\3+\3+\3+\3,\3,\3,\3,\3,\3,\3,\3,\3-\3-\3-\3-\3-\3-\3.\3"+
		".\3.\3.\3/\3/\3/\3/\3/\3/\3/\3\60\3\60\3\60\3\60\3\60\3\61\3\61\3\61\3"+
		"\61\3\61\3\61\3\62\3\62\3\62\3\62\3\62\3\62\3\63\3\63\3\63\3\63\3\63\3"+
		"\63\3\63\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\65\6\65\u0148"+
		"\n\65\r\65\16\65\u0149\3\65\3\65\6\65\u014e\n\65\r\65\16\65\u014f\3\66"+
		"\6\66\u0153\n\66\r\66\16\66\u0154\3\67\3\67\7\67\u0159\n\67\f\67\16\67"+
		"\u015c\13\67\38\38\58\u0160\n8\38\38\39\39\39\39\79\u0168\n9\f9\169\u016b"+
		"\139\39\39\3:\3:\3:\3:\3:\3:\7:\u0175\n:\f:\16:\u0178\13:\3:\3:\3:\3:"+
		"\3:\3;\3;\3;\3;\7;\u0183\n;\f;\16;\u0186\13;\3;\5;\u0189\n;\3;\3;\5;\u018d"+
		"\n;\3;\3;\3<\6<\u0192\n<\r<\16<\u0193\3<\3<\2\2=\3\3\5\4\7\5\t\6\13\7"+
		"\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25"+
		")\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O"+
		")Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65i\66k\67m8o9q:s;u<w=\3\2\13\3\2\62"+
		";\5\2C\\aac|\6\2\62;C\\aac|\5\2\13\f\17\17$$\5\2\f\f\17\17$$\4\2,,\61"+
		"\61\3\2\61\61\4\2\f\f\17\17\5\2\13\f\17\17\"\"\2\u01a3\2\3\3\2\2\2\2\5"+
		"\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2"+
		"\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33"+
		"\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2"+
		"\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2"+
		"\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2"+
		"\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K"+
		"\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2"+
		"\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2"+
		"\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\2q"+
		"\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\3y\3\2\2\2\5{\3\2\2\2\7}\3\2"+
		"\2\2\t\177\3\2\2\2\13\u0082\3\2\2\2\r\u0085\3\2\2\2\17\u0087\3\2\2\2\21"+
		"\u0089\3\2\2\2\23\u008f\3\2\2\2\25\u0091\3\2\2\2\27\u0093\3\2\2\2\31\u0095"+
		"\3\2\2\2\33\u0098\3\2\2\2\35\u009c\3\2\2\2\37\u009e\3\2\2\2!\u00a0\3\2"+
		"\2\2#\u00a3\3\2\2\2%\u00a5\3\2\2\2\'\u00a7\3\2\2\2)\u00a9\3\2\2\2+\u00ab"+
		"\3\2\2\2-\u00ad\3\2\2\2/\u00af\3\2\2\2\61\u00b2\3\2\2\2\63\u00b5\3\2\2"+
		"\2\65\u00b8\3\2\2\2\67\u00bb\3\2\2\29\u00be\3\2\2\2;\u00c1\3\2\2\2=\u00c6"+
		"\3\2\2\2?\u00cc\3\2\2\2A\u00d0\3\2\2\2C\u00d4\3\2\2\2E\u00da\3\2\2\2G"+
		"\u00e4\3\2\2\2I\u00eb\3\2\2\2K\u00f0\3\2\2\2M\u00f4\3\2\2\2O\u00f8\3\2"+
		"\2\2Q\u00fb\3\2\2\2S\u0100\3\2\2\2U\u0107\3\2\2\2W\u010c\3\2\2\2Y\u0114"+
		"\3\2\2\2[\u011a\3\2\2\2]\u011e\3\2\2\2_\u0125\3\2\2\2a\u012a\3\2\2\2c"+
		"\u0130\3\2\2\2e\u0136\3\2\2\2g\u013d\3\2\2\2i\u0147\3\2\2\2k\u0152\3\2"+
		"\2\2m\u0156\3\2\2\2o\u015d\3\2\2\2q\u0163\3\2\2\2s\u016e\3\2\2\2u\u017e"+
		"\3\2\2\2w\u0191\3\2\2\2yz\7<\2\2z\4\3\2\2\2{|\7?\2\2|\6\3\2\2\2}~\7A\2"+
		"\2~\b\3\2\2\2\177\u0080\7-\2\2\u0080\u0081\7?\2\2\u0081\n\3\2\2\2\u0082"+
		"\u0083\7/\2\2\u0083\u0084\7?\2\2\u0084\f\3\2\2\2\u0085\u0086\7}\2\2\u0086"+
		"\16\3\2\2\2\u0087\u0088\7\177\2\2\u0088\20\3\2\2\2\u0089\u008a\7r\2\2"+
		"\u008a\u008b\7t\2\2\u008b\u008c\7k\2\2\u008c\u008d\7p\2\2\u008d\u008e"+
		"\7v\2\2\u008e\22\3\2\2\2\u008f\u0090\7*\2\2\u0090\24\3\2\2\2\u0091\u0092"+
		"\7+\2\2\u0092\26\3\2\2\2\u0093\u0094\7.\2\2\u0094\30\3\2\2\2\u0095\u0096"+
		"\7k\2\2\u0096\u0097\7p\2\2\u0097\32\3\2\2\2\u0098\u0099\7\60\2\2\u0099"+
		"\u009a\7\60\2\2\u009a\u009b\7\60\2\2\u009b\34\3\2\2\2\u009c\u009d\7>\2"+
		"\2\u009d\36\3\2\2\2\u009e\u009f\7@\2\2\u009f \3\2\2\2\u00a0\u00a1\7/\2"+
		"\2\u00a1\u00a2\7@\2\2\u00a2\"\3\2\2\2\u00a3\u00a4\7#\2\2\u00a4$\3\2\2"+
		"\2\u00a5\u00a6\7/\2\2\u00a6&\3\2\2\2\u00a7\u00a8\7,\2\2\u00a8(\3\2\2\2"+
		"\u00a9\u00aa\7\61\2\2\u00aa*\3\2\2\2\u00ab\u00ac\7\'\2\2\u00ac,\3\2\2"+
		"\2\u00ad\u00ae\7-\2\2\u00ae.\3\2\2\2\u00af\u00b0\7@\2\2\u00b0\u00b1\7"+
		"?\2\2\u00b1\60\3\2\2\2\u00b2\u00b3\7>\2\2\u00b3\u00b4\7?\2\2\u00b4\62"+
		"\3\2\2\2\u00b5\u00b6\7?\2\2\u00b6\u00b7\7?\2\2\u00b7\64\3\2\2\2\u00b8"+
		"\u00b9\7#\2\2\u00b9\u00ba\7?\2\2\u00ba\66\3\2\2\2\u00bb\u00bc\7(\2\2\u00bc"+
		"\u00bd\7(\2\2\u00bd8\3\2\2\2\u00be\u00bf\7~\2\2\u00bf\u00c0\7~\2\2\u00c0"+
		":\3\2\2\2\u00c1\u00c2\7v\2\2\u00c2\u00c3\7t\2\2\u00c3\u00c4\7w\2\2\u00c4"+
		"\u00c5\7g\2\2\u00c5<\3\2\2\2\u00c6\u00c7\7h\2\2\u00c7\u00c8\7c\2\2\u00c8"+
		"\u00c9\7n\2\2\u00c9\u00ca\7u\2\2\u00ca\u00cb\7g\2\2\u00cb>\3\2\2\2\u00cc"+
		"\u00cd\7p\2\2\u00cd\u00ce\7k\2\2\u00ce\u00cf\7n\2\2\u00cf@\3\2\2\2\u00d0"+
		"\u00d1\7K\2\2\u00d1\u00d2\7p\2\2\u00d2\u00d3\7v\2\2\u00d3B\3\2\2\2\u00d4"+
		"\u00d5\7H\2\2\u00d5\u00d6\7n\2\2\u00d6\u00d7\7q\2\2\u00d7\u00d8\7c\2\2"+
		"\u00d8\u00d9\7v\2\2\u00d9D\3\2\2\2\u00da\u00db\7E\2\2\u00db\u00dc\7j\2"+
		"\2\u00dc\u00dd\7c\2\2\u00dd\u00de\7t\2\2\u00de\u00df\7c\2\2\u00df\u00e0"+
		"\7e\2\2\u00e0\u00e1\7v\2\2\u00e1\u00e2\7g\2\2\u00e2\u00e3\7t\2\2\u00e3"+
		"F\3\2\2\2\u00e4\u00e5\7U\2\2\u00e5\u00e6\7v\2\2\u00e6\u00e7\7t\2\2\u00e7"+
		"\u00e8\7k\2\2\u00e8\u00e9\7p\2\2\u00e9\u00ea\7i\2\2\u00eaH\3\2\2\2\u00eb"+
		"\u00ec\7D\2\2\u00ec\u00ed\7q\2\2\u00ed\u00ee\7q\2\2\u00ee\u00ef\7n\2\2"+
		"\u00efJ\3\2\2\2\u00f0\u00f1\7x\2\2\u00f1\u00f2\7c\2\2\u00f2\u00f3\7t\2"+
		"\2\u00f3L\3\2\2\2\u00f4\u00f5\7n\2\2\u00f5\u00f6\7g\2\2\u00f6\u00f7\7"+
		"v\2\2\u00f7N\3\2\2\2\u00f8\u00f9\7k\2\2\u00f9\u00fa\7h\2\2\u00faP\3\2"+
		"\2\2\u00fb\u00fc\7g\2\2\u00fc\u00fd\7n\2\2\u00fd\u00fe\7u\2\2\u00fe\u00ff"+
		"\7g\2\2\u00ffR\3\2\2\2\u0100\u0101\7u\2\2\u0101\u0102\7y\2\2\u0102\u0103"+
		"\7k\2\2\u0103\u0104\7v\2\2\u0104\u0105\7e\2\2\u0105\u0106\7j\2\2\u0106"+
		"T\3\2\2\2\u0107\u0108\7e\2\2\u0108\u0109\7c\2\2\u0109\u010a\7u\2\2\u010a"+
		"\u010b\7g\2\2\u010bV\3\2\2\2\u010c\u010d\7f\2\2\u010d\u010e\7g\2\2\u010e"+
		"\u010f\7h\2\2\u010f\u0110\7c\2\2\u0110\u0111\7w\2\2\u0111\u0112\7n\2\2"+
		"\u0112\u0113\7v\2\2\u0113X\3\2\2\2\u0114\u0115\7y\2\2\u0115\u0116\7j\2"+
		"\2\u0116\u0117\7k\2\2\u0117\u0118\7n\2\2\u0118\u0119\7g\2\2\u0119Z\3\2"+
		"\2\2\u011a\u011b\7h\2\2\u011b\u011c\7q\2\2\u011c\u011d\7t\2\2\u011d\\"+
		"\3\2\2\2\u011e\u011f\7x\2\2\u011f\u0120\7g\2\2\u0120\u0121\7e\2\2\u0121"+
		"\u0122\7v\2\2\u0122\u0123\7q\2\2\u0123\u0124\7t\2\2\u0124^\3\2\2\2\u0125"+
		"\u0126\7h\2\2\u0126\u0127\7w\2\2\u0127\u0128\7p\2\2\u0128\u0129\7e\2\2"+
		"\u0129`\3\2\2\2\u012a\u012b\7i\2\2\u012b\u012c\7w\2\2\u012c\u012d\7c\2"+
		"\2\u012d\u012e\7t\2\2\u012e\u012f\7f\2\2\u012fb\3\2\2\2\u0130\u0131\7"+
		"d\2\2\u0131\u0132\7t\2\2\u0132\u0133\7g\2\2\u0133\u0134\7c\2\2\u0134\u0135"+
		"\7m\2\2\u0135d\3\2\2\2\u0136\u0137\7t\2\2\u0137\u0138\7g\2\2\u0138\u0139"+
		"\7v\2\2\u0139\u013a\7w\2\2\u013a\u013b\7t\2\2\u013b\u013c\7p\2\2\u013c"+
		"f\3\2\2\2\u013d\u013e\7e\2\2\u013e\u013f\7q\2\2\u013f\u0140\7p\2\2\u0140"+
		"\u0141\7v\2\2\u0141\u0142\7k\2\2\u0142\u0143\7p\2\2\u0143\u0144\7w\2\2"+
		"\u0144\u0145\7g\2\2\u0145h\3\2\2\2\u0146\u0148\t\2\2\2\u0147\u0146\3\2"+
		"\2\2\u0148\u0149\3\2\2\2\u0149\u0147\3\2\2\2\u0149\u014a\3\2\2\2\u014a"+
		"\u014b\3\2\2\2\u014b\u014d\7\60\2\2\u014c\u014e\t\2\2\2\u014d\u014c\3"+
		"\2\2\2\u014e\u014f\3\2\2\2\u014f\u014d\3\2\2\2\u014f\u0150\3\2\2\2\u0150"+
		"j\3\2\2\2\u0151\u0153\t\2\2\2\u0152\u0151\3\2\2\2\u0153\u0154\3\2\2\2"+
		"\u0154\u0152\3\2\2\2\u0154\u0155\3\2\2\2\u0155l\3\2\2\2\u0156\u015a\t"+
		"\3\2\2\u0157\u0159\t\4\2\2\u0158\u0157\3\2\2\2\u0159\u015c\3\2\2\2\u015a"+
		"\u0158\3\2\2\2\u015a\u015b\3\2\2\2\u015bn\3\2\2\2\u015c\u015a\3\2\2\2"+
		"\u015d\u015f\7$\2\2\u015e\u0160\13\2\2\2\u015f\u015e\3\2\2\2\u015f\u0160"+
		"\3\2\2\2\u0160\u0161\3\2\2\2\u0161\u0162\7$\2\2\u0162p\3\2\2\2\u0163\u0169"+
		"\7$\2\2\u0164\u0165\7^\2\2\u0165\u0168\t\5\2\2\u0166\u0168\n\6\2\2\u0167"+
		"\u0164\3\2\2\2\u0167\u0166\3\2\2\2\u0168\u016b\3\2\2\2\u0169\u0167\3\2"+
		"\2\2\u0169\u016a\3\2\2\2\u016a\u016c\3\2\2\2\u016b\u0169\3\2\2\2\u016c"+
		"\u016d\7$\2\2\u016dr\3\2\2\2\u016e\u016f\7\61\2\2\u016f\u0170\7,\2\2\u0170"+
		"\u0176\3\2\2\2\u0171\u0175\n\7\2\2\u0172\u0173\7,\2\2\u0173\u0175\n\b"+
		"\2\2\u0174\u0171\3\2\2\2\u0174\u0172\3\2\2\2\u0175\u0178\3\2\2\2\u0176"+
		"\u0174\3\2\2\2\u0176\u0177\3\2\2\2\u0177\u0179\3\2\2\2\u0178\u0176\3\2"+
		"\2\2\u0179\u017a\7,\2\2\u017a\u017b\7\61\2\2\u017b\u017c\3\2\2\2\u017c"+
		"\u017d\b:\2\2\u017dt\3\2\2\2\u017e\u017f\7\61\2\2\u017f\u0180\7\61\2\2"+
		"\u0180\u0184\3\2\2\2\u0181\u0183\n\t\2\2\u0182\u0181\3\2\2\2\u0183\u0186"+
		"\3\2\2\2\u0184\u0182\3\2\2\2\u0184\u0185\3\2\2\2\u0185\u018c\3\2\2\2\u0186"+
		"\u0184\3\2\2\2\u0187\u0189\7\17\2\2\u0188\u0187\3\2\2\2\u0188\u0189\3"+
		"\2\2\2\u0189\u018a\3\2\2\2\u018a\u018d\7\f\2\2\u018b\u018d\7\2\2\3\u018c"+
		"\u0188\3\2\2\2\u018c\u018b\3\2\2\2\u018d\u018e\3\2\2\2\u018e\u018f\b;"+
		"\2\2\u018fv\3\2\2\2\u0190\u0192\t\n\2\2\u0191\u0190\3\2\2\2\u0192\u0193"+
		"\3\2\2\2\u0193\u0191\3\2\2\2\u0193\u0194\3\2\2\2\u0194\u0195\3\2\2\2\u0195"+
		"\u0196\b<\2\2\u0196x\3\2\2\2\20\2\u0149\u014f\u0154\u015a\u015f\u0167"+
		"\u0169\u0174\u0176\u0184\u0188\u018c\u0193\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
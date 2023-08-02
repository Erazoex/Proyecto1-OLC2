// Code generated from Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type GrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var GrammarLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func grammarlexerLexerInit() {
	staticData := &GrammarLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'let'", "'='", "'var'", "'print'", "'('", "')'", "'if'", "'{'",
		"'}'", "'for'", "'in'", "'!'", "'-'", "'*'", "'/'", "'%'", "'+'", "'>='",
		"'>'", "'<='", "'<'", "'=='", "'!='", "'&&'", "'||'", "'true'", "'false'",
		"'nil'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "DOUBLE", "INT", "ID",
		"STRING", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "DOUBLE", "INT", "ID", "STRING", "WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 33, 191, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 4,
		28, 152, 8, 28, 11, 28, 12, 28, 153, 1, 28, 1, 28, 4, 28, 158, 8, 28, 11,
		28, 12, 28, 159, 1, 29, 4, 29, 163, 8, 29, 11, 29, 12, 29, 164, 1, 30,
		1, 30, 5, 30, 169, 8, 30, 10, 30, 12, 30, 172, 9, 30, 1, 31, 1, 31, 1,
		31, 1, 31, 5, 31, 178, 8, 31, 10, 31, 12, 31, 181, 9, 31, 1, 31, 1, 31,
		1, 32, 4, 32, 186, 8, 32, 11, 32, 12, 32, 187, 1, 32, 1, 32, 0, 0, 33,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 61, 31, 63, 32, 65, 33, 1, 0, 5, 1, 0, 48, 57, 3, 0, 65, 90, 95,
		95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 10, 10, 13, 13,
		34, 34, 3, 0, 9, 10, 13, 13, 32, 32, 197, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0,
		0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0,
		0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1,
		0, 0, 0, 1, 67, 1, 0, 0, 0, 3, 71, 1, 0, 0, 0, 5, 73, 1, 0, 0, 0, 7, 77,
		1, 0, 0, 0, 9, 83, 1, 0, 0, 0, 11, 85, 1, 0, 0, 0, 13, 87, 1, 0, 0, 0,
		15, 90, 1, 0, 0, 0, 17, 92, 1, 0, 0, 0, 19, 94, 1, 0, 0, 0, 21, 98, 1,
		0, 0, 0, 23, 101, 1, 0, 0, 0, 25, 103, 1, 0, 0, 0, 27, 105, 1, 0, 0, 0,
		29, 107, 1, 0, 0, 0, 31, 109, 1, 0, 0, 0, 33, 111, 1, 0, 0, 0, 35, 113,
		1, 0, 0, 0, 37, 116, 1, 0, 0, 0, 39, 118, 1, 0, 0, 0, 41, 121, 1, 0, 0,
		0, 43, 123, 1, 0, 0, 0, 45, 126, 1, 0, 0, 0, 47, 129, 1, 0, 0, 0, 49, 132,
		1, 0, 0, 0, 51, 135, 1, 0, 0, 0, 53, 140, 1, 0, 0, 0, 55, 146, 1, 0, 0,
		0, 57, 151, 1, 0, 0, 0, 59, 162, 1, 0, 0, 0, 61, 166, 1, 0, 0, 0, 63, 173,
		1, 0, 0, 0, 65, 185, 1, 0, 0, 0, 67, 68, 5, 108, 0, 0, 68, 69, 5, 101,
		0, 0, 69, 70, 5, 116, 0, 0, 70, 2, 1, 0, 0, 0, 71, 72, 5, 61, 0, 0, 72,
		4, 1, 0, 0, 0, 73, 74, 5, 118, 0, 0, 74, 75, 5, 97, 0, 0, 75, 76, 5, 114,
		0, 0, 76, 6, 1, 0, 0, 0, 77, 78, 5, 112, 0, 0, 78, 79, 5, 114, 0, 0, 79,
		80, 5, 105, 0, 0, 80, 81, 5, 110, 0, 0, 81, 82, 5, 116, 0, 0, 82, 8, 1,
		0, 0, 0, 83, 84, 5, 40, 0, 0, 84, 10, 1, 0, 0, 0, 85, 86, 5, 41, 0, 0,
		86, 12, 1, 0, 0, 0, 87, 88, 5, 105, 0, 0, 88, 89, 5, 102, 0, 0, 89, 14,
		1, 0, 0, 0, 90, 91, 5, 123, 0, 0, 91, 16, 1, 0, 0, 0, 92, 93, 5, 125, 0,
		0, 93, 18, 1, 0, 0, 0, 94, 95, 5, 102, 0, 0, 95, 96, 5, 111, 0, 0, 96,
		97, 5, 114, 0, 0, 97, 20, 1, 0, 0, 0, 98, 99, 5, 105, 0, 0, 99, 100, 5,
		110, 0, 0, 100, 22, 1, 0, 0, 0, 101, 102, 5, 33, 0, 0, 102, 24, 1, 0, 0,
		0, 103, 104, 5, 45, 0, 0, 104, 26, 1, 0, 0, 0, 105, 106, 5, 42, 0, 0, 106,
		28, 1, 0, 0, 0, 107, 108, 5, 47, 0, 0, 108, 30, 1, 0, 0, 0, 109, 110, 5,
		37, 0, 0, 110, 32, 1, 0, 0, 0, 111, 112, 5, 43, 0, 0, 112, 34, 1, 0, 0,
		0, 113, 114, 5, 62, 0, 0, 114, 115, 5, 61, 0, 0, 115, 36, 1, 0, 0, 0, 116,
		117, 5, 62, 0, 0, 117, 38, 1, 0, 0, 0, 118, 119, 5, 60, 0, 0, 119, 120,
		5, 61, 0, 0, 120, 40, 1, 0, 0, 0, 121, 122, 5, 60, 0, 0, 122, 42, 1, 0,
		0, 0, 123, 124, 5, 61, 0, 0, 124, 125, 5, 61, 0, 0, 125, 44, 1, 0, 0, 0,
		126, 127, 5, 33, 0, 0, 127, 128, 5, 61, 0, 0, 128, 46, 1, 0, 0, 0, 129,
		130, 5, 38, 0, 0, 130, 131, 5, 38, 0, 0, 131, 48, 1, 0, 0, 0, 132, 133,
		5, 124, 0, 0, 133, 134, 5, 124, 0, 0, 134, 50, 1, 0, 0, 0, 135, 136, 5,
		116, 0, 0, 136, 137, 5, 114, 0, 0, 137, 138, 5, 117, 0, 0, 138, 139, 5,
		101, 0, 0, 139, 52, 1, 0, 0, 0, 140, 141, 5, 102, 0, 0, 141, 142, 5, 97,
		0, 0, 142, 143, 5, 108, 0, 0, 143, 144, 5, 115, 0, 0, 144, 145, 5, 101,
		0, 0, 145, 54, 1, 0, 0, 0, 146, 147, 5, 110, 0, 0, 147, 148, 5, 105, 0,
		0, 148, 149, 5, 108, 0, 0, 149, 56, 1, 0, 0, 0, 150, 152, 7, 0, 0, 0, 151,
		150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 154,
		1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 157, 5, 46, 0, 0, 156, 158, 7, 0,
		0, 0, 157, 156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0,
		159, 160, 1, 0, 0, 0, 160, 58, 1, 0, 0, 0, 161, 163, 7, 0, 0, 0, 162, 161,
		1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0,
		0, 0, 165, 60, 1, 0, 0, 0, 166, 170, 7, 1, 0, 0, 167, 169, 7, 2, 0, 0,
		168, 167, 1, 0, 0, 0, 169, 172, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 170,
		171, 1, 0, 0, 0, 171, 62, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 173, 179, 5,
		34, 0, 0, 174, 178, 8, 3, 0, 0, 175, 176, 5, 34, 0, 0, 176, 178, 5, 34,
		0, 0, 177, 174, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0,
		179, 177, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 182, 1, 0, 0, 0, 181,
		179, 1, 0, 0, 0, 182, 183, 5, 34, 0, 0, 183, 64, 1, 0, 0, 0, 184, 186,
		7, 4, 0, 0, 185, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 185, 1, 0,
		0, 0, 187, 188, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 190, 6, 32, 0, 0,
		190, 66, 1, 0, 0, 0, 8, 0, 153, 159, 164, 170, 177, 179, 187, 1, 6, 0,
		0,
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

// GrammarLexerInit initializes any static state used to implement GrammarLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGrammarLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GrammarLexerInit() {
	staticData := &GrammarLexerLexerStaticData
	staticData.once.Do(grammarlexerLexerInit)
}

// NewGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewGrammarLexer(input antlr.CharStream) *GrammarLexer {
	GrammarLexerInit()
	l := new(GrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GrammarLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Grammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GrammarLexer tokens.
const (
	GrammarLexerT__0       = 1
	GrammarLexerT__1       = 2
	GrammarLexerT__2       = 3
	GrammarLexerT__3       = 4
	GrammarLexerT__4       = 5
	GrammarLexerT__5       = 6
	GrammarLexerT__6       = 7
	GrammarLexerT__7       = 8
	GrammarLexerT__8       = 9
	GrammarLexerT__9       = 10
	GrammarLexerT__10      = 11
	GrammarLexerT__11      = 12
	GrammarLexerT__12      = 13
	GrammarLexerT__13      = 14
	GrammarLexerT__14      = 15
	GrammarLexerT__15      = 16
	GrammarLexerT__16      = 17
	GrammarLexerT__17      = 18
	GrammarLexerT__18      = 19
	GrammarLexerT__19      = 20
	GrammarLexerT__20      = 21
	GrammarLexerT__21      = 22
	GrammarLexerT__22      = 23
	GrammarLexerT__23      = 24
	GrammarLexerT__24      = 25
	GrammarLexerT__25      = 26
	GrammarLexerT__26      = 27
	GrammarLexerT__27      = 28
	GrammarLexerDOUBLE     = 29
	GrammarLexerINT        = 30
	GrammarLexerID         = 31
	GrammarLexerSTRING     = 32
	GrammarLexerWHITESPACE = 33
)

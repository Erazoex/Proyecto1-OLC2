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
		"", "':'", "'='", "'?'", "'+='", "'-='", "'{'", "'}'", "'print'", "'('",
		"')'", "','", "'in'", "'...'", "'<'", "'>'", "'!'", "'-'", "'*'", "'/'",
		"'%'", "'+'", "'>='", "'<='", "'=='", "'!='", "'&&'", "'||'", "'true'",
		"'false'", "'nil'", "'Int'", "'Float'", "'Character'", "'String'", "'Bool'",
		"'var'", "'let'", "'if'", "'else'", "'switch'", "'case'", "'default'",
		"'while'", "'for'", "'vector'", "'func'", "'guard'", "'break'", "'return'",
		"'continue'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "T_INT", "T_FLOAT",
		"T_CHAR", "T_STRING", "T_BOOL", "VAR", "LET", "IF", "ELSE", "SWITCH",
		"CASE", "DEFAULT", "WHILE", "FOR", "VECTOR", "FUNC", "GUARD", "BREAK",
		"RETURN", "CONTINUE", "DOUBLE", "INT", "ID", "CHAR", "STRING", "MULTCOMMENT",
		"COMMENT", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T_INT", "T_FLOAT", "T_CHAR",
		"T_STRING", "T_BOOL", "VAR", "LET", "IF", "ELSE", "SWITCH", "CASE",
		"DEFAULT", "WHILE", "FOR", "VECTOR", "FUNC", "GUARD", "BREAK", "RETURN",
		"CONTINUE", "DOUBLE", "INT", "ID", "CHAR", "STRING", "MULTCOMMENT",
		"COMMENT", "WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 58, 400, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22,
		1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1,
		25, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43,
		1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1,
		45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1,
		48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49,
		1, 50, 4, 50, 321, 8, 50, 11, 50, 12, 50, 322, 1, 50, 1, 50, 4, 50, 327,
		8, 50, 11, 50, 12, 50, 328, 1, 51, 4, 51, 332, 8, 51, 11, 51, 12, 51, 333,
		1, 52, 1, 52, 5, 52, 338, 8, 52, 10, 52, 12, 52, 341, 9, 52, 1, 53, 1,
		53, 3, 53, 345, 8, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 54, 5, 54,
		353, 8, 54, 10, 54, 12, 54, 356, 9, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1,
		55, 1, 55, 1, 55, 1, 55, 5, 55, 366, 8, 55, 10, 55, 12, 55, 369, 9, 55,
		1, 55, 1, 55, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56, 1, 56, 5, 56, 380,
		8, 56, 10, 56, 12, 56, 383, 9, 56, 1, 56, 3, 56, 386, 8, 56, 1, 56, 1,
		56, 3, 56, 390, 8, 56, 1, 56, 1, 56, 1, 57, 4, 57, 395, 8, 57, 11, 57,
		12, 57, 396, 1, 57, 1, 57, 0, 0, 58, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15,
		31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24,
		49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33,
		67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42,
		85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51,
		103, 52, 105, 53, 107, 54, 109, 55, 111, 56, 113, 57, 115, 58, 1, 0, 9,
		1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95,
		95, 97, 122, 3, 0, 9, 10, 13, 13, 34, 34, 3, 0, 10, 10, 13, 13, 34, 34,
		2, 0, 42, 42, 47, 47, 1, 0, 47, 47, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10,
		13, 13, 32, 32, 412, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0,
		37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0,
		0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0,
		0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0,
		0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1,
		0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75,
		1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0,
		83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0,
		0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0,
		0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1,
		0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0,
		113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 1, 117, 1, 0, 0, 0, 3, 119, 1, 0,
		0, 0, 5, 121, 1, 0, 0, 0, 7, 123, 1, 0, 0, 0, 9, 126, 1, 0, 0, 0, 11, 129,
		1, 0, 0, 0, 13, 131, 1, 0, 0, 0, 15, 133, 1, 0, 0, 0, 17, 139, 1, 0, 0,
		0, 19, 141, 1, 0, 0, 0, 21, 143, 1, 0, 0, 0, 23, 145, 1, 0, 0, 0, 25, 148,
		1, 0, 0, 0, 27, 152, 1, 0, 0, 0, 29, 154, 1, 0, 0, 0, 31, 156, 1, 0, 0,
		0, 33, 158, 1, 0, 0, 0, 35, 160, 1, 0, 0, 0, 37, 162, 1, 0, 0, 0, 39, 164,
		1, 0, 0, 0, 41, 166, 1, 0, 0, 0, 43, 168, 1, 0, 0, 0, 45, 171, 1, 0, 0,
		0, 47, 174, 1, 0, 0, 0, 49, 177, 1, 0, 0, 0, 51, 180, 1, 0, 0, 0, 53, 183,
		1, 0, 0, 0, 55, 186, 1, 0, 0, 0, 57, 191, 1, 0, 0, 0, 59, 197, 1, 0, 0,
		0, 61, 201, 1, 0, 0, 0, 63, 205, 1, 0, 0, 0, 65, 211, 1, 0, 0, 0, 67, 221,
		1, 0, 0, 0, 69, 228, 1, 0, 0, 0, 71, 233, 1, 0, 0, 0, 73, 237, 1, 0, 0,
		0, 75, 241, 1, 0, 0, 0, 77, 244, 1, 0, 0, 0, 79, 249, 1, 0, 0, 0, 81, 256,
		1, 0, 0, 0, 83, 261, 1, 0, 0, 0, 85, 269, 1, 0, 0, 0, 87, 275, 1, 0, 0,
		0, 89, 279, 1, 0, 0, 0, 91, 286, 1, 0, 0, 0, 93, 291, 1, 0, 0, 0, 95, 297,
		1, 0, 0, 0, 97, 303, 1, 0, 0, 0, 99, 310, 1, 0, 0, 0, 101, 320, 1, 0, 0,
		0, 103, 331, 1, 0, 0, 0, 105, 335, 1, 0, 0, 0, 107, 342, 1, 0, 0, 0, 109,
		348, 1, 0, 0, 0, 111, 359, 1, 0, 0, 0, 113, 375, 1, 0, 0, 0, 115, 394,
		1, 0, 0, 0, 117, 118, 5, 58, 0, 0, 118, 2, 1, 0, 0, 0, 119, 120, 5, 61,
		0, 0, 120, 4, 1, 0, 0, 0, 121, 122, 5, 63, 0, 0, 122, 6, 1, 0, 0, 0, 123,
		124, 5, 43, 0, 0, 124, 125, 5, 61, 0, 0, 125, 8, 1, 0, 0, 0, 126, 127,
		5, 45, 0, 0, 127, 128, 5, 61, 0, 0, 128, 10, 1, 0, 0, 0, 129, 130, 5, 123,
		0, 0, 130, 12, 1, 0, 0, 0, 131, 132, 5, 125, 0, 0, 132, 14, 1, 0, 0, 0,
		133, 134, 5, 112, 0, 0, 134, 135, 5, 114, 0, 0, 135, 136, 5, 105, 0, 0,
		136, 137, 5, 110, 0, 0, 137, 138, 5, 116, 0, 0, 138, 16, 1, 0, 0, 0, 139,
		140, 5, 40, 0, 0, 140, 18, 1, 0, 0, 0, 141, 142, 5, 41, 0, 0, 142, 20,
		1, 0, 0, 0, 143, 144, 5, 44, 0, 0, 144, 22, 1, 0, 0, 0, 145, 146, 5, 105,
		0, 0, 146, 147, 5, 110, 0, 0, 147, 24, 1, 0, 0, 0, 148, 149, 5, 46, 0,
		0, 149, 150, 5, 46, 0, 0, 150, 151, 5, 46, 0, 0, 151, 26, 1, 0, 0, 0, 152,
		153, 5, 60, 0, 0, 153, 28, 1, 0, 0, 0, 154, 155, 5, 62, 0, 0, 155, 30,
		1, 0, 0, 0, 156, 157, 5, 33, 0, 0, 157, 32, 1, 0, 0, 0, 158, 159, 5, 45,
		0, 0, 159, 34, 1, 0, 0, 0, 160, 161, 5, 42, 0, 0, 161, 36, 1, 0, 0, 0,
		162, 163, 5, 47, 0, 0, 163, 38, 1, 0, 0, 0, 164, 165, 5, 37, 0, 0, 165,
		40, 1, 0, 0, 0, 166, 167, 5, 43, 0, 0, 167, 42, 1, 0, 0, 0, 168, 169, 5,
		62, 0, 0, 169, 170, 5, 61, 0, 0, 170, 44, 1, 0, 0, 0, 171, 172, 5, 60,
		0, 0, 172, 173, 5, 61, 0, 0, 173, 46, 1, 0, 0, 0, 174, 175, 5, 61, 0, 0,
		175, 176, 5, 61, 0, 0, 176, 48, 1, 0, 0, 0, 177, 178, 5, 33, 0, 0, 178,
		179, 5, 61, 0, 0, 179, 50, 1, 0, 0, 0, 180, 181, 5, 38, 0, 0, 181, 182,
		5, 38, 0, 0, 182, 52, 1, 0, 0, 0, 183, 184, 5, 124, 0, 0, 184, 185, 5,
		124, 0, 0, 185, 54, 1, 0, 0, 0, 186, 187, 5, 116, 0, 0, 187, 188, 5, 114,
		0, 0, 188, 189, 5, 117, 0, 0, 189, 190, 5, 101, 0, 0, 190, 56, 1, 0, 0,
		0, 191, 192, 5, 102, 0, 0, 192, 193, 5, 97, 0, 0, 193, 194, 5, 108, 0,
		0, 194, 195, 5, 115, 0, 0, 195, 196, 5, 101, 0, 0, 196, 58, 1, 0, 0, 0,
		197, 198, 5, 110, 0, 0, 198, 199, 5, 105, 0, 0, 199, 200, 5, 108, 0, 0,
		200, 60, 1, 0, 0, 0, 201, 202, 5, 73, 0, 0, 202, 203, 5, 110, 0, 0, 203,
		204, 5, 116, 0, 0, 204, 62, 1, 0, 0, 0, 205, 206, 5, 70, 0, 0, 206, 207,
		5, 108, 0, 0, 207, 208, 5, 111, 0, 0, 208, 209, 5, 97, 0, 0, 209, 210,
		5, 116, 0, 0, 210, 64, 1, 0, 0, 0, 211, 212, 5, 67, 0, 0, 212, 213, 5,
		104, 0, 0, 213, 214, 5, 97, 0, 0, 214, 215, 5, 114, 0, 0, 215, 216, 5,
		97, 0, 0, 216, 217, 5, 99, 0, 0, 217, 218, 5, 116, 0, 0, 218, 219, 5, 101,
		0, 0, 219, 220, 5, 114, 0, 0, 220, 66, 1, 0, 0, 0, 221, 222, 5, 83, 0,
		0, 222, 223, 5, 116, 0, 0, 223, 224, 5, 114, 0, 0, 224, 225, 5, 105, 0,
		0, 225, 226, 5, 110, 0, 0, 226, 227, 5, 103, 0, 0, 227, 68, 1, 0, 0, 0,
		228, 229, 5, 66, 0, 0, 229, 230, 5, 111, 0, 0, 230, 231, 5, 111, 0, 0,
		231, 232, 5, 108, 0, 0, 232, 70, 1, 0, 0, 0, 233, 234, 5, 118, 0, 0, 234,
		235, 5, 97, 0, 0, 235, 236, 5, 114, 0, 0, 236, 72, 1, 0, 0, 0, 237, 238,
		5, 108, 0, 0, 238, 239, 5, 101, 0, 0, 239, 240, 5, 116, 0, 0, 240, 74,
		1, 0, 0, 0, 241, 242, 5, 105, 0, 0, 242, 243, 5, 102, 0, 0, 243, 76, 1,
		0, 0, 0, 244, 245, 5, 101, 0, 0, 245, 246, 5, 108, 0, 0, 246, 247, 5, 115,
		0, 0, 247, 248, 5, 101, 0, 0, 248, 78, 1, 0, 0, 0, 249, 250, 5, 115, 0,
		0, 250, 251, 5, 119, 0, 0, 251, 252, 5, 105, 0, 0, 252, 253, 5, 116, 0,
		0, 253, 254, 5, 99, 0, 0, 254, 255, 5, 104, 0, 0, 255, 80, 1, 0, 0, 0,
		256, 257, 5, 99, 0, 0, 257, 258, 5, 97, 0, 0, 258, 259, 5, 115, 0, 0, 259,
		260, 5, 101, 0, 0, 260, 82, 1, 0, 0, 0, 261, 262, 5, 100, 0, 0, 262, 263,
		5, 101, 0, 0, 263, 264, 5, 102, 0, 0, 264, 265, 5, 97, 0, 0, 265, 266,
		5, 117, 0, 0, 266, 267, 5, 108, 0, 0, 267, 268, 5, 116, 0, 0, 268, 84,
		1, 0, 0, 0, 269, 270, 5, 119, 0, 0, 270, 271, 5, 104, 0, 0, 271, 272, 5,
		105, 0, 0, 272, 273, 5, 108, 0, 0, 273, 274, 5, 101, 0, 0, 274, 86, 1,
		0, 0, 0, 275, 276, 5, 102, 0, 0, 276, 277, 5, 111, 0, 0, 277, 278, 5, 114,
		0, 0, 278, 88, 1, 0, 0, 0, 279, 280, 5, 118, 0, 0, 280, 281, 5, 101, 0,
		0, 281, 282, 5, 99, 0, 0, 282, 283, 5, 116, 0, 0, 283, 284, 5, 111, 0,
		0, 284, 285, 5, 114, 0, 0, 285, 90, 1, 0, 0, 0, 286, 287, 5, 102, 0, 0,
		287, 288, 5, 117, 0, 0, 288, 289, 5, 110, 0, 0, 289, 290, 5, 99, 0, 0,
		290, 92, 1, 0, 0, 0, 291, 292, 5, 103, 0, 0, 292, 293, 5, 117, 0, 0, 293,
		294, 5, 97, 0, 0, 294, 295, 5, 114, 0, 0, 295, 296, 5, 100, 0, 0, 296,
		94, 1, 0, 0, 0, 297, 298, 5, 98, 0, 0, 298, 299, 5, 114, 0, 0, 299, 300,
		5, 101, 0, 0, 300, 301, 5, 97, 0, 0, 301, 302, 5, 107, 0, 0, 302, 96, 1,
		0, 0, 0, 303, 304, 5, 114, 0, 0, 304, 305, 5, 101, 0, 0, 305, 306, 5, 116,
		0, 0, 306, 307, 5, 117, 0, 0, 307, 308, 5, 114, 0, 0, 308, 309, 5, 110,
		0, 0, 309, 98, 1, 0, 0, 0, 310, 311, 5, 99, 0, 0, 311, 312, 5, 111, 0,
		0, 312, 313, 5, 110, 0, 0, 313, 314, 5, 116, 0, 0, 314, 315, 5, 105, 0,
		0, 315, 316, 5, 110, 0, 0, 316, 317, 5, 117, 0, 0, 317, 318, 5, 101, 0,
		0, 318, 100, 1, 0, 0, 0, 319, 321, 7, 0, 0, 0, 320, 319, 1, 0, 0, 0, 321,
		322, 1, 0, 0, 0, 322, 320, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323, 324,
		1, 0, 0, 0, 324, 326, 5, 46, 0, 0, 325, 327, 7, 0, 0, 0, 326, 325, 1, 0,
		0, 0, 327, 328, 1, 0, 0, 0, 328, 326, 1, 0, 0, 0, 328, 329, 1, 0, 0, 0,
		329, 102, 1, 0, 0, 0, 330, 332, 7, 0, 0, 0, 331, 330, 1, 0, 0, 0, 332,
		333, 1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334, 104,
		1, 0, 0, 0, 335, 339, 7, 1, 0, 0, 336, 338, 7, 2, 0, 0, 337, 336, 1, 0,
		0, 0, 338, 341, 1, 0, 0, 0, 339, 337, 1, 0, 0, 0, 339, 340, 1, 0, 0, 0,
		340, 106, 1, 0, 0, 0, 341, 339, 1, 0, 0, 0, 342, 344, 5, 34, 0, 0, 343,
		345, 9, 0, 0, 0, 344, 343, 1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345, 346,
		1, 0, 0, 0, 346, 347, 5, 34, 0, 0, 347, 108, 1, 0, 0, 0, 348, 354, 5, 34,
		0, 0, 349, 350, 5, 92, 0, 0, 350, 353, 7, 3, 0, 0, 351, 353, 8, 4, 0, 0,
		352, 349, 1, 0, 0, 0, 352, 351, 1, 0, 0, 0, 353, 356, 1, 0, 0, 0, 354,
		352, 1, 0, 0, 0, 354, 355, 1, 0, 0, 0, 355, 357, 1, 0, 0, 0, 356, 354,
		1, 0, 0, 0, 357, 358, 5, 34, 0, 0, 358, 110, 1, 0, 0, 0, 359, 360, 5, 47,
		0, 0, 360, 361, 5, 42, 0, 0, 361, 367, 1, 0, 0, 0, 362, 366, 8, 5, 0, 0,
		363, 364, 5, 42, 0, 0, 364, 366, 8, 6, 0, 0, 365, 362, 1, 0, 0, 0, 365,
		363, 1, 0, 0, 0, 366, 369, 1, 0, 0, 0, 367, 365, 1, 0, 0, 0, 367, 368,
		1, 0, 0, 0, 368, 370, 1, 0, 0, 0, 369, 367, 1, 0, 0, 0, 370, 371, 5, 42,
		0, 0, 371, 372, 5, 47, 0, 0, 372, 373, 1, 0, 0, 0, 373, 374, 6, 55, 0,
		0, 374, 112, 1, 0, 0, 0, 375, 376, 5, 47, 0, 0, 376, 377, 5, 47, 0, 0,
		377, 381, 1, 0, 0, 0, 378, 380, 8, 7, 0, 0, 379, 378, 1, 0, 0, 0, 380,
		383, 1, 0, 0, 0, 381, 379, 1, 0, 0, 0, 381, 382, 1, 0, 0, 0, 382, 389,
		1, 0, 0, 0, 383, 381, 1, 0, 0, 0, 384, 386, 5, 13, 0, 0, 385, 384, 1, 0,
		0, 0, 385, 386, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 390, 5, 10, 0, 0,
		388, 390, 5, 0, 0, 1, 389, 385, 1, 0, 0, 0, 389, 388, 1, 0, 0, 0, 390,
		391, 1, 0, 0, 0, 391, 392, 6, 56, 0, 0, 392, 114, 1, 0, 0, 0, 393, 395,
		7, 8, 0, 0, 394, 393, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 394, 1, 0,
		0, 0, 396, 397, 1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 399, 6, 57, 0, 0,
		399, 116, 1, 0, 0, 0, 14, 0, 322, 328, 333, 339, 344, 352, 354, 365, 367,
		381, 385, 389, 396, 1, 6, 0, 0,
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
	GrammarLexerT__0        = 1
	GrammarLexerT__1        = 2
	GrammarLexerT__2        = 3
	GrammarLexerT__3        = 4
	GrammarLexerT__4        = 5
	GrammarLexerT__5        = 6
	GrammarLexerT__6        = 7
	GrammarLexerT__7        = 8
	GrammarLexerT__8        = 9
	GrammarLexerT__9        = 10
	GrammarLexerT__10       = 11
	GrammarLexerT__11       = 12
	GrammarLexerT__12       = 13
	GrammarLexerT__13       = 14
	GrammarLexerT__14       = 15
	GrammarLexerT__15       = 16
	GrammarLexerT__16       = 17
	GrammarLexerT__17       = 18
	GrammarLexerT__18       = 19
	GrammarLexerT__19       = 20
	GrammarLexerT__20       = 21
	GrammarLexerT__21       = 22
	GrammarLexerT__22       = 23
	GrammarLexerT__23       = 24
	GrammarLexerT__24       = 25
	GrammarLexerT__25       = 26
	GrammarLexerT__26       = 27
	GrammarLexerT__27       = 28
	GrammarLexerT__28       = 29
	GrammarLexerT__29       = 30
	GrammarLexerT_INT       = 31
	GrammarLexerT_FLOAT     = 32
	GrammarLexerT_CHAR      = 33
	GrammarLexerT_STRING    = 34
	GrammarLexerT_BOOL      = 35
	GrammarLexerVAR         = 36
	GrammarLexerLET         = 37
	GrammarLexerIF          = 38
	GrammarLexerELSE        = 39
	GrammarLexerSWITCH      = 40
	GrammarLexerCASE        = 41
	GrammarLexerDEFAULT     = 42
	GrammarLexerWHILE       = 43
	GrammarLexerFOR         = 44
	GrammarLexerVECTOR      = 45
	GrammarLexerFUNC        = 46
	GrammarLexerGUARD       = 47
	GrammarLexerBREAK       = 48
	GrammarLexerRETURN      = 49
	GrammarLexerCONTINUE    = 50
	GrammarLexerDOUBLE      = 51
	GrammarLexerINT         = 52
	GrammarLexerID          = 53
	GrammarLexerCHAR        = 54
	GrammarLexerSTRING      = 55
	GrammarLexerMULTCOMMENT = 56
	GrammarLexerCOMMENT     = 57
	GrammarLexerWHITESPACE  = 58
)

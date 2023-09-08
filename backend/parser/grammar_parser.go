// Code generated from Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Grammar

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GrammarParser struct {
	*antlr.BaseParser
}

var GrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func grammarParserInit() {
	staticData := &GrammarParserStaticData
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
		"init", "block", "stmt", "breakstmt", "continuestmt", "returnstmt",
		"declstmt", "asignstmt", "incstmt", "decstmt", "ifstmt", "switchstmt",
		"switchcase", "printlnstmt", "exprparams", "whilestmt", "forstmt", "forrange",
		"guardstmt", "array", "array_def", "vartype", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 270, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 51, 8, 1, 10, 1,
		12, 1, 54, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 69, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1,
		5, 3, 5, 77, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 96, 8, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 3, 10, 134, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 4, 11, 140, 8,
		11, 11, 11, 12, 11, 141, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 3, 12, 154, 8, 12, 1, 13, 1, 13, 1, 13, 3, 13, 159,
		8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 168, 8,
		14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 3, 16, 192, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 20, 1, 20, 5, 20, 214, 8, 20, 10, 20, 12, 20, 217, 9, 20,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 224, 8, 21, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 3, 22, 242, 8, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 265, 8, 22, 10, 22, 12,
		22, 268, 9, 22, 1, 22, 0, 1, 44, 23, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 0, 7, 1, 0, 36, 37,
		1, 0, 28, 29, 1, 0, 18, 20, 2, 0, 17, 17, 21, 21, 2, 0, 15, 15, 22, 22,
		2, 0, 14, 14, 23, 23, 1, 0, 24, 25, 290, 0, 46, 1, 0, 0, 0, 2, 52, 1, 0,
		0, 0, 4, 68, 1, 0, 0, 0, 6, 70, 1, 0, 0, 0, 8, 72, 1, 0, 0, 0, 10, 74,
		1, 0, 0, 0, 12, 95, 1, 0, 0, 0, 14, 97, 1, 0, 0, 0, 16, 101, 1, 0, 0, 0,
		18, 105, 1, 0, 0, 0, 20, 133, 1, 0, 0, 0, 22, 135, 1, 0, 0, 0, 24, 153,
		1, 0, 0, 0, 26, 155, 1, 0, 0, 0, 28, 167, 1, 0, 0, 0, 30, 169, 1, 0, 0,
		0, 32, 191, 1, 0, 0, 0, 34, 193, 1, 0, 0, 0, 36, 197, 1, 0, 0, 0, 38, 204,
		1, 0, 0, 0, 40, 211, 1, 0, 0, 0, 42, 223, 1, 0, 0, 0, 44, 241, 1, 0, 0,
		0, 46, 47, 3, 2, 1, 0, 47, 48, 5, 0, 0, 1, 48, 1, 1, 0, 0, 0, 49, 51, 3,
		4, 2, 0, 50, 49, 1, 0, 0, 0, 51, 54, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 52,
		53, 1, 0, 0, 0, 53, 3, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 55, 69, 3, 12, 6,
		0, 56, 69, 3, 14, 7, 0, 57, 69, 3, 16, 8, 0, 58, 69, 3, 18, 9, 0, 59, 69,
		3, 20, 10, 0, 60, 69, 3, 22, 11, 0, 61, 69, 3, 26, 13, 0, 62, 69, 3, 30,
		15, 0, 63, 69, 3, 32, 16, 0, 64, 69, 3, 36, 18, 0, 65, 69, 3, 6, 3, 0,
		66, 69, 3, 8, 4, 0, 67, 69, 3, 10, 5, 0, 68, 55, 1, 0, 0, 0, 68, 56, 1,
		0, 0, 0, 68, 57, 1, 0, 0, 0, 68, 58, 1, 0, 0, 0, 68, 59, 1, 0, 0, 0, 68,
		60, 1, 0, 0, 0, 68, 61, 1, 0, 0, 0, 68, 62, 1, 0, 0, 0, 68, 63, 1, 0, 0,
		0, 68, 64, 1, 0, 0, 0, 68, 65, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 67,
		1, 0, 0, 0, 69, 5, 1, 0, 0, 0, 70, 71, 5, 48, 0, 0, 71, 7, 1, 0, 0, 0,
		72, 73, 5, 50, 0, 0, 73, 9, 1, 0, 0, 0, 74, 76, 5, 49, 0, 0, 75, 77, 3,
		44, 22, 0, 76, 75, 1, 0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 11, 1, 0, 0, 0,
		78, 79, 7, 0, 0, 0, 79, 80, 5, 53, 0, 0, 80, 81, 5, 1, 0, 0, 81, 82, 3,
		42, 21, 0, 82, 83, 5, 2, 0, 0, 83, 84, 3, 44, 22, 0, 84, 96, 1, 0, 0, 0,
		85, 86, 7, 0, 0, 0, 86, 87, 5, 53, 0, 0, 87, 88, 5, 2, 0, 0, 88, 96, 3,
		44, 22, 0, 89, 90, 7, 0, 0, 0, 90, 91, 5, 53, 0, 0, 91, 92, 5, 1, 0, 0,
		92, 93, 3, 42, 21, 0, 93, 94, 5, 3, 0, 0, 94, 96, 1, 0, 0, 0, 95, 78, 1,
		0, 0, 0, 95, 85, 1, 0, 0, 0, 95, 89, 1, 0, 0, 0, 96, 13, 1, 0, 0, 0, 97,
		98, 5, 53, 0, 0, 98, 99, 5, 2, 0, 0, 99, 100, 3, 44, 22, 0, 100, 15, 1,
		0, 0, 0, 101, 102, 5, 53, 0, 0, 102, 103, 5, 4, 0, 0, 103, 104, 3, 44,
		22, 0, 104, 17, 1, 0, 0, 0, 105, 106, 5, 53, 0, 0, 106, 107, 5, 5, 0, 0,
		107, 108, 3, 44, 22, 0, 108, 19, 1, 0, 0, 0, 109, 110, 5, 38, 0, 0, 110,
		111, 3, 44, 22, 0, 111, 112, 5, 6, 0, 0, 112, 113, 3, 2, 1, 0, 113, 114,
		5, 7, 0, 0, 114, 134, 1, 0, 0, 0, 115, 116, 5, 38, 0, 0, 116, 117, 3, 44,
		22, 0, 117, 118, 5, 6, 0, 0, 118, 119, 3, 2, 1, 0, 119, 120, 5, 7, 0, 0,
		120, 121, 5, 39, 0, 0, 121, 122, 5, 6, 0, 0, 122, 123, 3, 2, 1, 0, 123,
		124, 5, 7, 0, 0, 124, 134, 1, 0, 0, 0, 125, 126, 5, 38, 0, 0, 126, 127,
		3, 44, 22, 0, 127, 128, 5, 6, 0, 0, 128, 129, 3, 2, 1, 0, 129, 130, 5,
		7, 0, 0, 130, 131, 5, 39, 0, 0, 131, 132, 3, 20, 10, 0, 132, 134, 1, 0,
		0, 0, 133, 109, 1, 0, 0, 0, 133, 115, 1, 0, 0, 0, 133, 125, 1, 0, 0, 0,
		134, 21, 1, 0, 0, 0, 135, 136, 5, 40, 0, 0, 136, 137, 3, 44, 22, 0, 137,
		139, 5, 6, 0, 0, 138, 140, 3, 24, 12, 0, 139, 138, 1, 0, 0, 0, 140, 141,
		1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 1, 0,
		0, 0, 143, 144, 5, 7, 0, 0, 144, 23, 1, 0, 0, 0, 145, 146, 5, 41, 0, 0,
		146, 147, 3, 44, 22, 0, 147, 148, 5, 1, 0, 0, 148, 149, 3, 2, 1, 0, 149,
		154, 1, 0, 0, 0, 150, 151, 5, 42, 0, 0, 151, 152, 5, 1, 0, 0, 152, 154,
		3, 2, 1, 0, 153, 145, 1, 0, 0, 0, 153, 150, 1, 0, 0, 0, 154, 25, 1, 0,
		0, 0, 155, 156, 5, 8, 0, 0, 156, 158, 5, 9, 0, 0, 157, 159, 3, 28, 14,
		0, 158, 157, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160,
		161, 5, 10, 0, 0, 161, 27, 1, 0, 0, 0, 162, 163, 3, 44, 22, 0, 163, 164,
		5, 11, 0, 0, 164, 165, 3, 28, 14, 0, 165, 168, 1, 0, 0, 0, 166, 168, 3,
		44, 22, 0, 167, 162, 1, 0, 0, 0, 167, 166, 1, 0, 0, 0, 168, 29, 1, 0, 0,
		0, 169, 170, 5, 43, 0, 0, 170, 171, 3, 44, 22, 0, 171, 172, 5, 6, 0, 0,
		172, 173, 3, 2, 1, 0, 173, 174, 5, 7, 0, 0, 174, 31, 1, 0, 0, 0, 175, 176,
		5, 44, 0, 0, 176, 177, 5, 53, 0, 0, 177, 178, 5, 12, 0, 0, 178, 179, 3,
		44, 22, 0, 179, 180, 5, 6, 0, 0, 180, 181, 3, 2, 1, 0, 181, 182, 5, 7,
		0, 0, 182, 192, 1, 0, 0, 0, 183, 184, 5, 44, 0, 0, 184, 185, 5, 53, 0,
		0, 185, 186, 5, 12, 0, 0, 186, 187, 3, 34, 17, 0, 187, 188, 5, 6, 0, 0,
		188, 189, 3, 2, 1, 0, 189, 190, 5, 7, 0, 0, 190, 192, 1, 0, 0, 0, 191,
		175, 1, 0, 0, 0, 191, 183, 1, 0, 0, 0, 192, 33, 1, 0, 0, 0, 193, 194, 3,
		44, 22, 0, 194, 195, 5, 13, 0, 0, 195, 196, 3, 44, 22, 0, 196, 35, 1, 0,
		0, 0, 197, 198, 5, 47, 0, 0, 198, 199, 3, 44, 22, 0, 199, 200, 5, 39, 0,
		0, 200, 201, 5, 6, 0, 0, 201, 202, 3, 2, 1, 0, 202, 203, 5, 7, 0, 0, 203,
		37, 1, 0, 0, 0, 204, 205, 5, 45, 0, 0, 205, 206, 5, 14, 0, 0, 206, 207,
		3, 42, 21, 0, 207, 208, 5, 15, 0, 0, 208, 209, 5, 53, 0, 0, 209, 210, 3,
		40, 20, 0, 210, 39, 1, 0, 0, 0, 211, 215, 5, 2, 0, 0, 212, 214, 3, 44,
		22, 0, 213, 212, 1, 0, 0, 0, 214, 217, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0,
		215, 216, 1, 0, 0, 0, 216, 41, 1, 0, 0, 0, 217, 215, 1, 0, 0, 0, 218, 224,
		5, 31, 0, 0, 219, 224, 5, 32, 0, 0, 220, 224, 5, 33, 0, 0, 221, 224, 5,
		34, 0, 0, 222, 224, 5, 35, 0, 0, 223, 218, 1, 0, 0, 0, 223, 219, 1, 0,
		0, 0, 223, 220, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 223, 222, 1, 0, 0, 0,
		224, 43, 1, 0, 0, 0, 225, 226, 6, 22, -1, 0, 226, 227, 5, 16, 0, 0, 227,
		242, 3, 44, 22, 17, 228, 229, 5, 17, 0, 0, 229, 242, 3, 44, 22, 16, 230,
		231, 5, 9, 0, 0, 231, 232, 3, 44, 22, 0, 232, 233, 5, 10, 0, 0, 233, 242,
		1, 0, 0, 0, 234, 242, 5, 52, 0, 0, 235, 242, 5, 51, 0, 0, 236, 242, 5,
		53, 0, 0, 237, 242, 5, 54, 0, 0, 238, 242, 5, 55, 0, 0, 239, 242, 7, 1,
		0, 0, 240, 242, 5, 30, 0, 0, 241, 225, 1, 0, 0, 0, 241, 228, 1, 0, 0, 0,
		241, 230, 1, 0, 0, 0, 241, 234, 1, 0, 0, 0, 241, 235, 1, 0, 0, 0, 241,
		236, 1, 0, 0, 0, 241, 237, 1, 0, 0, 0, 241, 238, 1, 0, 0, 0, 241, 239,
		1, 0, 0, 0, 241, 240, 1, 0, 0, 0, 242, 266, 1, 0, 0, 0, 243, 244, 10, 15,
		0, 0, 244, 245, 7, 2, 0, 0, 245, 265, 3, 44, 22, 16, 246, 247, 10, 14,
		0, 0, 247, 248, 7, 3, 0, 0, 248, 265, 3, 44, 22, 15, 249, 250, 10, 13,
		0, 0, 250, 251, 7, 4, 0, 0, 251, 265, 3, 44, 22, 14, 252, 253, 10, 12,
		0, 0, 253, 254, 7, 5, 0, 0, 254, 265, 3, 44, 22, 13, 255, 256, 10, 11,
		0, 0, 256, 257, 7, 6, 0, 0, 257, 265, 3, 44, 22, 12, 258, 259, 10, 10,
		0, 0, 259, 260, 5, 26, 0, 0, 260, 265, 3, 44, 22, 11, 261, 262, 10, 9,
		0, 0, 262, 263, 5, 27, 0, 0, 263, 265, 3, 44, 22, 10, 264, 243, 1, 0, 0,
		0, 264, 246, 1, 0, 0, 0, 264, 249, 1, 0, 0, 0, 264, 252, 1, 0, 0, 0, 264,
		255, 1, 0, 0, 0, 264, 258, 1, 0, 0, 0, 264, 261, 1, 0, 0, 0, 265, 268,
		1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 45, 1, 0,
		0, 0, 268, 266, 1, 0, 0, 0, 15, 52, 68, 76, 95, 133, 141, 153, 158, 167,
		191, 215, 223, 241, 264, 266,
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

// GrammarParserInit initializes any static state used to implement GrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GrammarParserInit() {
	staticData := &GrammarParserStaticData
	staticData.once.Do(grammarParserInit)
}

// NewGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGrammarParser(input antlr.TokenStream) *GrammarParser {
	GrammarParserInit()
	this := new(GrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Grammar.g4"

	return this
}

// GrammarParser tokens.
const (
	GrammarParserEOF         = antlr.TokenEOF
	GrammarParserT__0        = 1
	GrammarParserT__1        = 2
	GrammarParserT__2        = 3
	GrammarParserT__3        = 4
	GrammarParserT__4        = 5
	GrammarParserT__5        = 6
	GrammarParserT__6        = 7
	GrammarParserT__7        = 8
	GrammarParserT__8        = 9
	GrammarParserT__9        = 10
	GrammarParserT__10       = 11
	GrammarParserT__11       = 12
	GrammarParserT__12       = 13
	GrammarParserT__13       = 14
	GrammarParserT__14       = 15
	GrammarParserT__15       = 16
	GrammarParserT__16       = 17
	GrammarParserT__17       = 18
	GrammarParserT__18       = 19
	GrammarParserT__19       = 20
	GrammarParserT__20       = 21
	GrammarParserT__21       = 22
	GrammarParserT__22       = 23
	GrammarParserT__23       = 24
	GrammarParserT__24       = 25
	GrammarParserT__25       = 26
	GrammarParserT__26       = 27
	GrammarParserT__27       = 28
	GrammarParserT__28       = 29
	GrammarParserT__29       = 30
	GrammarParserT_INT       = 31
	GrammarParserT_FLOAT     = 32
	GrammarParserT_CHAR      = 33
	GrammarParserT_STRING    = 34
	GrammarParserT_BOOL      = 35
	GrammarParserVAR         = 36
	GrammarParserLET         = 37
	GrammarParserIF          = 38
	GrammarParserELSE        = 39
	GrammarParserSWITCH      = 40
	GrammarParserCASE        = 41
	GrammarParserDEFAULT     = 42
	GrammarParserWHILE       = 43
	GrammarParserFOR         = 44
	GrammarParserVECTOR      = 45
	GrammarParserFUNC        = 46
	GrammarParserGUARD       = 47
	GrammarParserBREAK       = 48
	GrammarParserRETURN      = 49
	GrammarParserCONTINUE    = 50
	GrammarParserDOUBLE      = 51
	GrammarParserINT         = 52
	GrammarParserID          = 53
	GrammarParserCHAR        = 54
	GrammarParserSTRING      = 55
	GrammarParserMULTCOMMENT = 56
	GrammarParserCOMMENT     = 57
	GrammarParserWHITESPACE  = 58
)

// GrammarParser rules.
const (
	GrammarParserRULE_init         = 0
	GrammarParserRULE_block        = 1
	GrammarParserRULE_stmt         = 2
	GrammarParserRULE_breakstmt    = 3
	GrammarParserRULE_continuestmt = 4
	GrammarParserRULE_returnstmt   = 5
	GrammarParserRULE_declstmt     = 6
	GrammarParserRULE_asignstmt    = 7
	GrammarParserRULE_incstmt      = 8
	GrammarParserRULE_decstmt      = 9
	GrammarParserRULE_ifstmt       = 10
	GrammarParserRULE_switchstmt   = 11
	GrammarParserRULE_switchcase   = 12
	GrammarParserRULE_printlnstmt  = 13
	GrammarParserRULE_exprparams   = 14
	GrammarParserRULE_whilestmt    = 15
	GrammarParserRULE_forstmt      = 16
	GrammarParserRULE_forrange     = 17
	GrammarParserRULE_guardstmt    = 18
	GrammarParserRULE_array        = 19
	GrammarParserRULE_array_def    = 20
	GrammarParserRULE_vartype      = 21
	GrammarParserRULE_expr         = 22
)

// IInitContext is an interface to support dynamic dispatch.
type IInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

	// IsInitContext differentiates from other interfaces.
	IsInitContext()
}

type InitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitContext() *InitContext {
	var p = new(InitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_init
	return p
}

func InitEmptyInitContext(p *InitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_init
}

func (*InitContext) IsInitContext() {}

func NewInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitContext {
	var p = new(InitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_init

	return p
}

func (s *InitContext) GetParser() antlr.Parser { return s.parser }

func (s *InitContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *InitContext) EOF() antlr.TerminalNode {
	return s.GetToken(GrammarParserEOF, 0)
}

func (s *InitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Init() (localctx IInitContext) {
	localctx = NewInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GrammarParserRULE_init)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Block()
	}
	{
		p.SetState(47)
		p.Match(GrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GrammarParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&11146230407102720) != 0 {
		{
			p.SetState(49)
			p.Stmt()
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declstmt() IDeclstmtContext
	Asignstmt() IAsignstmtContext
	Incstmt() IIncstmtContext
	Decstmt() IDecstmtContext
	Ifstmt() IIfstmtContext
	Switchstmt() ISwitchstmtContext
	Printlnstmt() IPrintlnstmtContext
	Whilestmt() IWhilestmtContext
	Forstmt() IForstmtContext
	Guardstmt() IGuardstmtContext
	Breakstmt() IBreakstmtContext
	Continuestmt() IContinuestmtContext
	Returnstmt() IReturnstmtContext

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Declstmt() IDeclstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclstmtContext)
}

func (s *StmtContext) Asignstmt() IAsignstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignstmtContext)
}

func (s *StmtContext) Incstmt() IIncstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncstmtContext)
}

func (s *StmtContext) Decstmt() IDecstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecstmtContext)
}

func (s *StmtContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *StmtContext) Switchstmt() ISwitchstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchstmtContext)
}

func (s *StmtContext) Printlnstmt() IPrintlnstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintlnstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintlnstmtContext)
}

func (s *StmtContext) Whilestmt() IWhilestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhilestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhilestmtContext)
}

func (s *StmtContext) Forstmt() IForstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForstmtContext)
}

func (s *StmtContext) Guardstmt() IGuardstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardstmtContext)
}

func (s *StmtContext) Breakstmt() IBreakstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakstmtContext)
}

func (s *StmtContext) Continuestmt() IContinuestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinuestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinuestmtContext)
}

func (s *StmtContext) Returnstmt() IReturnstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnstmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GrammarParserRULE_stmt)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.Declstmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.Asignstmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.Incstmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(58)
			p.Decstmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(59)
			p.Ifstmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(60)
			p.Switchstmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(61)
			p.Printlnstmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(62)
			p.Whilestmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(63)
			p.Forstmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(64)
			p.Guardstmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(65)
			p.Breakstmt()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(66)
			p.Continuestmt()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(67)
			p.Returnstmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakstmtContext is an interface to support dynamic dispatch.
type IBreakstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreakstmtContext differentiates from other interfaces.
	IsBreakstmtContext()
}

type BreakstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakstmtContext() *BreakstmtContext {
	var p = new(BreakstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_breakstmt
	return p
}

func InitEmptyBreakstmtContext(p *BreakstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_breakstmt
}

func (*BreakstmtContext) IsBreakstmtContext() {}

func NewBreakstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakstmtContext {
	var p = new(BreakstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_breakstmt

	return p
}

func (s *BreakstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakstmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(GrammarParserBREAK, 0)
}

func (s *BreakstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBreakstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Breakstmt() (localctx IBreakstmtContext) {
	localctx = NewBreakstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GrammarParserRULE_breakstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(GrammarParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinuestmtContext is an interface to support dynamic dispatch.
type IContinuestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinuestmtContext differentiates from other interfaces.
	IsContinuestmtContext()
}

type ContinuestmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinuestmtContext() *ContinuestmtContext {
	var p = new(ContinuestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_continuestmt
	return p
}

func InitEmptyContinuestmtContext(p *ContinuestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_continuestmt
}

func (*ContinuestmtContext) IsContinuestmtContext() {}

func NewContinuestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinuestmtContext {
	var p = new(ContinuestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_continuestmt

	return p
}

func (s *ContinuestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinuestmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(GrammarParserCONTINUE, 0)
}

func (s *ContinuestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinuestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinuestmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitContinuestmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Continuestmt() (localctx IContinuestmtContext) {
	localctx = NewContinuestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GrammarParserRULE_continuestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(GrammarParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnstmtContext is an interface to support dynamic dispatch.
type IReturnstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expr() IExprContext

	// IsReturnstmtContext differentiates from other interfaces.
	IsReturnstmtContext()
}

type ReturnstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnstmtContext() *ReturnstmtContext {
	var p = new(ReturnstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_returnstmt
	return p
}

func InitEmptyReturnstmtContext(p *ReturnstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_returnstmt
}

func (*ReturnstmtContext) IsReturnstmtContext() {}

func NewReturnstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnstmtContext {
	var p = new(ReturnstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_returnstmt

	return p
}

func (s *ReturnstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnstmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(GrammarParserRETURN, 0)
}

func (s *ReturnstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReturnstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitReturnstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Returnstmt() (localctx IReturnstmtContext) {
	localctx = NewReturnstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GrammarParserRULE_returnstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(GrammarParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(75)
			p.expr(0)
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclstmtContext is an interface to support dynamic dispatch.
type IDeclstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclstmtContext differentiates from other interfaces.
	IsDeclstmtContext()
}

type DeclstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclstmtContext() *DeclstmtContext {
	var p = new(DeclstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_declstmt
	return p
}

func InitEmptyDeclstmtContext(p *DeclstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_declstmt
}

func (*DeclstmtContext) IsDeclstmtContext() {}

func NewDeclstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclstmtContext {
	var p = new(DeclstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_declstmt

	return p
}

func (s *DeclstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclstmtContext) CopyAll(ctx *DeclstmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DeclstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclstmtWithTypeAndExprContext struct {
	DeclstmtContext
	vtype antlr.Token
}

func NewDeclstmtWithTypeAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclstmtWithTypeAndExprContext {
	var p = new(DeclstmtWithTypeAndExprContext)

	InitEmptyDeclstmtContext(&p.DeclstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclstmtContext))

	return p
}

func (s *DeclstmtWithTypeAndExprContext) GetVtype() antlr.Token { return s.vtype }

func (s *DeclstmtWithTypeAndExprContext) SetVtype(v antlr.Token) { s.vtype = v }

func (s *DeclstmtWithTypeAndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclstmtWithTypeAndExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *DeclstmtWithTypeAndExprContext) Vartype() IVartypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVartypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVartypeContext)
}

func (s *DeclstmtWithTypeAndExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclstmtWithTypeAndExprContext) VAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVAR, 0)
}

func (s *DeclstmtWithTypeAndExprContext) LET() antlr.TerminalNode {
	return s.GetToken(GrammarParserLET, 0)
}

func (s *DeclstmtWithTypeAndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitDeclstmtWithTypeAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclstmtWithExprContext struct {
	DeclstmtContext
	vtype antlr.Token
}

func NewDeclstmtWithExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclstmtWithExprContext {
	var p = new(DeclstmtWithExprContext)

	InitEmptyDeclstmtContext(&p.DeclstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclstmtContext))

	return p
}

func (s *DeclstmtWithExprContext) GetVtype() antlr.Token { return s.vtype }

func (s *DeclstmtWithExprContext) SetVtype(v antlr.Token) { s.vtype = v }

func (s *DeclstmtWithExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclstmtWithExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *DeclstmtWithExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclstmtWithExprContext) VAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVAR, 0)
}

func (s *DeclstmtWithExprContext) LET() antlr.TerminalNode {
	return s.GetToken(GrammarParserLET, 0)
}

func (s *DeclstmtWithExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitDeclstmtWithExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclstmtWithTypeContext struct {
	DeclstmtContext
	vtype antlr.Token
}

func NewDeclstmtWithTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclstmtWithTypeContext {
	var p = new(DeclstmtWithTypeContext)

	InitEmptyDeclstmtContext(&p.DeclstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclstmtContext))

	return p
}

func (s *DeclstmtWithTypeContext) GetVtype() antlr.Token { return s.vtype }

func (s *DeclstmtWithTypeContext) SetVtype(v antlr.Token) { s.vtype = v }

func (s *DeclstmtWithTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclstmtWithTypeContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *DeclstmtWithTypeContext) Vartype() IVartypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVartypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVartypeContext)
}

func (s *DeclstmtWithTypeContext) VAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVAR, 0)
}

func (s *DeclstmtWithTypeContext) LET() antlr.TerminalNode {
	return s.GetToken(GrammarParserLET, 0)
}

func (s *DeclstmtWithTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitDeclstmtWithType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Declstmt() (localctx IDeclstmtContext) {
	localctx = NewDeclstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GrammarParserRULE_declstmt)
	var _la int

	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclstmtWithTypeAndExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DeclstmtWithTypeAndExprContext).vtype = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GrammarParserVAR || _la == GrammarParserLET) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DeclstmtWithTypeAndExprContext).vtype = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(79)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(80)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(81)
			p.Vartype()
		}
		{
			p.SetState(82)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.expr(0)
		}

	case 2:
		localctx = NewDeclstmtWithExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DeclstmtWithExprContext).vtype = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GrammarParserVAR || _la == GrammarParserLET) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DeclstmtWithExprContext).vtype = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(86)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)
			p.expr(0)
		}

	case 3:
		localctx = NewDeclstmtWithTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(89)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*DeclstmtWithTypeContext).vtype = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == GrammarParserVAR || _la == GrammarParserLET) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*DeclstmtWithTypeContext).vtype = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(90)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(91)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.Vartype()
		}
		{
			p.SetState(93)
			p.Match(GrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAsignstmtContext is an interface to support dynamic dispatch.
type IAsignstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expr() IExprContext

	// IsAsignstmtContext differentiates from other interfaces.
	IsAsignstmtContext()
}

type AsignstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignstmtContext() *AsignstmtContext {
	var p = new(AsignstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_asignstmt
	return p
}

func InitEmptyAsignstmtContext(p *AsignstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_asignstmt
}

func (*AsignstmtContext) IsAsignstmtContext() {}

func NewAsignstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignstmtContext {
	var p = new(AsignstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_asignstmt

	return p
}

func (s *AsignstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *AsignstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsignstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitAsignstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Asignstmt() (localctx IAsignstmtContext) {
	localctx = NewAsignstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GrammarParserRULE_asignstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.Match(GrammarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIncstmtContext is an interface to support dynamic dispatch.
type IIncstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expr() IExprContext

	// IsIncstmtContext differentiates from other interfaces.
	IsIncstmtContext()
}

type IncstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncstmtContext() *IncstmtContext {
	var p = new(IncstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_incstmt
	return p
}

func InitEmptyIncstmtContext(p *IncstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_incstmt
}

func (*IncstmtContext) IsIncstmtContext() {}

func NewIncstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncstmtContext {
	var p = new(IncstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_incstmt

	return p
}

func (s *IncstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IncstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *IncstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IncstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIncstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Incstmt() (localctx IIncstmtContext) {
	localctx = NewIncstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GrammarParserRULE_incstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(102)
		p.Match(GrammarParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDecstmtContext is an interface to support dynamic dispatch.
type IDecstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expr() IExprContext

	// IsDecstmtContext differentiates from other interfaces.
	IsDecstmtContext()
}

type DecstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecstmtContext() *DecstmtContext {
	var p = new(DecstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_decstmt
	return p
}

func InitEmptyDecstmtContext(p *DecstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_decstmt
}

func (*DecstmtContext) IsDecstmtContext() {}

func NewDecstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecstmtContext {
	var p = new(DecstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_decstmt

	return p
}

func (s *DecstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DecstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *DecstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DecstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitDecstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Decstmt() (localctx IDecstmtContext) {
	localctx = NewDecstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GrammarParserRULE_decstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(GrammarParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfstmtContext is an interface to support dynamic dispatch.
type IIfstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstmtContext() *IfstmtContext {
	var p = new(IfstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_ifstmt
	return p
}

func InitEmptyIfstmtContext(p *IfstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_ifstmt
}

func (*IfstmtContext) IsIfstmtContext() {}

func NewIfstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstmtContext {
	var p = new(IfstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_ifstmt

	return p
}

func (s *IfstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstmtContext) CopyAll(ctx *IfstmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IfstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfWithElseIfContext struct {
	IfstmtContext
}

func NewIfWithElseIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfWithElseIfContext {
	var p = new(IfWithElseIfContext)

	InitEmptyIfstmtContext(&p.IfstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfstmtContext))

	return p
}

func (s *IfWithElseIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfWithElseIfContext) IF() antlr.TerminalNode {
	return s.GetToken(GrammarParserIF, 0)
}

func (s *IfWithElseIfContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfWithElseIfContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfWithElseIfContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GrammarParserELSE, 0)
}

func (s *IfWithElseIfContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *IfWithElseIfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIfWithElseIf(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfWithElseContext struct {
	IfstmtContext
	trueCondition  IBlockContext
	falseCondition IBlockContext
}

func NewIfWithElseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfWithElseContext {
	var p = new(IfWithElseContext)

	InitEmptyIfstmtContext(&p.IfstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfstmtContext))

	return p
}

func (s *IfWithElseContext) GetTrueCondition() IBlockContext { return s.trueCondition }

func (s *IfWithElseContext) GetFalseCondition() IBlockContext { return s.falseCondition }

func (s *IfWithElseContext) SetTrueCondition(v IBlockContext) { s.trueCondition = v }

func (s *IfWithElseContext) SetFalseCondition(v IBlockContext) { s.falseCondition = v }

func (s *IfWithElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfWithElseContext) IF() antlr.TerminalNode {
	return s.GetToken(GrammarParserIF, 0)
}

func (s *IfWithElseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfWithElseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GrammarParserELSE, 0)
}

func (s *IfWithElseContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfWithElseContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfWithElseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIfWithElse(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfSimpleContext struct {
	IfstmtContext
}

func NewIfSimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfSimpleContext {
	var p = new(IfSimpleContext)

	InitEmptyIfstmtContext(&p.IfstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfstmtContext))

	return p
}

func (s *IfSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfSimpleContext) IF() antlr.TerminalNode {
	return s.GetToken(GrammarParserIF, 0)
}

func (s *IfSimpleContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfSimpleContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfSimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIfSimple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Ifstmt() (localctx IIfstmtContext) {
	localctx = NewIfstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GrammarParserRULE_ifstmt)
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfSimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.Match(GrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.expr(0)
		}
		{
			p.SetState(111)
			p.Match(GrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Block()
		}
		{
			p.SetState(113)
			p.Match(GrammarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewIfWithElseContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.Match(GrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.expr(0)
		}
		{
			p.SetState(117)
			p.Match(GrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(118)

			var _x = p.Block()

			localctx.(*IfWithElseContext).trueCondition = _x
		}
		{
			p.SetState(119)
			p.Match(GrammarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Match(GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Match(GrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)

			var _x = p.Block()

			localctx.(*IfWithElseContext).falseCondition = _x
		}
		{
			p.SetState(123)
			p.Match(GrammarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewIfWithElseIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			p.Match(GrammarParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.expr(0)
		}
		{
			p.SetState(127)
			p.Match(GrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(128)
			p.Block()
		}
		{
			p.SetState(129)
			p.Match(GrammarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.Match(GrammarParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Ifstmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchstmtContext is an interface to support dynamic dispatch.
type ISwitchstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SWITCH() antlr.TerminalNode
	Expr() IExprContext
	AllSwitchcase() []ISwitchcaseContext
	Switchcase(i int) ISwitchcaseContext

	// IsSwitchstmtContext differentiates from other interfaces.
	IsSwitchstmtContext()
}

type SwitchstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchstmtContext() *SwitchstmtContext {
	var p = new(SwitchstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_switchstmt
	return p
}

func InitEmptySwitchstmtContext(p *SwitchstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_switchstmt
}

func (*SwitchstmtContext) IsSwitchstmtContext() {}

func NewSwitchstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchstmtContext {
	var p = new(SwitchstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_switchstmt

	return p
}

func (s *SwitchstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchstmtContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(GrammarParserSWITCH, 0)
}

func (s *SwitchstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchstmtContext) AllSwitchcase() []ISwitchcaseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISwitchcaseContext); ok {
			len++
		}
	}

	tst := make([]ISwitchcaseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISwitchcaseContext); ok {
			tst[i] = t.(ISwitchcaseContext)
			i++
		}
	}

	return tst
}

func (s *SwitchstmtContext) Switchcase(i int) ISwitchcaseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchcaseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchcaseContext)
}

func (s *SwitchstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitSwitchstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Switchstmt() (localctx ISwitchstmtContext) {
	localctx = NewSwitchstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GrammarParserRULE_switchstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(GrammarParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.expr(0)
	}
	{
		p.SetState(137)
		p.Match(GrammarParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GrammarParserCASE || _la == GrammarParserDEFAULT {
		{
			p.SetState(138)
			p.Switchcase()
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(143)
		p.Match(GrammarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchcaseContext is an interface to support dynamic dispatch.
type ISwitchcaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCasetype returns the casetype token.
	GetCasetype() antlr.Token

	// SetCasetype sets the casetype token.
	SetCasetype(antlr.Token)

	// Getter signatures
	Expr() IExprContext
	Block() IBlockContext
	CASE() antlr.TerminalNode
	DEFAULT() antlr.TerminalNode

	// IsSwitchcaseContext differentiates from other interfaces.
	IsSwitchcaseContext()
}

type SwitchcaseContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	casetype antlr.Token
}

func NewEmptySwitchcaseContext() *SwitchcaseContext {
	var p = new(SwitchcaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_switchcase
	return p
}

func InitEmptySwitchcaseContext(p *SwitchcaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_switchcase
}

func (*SwitchcaseContext) IsSwitchcaseContext() {}

func NewSwitchcaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchcaseContext {
	var p = new(SwitchcaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_switchcase

	return p
}

func (s *SwitchcaseContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchcaseContext) GetCasetype() antlr.Token { return s.casetype }

func (s *SwitchcaseContext) SetCasetype(v antlr.Token) { s.casetype = v }

func (s *SwitchcaseContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchcaseContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SwitchcaseContext) CASE() antlr.TerminalNode {
	return s.GetToken(GrammarParserCASE, 0)
}

func (s *SwitchcaseContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(GrammarParserDEFAULT, 0)
}

func (s *SwitchcaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchcaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchcaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitSwitchcase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Switchcase() (localctx ISwitchcaseContext) {
	localctx = NewSwitchcaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GrammarParserRULE_switchcase)
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserCASE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)

			var _m = p.Match(GrammarParserCASE)

			localctx.(*SwitchcaseContext).casetype = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.expr(0)
		}
		{
			p.SetState(147)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(148)
			p.Block()
		}

	case GrammarParserDEFAULT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)

			var _m = p.Match(GrammarParserDEFAULT)

			localctx.(*SwitchcaseContext).casetype = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.Block()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintlnstmtContext is an interface to support dynamic dispatch.
type IPrintlnstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exprparams() IExprparamsContext

	// IsPrintlnstmtContext differentiates from other interfaces.
	IsPrintlnstmtContext()
}

type PrintlnstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintlnstmtContext() *PrintlnstmtContext {
	var p = new(PrintlnstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_printlnstmt
	return p
}

func InitEmptyPrintlnstmtContext(p *PrintlnstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_printlnstmt
}

func (*PrintlnstmtContext) IsPrintlnstmtContext() {}

func NewPrintlnstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintlnstmtContext {
	var p = new(PrintlnstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_printlnstmt

	return p
}

func (s *PrintlnstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintlnstmtContext) Exprparams() IExprparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprparamsContext)
}

func (s *PrintlnstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintlnstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitPrintlnstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Printlnstmt() (localctx IPrintlnstmtContext) {
	localctx = NewPrintlnstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GrammarParserRULE_printlnstmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(GrammarParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(GrammarParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&69805796103488000) != 0 {
		{
			p.SetState(157)
			p.Exprparams()
		}

	}
	{
		p.SetState(160)
		p.Match(GrammarParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprparamsContext is an interface to support dynamic dispatch.
type IExprparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprparamsContext differentiates from other interfaces.
	IsExprparamsContext()
}

type ExprparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprparamsContext() *ExprparamsContext {
	var p = new(ExprparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprparams
	return p
}

func InitEmptyExprparamsContext(p *ExprparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_exprparams
}

func (*ExprparamsContext) IsExprparamsContext() {}

func NewExprparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprparamsContext {
	var p = new(ExprparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_exprparams

	return p
}

func (s *ExprparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprparamsContext) CopyAll(ctx *ExprparamsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExprWithParamsContext struct {
	ExprparamsContext
}

func NewExprWithParamsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprWithParamsContext {
	var p = new(ExprWithParamsContext)

	InitEmptyExprparamsContext(&p.ExprparamsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprparamsContext))

	return p
}

func (s *ExprWithParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprWithParamsContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprWithParamsContext) Exprparams() IExprparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprparamsContext)
}

func (s *ExprWithParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitExprWithParams(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprWithParamContext struct {
	ExprparamsContext
}

func NewExprWithParamContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprWithParamContext {
	var p = new(ExprWithParamContext)

	InitEmptyExprparamsContext(&p.ExprparamsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprparamsContext))

	return p
}

func (s *ExprWithParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprWithParamContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprWithParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitExprWithParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Exprparams() (localctx IExprparamsContext) {
	localctx = NewExprparamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GrammarParserRULE_exprparams)
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprWithParamsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.expr(0)
		}
		{
			p.SetState(163)
			p.Match(GrammarParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(164)
			p.Exprparams()
		}

	case 2:
		localctx = NewExprWithParamContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhilestmtContext is an interface to support dynamic dispatch.
type IWhilestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	Expr() IExprContext
	Block() IBlockContext

	// IsWhilestmtContext differentiates from other interfaces.
	IsWhilestmtContext()
}

type WhilestmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhilestmtContext() *WhilestmtContext {
	var p = new(WhilestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_whilestmt
	return p
}

func InitEmptyWhilestmtContext(p *WhilestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_whilestmt
}

func (*WhilestmtContext) IsWhilestmtContext() {}

func NewWhilestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhilestmtContext {
	var p = new(WhilestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_whilestmt

	return p
}

func (s *WhilestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhilestmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(GrammarParserWHILE, 0)
}

func (s *WhilestmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhilestmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhilestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhilestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhilestmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitWhilestmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Whilestmt() (localctx IWhilestmtContext) {
	localctx = NewWhilestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GrammarParserRULE_whilestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(GrammarParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.expr(0)
	}
	{
		p.SetState(171)
		p.Match(GrammarParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.Block()
	}
	{
		p.SetState(173)
		p.Match(GrammarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForstmtContext is an interface to support dynamic dispatch.
type IForstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsForstmtContext differentiates from other interfaces.
	IsForstmtContext()
}

type ForstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForstmtContext() *ForstmtContext {
	var p = new(ForstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forstmt
	return p
}

func InitEmptyForstmtContext(p *ForstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forstmt
}

func (*ForstmtContext) IsForstmtContext() {}

func NewForstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForstmtContext {
	var p = new(ForstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_forstmt

	return p
}

func (s *ForstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForstmtContext) CopyAll(ctx *ForstmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ForstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForWithRangeContext struct {
	ForstmtContext
}

func NewForWithRangeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForWithRangeContext {
	var p = new(ForWithRangeContext)

	InitEmptyForstmtContext(&p.ForstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForstmtContext))

	return p
}

func (s *ForWithRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForWithRangeContext) FOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserFOR, 0)
}

func (s *ForWithRangeContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *ForWithRangeContext) Forrange() IForrangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForrangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForrangeContext)
}

func (s *ForWithRangeContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForWithRangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitForWithRange(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForWithExprContext struct {
	ForstmtContext
}

func NewForWithExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForWithExprContext {
	var p = new(ForWithExprContext)

	InitEmptyForstmtContext(&p.ForstmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*ForstmtContext))

	return p
}

func (s *ForWithExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForWithExprContext) FOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserFOR, 0)
}

func (s *ForWithExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *ForWithExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForWithExprContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForWithExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitForWithExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Forstmt() (localctx IForstmtContext) {
	localctx = NewForstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GrammarParserRULE_forstmt)
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForWithExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)
			p.Match(GrammarParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Match(GrammarParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.expr(0)
		}
		{
			p.SetState(179)
			p.Match(GrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.Block()
		}
		{
			p.SetState(181)
			p.Match(GrammarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewForWithRangeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
			p.Match(GrammarParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.Match(GrammarParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Forrange()
		}
		{
			p.SetState(187)
			p.Match(GrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Block()
		}
		{
			p.SetState(189)
			p.Match(GrammarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForrangeContext is an interface to support dynamic dispatch.
type IForrangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBeginsWith returns the beginsWith rule contexts.
	GetBeginsWith() IExprContext

	// GetEndsWith returns the endsWith rule contexts.
	GetEndsWith() IExprContext

	// SetBeginsWith sets the beginsWith rule contexts.
	SetBeginsWith(IExprContext)

	// SetEndsWith sets the endsWith rule contexts.
	SetEndsWith(IExprContext)

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsForrangeContext differentiates from other interfaces.
	IsForrangeContext()
}

type ForrangeContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	beginsWith IExprContext
	endsWith   IExprContext
}

func NewEmptyForrangeContext() *ForrangeContext {
	var p = new(ForrangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forrange
	return p
}

func InitEmptyForrangeContext(p *ForrangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_forrange
}

func (*ForrangeContext) IsForrangeContext() {}

func NewForrangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForrangeContext {
	var p = new(ForrangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_forrange

	return p
}

func (s *ForrangeContext) GetParser() antlr.Parser { return s.parser }

func (s *ForrangeContext) GetBeginsWith() IExprContext { return s.beginsWith }

func (s *ForrangeContext) GetEndsWith() IExprContext { return s.endsWith }

func (s *ForrangeContext) SetBeginsWith(v IExprContext) { s.beginsWith = v }

func (s *ForrangeContext) SetEndsWith(v IExprContext) { s.endsWith = v }

func (s *ForrangeContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ForrangeContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForrangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForrangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForrangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitForrange(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Forrange() (localctx IForrangeContext) {
	localctx = NewForrangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GrammarParserRULE_forrange)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)

		var _x = p.expr(0)

		localctx.(*ForrangeContext).beginsWith = _x
	}
	{
		p.SetState(194)
		p.Match(GrammarParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)

		var _x = p.expr(0)

		localctx.(*ForrangeContext).endsWith = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGuardstmtContext is an interface to support dynamic dispatch.
type IGuardstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GUARD() antlr.TerminalNode
	Expr() IExprContext
	ELSE() antlr.TerminalNode
	Block() IBlockContext

	// IsGuardstmtContext differentiates from other interfaces.
	IsGuardstmtContext()
}

type GuardstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuardstmtContext() *GuardstmtContext {
	var p = new(GuardstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_guardstmt
	return p
}

func InitEmptyGuardstmtContext(p *GuardstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_guardstmt
}

func (*GuardstmtContext) IsGuardstmtContext() {}

func NewGuardstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardstmtContext {
	var p = new(GuardstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_guardstmt

	return p
}

func (s *GuardstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardstmtContext) GUARD() antlr.TerminalNode {
	return s.GetToken(GrammarParserGUARD, 0)
}

func (s *GuardstmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GuardstmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GrammarParserELSE, 0)
}

func (s *GuardstmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *GuardstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GuardstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitGuardstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Guardstmt() (localctx IGuardstmtContext) {
	localctx = NewGuardstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GrammarParserRULE_guardstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(GrammarParserGUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
		p.expr(0)
	}
	{
		p.SetState(199)
		p.Match(GrammarParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		p.Match(GrammarParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(201)
		p.Block()
	}
	{
		p.SetState(202)
		p.Match(GrammarParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VECTOR() antlr.TerminalNode
	Vartype() IVartypeContext
	ID() antlr.TerminalNode
	Array_def() IArray_defContext

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_array
	return p
}

func InitEmptyArrayContext(p *ArrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_array
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) VECTOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserVECTOR, 0)
}

func (s *ArrayContext) Vartype() IVartypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVartypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVartypeContext)
}

func (s *ArrayContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *ArrayContext) Array_def() IArray_defContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArray_defContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArray_defContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GrammarParserRULE_array)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(GrammarParserVECTOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Match(GrammarParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Vartype()
	}
	{
		p.SetState(207)
		p.Match(GrammarParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Array_def()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArray_defContext is an interface to support dynamic dispatch.
type IArray_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsArray_defContext differentiates from other interfaces.
	IsArray_defContext()
}

type Array_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_defContext() *Array_defContext {
	var p = new(Array_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_array_def
	return p
}

func InitEmptyArray_defContext(p *Array_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_array_def
}

func (*Array_defContext) IsArray_defContext() {}

func NewArray_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_defContext {
	var p = new(Array_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_array_def

	return p
}

func (s *Array_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_defContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Array_defContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Array_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitArray_def(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Array_def() (localctx IArray_defContext) {
	localctx = NewArray_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GrammarParserRULE_array_def)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.Match(GrammarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&69805796103488000) != 0 {
		{
			p.SetState(212)
			p.expr(0)
		}

		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVartypeContext is an interface to support dynamic dispatch.
type IVartypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVartypeContext differentiates from other interfaces.
	IsVartypeContext()
}

type VartypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVartypeContext() *VartypeContext {
	var p = new(VartypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_vartype
	return p
}

func InitEmptyVartypeContext(p *VartypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_vartype
}

func (*VartypeContext) IsVartypeContext() {}

func NewVartypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VartypeContext {
	var p = new(VartypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_vartype

	return p
}

func (s *VartypeContext) GetParser() antlr.Parser { return s.parser }

func (s *VartypeContext) CopyAll(ctx *VartypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VartypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VartypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrimitiveTypeContext struct {
	VartypeContext
	tipo antlr.Token
}

func NewPrimitiveTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)

	InitEmptyVartypeContext(&p.VartypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*VartypeContext))

	return p
}

func (s *PrimitiveTypeContext) GetTipo() antlr.Token { return s.tipo }

func (s *PrimitiveTypeContext) SetTipo(v antlr.Token) { s.tipo = v }

func (s *PrimitiveTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveTypeContext) T_INT() antlr.TerminalNode {
	return s.GetToken(GrammarParserT_INT, 0)
}

func (s *PrimitiveTypeContext) T_FLOAT() antlr.TerminalNode {
	return s.GetToken(GrammarParserT_FLOAT, 0)
}

func (s *PrimitiveTypeContext) T_CHAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserT_CHAR, 0)
}

func (s *PrimitiveTypeContext) T_STRING() antlr.TerminalNode {
	return s.GetToken(GrammarParserT_STRING, 0)
}

func (s *PrimitiveTypeContext) T_BOOL() antlr.TerminalNode {
	return s.GetToken(GrammarParserT_BOOL, 0)
}

func (s *PrimitiveTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitPrimitiveType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Vartype() (localctx IVartypeContext) {
	localctx = NewVartypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GrammarParserRULE_vartype)
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT_INT:
		localctx = NewPrimitiveTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(218)

			var _m = p.Match(GrammarParserT_INT)

			localctx.(*PrimitiveTypeContext).tipo = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserT_FLOAT:
		localctx = NewPrimitiveTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(219)

			var _m = p.Match(GrammarParserT_FLOAT)

			localctx.(*PrimitiveTypeContext).tipo = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserT_CHAR:
		localctx = NewPrimitiveTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(220)

			var _m = p.Match(GrammarParserT_CHAR)

			localctx.(*PrimitiveTypeContext).tipo = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserT_STRING:
		localctx = NewPrimitiveTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(221)

			var _m = p.Match(GrammarParserT_STRING)

			localctx.(*PrimitiveTypeContext).tipo = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserT_BOOL:
		localctx = NewPrimitiveTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(222)

			var _m = p.Match(GrammarParserT_BOOL)

			localctx.(*PrimitiveTypeContext).tipo = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DoubleExprContext struct {
	ExprContext
}

func NewDoubleExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoubleExprContext {
	var p = new(DoubleExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *DoubleExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleExprContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(GrammarParserDOUBLE, 0)
}

func (s *DoubleExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitDoubleExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolExprContext struct {
	ExprContext
}

func NewBoolExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExprContext {
	var p = new(BoolExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitBoolExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExprContext struct {
	ExprContext
}

func NewIdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExprContext {
	var p = new(IdExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExprContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParExprContext struct {
	ExprContext
}

func NewParExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParExprContext {
	var p = new(ParExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitParExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StrExprContext struct {
	ExprContext
}

func NewStrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StrExprContext {
	var p = new(StrExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(GrammarParserSTRING, 0)
}

func (s *StrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	ExprContext
	right IExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRight() IExprContext { return s.right }

func (s *NotExprContext) SetRight(v IExprContext) { s.right = v }

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntExprContext struct {
	ExprContext
}

func NewIntExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntExprContext {
	var p = new(IntExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IntExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntExprContext) INT() antlr.TerminalNode {
	return s.GetToken(GrammarParserINT, 0)
}

func (s *IntExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIntExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilExprContext struct {
	ExprContext
}

func NewNilExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilExprContext {
	var p = new(NilExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NilExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitNilExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewOpExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpExprContext {
	var p = new(OpExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OpExprContext) GetOp() antlr.Token { return s.op }

func (s *OpExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *OpExprContext) GetLeft() IExprContext { return s.left }

func (s *OpExprContext) GetRight() IExprContext { return s.right }

func (s *OpExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *OpExprContext) SetRight(v IExprContext) { s.right = v }

func (s *OpExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OpExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OpExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitOpExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExprContext struct {
	ExprContext
	right IExprContext
}

func NewUnaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExprContext {
	var p = new(UnaryExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *UnaryExprContext) GetRight() IExprContext { return s.right }

func (s *UnaryExprContext) SetRight(v IExprContext) { s.right = v }

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitUnaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CharExprContext struct {
	ExprContext
}

func NewCharExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CharExprContext {
	var p = new(CharExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CharExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharExprContext) CHAR() antlr.TerminalNode {
	return s.GetToken(GrammarParserCHAR, 0)
}

func (s *CharExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitCharExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *GrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, GrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT__15:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(226)
			p.Match(GrammarParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)

			var _x = p.expr(17)

			localctx.(*NotExprContext).right = _x
		}

	case GrammarParserT__16:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(228)
			p.Match(GrammarParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)

			var _x = p.expr(16)

			localctx.(*UnaryExprContext).right = _x
		}

	case GrammarParserT__8:
		localctx = NewParExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(230)
			p.Match(GrammarParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.expr(0)
		}
		{
			p.SetState(232)
			p.Match(GrammarParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserINT:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(234)
			p.Match(GrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserDOUBLE:
		localctx = NewDoubleExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(235)
			p.Match(GrammarParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserID:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(236)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserCHAR:
		localctx = NewCharExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(237)
			p.Match(GrammarParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserSTRING:
		localctx = NewStrExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(238)
			p.Match(GrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserT__27, GrammarParserT__28:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(239)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GrammarParserT__27 || _la == GrammarParserT__28) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case GrammarParserT__29:
		localctx = NewNilExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(240)
			p.Match(GrammarParserT__29)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(264)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(243)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(244)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1835008) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(245)

					var _x = p.expr(16)

					localctx.(*OpExprContext).right = _x
				}

			case 2:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(246)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(247)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__16 || _la == GrammarParserT__20) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(248)

					var _x = p.expr(15)

					localctx.(*OpExprContext).right = _x
				}

			case 3:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(249)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(250)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__14 || _la == GrammarParserT__21) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(251)

					var _x = p.expr(14)

					localctx.(*OpExprContext).right = _x
				}

			case 4:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(252)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(253)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__13 || _la == GrammarParserT__22) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(254)

					var _x = p.expr(13)

					localctx.(*OpExprContext).right = _x
				}

			case 5:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(255)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(256)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__23 || _la == GrammarParserT__24) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(257)

					var _x = p.expr(12)

					localctx.(*OpExprContext).right = _x
				}

			case 6:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(258)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(259)

					var _m = p.Match(GrammarParserT__25)

					localctx.(*OpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(260)

					var _x = p.expr(11)

					localctx.(*OpExprContext).right = _x
				}

			case 7:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(261)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(262)

					var _m = p.Match(GrammarParserT__26)

					localctx.(*OpExprContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(263)

					var _x = p.expr(10)

					localctx.(*OpExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *GrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 22:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
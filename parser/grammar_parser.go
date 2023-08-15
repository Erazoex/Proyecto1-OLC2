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
		"", "':'", "'='", "'?'", "'!'", "'-'", "'*'", "'/'", "'%'", "'+'", "'>='",
		"'>'", "'<='", "'<'", "'=='", "'!='", "'&&'", "'||'", "'('", "')'",
		"'true'", "'false'", "'nil'", "'Int'", "'Float'", "'Character'", "'String'",
		"'Bool'", "'var'", "'let'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "T_INT", "T_FLOAT", "T_CHAR", "T_STRING", "T_BOOL",
		"VAR", "LET", "DOUBLE", "INT", "ID", "CHAR", "STRING", "MULTCOMMENT",
		"COMMENT", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"init", "block", "stmt", "declstmt", "asignstmt", "vartype", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 37, 102, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 0, 1, 1, 5, 1, 19, 8, 1, 10,
		1, 12, 1, 22, 9, 1, 1, 2, 1, 2, 3, 2, 26, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 3, 3, 45, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 3, 5, 56, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 74, 8, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 97, 8, 6, 10, 6, 12, 6, 100,
		9, 6, 1, 6, 0, 1, 12, 7, 0, 2, 4, 6, 8, 10, 12, 0, 7, 1, 0, 28, 29, 1,
		0, 20, 21, 1, 0, 6, 8, 2, 0, 5, 5, 9, 9, 1, 0, 10, 11, 1, 0, 12, 13, 1,
		0, 14, 15, 118, 0, 14, 1, 0, 0, 0, 2, 20, 1, 0, 0, 0, 4, 25, 1, 0, 0, 0,
		6, 44, 1, 0, 0, 0, 8, 46, 1, 0, 0, 0, 10, 55, 1, 0, 0, 0, 12, 73, 1, 0,
		0, 0, 14, 15, 3, 2, 1, 0, 15, 16, 5, 0, 0, 1, 16, 1, 1, 0, 0, 0, 17, 19,
		3, 4, 2, 0, 18, 17, 1, 0, 0, 0, 19, 22, 1, 0, 0, 0, 20, 18, 1, 0, 0, 0,
		20, 21, 1, 0, 0, 0, 21, 3, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 23, 26, 3, 6,
		3, 0, 24, 26, 3, 8, 4, 0, 25, 23, 1, 0, 0, 0, 25, 24, 1, 0, 0, 0, 26, 5,
		1, 0, 0, 0, 27, 28, 7, 0, 0, 0, 28, 29, 5, 32, 0, 0, 29, 30, 5, 1, 0, 0,
		30, 31, 3, 10, 5, 0, 31, 32, 5, 2, 0, 0, 32, 33, 3, 12, 6, 0, 33, 45, 1,
		0, 0, 0, 34, 35, 7, 0, 0, 0, 35, 36, 5, 32, 0, 0, 36, 37, 5, 2, 0, 0, 37,
		45, 3, 12, 6, 0, 38, 39, 7, 0, 0, 0, 39, 40, 5, 32, 0, 0, 40, 41, 5, 1,
		0, 0, 41, 42, 3, 10, 5, 0, 42, 43, 5, 3, 0, 0, 43, 45, 1, 0, 0, 0, 44,
		27, 1, 0, 0, 0, 44, 34, 1, 0, 0, 0, 44, 38, 1, 0, 0, 0, 45, 7, 1, 0, 0,
		0, 46, 47, 5, 32, 0, 0, 47, 48, 5, 2, 0, 0, 48, 49, 3, 12, 6, 0, 49, 9,
		1, 0, 0, 0, 50, 56, 5, 23, 0, 0, 51, 56, 5, 24, 0, 0, 52, 56, 5, 25, 0,
		0, 53, 56, 5, 26, 0, 0, 54, 56, 5, 27, 0, 0, 55, 50, 1, 0, 0, 0, 55, 51,
		1, 0, 0, 0, 55, 52, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 54, 1, 0, 0, 0,
		56, 11, 1, 0, 0, 0, 57, 58, 6, 6, -1, 0, 58, 59, 5, 4, 0, 0, 59, 74, 3,
		12, 6, 17, 60, 61, 5, 5, 0, 0, 61, 74, 3, 12, 6, 16, 62, 63, 5, 18, 0,
		0, 63, 64, 3, 12, 6, 0, 64, 65, 5, 19, 0, 0, 65, 74, 1, 0, 0, 0, 66, 74,
		5, 31, 0, 0, 67, 74, 5, 30, 0, 0, 68, 74, 5, 32, 0, 0, 69, 74, 5, 33, 0,
		0, 70, 74, 5, 34, 0, 0, 71, 74, 7, 1, 0, 0, 72, 74, 5, 22, 0, 0, 73, 57,
		1, 0, 0, 0, 73, 60, 1, 0, 0, 0, 73, 62, 1, 0, 0, 0, 73, 66, 1, 0, 0, 0,
		73, 67, 1, 0, 0, 0, 73, 68, 1, 0, 0, 0, 73, 69, 1, 0, 0, 0, 73, 70, 1,
		0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 72, 1, 0, 0, 0, 74, 98, 1, 0, 0, 0, 75,
		76, 10, 15, 0, 0, 76, 77, 7, 2, 0, 0, 77, 97, 3, 12, 6, 16, 78, 79, 10,
		14, 0, 0, 79, 80, 7, 3, 0, 0, 80, 97, 3, 12, 6, 15, 81, 82, 10, 13, 0,
		0, 82, 83, 7, 4, 0, 0, 83, 97, 3, 12, 6, 14, 84, 85, 10, 12, 0, 0, 85,
		86, 7, 5, 0, 0, 86, 97, 3, 12, 6, 13, 87, 88, 10, 11, 0, 0, 88, 89, 7,
		6, 0, 0, 89, 97, 3, 12, 6, 12, 90, 91, 10, 10, 0, 0, 91, 92, 5, 16, 0,
		0, 92, 97, 3, 12, 6, 11, 93, 94, 10, 9, 0, 0, 94, 95, 5, 17, 0, 0, 95,
		97, 3, 12, 6, 10, 96, 75, 1, 0, 0, 0, 96, 78, 1, 0, 0, 0, 96, 81, 1, 0,
		0, 0, 96, 84, 1, 0, 0, 0, 96, 87, 1, 0, 0, 0, 96, 90, 1, 0, 0, 0, 96, 93,
		1, 0, 0, 0, 97, 100, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0,
		99, 13, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 7, 20, 25, 44, 55, 73, 96, 98,
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
	GrammarParserT_INT       = 23
	GrammarParserT_FLOAT     = 24
	GrammarParserT_CHAR      = 25
	GrammarParserT_STRING    = 26
	GrammarParserT_BOOL      = 27
	GrammarParserVAR         = 28
	GrammarParserLET         = 29
	GrammarParserDOUBLE      = 30
	GrammarParserINT         = 31
	GrammarParserID          = 32
	GrammarParserCHAR        = 33
	GrammarParserSTRING      = 34
	GrammarParserMULTCOMMENT = 35
	GrammarParserCOMMENT     = 36
	GrammarParserWHITESPACE  = 37
)

// GrammarParser rules.
const (
	GrammarParserRULE_init      = 0
	GrammarParserRULE_block     = 1
	GrammarParserRULE_stmt      = 2
	GrammarParserRULE_declstmt  = 3
	GrammarParserRULE_asignstmt = 4
	GrammarParserRULE_vartype   = 5
	GrammarParserRULE_expr      = 6
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
		p.SetState(14)
		p.Block()
	}
	{
		p.SetState(15)
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
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&5100273664) != 0 {
		{
			p.SetState(17)
			p.Stmt()
		}

		p.SetState(22)
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
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserVAR, GrammarParserLET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)
			p.Declstmt()
		}

	case GrammarParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.Asignstmt()
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
	p.EnterRule(localctx, 6, GrammarParserRULE_declstmt)
	var _la int

	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclstmtWithTypeAndExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)

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
			p.SetState(28)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(29)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(30)
			p.Vartype()
		}
		{
			p.SetState(31)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(32)
			p.expr(0)
		}

	case 2:
		localctx = NewDeclstmtWithExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(34)

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
			p.SetState(35)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(36)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.expr(0)
		}

	case 3:
		localctx = NewDeclstmtWithTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(38)

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
			p.SetState(39)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.Match(GrammarParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.Vartype()
		}
		{
			p.SetState(42)
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
	p.EnterRule(localctx, 8, GrammarParserRULE_asignstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.Match(GrammarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
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
	p.EnterRule(localctx, 10, GrammarParserRULE_vartype)
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT_INT:
		localctx = NewPrimitiveTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)

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
			p.SetState(51)

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
			p.SetState(52)

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
			p.SetState(53)

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
			p.SetState(54)

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
	_startState := 12
	p.EnterRecursionRule(localctx, 12, GrammarParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT__3:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(58)
			p.Match(GrammarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(59)

			var _x = p.expr(17)

			localctx.(*NotExprContext).right = _x
		}

	case GrammarParserT__4:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(60)
			p.Match(GrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)

			var _x = p.expr(16)

			localctx.(*UnaryExprContext).right = _x
		}

	case GrammarParserT__17:
		localctx = NewParExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(62)
			p.Match(GrammarParserT__17)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.expr(0)
		}
		{
			p.SetState(64)
			p.Match(GrammarParserT__18)
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
			p.SetState(66)
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
			p.SetState(67)
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
			p.SetState(68)
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
			p.SetState(69)
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
			p.SetState(70)
			p.Match(GrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserT__19, GrammarParserT__20:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(71)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GrammarParserT__19 || _la == GrammarParserT__20) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case GrammarParserT__21:
		localctx = NewNilExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(72)
			p.Match(GrammarParserT__21)
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
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(75)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(76)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&448) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(77)

					var _x = p.expr(16)

					localctx.(*OpExprContext).right = _x
				}

			case 2:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(78)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(79)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__4 || _la == GrammarParserT__8) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(80)

					var _x = p.expr(15)

					localctx.(*OpExprContext).right = _x
				}

			case 3:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(82)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__9 || _la == GrammarParserT__10) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(83)

					var _x = p.expr(14)

					localctx.(*OpExprContext).right = _x
				}

			case 4:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(85)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__11 || _la == GrammarParserT__12) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(86)

					var _x = p.expr(13)

					localctx.(*OpExprContext).right = _x
				}

			case 5:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(88)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__13 || _la == GrammarParserT__14) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(89)

					var _x = p.expr(12)

					localctx.(*OpExprContext).right = _x
				}

			case 6:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(91)
					p.Match(GrammarParserT__15)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(92)

					var _x = p.expr(11)

					localctx.(*OpExprContext).right = _x
				}

			case 7:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_expr)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(94)
					p.Match(GrammarParserT__16)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(95)

					var _x = p.expr(10)

					localctx.(*OpExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
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
	case 6:
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

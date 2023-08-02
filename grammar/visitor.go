package grammar

import (
	"proyecto2/parser"

	"github.com/antlr4-go/antlr/v4"
)

// Funcion Principal del objeto Visitor
func (v *Visitor) Visit(tree antlr.ParseTree) Value {
	switch val := tree.(type) {
	case *parser.InitContext:
		return v.VisitInit(val)
	case *parser.BlockContext:
		return v.VisitBlock(val)
	case *parser.StmtContext:
		return v.VisitStmt(val)
	case *parser.DeclstmtContext:
		return v.VisitDeclstmt(val)
	case *parser.PrintstmtContext:
		return v.VisitPrintstmt(val)
	case *parser.IfstmtContext:
		return v.VisitIfstmt(val)
	case *parser.ForstmtContext:
		return v.VisitForstmt(val)
	case *parser.IntExprContext:
		return v.VisitIntExpr(val)
	case *parser.DoubleExprContext:
		return v.VisitDoubleExpr(val)
	case *parser.NilExprContext:
		return v.VisitNilExpr(val)
	case *parser.IdExprContext:
		return v.VisitIdExpr(val)
	case *parser.StrExprContext:
		return v.VisitStrExpr(val)
	case *parser.BoolExprContext:
		return v.VisitBoolExpr(val)
	case *parser.NotExprContext:
		return v.VisitNotExpr(val)
	case *parser.UnaryExprContext:
		return v.VisitUnaryExpr(val)
	}
	return Value{value: true}
}

func (v *Visitor) VisitInit(ctx *parser.InitContext) Value {
	return v.Visit(ctx.Block())
}

func Execute() {
	input := `
	let varInt = 4
	let varDouble = 3.1416
	let varString = "Hello, world!"
	let varBoolean = true
	let varNil = nil
	var b = 18
	if varBoolean {
		print(varInt)
		print(varDouble)
		print(varString)
		print(varBoolean)
		print(varNil)
	}
	`
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewGrammarLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGrammarParser(tokenStream)
	p.BuildParseTrees = true
	tree := p.Init()
	eval := Visitor{
		memory: make(map[string]Value),
	}
	eval.Visit(tree)
}

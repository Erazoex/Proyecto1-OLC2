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
	case *parser.CharExprContext:
		return v.VisitCharExpr(val)
	case *parser.BoolExprContext:
		return v.VisitBoolExpr(val)
	case *parser.NotExprContext:
		return v.VisitNotExpr(val)
	case *parser.UnaryExprContext:
		return v.VisitUnaryExpr(val)
	case *parser.ParExprContext:
		return v.VisitParExpr(val)
	case *parser.OpExprContext:
		return v.VisitOpExpr(val)
	case *parser.ExprWithParamContext:
		return v.VisitExprWithParam(val)
	case *parser.ExprWithParamsContext:
		return v.VisitExprWithParams(val)
	case *parser.PrimitiveTypeContext:
		return v.VisitPrimitiveType(val)
	case *parser.DeclstmtWithTypeAndExprContext:
		return v.VisitDeclstmtWithTypeAndExpr(val)
	case *parser.DeclstmtWithExprContext:
		return v.VisitDeclstmtWithExpr(val)
	case *parser.DeclstmtWithTypeContext:
		return v.VisitDeclstmtWithType(val)
	case *parser.AsignstmtContext:
		return v.VisitAsignstmt(val)
	case *parser.IncstmtContext:
		return v.VisitIncstmt(val)
	case *parser.DecstmtContext:
		return v.VisitDecstmt(val)
	case *parser.IfSimpleContext:
		return v.VisitIfSimple(val)
	case *parser.IfWithElseContext:
		return v.VisitIfWithElse(val)
	case *parser.IfWithElseIfContext:
		return v.VisitIfWithElseIf(val)
	case *parser.SwitchstmtContext:
		return v.VisitSwitchstmt(val)
	case *parser.PrintlnstmtContext:
		return v.VisitPrintlnstmt(val)
	case *parser.ForWithRangeContext:
		return v.VisitForWithRange(val)
	}
	return Value{value: "la funcion todavia no se ha implementado"}
}

func (v *Visitor) VisitInit(ctx *parser.InitContext) Value {
	return v.Visit(ctx.Block())
}

func Execute() {
	input := `
	for item in 0...10+2 {
		print(item)
	}
	`
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewGrammarLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGrammarParser(tokenStream)
	p.BuildParseTrees = true
	tree := p.Init()
	eval := Visitor{
		environment: NewEnvironment(nil),
	}
	eval.Visit(tree)
}

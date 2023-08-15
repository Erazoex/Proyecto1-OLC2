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
	}
	return Value{value: "la funcion todavia no se ha implementado"}
}

func (v *Visitor) VisitInit(ctx *parser.InitContext) Value {
	return v.Visit(ctx.Block())
}

func Execute() {
	input := `
	/*
		este es mi comentario multilinea final
	*/
	var valor: String?	//correcto, declaracion sin valor

	// correcto, declaracion de una variable tipo Int con valor
	var valor1 = 10

	var valor11:Int = 10.01	// Error: no se puede asignar un Float a un Int
	
	var valor2:Float = 10.2 // correcto

	var valor2_1:Float= 10+1 // correcto sera 11.0

	var valor3 = "esto es una variable" // correcto variable tipo string

	var char:Character = "A" // correcto variable tipo Character

	var valor4:Bool = true // correcto

	// debe ser error ya que los tipos no son compatibles
	var valor4_1:String = true

	// debe ser un error ya que existe otra variable valor3 definida previamente
	var valor3:Int = 10

	var .58 = 4 // debe ser error porque .58 no es un nombre valido

	var if = "10"// debe ser un error porque "if" es una palabra reservada

	// ejemplo de asignaciones

	valor1 = 200 // correcto

	valor3 = "otra cadena" //correcto

	valor4 = 10 //error tipos incompatibles (Iool, Int)

	valor2 = 200 // error tipos incompatibles (Float, Int)

	char = "otra cadena" //error tipos incompatibles (Character, String)

	`
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewGrammarLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGrammarParser(tokenStream)
	p.BuildParseTrees = true
	tree := p.Init()
	eval := Visitor{
		environment: *NewEnvironment(nil),
	}
	eval.Visit(tree)
}

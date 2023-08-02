package grammar

import (
	"fmt"
	"proyecto2/parser"
)

func (v *Visitor) VisitStmt(ctx *parser.StmtContext) Value {
	if ctx.Declstmt() != nil {
		return v.Visit(ctx.Declstmt())
	}
	if ctx.Printstmt() != nil {
		return v.Visit(ctx.Printstmt())
	}
	if ctx.Ifstmt() != nil {
		return v.Visit(ctx.Ifstmt())
	}
	if ctx.Forstmt() != nil {
		return v.Visit(ctx.Forstmt())
	}
	return Value{value: true}
}

func (v *Visitor) VisitForstmt(ctx *parser.ForstmtContext) Value {
	return Value{value: true}
}

func (v *Visitor) VisitDeclstmt(ctx *parser.DeclstmtContext) Value {
	switch ctx.GetTypeVar().GetText() {
	case "let":
		// constante
	case "var":
		// variable
	}
	// hacer comprobacion de que la variable no sea una constante
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr())
	v.memory[id] = value
	return Value{value: true}
}

func (v *Visitor) VisitPrintstmt(ctx *parser.PrintstmtContext) Value {
	fmt.Print(v.Visit(ctx.Expr()).value)
	return Value{value: true}
}

func (v *Visitor) VisitIfstmt(ctx *parser.IfstmtContext) Value {
	value, ok := v.Visit(ctx.Expr()).value.(bool)
	if ok && value {
		return v.Visit(ctx.Block())
	}
	return Value{value: true}
}

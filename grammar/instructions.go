package grammar

import (
	"fmt"
	"proyecto2/parser"
)

func (v *Visitor) VisitStmt(ctx *parser.StmtContext) Value {
	if ctx.Declstmt() != nil {
		return v.Visit(ctx.Declstmt())
	}
	if ctx.Asignstmt() != nil {
		return v.Visit(ctx.Asignstmt())
	}
	return Value{value: true}
}

func (v *Visitor) VisitDeclstmtWithTypeAndExpr(ctx *parser.DeclstmtWithTypeAndExprContext) Value {
	id := ctx.ID().GetText()
	// comprobar que la variable no exista anteriormente
	_, existe := v.environment.tablaSimbolos[id]
	if existe {
		fmt.Println("la variable ya existe", id)
		// TODO: comprobar errores aqui
		return Value{value: true, Type: ERROR}
	}
	value := v.Visit(ctx.Expr())
	value.Id = id
	// asignar si es constante o no
	constante := ctx.GetVtype().GetText()
	switch constante {
	case "let":
		value.Editable = false
	case "var":
		value.Editable = true
	default:
		value.Editable = true
	}
	tipo := v.Visit(ctx.Vartype())
	if value.Type != tipo.Type && !(value.Type == INT && tipo.Type == FLOAT) {
		fmt.Println("la variable no es del mismo tipo que la expresion", id, "tipo:", tipo.Type)
		// TODO: comprobar errores aqui
		return Value{value: true, Type: ERROR}
	}
	// cambiando de manera nativa el valor de la asignacion de un int a un float
	if value.Type == INT && tipo.Type == FLOAT {
		value.value = float64(value.value.(int64))
		value.Type = FLOAT
	}
	v.environment.tablaSimbolos[id] = value
	return Value{value: nil, Type: ACCEPTED}
}

func (v *Visitor) VisitDeclstmtWithExpr(ctx *parser.DeclstmtWithExprContext) Value {
	id := ctx.ID().GetText()
	// comprobar que la variable no exista anteriormente
	_, existe := v.environment.tablaSimbolos[id]
	if existe {
		fmt.Println("la variable ya existe", id)
		// TODO: comprobar errores aqui
		return Value{value: true, Type: ERROR}
	}
	value := v.Visit(ctx.Expr())
	value.Id = id
	// asignar si es constante o no
	constante := ctx.GetVtype().GetText()
	switch constante {
	case "let":
		value.Editable = false
	case "var":
		value.Editable = true
	default:
		value.Editable = true
	}
	v.environment.tablaSimbolos[id] = value
	return Value{value: nil, Type: ACCEPTED}
}

func (v *Visitor) VisitDeclstmtWithType(ctx *parser.DeclstmtWithTypeContext) Value {
	id := ctx.ID().GetText()
	// comprobar que la variable no exista anteriormente
	_, existe := v.environment.tablaSimbolos[id]
	if existe {
		fmt.Println("la variable ya existe", id)
		// TODO: comprobar errores aqui
		return Value{value: true, Type: ERROR}
	}
	value := Value{value: nil, Type: v.Visit(ctx.Vartype()).Type}
	value.Id = id
	// asignar si es constante o no
	constante := ctx.GetVtype().GetText()
	switch constante {
	case "let":
		value.Editable = false
	case "var":
		value.Editable = true
	default:
		value.Editable = true
	}
	v.environment.tablaSimbolos[id] = value
	return Value{value: nil, Type: ACCEPTED}
}

func (v *Visitor) VisitAsignstmt(ctx *parser.AsignstmtContext) Value {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr())
	value.Id = id
	updated := v.environment.UpdateValue(value)
	if !updated {
		// TODO: errores mas adelante
		// fmt.Println("la variable no se pudo actualizar", value.Id)
		return Value{value: false, Type: ERROR}
	}
	return Value{value: nil, Type: ACCEPTED}
}

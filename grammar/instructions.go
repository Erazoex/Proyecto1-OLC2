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
	if ctx.Incstmt() != nil {
		return v.Visit(ctx.Asignstmt())
	}
	if ctx.Decstmt() != nil {
		return v.Visit(ctx.Decstmt())
	}
	return Value{value: true, Type: ACCEPTED}
}

// DECLARACION
func (v *Visitor) VisitDeclstmtWithTypeAndExpr(ctx *parser.DeclstmtWithTypeAndExprContext) Value {
	id := ctx.ID().GetText()
	// comprobar que la variable no exista anteriormente
	_, existe := v.environment.GetValue(id)
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
	_, existe := v.environment.GetValue(id)
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
	value := Value{value: nil, Type: v.Visit(ctx.Vartype()).Type}
	value.Id = id
	// asignar si es constante o no
	constante := ctx.GetVtype().GetText()
	switch constante {
	case "let":
		fmt.Println("La constante debe tener un valor asignado", value.Id)
		return Value{value: true, Type: ERROR}
	case "var":
		value.Editable = true
	default:
		value.Editable = true
	}
	// comprobar que la variable no exista anteriormente
	_, existe := v.environment.GetValue(id)
	if existe {
		fmt.Println("la variable ya existe", id)
		// TODO: comprobar errores aqui
		return Value{value: true, Type: ERROR}
	}
	v.environment.tablaSimbolos[id] = value
	return Value{value: nil, Type: ACCEPTED}
}

// ASIGNACION
func (v *Visitor) VisitAsignstmt(ctx *parser.AsignstmtContext) Value {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr())
	value.Id = id
	updated := v.environment.UpdateValue(value)
	if !updated {
		// TODO: errores mas adelante
		// fmt.Println("la variable no se pudo actualizar", value.Id)
		return Value{value: false}
	}
	return Value{value: nil, Type: ACCEPTED}
}

// INCREMENTO Y DECREMENTO
func (v *Visitor) VisitIncstmt(ctx *parser.IncstmtContext) Value {
	id := ctx.ID().GetText()
	ExprValue := v.Visit(ctx.Expr())
	value, ok := v.environment.GetValue(id)
	if !ok {
		// TODO: implementar error mas adelante
		return Value{value: false}
	}
	if value.Type != INT && value.Type != FLOAT && value.Type != STRING {
		// TODO: implementar error mas adelante
		return Value{value: false}
	}
	if value.Type != ExprValue.Type {
		// TODO: implementar error mas adelante
		return Value{value: false}
	}
	if !value.Editable {
		// TODO: implementar error mas adelante
		return Value{value: false}
	}
	switch value.Type {
	case INT:
		value.value = value.value.(int64) + ExprValue.value.(int64)
	case FLOAT:
		value.value = value.value.(float64) + ExprValue.value.(float64)
	case STRING:
		value.value = value.value.(string) + ExprValue.value.(string)
	}
	v.environment.UpdateValue(value)
	return Value{value: true, Type: ACCEPTED}
}

func (v *Visitor) VisitDecstmt(ctx *parser.DecstmtContext) Value {
	id := ctx.ID().GetText()
	ExprValue := v.Visit(ctx.Expr())
	value, ok := v.environment.GetValue(id)
	if !ok {
		// TODO: implementar error mas adelante
		return Value{value: false}
	}
	if value.Type != INT && value.Type != FLOAT {
		// TODO: implementar error mas adelante
		return Value{value: false}
	}
	if value.Type != ExprValue.Type {
		// TODO: implementar error mas adelante
		return Value{value: false}
	}
	if !value.Editable {
		// TODO: implementar error mas adelante
		return Value{value: false}
	}
	switch value.Type {
	case INT:
		value.value = value.value.(int64) - ExprValue.value.(int64)
	case FLOAT:
		value.value = value.value.(float64) - ExprValue.value.(float64)
	}
	v.environment.UpdateValue(value)
	return Value{value: true, Type: ACCEPTED}
}

// IF

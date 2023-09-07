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
	if ctx.Ifstmt() != nil {
		return v.Visit(ctx.Ifstmt())
	}
	if ctx.Switchstmt() != nil {
		return v.Visit(ctx.Switchstmt())
	}
	if ctx.Printlnstmt() != nil {
		return v.Visit(ctx.Printlnstmt())
	}
	if ctx.Whilestmt() != nil {
		return v.Visit(ctx.Whilestmt())
	}
	if ctx.Forstmt() != nil {
		return v.Visit(ctx.Forstmt())
	}
	return Value{value: true, Type: ACCEPTED}
}

// DECLARACION
func (v *Visitor) VisitDeclstmtWithTypeAndExpr(ctx *parser.DeclstmtWithTypeAndExprContext) Value {
	id := ctx.ID().GetText()
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
	value.Type = v.Visit(ctx.Expr()).Type
	// cambiando de manera nativa el valor de la asignacion de un int a un float
	if value.Type == INT && tipo.Type == FLOAT {
		value.value = float64(value.value.(int64))
		value.Type = FLOAT
	}
	v.environment.SaveValue(value)
	return Value{value: nil, Type: ACCEPTED}
}

func (v *Visitor) VisitDeclstmtWithExpr(ctx *parser.DeclstmtWithExprContext) Value {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr())
	value.Type = v.Visit(ctx.Expr()).Type
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
	v.environment.SaveValue(value)
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
	value.value = nil
	value.Type = NIL
	v.environment.SaveValue(value)
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
func (v *Visitor) VisitIfSimple(ctx *parser.IfSimpleContext) Value {
	condition := v.Visit(ctx.Expr())
	if condition.Type != BOOL {
		// TODO: implementar error aqui
		return Value{value: false}
	}
	if condition.value == true {
		v.Visit(ctx.Block())
	}
	return Value{value: true, Type: ACCEPTED}
}

func (v *Visitor) VisitIfWithElse(ctx *parser.IfWithElseContext) Value {
	condition := v.Visit(ctx.Expr())
	if condition.Type != BOOL {
		// TODO: implementar error aqui
		return Value{value: false}
	}
	if condition.value == true {
		v.Visit(ctx.GetTrueCondition())
	} else {
		v.Visit(ctx.GetFalseCondition())
	}
	return Value{value: true, Type: ACCEPTED}
}

func (v *Visitor) VisitIfWithElseIf(ctx *parser.IfWithElseIfContext) Value {
	condition := v.Visit(ctx.Expr())
	if condition.Type != BOOL {
		// TODO: implementar error aqui
		return Value{value: false}
	}
	if condition.value == true {
		v.Visit(ctx.Block())
	} else {
		v.Visit(ctx.Ifstmt())
	}
	return Value{value: true, Type: ACCEPTED}
}

// SWITCH
func (v *Visitor) VisitSwitchstmt(ctx *parser.SwitchstmtContext) Value {
	exprValue := v.Visit(ctx.Expr())
	isThereAnyDefault := false
	for i := 0; ctx.Switchcase(i) != nil; i++ {
		var transfer Value
		if ctx.Switchcase(i).GetCasetype().GetText() == "default" {
			isThereAnyDefault = true
		}
		if isThereAnyDefault && ctx.Switchcase(i) != nil {
			// TODO: implementar error aqui
			// Error: no se puede declarar un case o otro default despues de un default
			return Value{value: false}
		}
		option := ctx.Switchcase(i).GetCasetype().GetText()
		switch option {
		case "case":
			if exprValue.value == v.Visit(ctx.Switchcase(i).Expr()).value {
				transfer = v.Visit(ctx.Switchcase(i).Block())
			}
		case "default":
			transfer = v.Visit(ctx.Switchcase(i).Block())
		}
		if transfer.Type == BREAK {
			break
		}
	}
	return Value{value: true, Type: ACCEPTED}
}

// PRINTLN
func (v *Visitor) VisitPrintlnstmt(ctx *parser.PrintlnstmtContext) Value {
	if ctx.Exprparams() != nil {
		for _, item := range v.Visit(ctx.Exprparams()).value.([]Value) {
			fmt.Print(item.value, " ")
		}
		fmt.Println()
	} else {
		fmt.Println()
	}
	return Value{value: true, Type: ACCEPTED}
}

// WHILE
func (v *Visitor) VisitWhilestmt(ctx *parser.WhilestmtContext) Value {
	v.environment = NewEnvironment(v.environment)
	for {
		condition := v.Visit(ctx.Expr())
		if condition.Type != BOOL {
			// TODO: implementar error aqui
			// la condicion no es de tipo boolean
			return Value{value: true}
		}
		if condition.value == true {
			transfer := v.Visit(ctx.Block())
			if transfer.Type == BREAK {
				return Value{value: true, Type: ACCEPTED}
			}
			if transfer.Type == CONTINUE {
				continue
			}
		} else {
			break
		}
	}
	return Value{value: true, Type: ACCEPTED}
}

// TODO: fors no implementados
// FOR
func (v *Visitor) VisitForWithExpr(ctx *parser.ForWithExprContext) Value {
	// TODO: realizar los cambios correspondientes aqui para que se puedan usar strings y
	// vectores
	v.environment = NewEnvironment(v.environment)
	for {
		condition := v.Visit(ctx.Expr())
		if condition.Type != BOOL {
			// TODO: implementar error
			// la condicion no es de tipo boolean
			return Value{value: true}
		}
		if condition.value == true {
			transfer := v.Visit(ctx.Block())
			if transfer.Type == BREAK {
				return Value{value: true, Type: ACCEPTED}
			}
		} else {
			break
		}
	}
	return Value{value: true, Type: ACCEPTED}
}

func (v *Visitor) VisitForWithRange(ctx *parser.ForWithRangeContext) Value {
	v.environment = NewEnvironment(v.environment)
	id := ctx.ID().GetText()
	empty := false
	if id == "_" {
		empty = true
	}
	left := v.Visit(ctx.Forrange().GetBeginsWith())
	temp := Value{
		value:    left.value,
		Editable: true,
		Id:       id,
		Type:     left.Type,
	}
	if !empty {
		v.environment.SaveValue(temp)
	}
	for {
		temp, _ = v.environment.GetValue(temp.Id)
		right := v.Visit(ctx.Forrange().GetEndsWith())
		if left.Type != INT || right.Type != INT {
			// TODO: implementar error aqui
			// el rango no contiene expresiones de tipo INT
			return Value{Type: ERROR}
		}
		if left.value.(int64) > right.value.(int64) {
			// TODO: implementar error aqui
			// la expresion en la izquierda no puede ser mayor a la expresion en la derecha
			return Value{Type: ERROR}
		}
		var transfer Value
		if temp.value.(int64) <= right.value.(int64) {
			transfer = v.Visit(ctx.Block())
		} else {
			break
		}
		if transfer.Type == CONTINUE {
			continue
		}
		if transfer.Type == BREAK {
			break
		}
		if transfer.Type == RETURN {
			return transfer
		}
		temp.value = temp.value.(int64) + 1
		if !empty {
			v.environment.UpdateValue(temp)
		}
	}

	return Value{Type: ACCEPTED}
}

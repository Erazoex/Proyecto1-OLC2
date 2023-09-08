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
	if ctx.Guardstmt() != nil {
		return v.Visit(ctx.Guardstmt())
	}
	if ctx.Breakstmt() != nil {
		return v.Visit(ctx.Breakstmt())
	}
	if ctx.Continuestmt() != nil {
		return v.Visit(ctx.Continuestmt())
	}
	if ctx.Returnstmt() != nil {
		return v.Visit(ctx.Returnstmt())
	}
	return Value{value: true, Type: ACCEPTED}
}

func (v *Visitor) VisitBreakstmt(ctx *parser.BreakstmtContext) Value {
	return Value{
		Type: BREAK,
	}
}

func (v *Visitor) VisitContinuestmt(ctx *parser.ContinuestmtContext) Value {
	return Value{
		Type: CONTINUE,
	}
}

func (v *Visitor) VisitReturnstmt(ctx *parser.ReturnstmtContext) Value {
	return Value{
		Type: RETURN,
	}
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
	if value.Type != tipo.Type && !(value.Type == INT && tipo.Type == FLOAT) && !(tipo.Type == STRING && value.Type == CHAR) {
		fmt.Println("la variable no es del mismo tipo que la expresion", id, value.Type, "tipo:", tipo.Type)
		// TODO: comprobar errores aqui
		v.push(error{
			desc:   "la variable no es del mismo tipo que la expresion",
			line:   ctx.GetStart().GetLine(),
			column: ctx.GetStart().GetColumn(),
		})
		return Value{value: true, Type: ERROR}
	}
	value.Type = tipo.Type
	// cambiando de manera nativa el valor de la asignacion de un int a un float
	if value.Type == INT && tipo.Type == FLOAT {
		value.value = float64(value.value.(int64))
		value.Type = FLOAT
	}
	v.Environment.SaveValue(value)
	return Value{value: nil, Type: ACCEPTED}
}

func (v *Visitor) VisitDeclstmtWithExpr(ctx *parser.DeclstmtWithExprContext) Value {
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
	v.Environment.SaveValue(value)
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
		// TODO: implementar error aqui
		// la constante debe tener un valor asignado
		fmt.Println("La constante debe tener un valor asignado", value.Id)
		v.push(error{
			desc:   "La constante debe tener un valor asignado",
			line:   ctx.GetStart().GetLine(),
			column: ctx.GetStart().GetColumn(),
		})
		return Value{value: true, Type: ERROR}
	case "var":
		value.Editable = true
	default:
		value.Editable = true
	}
	value.value = nil
	value.Type = NIL
	v.Environment.SaveValue(value)
	return Value{value: nil, Type: ACCEPTED}
}

// ASIGNACION
func (v *Visitor) VisitAsignstmt(ctx *parser.AsignstmtContext) Value {
	id := ctx.ID().GetText()
	value := v.Visit(ctx.Expr())
	value.Id = id
	updated := v.Environment.UpdateValue(value)
	if !updated {
		// TODO: errores mas adelante
		// fmt.Println("la variable no se pudo actualizar", value.Id)
		v.push(error{
			desc:   "La variable no se pudo actualizar",
			line:   ctx.GetStart().GetLine(),
			column: ctx.GetStart().GetColumn(),
		})
		return Value{value: false}
	}
	return Value{value: nil, Type: ACCEPTED}
}

// INCREMENTO Y DECREMENTO
func (v *Visitor) VisitIncstmt(ctx *parser.IncstmtContext) Value {
	id := ctx.ID().GetText()
	ExprValue := v.Visit(ctx.Expr())
	value, ok := v.Environment.GetValue(id)
	if !ok {
		// TODO: implementar error mas adelante
		v.push(error{
			desc:   "la variable a incrementar el valor no existe",
			line:   ctx.GetStart().GetLine(),
			column: ctx.GetStart().GetColumn(),
		})
		return Value{value: false}
	}
	if value.Type != INT && value.Type != FLOAT && value.Type != STRING {
		// TODO: implementar error mas adelante
		v.push(error{
			desc:   "La variable a incrementar no es de un tipo permitido",
			line:   ctx.GetStart().GetLine(),
			column: ctx.GetStart().GetColumn(),
		})
		return Value{value: false}
	}
	if value.Type != ExprValue.Type {
		// TODO: implementar error mas adelante
		v.push(error{
			desc:   "La expresion no es del mismo tipo que la variable",
			line:   ctx.GetStart().GetLine(),
			column: ctx.GetStart().GetColumn(),
		})
		return Value{value: false}
	}
	if !value.Editable {
		v.push(error{
			desc:   "No se puede incrementar una constante",
			line:   ctx.GetStart().GetLine(),
			column: ctx.GetStart().GetColumn(),
		})
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
	v.Environment.UpdateValue(value)
	return Value{value: true, Type: ACCEPTED}
}

func (v *Visitor) VisitDecstmt(ctx *parser.DecstmtContext) Value {
	id := ctx.ID().GetText()
	ExprValue := v.Visit(ctx.Expr())
	value, ok := v.Environment.GetValue(id)
	if !ok {
		// TODO: implementar error mas adelante
		v.push(error{
			desc:   "La variable a decrementar no existe",
			line:   ctx.GetStart().GetLine(),
			column: ctx.GetStart().GetColumn(),
		})
		return Value{value: false}
	}
	if value.Type != INT && value.Type != FLOAT {
		// TODO: implementar error mas adelante
		v.push(error{
			desc:   "El tipo de la variable no es de un tipo permitido a decrementar",
			line:   ctx.GetStart().GetLine(),
			column: ctx.GetStart().GetColumn(),
		})
		return Value{value: false}
	}
	if value.Type != ExprValue.Type {
		// TODO: implementar error mas adelante
		v.push(error{
			desc:   "La expresion no es del mismo tipo que la variable",
			line:   ctx.GetStart().GetLine(),
			column: ctx.GetStart().GetColumn(),
		})
		return Value{value: false}
	}
	if !value.Editable {
		// TODO: implementar error mas adelante
		v.push(error{
			desc:   "No se puede decrementar una constante",
			line:   ctx.GetStart().GetLine(),
			column: ctx.GetStart().GetColumn(),
		})
		return Value{value: false}
	}
	switch value.Type {
	case INT:
		value.value = value.value.(int64) - ExprValue.value.(int64)
	case FLOAT:
		value.value = value.value.(float64) - ExprValue.value.(float64)
	}
	v.Environment.UpdateValue(value)
	return Value{value: true, Type: ACCEPTED}
}

// IF
func (v *Visitor) VisitIfSimple(ctx *parser.IfSimpleContext) Value {
	condition := v.Visit(ctx.Expr())
	if condition.Type != BOOL {
		// TODO: implementar error aqui
		v.push(error{
			desc:   "La condicion no es de tipo Boolean",
			line:   ctx.GetStart().GetLine(),
			column: ctx.GetStart().GetColumn(),
		})
		return Value{value: false}
	}
	if condition.value == true {
		return v.Visit(ctx.Block())
	}
	return Value{value: true, Type: ACCEPTED}
}

func (v *Visitor) VisitIfWithElse(ctx *parser.IfWithElseContext) Value {
	condition := v.Visit(ctx.Expr())
	if condition.Type != BOOL {
		// TODO: implementar error aqui
		v.push(error{
			desc:   "La condicion no es de tipo Boolean",
			line:   ctx.GetStart().GetLine(),
			column: ctx.GetStart().GetColumn(),
		})
		return Value{value: false}
	}
	if condition.value == true {
		return v.Visit(ctx.GetTrueCondition())
	} else {
		return v.Visit(ctx.GetFalseCondition())
	}
}

func (v *Visitor) VisitIfWithElseIf(ctx *parser.IfWithElseIfContext) Value {
	condition := v.Visit(ctx.Expr())
	if condition.Type != BOOL {
		// TODO: implementar error aqui
		v.push(error{
			desc:   "La condicion no es de tipo Boolean",
			line:   ctx.GetStart().GetLine(),
			column: ctx.GetStart().GetColumn(),
		})
		return Value{value: false}
	}
	if condition.value == true {
		return v.Visit(ctx.Block())
	} else {
		return v.Visit(ctx.Ifstmt())
	}
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
			v.push(error{
				desc:   "No se puede declarar un case u otro default despues de un default",
				line:   ctx.GetStart().GetLine(),
				column: ctx.GetStart().GetColumn(),
			})
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
			v.Print(fmt.Sprintf("%v", item.value))
			fmt.Print(item.value, " ")
		}
		v.Print("\n")
		fmt.Println()
	} else {
		v.Print("\n")
		fmt.Println()
	}
	return Value{value: true, Type: ACCEPTED}
}

// WHILE
func (v *Visitor) VisitWhilestmt(ctx *parser.WhilestmtContext) Value {
	newEnv := NewEnvironment(v.Environment)
	v.Environment.push(newEnv)
	v.Environment = newEnv
	for {
		condition := v.Visit(ctx.Expr())
		if condition.Type != BOOL {
			// TODO: implementar error aqui
			// la condicion no es de tipo boolean
			v.push(error{
				desc:   "La condicion no es de tipo boolean",
				line:   ctx.GetStart().GetLine(),
				column: ctx.GetStart().GetColumn(),
			})
			return Value{Type: ERROR}
		}
		if condition.value.(bool) {
			transfer := v.Visit(ctx.Block())
			if transfer.Type == BREAK {
				break
			}
			if transfer.Type == CONTINUE {
				continue
			}
			if transfer.Type == RETURN {
				return transfer
			}
		} else {
			break
		}
	}
	return Value{value: true, Type: ACCEPTED}
}

// TODO: for con expresion no implementado
// FOR
func (v *Visitor) VisitForWithExpr(ctx *parser.ForWithExprContext) Value {
	// TODO: realizar los cambios correspondientes aqui para que se puedan usar strings y
	// vectores
	newEnv := NewEnvironment(v.Environment)
	v.Environment.push(newEnv)
	v.Environment = newEnv
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
	newEnv := NewEnvironment(v.Environment)
	v.Environment.push(newEnv)
	v.Environment = newEnv
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
		Type:     INT,
	}
	if !empty {
		v.Environment.SaveValue(temp)
	}
	for {
		if !empty {
			temp, _ = v.Environment.GetValue(temp.Id)
		}
		right := v.Visit(ctx.Forrange().GetEndsWith())
		if left.Type != INT || right.Type != INT {
			// TODO: implementar error aqui
			// el rango no contiene expresiones de tipo INT
			v.push(error{
				desc:   "El rango no contiene expresiones de tipo INT",
				line:   ctx.GetStart().GetLine(),
				column: ctx.GetStart().GetColumn(),
			})
			return Value{Type: ERROR}
		}
		if temp.value.(int64) > right.value.(int64) {
			// TODO: implementar error aqui
			// la expresion en la izquierda no puede ser mayor a la expresion en la derecha
			v.push(error{
				desc:   "La expresion en la izquierda no puede ser mayor a la expresion en la derecha",
				line:   ctx.GetStart().GetLine(),
				column: ctx.GetStart().GetColumn(),
			})
			return Value{Type: ERROR}
		}
		if temp.value.(int64) < right.value.(int64) {
			transfer := v.Visit(ctx.Block())
			if transfer.Type == CONTINUE {
				continue
			}
			if transfer.Type == BREAK {
				break
			}
			if transfer.Type == RETURN {
				return transfer
			}
		} else {
			break
		}
		temp.value = temp.value.(int64) + 1
		if !empty {
			v.Environment.UpdateValue(temp)
		}
	}

	return Value{Type: ACCEPTED}
}

// GUARD
func (v *Visitor) VisitGuardstmt(ctx *parser.GuardstmtContext) Value {
	condition := v.Visit(ctx.Expr())
	if condition.Type != BOOL {
		// TODO: implementar error aqui
		// la condicion no es de tipo Boolean
		v.push(error{
			desc:   "La condicion no es de tipo Boolean",
			line:   ctx.GetStart().GetLine(),
			column: ctx.GetStart().GetColumn(),
		})
		return Value{Type: ERROR}
	}
	if !condition.value.(bool) {
		transfer := v.Visit(ctx.Block())
		if transfer.Type == BREAK {
			return transfer
		}
		if transfer.Type == CONTINUE {
			return transfer
		}
		if transfer.Type == RETURN {
			return transfer
		}
		// TODO: implementar error aqui
		// la instruccion guard no contiene una sentencia de transferencia
		v.push(error{
			desc:   "La instruccion guard no contiene una sentencia de transferencia",
			line:   ctx.GetStart().GetLine(),
			column: ctx.GetStart().GetColumn(),
		})
		return Value{Type: ERROR}
	}
	return Value{Type: ACCEPTED}
}

package grammar

import (
	"fmt"
	"proyecto2/parser"
	"strconv"
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
	if ctx.Funcstmt() != nil {
		return v.Visit(ctx.Funcstmt())
	}
	if ctx.Callstmt() != nil {
		return v.Visit(ctx.Callstmt())
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
	if ctx.Expr() != nil {
		return Value{
			value: v.Visit(ctx.Expr()),
			Type:  RETURN,
		}
	}
	return Value{
		value: nil,
		Type:  RETURN,
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
		myValue := v.Visit(ctx.Exprparams())
		valueArray := myValue.value.([]Value)
		for _, item := range valueArray {
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

// FUNCIONES EMBEBIDAS
func (v *Visitor) VisitIntstmt(ctx *parser.IntstmtContext) Value {
	expr := v.Visit(ctx.Expr())
	if expr.Type == STRING {
		value, _ := strconv.Atoi(expr.value.(string))
		return Value{
			value: int64(value),
			Type:  INT,
		}
	}
	if expr.Type == FLOAT {
		return Value{
			value: int64(expr.value.(float64)),
			Type:  INT,
		}
	}
	v.push(error{
		desc:   "el tipo de expresion no se puede convertir en int",
		line:   ctx.GetStart().GetLine(),
		column: ctx.GetStart().GetColumn(),
	})
	return Value{Type: NIL}
}

func (v *Visitor) VisitFloatstmt(ctx *parser.FloatstmtContext) Value {
	expr := v.Visit(ctx.Expr())
	if expr.Type == STRING {
		value, _ := strconv.ParseFloat(expr.value.(string), 64)
		return Value{
			value: value,
			Type:  FLOAT,
		}
	}
	if expr.Type == INT {
		return Value{
			value: float64(expr.value.(int64)),
			Type:  FLOAT,
		}
	}
	v.push(error{
		desc:   "el tipo de expresion no se puede convertir en float",
		line:   ctx.GetStart().GetLine(),
		column: ctx.GetStart().GetColumn(),
	})
	return Value{Type: NIL}
}

func (v *Visitor) VisitStringstmt(ctx *parser.StringstmtContext) Value {
	expr := v.Visit(ctx.Expr())
	if expr.Type == INT || expr.Type == FLOAT || expr.Type == BOOL {
		return Value{
			value: fmt.Sprintf("%v", expr.value),
			Type:  STRING,
		}
	}
	v.push(error{
		desc:   "el tipo de expresion no se puede convertir en string",
		line:   ctx.GetStart().GetLine(),
		column: ctx.GetStart().GetColumn(),
	})
	return Value{Type: NIL}
}

// FUNCION
func (v *Visitor) VisitFuncstmt(ctx *parser.FuncstmtContext) Value {
	id := ctx.ID().GetText()
	var returnType DataType
	returnType = VOID
	if ctx.Vartype() != nil {
		funcType := ctx.Vartype().GetText()
		switch funcType {
		case "Int":
			returnType = INT
		case "Float":
			returnType = FLOAT
		case "String":
			returnType = STRING
		case "Bool":
			returnType = BOOL
		case "Character":
			returnType = CHAR
		default:
			returnType = VOID
		}
	}

	result := v.Environment.SaveFunc(Function{
		Id:   id,
		Type: returnType,
		Ctx:  ctx,
	})
	if result {
		return Value{Type: ACCEPTED}
	}
	v.push(error{
		desc:   "ya existe una funcion con ese nombre",
		line:   ctx.GetStart().GetLine(),
		column: ctx.GetStart().GetColumn(),
	})
	return Value{Type: ERROR}
}

func (v *Visitor) VisitCallstmt(ctx *parser.CallstmtContext) Value {
	id := ctx.ID().GetText()
	funcion, ok := v.Environment.GetFunc(id)
	if !ok {
		v.push(error{
			desc:   "la funcion no existe",
			line:   ctx.GetStart().GetLine(),
			column: ctx.GetStart().GetColumn(),
		})
		return Value{}
	}
	Llamada := v.Visit(ctx.Exprparams())
	parametrosLlamada := Llamada.value.([]Value)
	newEnv := NewEnvironment(v.Environment)
	v.Environment.push(newEnv)
	v.Environment = newEnv
	var parametros []Parameter
	if funcion.Ctx.Listaparametros() != nil {
		parametros = v.Visit(funcion.Ctx.Listaparametros()).value.([]Parameter)
	}
	if !(parametrosLlamada != nil && parametros != nil) {
		if ctx.Exprparams() != nil {
			v.push(error{
				desc:   "la llamada tiene parametros pero la funcion no",
				line:   ctx.GetStart().GetLine(),
				column: ctx.GetStart().GetColumn(),
			})
			return Value{}
		}
		if parametros != nil {
			v.push(error{
				desc:   "la llamada tiene parametros pero la funcion no",
				line:   ctx.GetStart().GetLine(),
				column: ctx.GetStart().GetColumn(),
			})
			return Value{}
		}
	}
	if len(parametros) != len(parametrosLlamada) {
		v.push(error{
			desc:   "la llamada tiene parametros pero la funcion no",
			line:   ctx.GetStart().GetLine(),
			column: ctx.GetStart().GetColumn(),
		})
		return Value{}
	}
	for i := 0; i < len(parametros); i++ {
		if parametros[i].Type != parametrosLlamada[i].Type {
			v.push(error{
				desc:   "el tipo de los parametros no coincide",
				line:   ctx.GetStart().GetLine(),
				column: ctx.GetStart().GetColumn(),
			})
			return Value{}
		}
		v.Environment.SaveValue(Value{
			Id:    parametros[i].Id,
			value: parametrosLlamada[i].value,
			Type:  parametros[i].Type,
		})
	}
	if funcion.Type != VOID {
		result := v.Visit(funcion.Ctx.Block())
		if result.Type != RETURN {
			v.push(error{
				desc:   "no se encontro un return en la funcion",
				line:   ctx.GetStart().GetLine(),
				column: ctx.GetStart().GetColumn(),
			})
			return Value{}
		}
		if result.value == nil {
			v.push(error{
				desc:   "no se encontro un valor para retornar",
				line:   ctx.GetStart().GetLine(),
				column: ctx.GetStart().GetColumn(),
			})
			return Value{}
		}
		if result.value.(Value).Type != funcion.Type {
			v.push(error{
				desc:   "el valor de retorno no es del mismo tipo que ",
				line:   ctx.GetStart().GetLine(),
				column: ctx.GetStart().GetColumn(),
			})
			return Value{}
		}
		return Value{
			value: result.value.(Value).value,
			Type:  funcion.Type,
		}
	}

	return Value{}
}

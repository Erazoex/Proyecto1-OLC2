package grammar

import (
	"proyecto2/parser"
	"strconv"
	"strings"
)

func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) Value {
	val, ok := v.Environment.GetValue(ctx.ID().GetText())
	if ok {
		return val
	}
	// TODO: implementar error aqui
	// no se encontro la variable
	v.push(error{
		desc:   "No se encontro la variable",
		line:   ctx.GetStart().GetLine(),
		column: ctx.GetStart().GetColumn(),
	})
	return Value{Type: ERROR}
}

func (v *Visitor) VisitNotExpr(ctx *parser.NotExprContext) Value {
	right := v.Visit(ctx.GetRight())
	if right.Type == BOOL {
		if right.value.(bool) {
			right.value = false
			return right
		} else {
			right.value = true
			return right
		}
	}
	// TODO: manejo de errores aqui
	v.push(error{
		desc:   "No se puede negar la expresion",
		line:   ctx.GetStart().GetLine(),
		column: ctx.GetStart().GetColumn(),
	})
	return Value{value: false}
}

func (v *Visitor) VisitUnaryExpr(ctx *parser.UnaryExprContext) Value {
	right := v.Visit(ctx.GetRight())
	if right.Type == INT {
		return Value{value: right.value.(int64) * -1, Type: INT}
	}
	if right.Type == FLOAT {
		return Value{value: right.value.(float64) * -1, Type: FLOAT}
	}
	// TODO: manejo de errores aqui
	v.push(error{
		desc:   "Tipo de expresion no valido para expresion unaria",
		line:   ctx.GetStart().GetLine(),
		column: ctx.GetStart().GetColumn(),
	})
	return Value{value: false}
}

func (v *Visitor) VisitParExpr(ctx *parser.ParExprContext) Value {
	result := v.Visit(ctx.Expr())
	return result
}

func (v *Visitor) VisitIntExpr(ctx *parser.IntExprContext) Value {
	i, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return Value{
		value: i,
		Type:  INT,
	}
}

func (v *Visitor) VisitDoubleExpr(ctx *parser.DoubleExprContext) Value {
	i, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return Value{
		value: i,
		Type:  FLOAT,
	}
}

func (v *Visitor) VisitNilExpr(ctx *parser.NilExprContext) Value {
	return Value{
		value: ctx.GetText(),
		Type:  NIL,
	}
}

func (v *Visitor) VisitStrExpr(ctx *parser.StrExprContext) Value {
	value := strings.TrimRight(strings.TrimLeft(ctx.GetText(), "\""), "\"")
	value = strings.ReplaceAll(value, "\\\"", "\"")
	value = strings.ReplaceAll(value, "\\\\", "\\")
	value = strings.ReplaceAll(value, "\\n", "\n")
	value = strings.ReplaceAll(value, "\\r", "\r")
	value = strings.ReplaceAll(value, "\\t", "\t")
	return Value{
		value: value,
		Type:  STRING,
	}
}

func (v *Visitor) VisitCharExpr(ctx *parser.CharExprContext) Value {
	value := strings.Trim(ctx.GetText(), "\"")
	return Value{
		value: value,
		Type:  CHAR,
	}
}

func (v *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) Value {
	value, _ := strconv.ParseBool(ctx.GetText())
	return Value{
		value: value,
		Type:  BOOL,
	}
}

func (v *Visitor) VisitExprWithParams(ctx *parser.ExprWithParamsContext) Value {
	expr := v.Visit(ctx.Expr())
	var newSlice []Value
	newSlice = append(newSlice, expr)
	exprSlice := v.Visit(ctx.Exprparams())
	newSlice = append(newSlice, exprSlice.value.([]Value)...)
	return Value{
		value: newSlice,
		Type:  expr.Type,
	}
}

func (v *Visitor) VisitExprWithParam(ctx *parser.ExprWithParamContext) Value {
	expr := v.Visit(ctx.Expr())
	var array []Value
	array = append(array, expr)
	return Value{
		value: array,
		Type:  expr.Type,
	}
}

func (v *Visitor) VisitFuncParameters(ctx *parser.FuncParametersContext) Value {
	exprWithParameter := v.Visit(ctx.Parametro())
	var newSlice []Parameter
	newSlice = append(newSlice, exprWithParameter.value.(Parameter))
	exprSlice := v.Visit(ctx.Listaparametros())
	newSlice = append(newSlice, exprSlice.value.([]Parameter)...)
	return Value{
		value: newSlice,
		Type:  ACCEPTED,
	}
}

func (v *Visitor) VisitFuncParameter(ctx *parser.FuncParameterContext) Value {
	exprWithParameter := v.Visit(ctx.Parametro())
	var array []Parameter
	array = append(array, exprWithParameter.value.(Parameter))
	return Value{
		value: array,
		Type:  ACCEPTED,
	}
}

// PARAMETRO
func (v *Visitor) VisitParametro(ctx *parser.ParametroContext) Value {
	varId := ctx.GetVarId().GetText()
	varType := v.Visit(ctx.Vartype())
	newParam := Parameter{
		Id:   varId,
		Type: varType.Type,
	}
	return Value{
		value: newParam,
	}

}

func (v *Visitor) VisitIntConvExpr(ctx *parser.IntConvExprContext) Value {
	return v.Visit(ctx.Intstmt())
}

func (v *Visitor) VisitFloatConvExpr(ctx *parser.FloatConvExprContext) Value {
	return v.Visit(ctx.Floatstmt())
}

func (v *Visitor) VisitStringConvExpr(ctx *parser.StringConvExprContext) Value {
	return v.Visit(ctx.Stringstmt())
}

func (v *Visitor) VisitCallExpr(ctx *parser.CallExprContext) Value {
	return v.Visit(ctx.Callstmt())
}

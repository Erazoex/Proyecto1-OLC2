package grammar

import (
	"fmt"
	"proyecto2/parser"
	"strconv"
	"strings"
)

func (v *Visitor) VisitNotExpr(ctx *parser.NotExprContext) Value {
	right := v.Visit(ctx.GetRight())
	if right.Type == BOOL {
		return Value{value: !right.value.(bool)}
	}
	// TODO: manejo de errores aqui
	return Value{value: false}
}

func (v *Visitor) VisitUnaryExpr(ctx *parser.UnaryExprContext) Value {
	right := v.Visit(ctx.GetRight())
	if right.Type == INT {
		return Value{value: -right.value.(int)}
	}
	if right.Type == FLOAT {
		return Value{value: -right.value.(float64)}
	}
	// TODO: manejo de errores aqui
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
	value := strings.Trim(ctx.GetText(), "\"")[0]
	return Value{
		value: value,
		Type:  CHAR,
	}
}

func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) Value {
	id := ctx.GetText()
	value, ok := v.environment.tablaSimbolos[id]
	if ok {
		return value
	} else {
		fmt.Println("la variable no existe: " + id)
	}
	return Value{value: false}
}

func (v *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) Value {
	value, _ := strconv.ParseBool(ctx.GetText())
	return Value{
		value: value,
		Type:  BOOL,
	}
}

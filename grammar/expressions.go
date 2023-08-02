package grammar

import (
	"fmt"
	"proyecto2/parser"
	"strconv"
	"strings"
)

func (v *Visitor) VisitNotExpr(ctx *parser.NotExprContext) Value {
	right := v.Visit(ctx.GetRight())
	switch right.value.(type) {
	case bool:
		return Value{value: !right.value.(bool)}
	default:
		return Value{value: false}
	}
}

func (v *Visitor) VisitUnaryExpr(ctx *parser.UnaryExprContext) Value {
	right := v.Visit(ctx.GetRight())
	switch right.value.(type) {
	case int:
		return Value{value: -right.value.(int)}
	case float64:
		return Value{value: -right.value.(float64)}
	}
	return Value{value: false}
}

func (v *Visitor) VisitIntExpr(ctx *parser.IntExprContext) Value {
	i, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return Value{value: i}
}

func (v *Visitor) VisitDoubleExpr(ctx *parser.DoubleExprContext) Value {
	i, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return Value{value: i}
}

func (v *Visitor) VisitNilExpr(ctx *parser.NilExprContext) Value {
	return Value{value: ctx.GetText()}
}

func (v *Visitor) VisitStrExpr(ctx *parser.StrExprContext) Value {
	value := strings.TrimRight(strings.TrimLeft(ctx.GetText(), "\""), "\"")
	return Value{value: value}
}

func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) Value {
	id := ctx.GetText()
	value, ok := v.memory[id]
	if ok {
		return value
	} else {
		fmt.Println("la variable no existe: " + id)
	}
	return Value{value: false}
}

func (v *Visitor) VisitBoolExpr(ctx *parser.BoolExprContext) Value {
	value, _ := strconv.ParseBool(ctx.GetText())
	return Value{value: value}
}

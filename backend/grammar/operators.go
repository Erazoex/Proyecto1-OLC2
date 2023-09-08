package grammar

import (
	"proyecto2/parser"
)

/*
Funcion principal para la utilizacion de operadores
*/
func (v *Visitor) VisitOpExpr(ctx *parser.OpExprContext) Value {
	left := v.Visit(ctx.GetLeft())
	right := v.Visit(ctx.GetRight())
	op := ctx.GetOp().GetText()
	switch op {
	case "+":
		return v.Addition(left, right, ctx)
	case "-":
		return v.Subtraction(left, right, ctx)
	case "*":
		return v.Multiplication(left, right, ctx)
	case "/":
		return v.Division(left, right, ctx)
	case "%":
		return v.Module(left, right, ctx)
	case ">=":
		return v.GreaterOrEqual(left, right, ctx)
	case ">":
		return v.Greater(left, right, ctx)
	case "<=":
		return v.LessOrEqual(left, right, ctx)
	case "<":
		return v.Less(left, right, ctx)
	case "==":
		return v.Equal(left, right)
	case "!=":
		return v.NotEqual(left, right)
	case "&&":
		return v.And(left, right, ctx)
	case "||":
		return v.Or(left, right, ctx)
	default:
		return Value{value: false}
	}
}

func (v *Visitor) Addition(leftValue, rightValue Value, ctx *parser.OpExprContext) Value {
	switch leftValue.value.(type) {
	case int64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(int64) + rightValue.value.(int64), Type: INT}
		case float64:
			return Value{value: float64(leftValue.value.(int64)) + rightValue.value.(float64), Type: FLOAT}
		default:
			v.push(error{
				desc:   "no se puede realizar la suma con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	case float64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(float64) + float64(rightValue.value.(int64)), Type: FLOAT}
		case float64:
			return Value{value: leftValue.value.(float64) + rightValue.value.(float64), Type: FLOAT}
		default:
			v.push(error{
				desc:   "no se puede realizar la suma con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	case string:
		switch rightValue.value.(type) {
		case string:
			return Value{value: leftValue.value.(string) + rightValue.value.(string), Type: STRING}
		case byte:
			return Value{value: leftValue.value.(string) + string(rightValue.value.(byte)), Type: STRING}
		case rune:
			return Value{value: leftValue.value.(string) + string(rightValue.value.(rune)), Type: STRING}
		default:
			v.push(error{
				desc:   "no se puede realizar la suma con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	case byte:
		switch rightValue.value.(type) {
		case string:
			return Value{value: string(leftValue.value.(byte)) + rightValue.value.(string), Type: STRING}
		case byte:
			return Value{value: string(leftValue.value.(byte)) + string(rightValue.value.(byte)), Type: STRING}
		case rune:
			return Value{value: string(leftValue.value.(byte)) + string(rightValue.value.(rune)), Type: STRING}
		default:
			v.push(error{
				desc:   "no se puede realizar la suma con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	case rune:
		switch rightValue.value.(type) {
		case string:
			return Value{value: string(leftValue.value.(rune)) + rightValue.value.(string), Type: STRING}
		case rune:
			return Value{value: string(leftValue.value.(rune)) + string(rightValue.value.(rune)), Type: STRING}
		case byte:
			return Value{value: string(leftValue.value.(rune)) + string(rightValue.value.(byte)), Type: STRING}
		default:
			v.push(error{
				desc:   "no se puede realizar la suma con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	default:
		v.push(error{
			desc:   "no se puede realizar suma con expresiones no permitidas",
			line:   ctx.GetOp().GetLine(),
			column: ctx.GetOp().GetColumn(),
		})
		return Value{value: false}
	}
}

func (v *Visitor) Subtraction(leftValue, rightValue Value, ctx *parser.OpExprContext) Value {
	switch leftValue.value.(type) {
	case int64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(int64) - rightValue.value.(int64), Type: INT}
		case float64:
			return Value{value: float64(leftValue.value.(int64)) - rightValue.value.(float64), Type: FLOAT}
		default:
			v.push(error{
				desc:   "no se puede realizar la resta con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	case float64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(float64) - float64(rightValue.value.(int64)), Type: FLOAT}
		case float64:
			return Value{value: leftValue.value.(float64) - rightValue.value.(float64), Type: FLOAT}
		default:
			v.push(error{
				desc:   "no se puede realizar la resta con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	default:
		v.push(error{
			desc:   "no se puede realizar resta con expresiones no permitidas",
			line:   ctx.GetOp().GetLine(),
			column: ctx.GetOp().GetColumn(),
		})
		return Value{value: false}
	}
}

func (v *Visitor) Multiplication(leftValue, rightValue Value, ctx *parser.OpExprContext) Value {
	switch leftValue.value.(type) {
	case int64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(int64) * rightValue.value.(int64), Type: INT}
		case float64:
			return Value{value: float64(leftValue.value.(int64)) * rightValue.value.(float64), Type: FLOAT}
		default:
			v.push(error{
				desc:   "no se puede realizar la multiplicacion con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	case float64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(float64) * float64(rightValue.value.(int64)), Type: FLOAT}
		case float64:
			return Value{value: leftValue.value.(float64) * rightValue.value.(float64), Type: FLOAT}
		default:
			v.push(error{
				desc:   "no se puede realizar la multiplicacion con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	default:
		v.push(error{
			desc:   "no se puede realizar multiplicacion con expresiones no permitidas",
			line:   ctx.GetOp().GetLine(),
			column: ctx.GetOp().GetColumn(),
		})
		return Value{value: false}
	}
}

func (v *Visitor) Division(leftValue, rightValue Value, ctx *parser.OpExprContext) Value {
	switch leftValue.value.(type) {
	case int64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(int64) / rightValue.value.(int64), Type: INT}
		case float64:
			return Value{value: float64(leftValue.value.(int64)) / rightValue.value.(float64), Type: FLOAT}
		default:
			v.push(error{
				desc:   "no se puede realizar la division con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	case float64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(float64) / float64(rightValue.value.(int64)), Type: FLOAT}
		case float64:
			return Value{value: leftValue.value.(float64) / rightValue.value.(float64), Type: FLOAT}
		default:
			v.push(error{
				desc:   "no se puede realizar la division con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	default:
		v.push(error{
			desc:   "no se puede realizar division con expresiones no permitidas",
			line:   ctx.GetOp().GetLine(),
			column: ctx.GetOp().GetColumn(),
		})
		return Value{value: false}
	}
}

func (v *Visitor) Module(leftValue, rightValue Value, ctx *parser.OpExprContext) Value {
	switch leftValue.value.(type) {
	case int64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(int64) % rightValue.value.(int64), Type: INT}
		default:
			v.push(error{
				desc:   "no se puede realizar modulo con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	default:
		v.push(error{
			desc:   "no se puede realizar modulo con expresiones no permitidas",
			line:   ctx.GetOp().GetLine(),
			column: ctx.GetOp().GetColumn(),
		})
		return Value{value: false}
	}
}

func (v *Visitor) GreaterOrEqual(leftValue, rightValue Value, ctx *parser.OpExprContext) Value {
	switch leftValue.value.(type) {
	case int64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(int64) >= rightValue.value.(int64), Type: BOOL}
		case float64:
			return Value{value: float64(leftValue.value.(int64)) >= rightValue.value.(float64), Type: BOOL}
		default:
			v.push(error{
				desc:   "no se puede realizar mayor que o igual con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	case float64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(float64) >= float64(rightValue.value.(int64)), Type: BOOL}
		case float64:
			return Value{value: leftValue.value.(float64) >= rightValue.value.(float64), Type: BOOL}
		default:
			v.push(error{
				desc:   "no se puede realizar mayor que o igual con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	default:
		v.push(error{
			desc:   "no se puede realizar mayor que o igual con expresiones no permitidas",
			line:   ctx.GetOp().GetLine(),
			column: ctx.GetOp().GetColumn(),
		})
		return Value{value: false}
	}
}

func (v *Visitor) Greater(leftValue, rightValue Value, ctx *parser.OpExprContext) Value {
	switch leftValue.value.(type) {
	case int64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(int64) > rightValue.value.(int64), Type: BOOL}
		case float64:
			return Value{value: float64(leftValue.value.(int64)) > rightValue.value.(float64), Type: BOOL}
		default:
			v.push(error{
				desc:   "no se puede realizar mayor con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	case float64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(float64) > float64(rightValue.value.(int64)), Type: BOOL}
		case float64:
			return Value{value: leftValue.value.(float64) > rightValue.value.(float64), Type: BOOL}
		default:
			v.push(error{
				desc:   "no se puede realizar mayor con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	default:
		v.push(error{
			desc:   "no se puede realizar mayor con expresiones no permitidas",
			line:   ctx.GetOp().GetLine(),
			column: ctx.GetOp().GetColumn(),
		})
		return Value{value: false}
	}
}

func (v *Visitor) LessOrEqual(leftValue, rightValue Value, ctx *parser.OpExprContext) Value {
	switch leftValue.value.(type) {
	case int64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(int64) <= rightValue.value.(int64), Type: BOOL}
		case float64:
			return Value{value: float64(leftValue.value.(int64)) <= rightValue.value.(float64), Type: BOOL}
		default:
			v.push(error{
				desc:   "no se puede realizar menor o igual con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	case float64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(float64) <= float64(rightValue.value.(int64)), Type: BOOL}
		case float64:
			return Value{value: leftValue.value.(float64) <= rightValue.value.(float64), Type: BOOL}
		default:
			v.push(error{
				desc:   "no se puede realizar menor o igual con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	default:
		v.push(error{
			desc:   "no se puede realizar menor o igual con expresiones no permitidas",
			line:   ctx.GetOp().GetLine(),
			column: ctx.GetOp().GetColumn(),
		})
		return Value{value: false}
	}
}

func (v *Visitor) Less(leftValue, rightValue Value, ctx *parser.OpExprContext) Value {
	switch leftValue.value.(type) {
	case int64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(int64) < rightValue.value.(int64), Type: BOOL}
		case float64:
			return Value{value: float64(leftValue.value.(int64)) < rightValue.value.(float64), Type: BOOL}
		default:
			v.push(error{
				desc:   "no se puede realizar menor con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	case float64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(float64) < float64(rightValue.value.(int64)), Type: BOOL}
		case float64:
			return Value{value: leftValue.value.(float64) < rightValue.value.(float64), Type: BOOL}
		default:
			v.push(error{
				desc:   "no se puede realizar menor con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{value: false}
		}
	default:
		v.push(error{
			desc:   "no se puede realizar menor con expresiones no permitidas",
			line:   ctx.GetOp().GetLine(),
			column: ctx.GetOp().GetColumn(),
		})
		return Value{value: false}
	}
}

func (v *Visitor) Equal(leftValue, rightValue Value) Value {
	return Value{value: leftValue.value == rightValue.value, Type: BOOL}
}

func (v *Visitor) NotEqual(leftValue, rightValue Value) Value {
	return Value{value: leftValue.value != rightValue.value, Type: BOOL}
}

func (v *Visitor) And(leftValue, rightValue Value, ctx *parser.OpExprContext) Value {
	switch leftValue.value.(type) {
	case bool:
		switch rightValue.value.(type) {
		case bool:
			return Value{value: leftValue.value.(bool) && rightValue.value.(bool), Type: BOOL}
		default:
			// TODO: implementar error
			// la expresion no es de tipo bool
			v.push(error{
				desc:   "no se puede realizar and con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{Type: ERROR}
		}
	default:
		// TODO: implementar error
		// la expresion no es de tipo bool
		v.push(error{
			desc:   "no se puede realizar and con expresiones no permitidas",
			line:   ctx.GetOp().GetLine(),
			column: ctx.GetOp().GetColumn(),
		})
		return Value{Type: ERROR}
	}
}

func (v *Visitor) Or(leftValue, rightValue Value, ctx *parser.OpExprContext) Value {
	switch leftValue.value.(type) {
	case bool:
		switch rightValue.value.(type) {
		case bool:
			return Value{value: leftValue.value.(bool) || rightValue.value.(bool), Type: BOOL}
		default:
			// TODO: implementar error
			// la expresion no es de tipo bool
			v.push(error{
				desc:   "no se puede realizar or con expresiones no permitidas",
				line:   ctx.GetOp().GetLine(),
				column: ctx.GetOp().GetColumn(),
			})
			return Value{Type: ERROR}
		}
	default:
		// TODO: implementar error
		// la expresion no es de tipo bool
		v.push(error{
			desc:   "no se puede realizar or con expresiones no permitidas",
			line:   ctx.GetOp().GetLine(),
			column: ctx.GetOp().GetColumn(),
		})
		return Value{Type: ERROR}
	}
}

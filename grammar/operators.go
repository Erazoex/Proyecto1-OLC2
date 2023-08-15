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
		return Addition(left, right)
	case "-":
		return Subtraction(left, right)
	case "*":
		return Multiplication(left, right)
	case "/":
		return Division(left, right)
	case "%":
		return Module(left, right)
	case ">=":
		return GreaterOrEqual(left, right)
	case ">":
		return Greater(left, right)
	case "<=":
		return LessOrEqual(left, right)
	case "<":
		return Less(left, right)
	case "==":
		return Equal(left, right)
	case "!=":
		return NotEqual(left, right)
	case "&&":
		return And(left, right)
	case "||":
		return Or(left, right)
	default:
		return Value{value: false}
	}
}

func Addition(leftValue, rightValue Value) Value {
	switch leftValue.value.(type) {
	case int64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(int64) + rightValue.value.(int64), Type: INT}
		case float64:
			return Value{value: float64(leftValue.value.(int64)) + rightValue.value.(float64), Type: FLOAT}
		default:
			return Value{value: false}
		}
	case float64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(float64) + float64(rightValue.value.(int64)), Type: FLOAT}
		case float64:
			return Value{value: leftValue.value.(float64) + rightValue.value.(float64), Type: FLOAT}
		default:
			return Value{value: false}
		}
	case string:
		switch rightValue.value.(type) {
		case string:
			return Value{value: leftValue.value.(string) + rightValue.value.(string), Type: STRING}
		default:
			return Value{value: false}
		}
	default:
		return Value{value: false}
	}
}

func Subtraction(leftValue, rightValue Value) Value {
	switch leftValue.value.(type) {
	case int64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(int64) - rightValue.value.(int64), Type: INT}
		case float64:
			return Value{value: float64(leftValue.value.(int64)) - rightValue.value.(float64), Type: FLOAT}
		default:
			return Value{value: false}
		}
	case float64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(float64) - float64(rightValue.value.(int64)), Type: FLOAT}
		case float64:
			return Value{value: leftValue.value.(float64) - rightValue.value.(float64), Type: FLOAT}
		default:
			return Value{value: false}
		}
	default:
		return Value{value: false}
	}
}

func Multiplication(leftValue, rightValue Value) Value {
	switch leftValue.value.(type) {
	case int64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(int64) * rightValue.value.(int64), Type: INT}
		case float64:
			return Value{value: float64(leftValue.value.(int64)) * rightValue.value.(float64), Type: FLOAT}
		default:
			return Value{value: false}
		}
	case float64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(float64) * float64(rightValue.value.(int64)), Type: FLOAT}
		case float64:
			return Value{value: leftValue.value.(float64) * rightValue.value.(float64), Type: FLOAT}
		default:
			return Value{value: false}
		}
	default:
		return Value{value: false}
	}
}

func Division(leftValue, rightValue Value) Value {
	switch leftValue.value.(type) {
	case int64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: float64(leftValue.value.(int64)) / float64(rightValue.value.(int64)), Type: INT}
		case float64:
			return Value{value: float64(leftValue.value.(int64)) / rightValue.value.(float64), Type: FLOAT}
		default:
			return Value{value: false}
		}
	case float64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(float64) / float64(rightValue.value.(int64)), Type: FLOAT}
		case float64:
			return Value{value: leftValue.value.(float64) / rightValue.value.(float64), Type: FLOAT}
		default:
			return Value{value: false}
		}
	default:
		return Value{value: false}
	}
}

func Module(leftValue, rightValue Value) Value {
	switch leftValue.value.(type) {
	case int64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(int64) % rightValue.value.(int64), Type: INT}
		default:
			return Value{value: false}
		}
	default:
		return Value{value: false}
	}
}

func GreaterOrEqual(leftValue, rightValue Value) Value {
	switch leftValue.value.(type) {
	case int64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(int64) >= rightValue.value.(int64), Type: BOOL}
		case float64:
			return Value{value: float64(leftValue.value.(int64)) >= rightValue.value.(float64), Type: BOOL}
		default:
			return Value{value: false}
		}
	case float64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(float64) >= float64(rightValue.value.(int64)), Type: BOOL}
		case float64:
			return Value{value: leftValue.value.(float64) >= rightValue.value.(float64), Type: BOOL}
		default:
			return Value{value: false}
		}
	default:
		return Value{value: false}
	}
}

func Greater(leftValue, rightValue Value) Value {
	switch leftValue.value.(type) {
	case int64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(int64) > rightValue.value.(int64), Type: BOOL}
		case float64:
			return Value{value: float64(leftValue.value.(int64)) > rightValue.value.(float64), Type: BOOL}
		default:
			return Value{value: false}
		}
	case float64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(float64) > float64(rightValue.value.(int64)), Type: BOOL}
		case float64:
			return Value{value: leftValue.value.(float64) > rightValue.value.(float64), Type: BOOL}
		default:
			return Value{value: false}
		}
	default:
		return Value{value: false}
	}
}

func LessOrEqual(leftValue, rightValue Value) Value {
	switch leftValue.value.(type) {
	case int64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(int64) <= rightValue.value.(int64), Type: BOOL}
		case float64:
			return Value{value: float64(leftValue.value.(int64)) <= rightValue.value.(float64), Type: BOOL}
		default:
			return Value{value: false}
		}
	case float64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(float64) <= float64(rightValue.value.(int64)), Type: BOOL}
		case float64:
			return Value{value: leftValue.value.(float64) <= rightValue.value.(float64), Type: BOOL}
		default:
			return Value{value: false}
		}
	default:
		return Value{value: false}
	}
}

func Less(leftValue, rightValue Value) Value {
	switch leftValue.value.(type) {
	case int64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(int64) < rightValue.value.(int64), Type: BOOL}
		case float64:
			return Value{value: float64(leftValue.value.(int64)) < rightValue.value.(float64), Type: BOOL}
		default:
			return Value{value: false}
		}
	case float64:
		switch rightValue.value.(type) {
		case int64:
			return Value{value: leftValue.value.(float64) < float64(rightValue.value.(int64)), Type: BOOL}
		case float64:
			return Value{value: leftValue.value.(float64) < rightValue.value.(float64), Type: BOOL}
		default:
			return Value{value: false}
		}
	default:
		return Value{value: false}
	}
}

func Equal(leftValue, rightValue Value) Value {
	return Value{value: leftValue.value == rightValue, Type: BOOL}
}

func NotEqual(leftValue, rightValue Value) Value {
	return Value{value: leftValue.value != rightValue, Type: BOOL}
}

func And(leftValue, rightValue Value) Value {
	switch leftValue.value.(type) {
	case bool:
		switch rightValue.value.(type) {
		case bool:
			return Value{value: leftValue.value.(bool) && rightValue.value.(bool), Type: BOOL}
		default:
			return Value{value: false}
		}
	default:
		return Value{value: false}
	}
}

func Or(leftValue, rightValue Value) Value {
	switch leftValue.value.(type) {
	case bool:
		switch rightValue.value.(type) {
		case bool:
			return Value{value: leftValue.value.(bool) || rightValue.value.(bool), Type: BOOL}
		default:
			return Value{value: false}
		}
	default:
		return Value{value: false}
	}
}
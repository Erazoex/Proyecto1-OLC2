package grammar

import "proyecto2/parser"

func (v *Visitor) VisitPrimitiveType(ctx *parser.PrimitiveTypeContext) Value {
	tipo := ctx.GetTipo().GetText()
	switch tipo {
	case "Int":
		return Value{Type: INT}
	case "Float":
		return Value{Type: FLOAT}
	case "String":
		return Value{Type: STRING}
	case "Character":
		return Value{Type: CHAR}
	case "Bool":
		return Value{Type: BOOL}
	case "nil":
		return Value{Type: NIL}
	}
	return Value{Type: ERROR}
}

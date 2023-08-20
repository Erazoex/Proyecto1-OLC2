package grammar

import "proyecto2/parser"

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) Value {
	for i := 0; ctx.Stmt(i) != nil; i++ {
		transfer := v.Visit(ctx.Stmt(i))
		if transfer.Type == BREAK || transfer.Type == CONTINUE {
			return transfer
		}
	}
	return Value{value: true, Type: ACCEPTED}
}

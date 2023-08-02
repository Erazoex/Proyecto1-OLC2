package grammar

import "proyecto2/parser"

func (v *Visitor) VisitBlock(ctx *parser.BlockContext) Value {
	for i := 0; ctx.Stmt(i) != nil; i++ {
		v.Visit(ctx.Stmt(i))
	}
	return Value{value: true}
}

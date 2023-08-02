package grammar

import "proyecto2/parser"

type Visitor struct {
	parser.BaseGrammarVisitor
	memory map[string]Value
}

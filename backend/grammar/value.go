package grammar

import "proyecto2/parser"

type Value struct {
	value    interface{}
	Id       string
	Editable bool
	Type     DataType
}

type Function struct {
	Type DataType
	Id   string
	Ctx  *parser.FuncstmtContext
}

type Parameter struct {
	Type DataType
	Id   string
}

type error struct {
	desc   string
	line   int
	column int
}

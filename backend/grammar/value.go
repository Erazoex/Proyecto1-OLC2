package grammar

type Value struct {
	value    interface{}
	Id       string
	Editable bool
	Type     DataType
}

type Function struct {
	Type       DataType
	Id         string
	parametros []any
}

type error struct {
	desc   string
	line   int
	column int
}

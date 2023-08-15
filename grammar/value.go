package grammar

type Value struct {
	value    interface{}
	Id       string
	Editable bool
	Type     DataType
}

type Function func(...interface{}) interface{}

type error struct {
	descripcion string
}

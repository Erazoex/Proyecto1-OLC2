package grammar

type DataType int

const (
	ERROR DataType = iota
	INT
	FLOAT
	STRING
	BOOL
	CHAR
	NIL
	ACCEPTED
)

package grammar

type DataType int

const (
	ERROR DataType = iota
	ACCEPTED
	INT
	FLOAT
	STRING
	BOOL
	CHAR
	NIL
	BREAK
	CONTINUE
	RETURN
	TUPLE
)

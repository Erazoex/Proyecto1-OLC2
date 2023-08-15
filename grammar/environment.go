package grammar

import (
	"fmt"
	"proyecto2/parser"
)

type Visitor struct {
	parser.BaseGrammarVisitor
	environment Environment
	errores     []error
}

type Environment struct {
	tablaSimbolos         map[string]Value
	tablaSimbolos_metodos map[string]Function
	padre                 *Environment
}

func NewEnvironment(padre *Environment) *Environment {
	return &Environment{
		tablaSimbolos:         make(map[string]Value),
		tablaSimbolos_metodos: make(map[string]Function),
		padre:                 padre,
	}
}

func (e *Environment) GetEnvironment() map[string]Value {
	return e.tablaSimbolos
}

func (e *Environment) SaveValue(val Value) bool {
	if e.SearchValue(val.Id) {
		// TODO: implementar error aqui
		fmt.Println("la variable ya existe", val.Id)
		return false
	}
	e.tablaSimbolos[val.Id] = val
	return true
}

func (e *Environment) SearchValue(id string) bool {
	_, ok := e.tablaSimbolos[id]
	return ok
}

func (e *Environment) EditableValue(id string) bool {
	value, ok := e.tablaSimbolos[id]
	if ok && value.Editable {
		return true
	}
	return false
}

func (e *Environment) UpdateValue(val Value) bool {
	if !e.EditableValue(val.Id) {
		// TODO: implementar error aqui
		fmt.Println("la variable es una constante", val.Id)
		return false
	}
	temp := e.tablaSimbolos[val.Id]
	if !(temp.Type == val.Type) && val.Type != NIL {
		// TODO: implementar error aqui
		fmt.Println("la variable", val.Id, "no es del mismo tipo", val.Type)
		return false
	}
	temp.value = val.value
	e.tablaSimbolos[val.Id] = temp
	return true
}

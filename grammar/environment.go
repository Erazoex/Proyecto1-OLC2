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

func (e *Environment) UpdateValue(newValue Value) bool {
	temp := e
	for temp != nil {
		val, ok := temp.tablaSimbolos[newValue.Id]
		if ok && !temp.EditableValue(newValue.Id) {
			// TODO: implementar error aqui
			fmt.Println("la variable es una constante", newValue.Id)
			return false
		}
		if ok && !(val.Type == newValue.Type) && val.Type != NIL {
			// TODO: implememntar error aqui
			fmt.Println("la variable", newValue.Id, "no es del mismo tipo", newValue.Type)
			return false
		}
		if ok {
			val.value = newValue.value
			temp.tablaSimbolos[newValue.Id] = val
			return true
		}
		temp = temp.padre
	}
	// TODO: implementar error aqui
	return false
}

func (e *Environment) UpdateForValue(newValue Value) bool {
	temp := e
	for temp != nil {
		val, ok := temp.tablaSimbolos[newValue.Id]
		if ok && !(val.Type == newValue.Type) && val.Type != NIL {
			// TODO: implememntar error aqui
			fmt.Println("la variable", newValue.Id, "no es del mismo tipo", newValue.Type)
			return false
		}
		if ok {
			val.value = newValue.value
			temp.tablaSimbolos[newValue.Id] = val
			return true
		}
		temp = temp.padre
	}
	// TODO: implementar error aqui
	return false
}

func (e *Environment) GetValue(id string) (Value, bool) {
	var val Value
	var ok bool
	temp := e
	for temp != nil {
		val, ok = temp.tablaSimbolos[id]
		if ok {
			return val, ok
		}
		temp = temp.padre
	}
	return val, ok
}

package grammar

import (
	"fmt"
	"proyecto2/parser"
	"strconv"
	"strings"
)

type Visitor struct {
	parser.BaseGrammarVisitor
	Environment *Environment
	Errores     []error
	Consola     string
}

func (v *Visitor) push(newError error) {
	v.Errores = append(v.Errores, newError)
}

func (v *Visitor) Print(str string) {
	v.Consola += str
}

func (v *Visitor) GetConsole() string {
	return v.Consola
}

type Environment struct {
	tablaSimbolos         map[string]Value
	tablaSimbolos_metodos map[string]Function
	padre                 *Environment
	hijos                 []*Environment
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

func (e *Environment) push(nuevoHijo *Environment) {
	e.hijos = append(e.hijos, nuevoHijo)
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
			fmt.Println("la variable", newValue.Id, "no es del mismo tipo", val.Type, newValue.Type)
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

func (e *Environment) GetValueInCurrentEnvironment(id string) (Value, bool) {
	val, ok := e.tablaSimbolos[id]
	return val, ok
}

// funciones de funcion
func (e *Environment) SaveFunc(newFunc Function) bool {
	if e.SearchFunc(newFunc.Id) {
		// la funcion ya existe
		return false
	}
	e.tablaSimbolos_metodos[newFunc.Id] = newFunc
	return true
}

func (e *Environment) GetFunc(id string) (Function, bool) {
	var val Function
	var ok bool
	temp := e
	for temp != nil {
		val, ok = temp.tablaSimbolos_metodos[id]
		if ok {
			return val, ok
		}
		temp = temp.padre
	}
	return val, ok
}

func (e *Environment) SearchFunc(id string) bool {
	_, ok := e.tablaSimbolos_metodos[id]
	return ok
}

func (e *Environment) generateSymbolTable() string {
	var graph strings.Builder
	if e == nil {
		return ""
	}
	for key, value := range e.GetEnvironment() {
		graph.WriteString("<tr>\n")
		graph.WriteString("<th>")
		graph.WriteString(key)
		graph.WriteString("</th>\n")
		graph.WriteString("<th>")
		graph.WriteString(fmt.Sprintf("%v", value.Type))
		graph.WriteString("</th>\n")
		graph.WriteString("<th>")
		graph.WriteString(strconv.FormatBool(value.Editable))
		graph.WriteString("</th>\n")
		graph.WriteString("<th>")
		graph.WriteString(fmt.Sprintf("%v", value.value))
		graph.WriteString("</th>\n")
		graph.WriteString("</tr>\n")
	}
	for _, newEnv := range e.hijos {
		graph.WriteString("<tr><th colspan=\"4\">Nuevo ambiente generado</th></tr>\n")
		graph.WriteString(newEnv.generateSymbolTable())
	}
	return graph.String()
}

package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type Asignation struct {
	Lin        int
	Col        int
	Id         string
	Expression interfaces.Expression
}

func NewAsignation(lin int, col int, id string, val interfaces.Expression) Asignation {
	asig := Asignation{lin, col, id, val}
	return asig
}

func (p Asignation) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return nil
}

func DeepCopyArray(source interface{}) interface{} {
	switch source := source.(type) {
	case []interface{}:
		copyArray := make([]interface{}, len(source))
		for i, val := range source {
			copyArray[i] = DeepCopyArray(val)
		}
		return copyArray
	default:
		return source
	}
}

package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
	"strconv"
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

func (p Asignation) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	gen.AddComment("Generando asignacion")
	//buscando variable en entorno
	variable := env.(environment.Environment).FindVar(p.Id)
	//ejecutando valor
	result = p.Expression.Execute(ast, env, gen)
	//realizando asignacion
	gen.AddSetStack(strconv.Itoa(variable.Position), result.Value)
	gen.AddBr()
	return result
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

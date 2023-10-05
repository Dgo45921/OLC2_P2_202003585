package expressions

import (
	"PY1/environment"
	"PY1/generator"
)

type StructExp struct {
	Lin    int
	Col    int
	ID     string
	Fields []environment.KeyValue
}

func NewStructExp(lin int, col int, id string, accesses []environment.KeyValue) StructExp {
	structaccess := StructExp{lin, col, id, accesses}
	return structaccess
}

func (p StructExp) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}

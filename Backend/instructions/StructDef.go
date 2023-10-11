package instructions

import (
	"PY1/environment"
	"PY1/generator"
)

type StructDef struct {
	Lin      int
	Col      int
	Id       string
	insBlock []interface{}
	Type     environment.TipoExpresion
}

func NewStructDef(lin int, col int, id string, insBlock []interface{}) StructDef {
	return StructDef{lin, col, id, insBlock, environment.STRUCT_DEF}
}

func (p StructDef) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	return environment.Value{}
}

func repeatedValue(id string, arraykevalue []environment.KeyValue) bool {
	for _, kv := range arraykevalue {
		if id == kv.Key {
			return true
		}
	}

	return false
}

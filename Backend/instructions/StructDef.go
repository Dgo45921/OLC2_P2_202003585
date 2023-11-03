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
	if env.(environment.Environment).Prev != nil {
		ast.SetError(p.Lin, p.Col, "Los structs solo pueden ser declarados en el ambito global")
		return environment.Value{}
	}

	structMap := []environment.KeyValue{}

	for _, inst := range p.insBlock {

		if _, isVarDec := inst.(VarDec); isVarDec {
			response := inst.(VarDec).GetVarDec(ast, env)
			if response != nil {
				if !repeatedValue(inst.(VarDec).Id, structMap) {
					newKeyValuePair := environment.KeyValue{inst.(VarDec).Id, response}
					structMap = append(structMap, newKeyValuePair)
				} else {
					ast.SetError(p.Lin, p.Col, "atributo var en struct repetido")
					return environment.Value{}
				}

			} else {
				ast.SetError(p.Lin, p.Col, "El tipo de asignacion a un atributo var no fue válida")
				return environment.Value{}

			}

		}

	}

	var newEnv = environment.NewEnvironment(env, environment.STRUCT)
	newstruct := environment.Symbol{
		Lin:        p.Lin,
		Col:        p.Col,
		Value:      structMap,
		Type:       p.Type,
		StructType: p.Id,
		StructEnv:  newEnv,
	}

	env.(environment.Environment).SaveStruct(p.Id, newstruct)
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

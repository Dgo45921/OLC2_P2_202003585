package expressions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
	"strconv"
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

	typeStruct := env.(environment.Environment).FindVar(p.ID)
	envStruct := typeStruct.StructEnv

	finalStruct := []environment.KeyValue{}

	if typeStruct.Type == environment.STRUCT_DEF {

		newfields := make([]environment.KeyValue, len(p.Fields))

		gen.AddComment("definiendo valores de instancia: " + p.ID)

		for index, _ := range p.Fields {
			if _, isBreak := p.Fields[index].Value.(interfaces.Expression); isBreak {
				keyy := p.Fields[index].Key

				dec := CreateVarDec(ast, envStruct, gen, keyy, p.Lin, p.Col, p.Fields[index].Value.(interfaces.Expression))

				newfields[index].Key = keyy
				newfields[index].Value = dec
			}
		}

		for index, kv := range newfields {
			if _, isBreak := kv.Value.(StructExp); isBreak {
				skrr := kv.Value.(interfaces.Expression).Execute(ast, envStruct, gen)
				if !valTypeOk(kv.Key, typeStruct.Value.([]environment.KeyValue), skrr.Type) {
					ast.SetPrint("Error: Argumento pasado no coincide con el definido por el struct!\n")
					return environment.Value{}
				}
				newfields[index].Value = skrr
			} else if _, isBreak := kv.Value.(interfaces.Expression); isBreak {
				skrr := kv.Value.(interfaces.Expression).Execute(ast, envStruct, gen)
				if !valTypeOk(kv.Key, typeStruct.Value.([]environment.KeyValue), skrr.Type) {
					ast.SetPrint("Error: Argumento pasado no coincide con el definido por el struct!\n")
					return environment.Value{}
				}
			}

		}
		xd := addDefaultValues(newfields, typeStruct.Value.([]environment.KeyValue))
		finalStruct = append(finalStruct, xd...)
		// envStruct.SaveVariableStruct(p.ID, environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.STRUCT_IMP, Value: finalStruct, StructType: p.ID})
		return environment.Value{StructValues: finalStruct, Type: environment.STRUCT_IMP, Id: p.ID}

	} else {
		ast.SetPrint("Error: el struct: " + p.ID + " no existe!\n")
	}

	return result
}

func executeAllKeyValues(val []environment.KeyValue, ast *environment.AST, env interface{}, gen *generator.Generator) {
	for index, _ := range val {
		val[index].Value = val[index].Value.(interfaces.Expression).Execute(ast, env, gen).Value
	}

}

func addDefaultValues(fields []environment.KeyValue, defaults []environment.KeyValue) []environment.KeyValue {

	for index, d := range defaults {
		found := false
		for _, f := range fields {
			if f.Key == d.Key {
				found = true
				break
			}
		}
		if !found {

			valToAdd := defaults[index]
			fields = append(fields, valToAdd)
		}

	}

	newFields := fields
	return newFields

}

func valTypeOk(targetField string, fields []environment.KeyValue, typee environment.TipoExpresion) bool {
	for _, kv := range fields {

		if kv.Key == targetField {
			if kv.Value.(environment.Symbol).Type != typee {
				return false
			}

			if kv.Value.(environment.Symbol).Const {
				if kv.Value.(environment.Symbol).Value == nil {
					return true
				}
				return false

			}
			return true

		}

	}

	return true
}

func CreateVarDec(ast *environment.AST, env interface{}, gen *generator.Generator, Id string, Lin int, Col int, exp interfaces.Expression) environment.Value {
	if env.(environment.Environment).VariableExists(Id) {
		ast.SetError(Lin, Col, "Error, variable ya declarada!")
		return environment.Value{}
	}
	var result environment.Value
	var newVar environment.Symbol
	result = exp.Execute(ast, env, gen)
	if result.Type == environment.STRUCT_IMP {

		return result
	}
	gen.AddComment("Agregando una declaracion")
	newVar = env.(environment.Environment).SaveVariableStructArg(Id, result.Type)
	extra := result
	extra.Id = Id
	extra.Scope = env.(environment.Environment).Scope
	extra.Lin = Lin
	extra.Col = Col
	ast.SaveSymbol(Id, extra)

	if result.Type == environment.BOOLEAN {
		//si no es temp (boolean)
		newLabel := gen.NewLabel()
		//add labels
		for _, lvl := range result.TrueLabel {
			gen.AddLabel(lvl.(string))
		}
		gen.AddSetStack(strconv.Itoa(newVar.Position), "1")
		gen.AddGoto(newLabel)
		//add labels
		for _, lvl := range result.FalseLabel {
			gen.AddLabel(lvl.(string))
		}
		gen.AddSetStack(strconv.Itoa(newVar.Position), "0")
		gen.AddGoto(newLabel)
		gen.AddLabel(newLabel)
		gen.AddBr()
	} else {
		//si es temp (num,string,etc)
		gen.AddSetStack(strconv.Itoa(newVar.Position), result.Value)
		gen.AddBr()
	}

	return result
}

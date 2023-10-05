package expressions

import (
	"PY1/environment"
	"PY1/interfaces"
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

func (p StructExp) Execute(ast *environment.AST, env interface{}) environment.Symbol {

	typeStruct := env.(environment.Environment).FindVar(p.ID)
	finalStruct := []environment.KeyValue{}

	if typeStruct.Type == environment.STRUCT_DEF {
		correctOrder := checkResponseOrder(p.Fields, typeStruct.Value.([]environment.KeyValue))
		if !correctOrder {
			ast.SetPrint("Error: los argumentos para crear el struct no estan en orden o hace falta uno!\n")
			return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: nil}
		}

		newfields := make([]environment.KeyValue, len(p.Fields))

		for index, _ := range p.Fields {
			if _, isBreak := p.Fields[index].Value.(interfaces.Expression); isBreak {
				vall := p.Fields[index].Value.(interfaces.Expression).Execute(ast, env)
				keyy := p.Fields[index].Key
				newfields[index].Key = keyy
				newfields[index].Value = vall
			}
		}

		for index, kv := range newfields {
			if _, isBreak := kv.Value.(StructExp); isBreak {
				skrr := kv.Value.(interfaces.Expression).Execute(ast, env)
				if !valTypeOk(kv.Key, typeStruct.Value.([]environment.KeyValue), skrr.Type) {
					ast.SetPrint("Error: Argumento pasado no coincide con el definido por el struct!\n")
					return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: nil}
				}
				newfields[index].Value = skrr
			} else if _, isBreak := kv.Value.(interfaces.Expression); isBreak {
				skrr := kv.Value.(interfaces.Expression).Execute(ast, env)
				if !valTypeOk(kv.Key, typeStruct.Value.([]environment.KeyValue), skrr.Type) {
					ast.SetPrint("Error: Argumento pasado no coincide con el definido por el struct!\n")
					return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: nil}
				}
			}

		}
		xd := addDefaultValues(newfields, typeStruct.Value.([]environment.KeyValue))
		finalStruct = append(finalStruct, xd...)
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.STRUCT_IMP, Value: finalStruct, StructType: p.ID}

	} else {
		ast.SetPrint("Error: el struct: " + p.ID + " no existe!\n")
	}

	return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: nil}
}

func executeAllKeyValues(val []environment.KeyValue, ast *environment.AST, env interface{}) {
	for index, _ := range val {
		val[index].Value = val[index].Value.(interfaces.Expression).Execute(ast, env).Value
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

func checkResponseOrder(responses []environment.KeyValue, fields []environment.KeyValue) bool {
	requiredIndex := 0

	for _, r := range responses {
		if requiredIndex >= len(fields) {
			return false
		}

		for fields[requiredIndex].Key != r.Key {
			if _, isSymbol := fields[requiredIndex].Value.(environment.Symbol); isSymbol {
				if fields[requiredIndex].Value.(environment.Symbol).Value == nil {
					return false
				}

			} else if _, isPrimitive := fields[requiredIndex].Value.(Primitive); isPrimitive {
				if fields[requiredIndex].Value.(Primitive).Value == nil {
					return false
				}
			}

			requiredIndex++
		}

		requiredIndex++
	}

	for _, f := range fields[requiredIndex:] {
		if _, isSymbol := f.Value.(environment.Symbol); isSymbol {
			if f.Value.(environment.Symbol).Value == nil {
				return false
			}

		} else if _, isPrimitive := f.Value.(Primitive); isPrimitive {
			if f.Value.(Primitive).Value == nil {
				return false
			}

		}

	}

	return true
}

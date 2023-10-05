package instructions

import (
	"PY1/environment"
	"PY1/expressions"
	"PY1/interfaces"
)

type StructFuncCall struct {
	Lin    int
	Col    int
	Id     string
	FuncId string
}

func NewStructFuncCall(lin int, col int, id string, funcid string) StructFuncCall {
	return StructFuncCall{lin, col, id, funcid}
}

func (p StructFuncCall) Execute(ast *environment.AST, env interface{}) interface{} {
	if !env.(environment.Environment).VariableExists(p.Id) {
		ast.SetError(p.Lin, p.Col, "Struct no encontrado para ejecutar metodo")
	}

	foundStruct := env.(environment.Environment).FindVar(p.Id)

	accessArray := []string{p.FuncId}
	foundFunc := GetValueByArray(accessArray, foundStruct)
	if foundFunc != nil {

		if _, isBreak := foundFunc.(environment.FunctionSymbol); isBreak {
			// run the function code
			for _, inst := range foundFunc.(environment.FunctionSymbol).InsBlock {

				if _, isPrint := inst.(Print); isPrint {
					for index, val := range inst.(Print).Value {
						if _, isselfaccess := val.(expressions.SelfAccess); isselfaccess {
							symbolin := val.(expressions.SelfAccess).Execute(ast, env)
							ahh := GetValueByArray([]string{symbolin.Value.(string)}, foundStruct)
							if ahh != nil {

								inst.(Print).Value[index] = ahh

							}

						}
					}
				}

				response := inst.(interfaces.Instruction).Execute(ast, env)
				if _, ismod := response.(SelfModification); ismod {
					if !foundFunc.(environment.FunctionSymbol).Mutating {
						ast.SetError(p.Lin, p.Col, "No se puede modificar atributo de struct en funcion no mutating")
						return nil
					}

					accessArray2 := []string{response.(SelfModification).Id}
					err := UpdateValueByArray(accessArray2, foundStruct, response.(SelfModification).NewValue.Execute(ast, env), env)
					if err != nil {
						ast.SetError(p.Lin, p.Col, "No se pudo modificar el atributo")
					}

				}

			}

		} else {
			ast.SetError(p.Lin, p.Col, "El atributo de struct no es una funcion")
		}

	} else {
		ast.SetError(p.Lin, p.Col, "No se hallo ninguna funcino en el struct con ese nombre")
	}

	return nil
}

func GetValueByArray(arr []string, symbol environment.Symbol) interface{} {
	var currentSymbol = symbol

	for _, key := range arr {
		found := false
		if kvArr, ok := currentSymbol.Value.([]environment.KeyValue); ok {
			for _, kv := range kvArr {
				if kv.Key == key {
					if _, isBreak := kv.Value.(environment.Symbol); isBreak {
						currentSymbol = kv.Value.(environment.Symbol)
					} else if _, isBreak := kv.Value.(environment.FunctionSymbol); isBreak {
						return kv.Value.(environment.FunctionSymbol)
					}
					found = true
					break
				}
			}
		}

		if !found {
			return nil
		}
	}

	return currentSymbol
}

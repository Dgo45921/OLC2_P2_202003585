package expressions

import (
	"PY1/environment"
	"PY1/interfaces"
	"strings"
)

type CallFuncExp struct {
	Lin        int
	Col        int
	Id         string
	Parameters []environment.FuncArg
}

func NewCallFuncExp(lin int, col int, val string, param []environment.FuncArg) CallFuncExp {
	exp := CallFuncExp{lin, col, val, param}
	return exp
}

func (p CallFuncExp) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	foundFunc, exists := env.(environment.Environment).FindFunc(p.Id)

	if !exists {
		ast.SetError(p.Lin, p.Col, "Funcion no existente")
		return environment.Symbol{
			Lin:   p.Lin,
			Col:   p.Col,
			Value: nil,
		}

	}
	// checking length of parameters and arguments
	if len(p.Parameters) != len(foundFunc.Args) {
		ast.SetError(p.Lin, p.Col, "Cantidad de argumentos no coincide con la cantidad de parametros")
		return environment.Symbol{
			Lin:   p.Lin,
			Col:   p.Col,
			Value: nil,
		}
	}
	newEnv := environment.NewEnvironment(env, environment.FUNC)
	// check array of values and types
	for index, val := range p.Parameters {
		if val.Id == "" && val.Reference == true && val.RealId != "" {
			val.Id = foundFunc.Args[index].SID
			p.Parameters[index] = val
		}

	}
	// check array of values and types
	for index := range p.Parameters {
		valParameter := p.Parameters[index].Value.(interfaces.Expression).Execute(ast, env)
		valParameter.Scope = newEnv.Scope
		if valParameter.Type == environment.VECTOR_STRING || valParameter.Type == environment.VECTOR_STRUCT || valParameter.Type == environment.VECTOR_CHAR || valParameter.Type == environment.VECTOR_FLOAT || valParameter.Type == environment.VECTOR_BOOLEAN || valParameter.Type == environment.VECTOR_INT || valParameter.Type == environment.VECTOR || valParameter.Type == environment.MATRIX_INT || valParameter.Type == environment.MATRIX_FLOAT || valParameter.Type == environment.MATRIX_BOOLEAN || valParameter.Type == environment.MATRIX_CHAR {
			valParameter.Value = DeepCopyArray(valParameter.Value)
		}
		if foundFunc.Args[index].Id == "_" {
			if getTypeByString(p.Lin, p.Col, foundFunc.Args[index].Type, ast, env, p.Parameters[index].Value.(interfaces.Expression)) == valParameter.Type {
				isByReference := foundFunc.Args[index].Reference
				if isByReference == p.Parameters[index].Reference {
					if isByReference {
						if env.(environment.Environment).VariableExists(p.Parameters[index].RealId) || env.(environment.Environment).ReferenceExists(p.Parameters[index].RealId) {
							newEnv.SaveReference(foundFunc.Args[index].SID, valParameter)
							ast.SaveSymbol(foundFunc.Args[index].SID, valParameter)
						} else {
							ast.SetError(p.Lin, p.Col, "La referencia solo funciona con variables")
						}

					} else {
						pivote := valParameter
						pivote.Scope = newEnv.Scope
						newEnv.SaveVariable(foundFunc.Args[index].SID, pivote)
						ast.SaveSymbol(foundFunc.Args[index].SID, pivote)

					}

				} else {
					ast.SetError(p.Lin, p.Col, "atributos definidos como valor por ref o por valor, no coinciden")
					return environment.Symbol{
						Lin:   p.Lin,
						Col:   p.Col,
						Value: nil,
					}
				}

			} else {
				ast.SetError(p.Lin, p.Col, "tipo de parametro no coincide con el argumento enviado")
				return environment.Symbol{
					Lin:   p.Lin,
					Col:   p.Col,
					Value: nil,
				}

			}

		} else {
			exists, indexx := checkIfParameterExists(foundFunc.Args, p.Parameters[index].Id)
			if exists {

				if getTypeByString(p.Lin, p.Col, foundFunc.Args[indexx].Type, ast, env, p.Parameters[index].Value.(interfaces.Expression)) == valParameter.Type {
					isByReference := foundFunc.Args[indexx].Reference
					if isByReference == p.Parameters[index].Reference {
						if isByReference {
							if env.(environment.Environment).VariableExists(p.Parameters[index].RealId) || env.(environment.Environment).ReferenceExists(p.Parameters[index].RealId) {
								newEnv.SaveReference(foundFunc.Args[index].SID, valParameter)
								ast.SaveSymbol(foundFunc.Args[index].SID, valParameter)
							} else {
								ast.SetError(p.Lin, p.Col, "la referencia solo funciona con variables")
							}

						} else {
							pivote := valParameter
							pivote.Scope = newEnv.Scope
							newEnv.SaveVariable(foundFunc.Args[index].SID, pivote)
							ast.SaveSymbol(foundFunc.Args[index].SID, pivote)

						}

					} else {
						ast.SetError(p.Lin, p.Col, "atributos definidos como valor por ref o por valor, no coinciden")
						return environment.Symbol{
							Lin:   p.Lin,
							Col:   p.Col,
							Value: nil,
						}
					}

				} else {
					ast.SetError(p.Lin, p.Col, "tipo de parametro no coincide con el argumento enviado")
					return environment.Symbol{
						Lin:   p.Lin,
						Col:   p.Col,
						Value: nil,
					}

				}

			}

		}

	}

	// setting up the function
	for _, inst := range foundFunc.InsBlock {
		// is not any of that cases
		var response = inst.(interfaces.Instruction).Execute(ast, newEnv)
		if response != nil {
			if _, isReturn := response.(environment.Symbol); isReturn {
				valretorno := response.(environment.Symbol)
				if valretorno.Type == foundFunc.ReturnType {
					if valretorno.Type == environment.STRUCT_IMP {
						founstructdef := newEnv.FindVar(valretorno.StructType)
						if founstructdef.Type == environment.STRUCT_DEF && founstructdef.StructType == valretorno.StructType {
							for index, parameter := range p.Parameters {
								if parameter.Reference {
									_, indexx := checkIfParameterExists(foundFunc.Args, p.Parameters[index].Id)
									newEnv.SetReferenceValues(parameter.RealId, foundFunc.Args[indexx].SID)
								}

							}
							return valretorno
						} else {
							ast.SetError(p.Lin, p.Col, "tipo de retorno no coincide por el definido para la funcion")
							return environment.Symbol{
								Lin:   p.Lin,
								Col:   p.Col,
								Value: nil,
							}
						}
					}

					for index, parameter := range p.Parameters {
						if parameter.Reference {
							_, indexx := checkIfParameterExists(foundFunc.Args, p.Parameters[index].Id)
							newEnv.SetReferenceValues(parameter.RealId, foundFunc.Args[indexx].SID)
						}
					}

					return valretorno

				} else {
					ast.SetError(p.Lin, p.Col, "El tipo de retorno definido en la funcion no coincide con el valor del return")
					return environment.Symbol{
						Lin:   p.Lin,
						Col:   p.Col,
						Value: nil,
					}

				}

			}
		} else {
			continue
		}

	}
	if foundFunc.ReturnType == environment.NULL {
		for index, parameter := range p.Parameters {
			if parameter.Reference {
				_, indexx := checkIfParameterExists(foundFunc.Args, p.Parameters[index].Id)
				newEnv.SetReferenceValues(parameter.RealId, foundFunc.Args[indexx].SID)
			}
		}

	} else {
		ast.SetError(p.Lin, p.Col, "La funcion no tiene un return con el tipo de dato definido")
	}

	return environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Value: nil,
	}

}

func getTypeByString(lin int, col int, val string, ast *environment.AST, env interface{}, expression interfaces.Expression) environment.TipoExpresion {
	if val == "Int" {
		return environment.INTEGER
	} else if val == "Float" {
		return environment.FLOAT
	} else if val == "Bool" {
		return environment.BOOLEAN
	} else if val == "String" {
		return environment.STRING
	} else if val == "Character" {
		return environment.CHAR
	} else if strings.Contains(val, "[") {
		structExp := expression.Execute(ast, env)
		if _, isBreak := structExp.Value.([]interface{}); !isBreak {
			ast.SetError(lin, col, "tipo de parametro no coincide con el argumento enviado")
			return environment.NULL

		}
		depth := GetDepth(structExp.Value.([]interface{}))

		if depth == countCharOccurrences(val, '[') {
			if depth == 1 {
				if strings.Contains(val, "Int") {
					return environment.VECTOR_INT
				} else if strings.Contains(val, "Float") {
					return environment.VECTOR_FLOAT
				} else if strings.Contains(val, "Bool") {
					return environment.VECTOR_BOOLEAN
				} else if strings.Contains(val, "String") {
					return environment.VECTOR_STRING
				} else if strings.Contains(val, "Character") {
					return environment.VECTOR_CHAR
				} else {
					StructType := strings.ReplaceAll(strings.ReplaceAll(val, "[", ""), "]", "")
					structcase := env.(environment.Environment).FindVar(StructType)
					if structcase.Type == environment.STRUCT_DEF && structExp.StructType == val {
						return environment.STRUCT_IMP
					} else {
						return environment.NULL
					}
				}

			} else {

				if strings.Contains(val, "Int") {
					return environment.MATRIX_INT
				} else if strings.Contains(val, "Float") {
					return environment.MATRIX_FLOAT
				} else if strings.Contains(val, "Bool") {
					return environment.MATRIX_BOOLEAN
				} else if strings.Contains(val, "String") {
					return environment.MATRIX_STRING
				} else if strings.Contains(val, "Character") {
					return environment.MATRIX_CHAR
				} else {
					return environment.NULL
				}
			}

		}
		return environment.NULL

	} else {
		structExp := expression.Execute(ast, env)
		structcase := env.(environment.Environment).FindVar(val)
		if structcase.Type == environment.STRUCT_DEF && structExp.StructType == val {
			return environment.STRUCT_IMP
		} else {
			return environment.NULL
		}
	}

}
func GetDepth(arr []interface{}) int {
	if len(arr) == 0 {
		return 1
	}

	maxDepth := 0
	for _, item := range arr {
		if nestedArr, ok := item.([]interface{}); ok {
			depth := GetDepth(nestedArr)
			if depth > maxDepth {
				maxDepth = depth
			}
		}
	}

	return maxDepth + 1
}

func checkIfParameterExists(arr []environment.FuncParam, str string) (bool, int) {
	for index, element := range arr {
		if element.Id == str {
			return true, index
		}

	}

	return false, 0
}

func checkIfParameterReferenceExists(arr []environment.FuncParam, str string) (bool, int) {
	for index, element := range arr {
		if element.SID == str {
			return true, index
		}

	}

	return false, 0
}

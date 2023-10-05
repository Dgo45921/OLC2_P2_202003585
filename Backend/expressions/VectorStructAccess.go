package expressions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type VectorStructAccess struct {
	Lin      int
	Col      int
	IDvector string
	Index    []interface{}
	Accesses []string
}

func NewVectorStructAccess(lin int, col int, id string, index []interface{}, acc []string) VectorStructAccess {
	exp := VectorStructAccess{lin, col, id, index, acc}
	return exp
}

func (p VectorStructAccess) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	if env.(environment.Environment).VariableExists(p.IDvector) {
		foundVec := env.(environment.Environment).FindVar(p.IDvector)
		if foundVec.Type == environment.VECTOR {
			ast.SetError(p.Lin, p.Col, "Vector vacio")
			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Value: nil,
			}
		}

		var indexes = GetIndexes(p.Index, ast, env)

		if !AllNonNegativeIntegers(indexes) {
			ast.SetError(p.Lin, p.Col, "indices deben de ser enteros mayores o iguales a 0")
			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Value: nil,
			}
		}

		if _, isArray := foundVec.Value.([]interface{}); isArray {
			if foundVec.Type == environment.VECTOR_STRING || foundVec.Type == environment.VECTOR_STRUCT || foundVec.Type == environment.VECTOR_CHAR || foundVec.Type == environment.VECTOR_FLOAT || foundVec.Type == environment.VECTOR_BOOLEAN || foundVec.Type == environment.VECTOR_INT || foundVec.Type == environment.MATRIX_INT || foundVec.Type == environment.MATRIX_FLOAT || foundVec.Type == environment.MATRIX_STRING || foundVec.Type == environment.MATRIX_BOOLEAN || foundVec.Type == environment.MATRIX_CHAR {

				val, exists := GetIndexValue(foundVec.Value, indexes)

				if !exists {
					ast.SetError(p.Lin, p.Col, "indice no existente")
					return environment.Symbol{
						Lin:   p.Lin,
						Col:   p.Col,
						Value: nil,
					}
				}

				var accesstype = environment.INTEGER
				if foundVec.Value == environment.VECTOR_INT || foundVec.Value == environment.MATRIX_INT {
					accesstype = environment.INTEGER
				} else if foundVec.Value == environment.VECTOR_FLOAT || foundVec.Value == environment.MATRIX_FLOAT {
					accesstype = environment.FLOAT
				} else if foundVec.Value == environment.VECTOR_STRING || foundVec.Value == environment.MATRIX_STRING {
					accesstype = environment.STRING
				} else if foundVec.Value == environment.VECTOR_CHAR || foundVec.Value == environment.MATRIX_CHAR {
					accesstype = environment.CHAR
				} else if foundVec.Value == environment.VECTOR_BOOLEAN || foundVec.Value == environment.MATRIX_BOOLEAN {
					accesstype = environment.BOOLEAN
				} else if foundVec.Value == environment.VECTOR_STRUCT {
					accesstype = environment.STRUCT_IMP
				}

				if _, isBreak := val.([]environment.KeyValue); isBreak {

					for _, acc := range p.Accesses {
						if _, isArrayKeyValue := val.([]environment.KeyValue); isArrayKeyValue {
							for _, kv := range val.([]environment.KeyValue) {
								if kv.Key == acc {
									result, err := searchNestedValue(kv, p.Accesses)
									if err != nil {
										return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: nil}
									} else if _, isSymbol := result.(environment.Symbol); isSymbol {
										return result.(environment.Symbol)
									} else {
										result = result.(interfaces.Expression).Execute(ast, env)
										return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: result.(environment.Symbol).Type, Value: result.(environment.Symbol).Value}
									}

								}

							}
						}

					}

				} else {
					return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: nil}

				}
				return environment.Symbol{
					Lin:   p.Lin,
					Col:   p.Col,
					Value: val,
					Type:  accesstype,
					Const: false,
				}

			} else {
				ast.SetError(p.Lin, p.Col, "el acceso [] solo funciona con vectores o matrices")
				return environment.Symbol{
					Lin:   p.Lin,
					Col:   p.Col,
					Value: nil,
				}

			}

		}
		ast.SetError(p.Lin, p.Col, "el acceso [] solo funciona con vectores o matrices")
		return environment.Symbol{
			Lin:   p.Lin,
			Col:   p.Col,
			Value: nil,
		}

	} else if env.(environment.Environment).ReferenceExists(p.IDvector) {
		foundVec := env.(environment.Environment).FindReference(p.IDvector)
		if foundVec.Type == environment.VECTOR {
			ast.SetError(p.Lin, p.Col, "Vector vacio")
			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Value: nil,
			}
		}

		var indexes = GetIndexes(p.Index, ast, env)

		if !AllNonNegativeIntegers(indexes) {
			ast.SetError(p.Lin, p.Col, "indices deben de ser enteros mayores o iguales a 0")
			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Value: nil,
			}
		}

		if _, isArray := foundVec.Value.([]interface{}); isArray {
			if foundVec.Type == environment.VECTOR_STRING || foundVec.Type == environment.VECTOR_STRUCT || foundVec.Type == environment.VECTOR_CHAR || foundVec.Type == environment.VECTOR_FLOAT || foundVec.Type == environment.VECTOR_BOOLEAN || foundVec.Type == environment.VECTOR_INT || foundVec.Type == environment.MATRIX_INT || foundVec.Type == environment.MATRIX_FLOAT || foundVec.Type == environment.MATRIX_STRING || foundVec.Type == environment.MATRIX_BOOLEAN || foundVec.Type == environment.MATRIX_CHAR {

				val, exists := GetIndexValue(foundVec.Value, indexes)

				if !exists {
					ast.SetError(p.Lin, p.Col, "indice no existente")
					return environment.Symbol{
						Lin:   p.Lin,
						Col:   p.Col,
						Value: nil,
					}
				}

				var accesstype = environment.INTEGER
				if foundVec.Value == environment.VECTOR_INT || foundVec.Value == environment.MATRIX_INT {
					accesstype = environment.INTEGER
				} else if foundVec.Value == environment.VECTOR_FLOAT || foundVec.Value == environment.MATRIX_FLOAT {
					accesstype = environment.FLOAT
				} else if foundVec.Value == environment.VECTOR_STRING || foundVec.Value == environment.MATRIX_STRING {
					accesstype = environment.STRING
				} else if foundVec.Value == environment.VECTOR_CHAR || foundVec.Value == environment.MATRIX_CHAR {
					accesstype = environment.CHAR
				} else if foundVec.Value == environment.VECTOR_BOOLEAN || foundVec.Value == environment.MATRIX_BOOLEAN {
					accesstype = environment.BOOLEAN
				} else if foundVec.Value == environment.VECTOR_STRUCT {
					accesstype = environment.STRUCT_IMP
				}

				if _, isBreak := val.([]environment.KeyValue); isBreak {

					for _, acc := range p.Accesses {
						if _, isArrayKeyValue := val.([]environment.KeyValue); isArrayKeyValue {
							for _, kv := range val.([]environment.KeyValue) {
								if kv.Key == acc {
									result, err := searchNestedValue(kv, p.Accesses)
									if err != nil {
										return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: nil}
									} else if _, isSymbol := result.(environment.Symbol); isSymbol {
										return result.(environment.Symbol)
									} else {
										result = result.(interfaces.Expression).Execute(ast, env)
										return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: result.(environment.Symbol).Type, Value: result.(environment.Symbol).Value}
									}

								}

							}
						}

					}

				} else {
					return environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.NULL, Value: nil}

				}
				return environment.Symbol{
					Lin:   p.Lin,
					Col:   p.Col,
					Value: val,
					Type:  accesstype,
					Const: false,
				}

			} else {
				ast.SetError(p.Lin, p.Col, "el acceso [] solo funciona con vectores o matrices")
				return environment.Symbol{
					Lin:   p.Lin,
					Col:   p.Col,
					Value: nil,
				}

			}

		}
		ast.SetError(p.Lin, p.Col, "el acceso [] solo funciona con vectores o matrices")
		return environment.Symbol{
			Lin:   p.Lin,
			Col:   p.Col,
			Value: nil,
		}
	}

	return environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Value: nil,
	}

}

package expressions

import (
	"PY1/environment"
	"PY1/interfaces"
	"reflect"
)

type VectorAccess struct {
	Lin   int
	Col   int
	Id    string
	Index []interface{}
}

func NewVectorAccess(lin int, col int, id string, index []interface{}) VectorAccess {
	exp := VectorAccess{lin, col, id, index}
	return exp
}

func (p VectorAccess) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	if env.(environment.Environment).VariableExists(p.Id) {
		foundVar := env.(environment.Environment).FindVar(p.Id)
		if foundVar.Type == environment.VECTOR {
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

		if _, isArray := foundVar.Value.([]interface{}); isArray {
			if foundVar.Type == environment.VECTOR_STRING || foundVar.Type == environment.VECTOR_STRUCT || foundVar.Type == environment.VECTOR_CHAR || foundVar.Type == environment.VECTOR_FLOAT || foundVar.Type == environment.VECTOR_BOOLEAN || foundVar.Type == environment.VECTOR_INT || foundVar.Type == environment.MATRIX_INT || foundVar.Type == environment.MATRIX_FLOAT || foundVar.Type == environment.MATRIX_STRING || foundVar.Type == environment.MATRIX_BOOLEAN || foundVar.Type == environment.MATRIX_CHAR {

				val, exists := GetIndexValue(foundVar.Value, indexes)

				if !exists {
					ast.SetError(p.Lin, p.Col, "indice no existente")
					return environment.Symbol{
						Lin:   p.Lin,
						Col:   p.Col,
						Value: nil,
					}
				}

				var accesstype = environment.INTEGER
				if foundVar.Type == environment.VECTOR_INT {
					accesstype = environment.INTEGER
				} else if foundVar.Type == environment.VECTOR_FLOAT {
					accesstype = environment.FLOAT
				} else if foundVar.Type == environment.VECTOR_STRING {
					accesstype = environment.STRING
				} else if foundVar.Type == environment.VECTOR_CHAR {
					accesstype = environment.CHAR
				} else if foundVar.Type == environment.VECTOR_BOOLEAN {
					accesstype = environment.BOOLEAN
				} else if foundVar.Type == environment.VECTOR_STRUCT {
					accesstype = environment.STRUCT_IMP
					//sasasa
				} else if foundVar.Type == environment.MATRIX_INT {
					if _, isBreak := val.([]interface{}); isBreak {
						depth := GetDepth(val.([]interface{}))
						if depth == 1 {
							accesstype = environment.VECTOR_INT
						} else {
							accesstype = environment.MATRIX_INT
						}

					} else {
						accesstype = environment.INTEGER
					}

				} else if foundVar.Type == environment.MATRIX_FLOAT {
					if _, isBreak := val.([]interface{}); isBreak {
						depth := GetDepth(val.([]interface{}))
						if depth == 1 {
							accesstype = environment.VECTOR_FLOAT
						} else {
							accesstype = environment.MATRIX_FLOAT
						}

					} else {
						accesstype = environment.FLOAT
					}

				} else if foundVar.Type == environment.MATRIX_STRING {
					if _, isBreak := val.([]interface{}); isBreak {
						depth := GetDepth(val.([]interface{}))
						if depth == 1 {
							accesstype = environment.VECTOR_STRING
						} else {
							accesstype = environment.MATRIX_STRING
						}

					} else {
						accesstype = environment.STRING
					}

				} else if foundVar.Type == environment.MATRIX_CHAR {
					if _, isBreak := val.([]interface{}); isBreak {
						depth := GetDepth(val.([]interface{}))
						if depth == 1 {
							accesstype = environment.VECTOR_CHAR
						} else {
							accesstype = environment.MATRIX_CHAR
						}

					} else {
						accesstype = environment.CHAR
					}

				} else if foundVar.Type == environment.MATRIX_BOOLEAN {
					if _, isBreak := val.([]interface{}); isBreak {
						depth := GetDepth(val.([]interface{}))
						if depth == 1 {
							accesstype = environment.VECTOR_BOOLEAN
						} else {
							accesstype = environment.MATRIX_BOOLEAN
						}

					} else {
						accesstype = environment.BOOLEAN
					}

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

	} else if env.(environment.Environment).ReferenceExists(p.Id) {
		foundVar := env.(environment.Environment).FindReference(p.Id)
		if foundVar.Type == environment.VECTOR {
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

		if _, isArray := foundVar.Value.([]interface{}); isArray {
			if foundVar.Type == environment.VECTOR_STRING || foundVar.Type == environment.VECTOR_STRUCT || foundVar.Type == environment.VECTOR_CHAR || foundVar.Type == environment.VECTOR_FLOAT || foundVar.Type == environment.VECTOR_BOOLEAN || foundVar.Type == environment.VECTOR_INT || foundVar.Type == environment.MATRIX_INT || foundVar.Type == environment.MATRIX_FLOAT || foundVar.Type == environment.MATRIX_STRING || foundVar.Type == environment.MATRIX_BOOLEAN || foundVar.Type == environment.MATRIX_CHAR {

				val, exists := GetIndexValue(foundVar.Value, indexes)

				if !exists {
					ast.SetError(p.Lin, p.Col, "indice no existente")
					return environment.Symbol{
						Lin:   p.Lin,
						Col:   p.Col,
						Value: nil,
					}
				}

				var accesstype = environment.INTEGER
				if foundVar.Type == environment.VECTOR_INT {
					accesstype = environment.INTEGER
				} else if foundVar.Type == environment.VECTOR_FLOAT {
					accesstype = environment.FLOAT
				} else if foundVar.Type == environment.VECTOR_STRING {
					accesstype = environment.STRING
				} else if foundVar.Type == environment.VECTOR_CHAR {
					accesstype = environment.CHAR
				} else if foundVar.Type == environment.VECTOR_BOOLEAN {
					accesstype = environment.BOOLEAN
				} else if foundVar.Type == environment.VECTOR_STRUCT {
					accesstype = environment.STRUCT_IMP
					//sasasa
				} else if foundVar.Type == environment.MATRIX_INT {
					if _, isBreak := val.([]interface{}); isBreak {
						depth := GetDepth(val.([]interface{}))
						if depth == 1 {
							accesstype = environment.VECTOR_INT
						} else {
							accesstype = environment.MATRIX_INT
						}

					} else {
						accesstype = environment.INTEGER
					}

				} else if foundVar.Type == environment.MATRIX_FLOAT {
					if _, isBreak := val.([]interface{}); isBreak {
						depth := GetDepth(val.([]interface{}))
						if depth == 1 {
							accesstype = environment.VECTOR_FLOAT
						} else {
							accesstype = environment.MATRIX_FLOAT
						}

					} else {
						accesstype = environment.FLOAT
					}

				} else if foundVar.Type == environment.MATRIX_STRING {
					if _, isBreak := val.([]interface{}); isBreak {
						depth := GetDepth(val.([]interface{}))
						if depth == 1 {
							accesstype = environment.VECTOR_STRING
						} else {
							accesstype = environment.MATRIX_STRING
						}

					} else {
						accesstype = environment.STRING
					}

				} else if foundVar.Type == environment.MATRIX_CHAR {
					if _, isBreak := val.([]interface{}); isBreak {
						depth := GetDepth(val.([]interface{}))
						if depth == 1 {
							accesstype = environment.VECTOR_CHAR
						} else {
							accesstype = environment.MATRIX_CHAR
						}

					} else {
						accesstype = environment.CHAR
					}

				} else if foundVar.Type == environment.MATRIX_BOOLEAN {
					if _, isBreak := val.([]interface{}); isBreak {
						depth := GetDepth(val.([]interface{}))
						if depth == 1 {
							accesstype = environment.VECTOR_BOOLEAN
						} else {
							accesstype = environment.MATRIX_BOOLEAN
						}

					} else {
						accesstype = environment.BOOLEAN
					}

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

func GetIndexValue(arr interface{}, indexes []int) (interface{}, bool) {
	if len(indexes) == 0 || reflect.ValueOf(arr).Kind() != reflect.Slice {
		return nil, false
	}

	index := indexes[0]
	if index < 0 || index >= reflect.ValueOf(arr).Len() {
		return nil, false
	}

	if len(indexes) == 1 {
		return reflect.ValueOf(arr).Index(index).Interface(), true
	}

	nextLevel := reflect.ValueOf(arr).Index(index).Interface()
	return GetIndexValue(nextLevel, indexes[1:])
}

func GetIndexes(val []interface{}, ast *environment.AST, env interface{}) []int {
	indexes := make([]int, len(val))
	for i, index := range val {
		var retrievedVal = index.(interfaces.Expression).Execute(ast, env).Value
		indexes[i] = retrievedVal.(int)
	}

	return indexes
}

func AllNonNegativeIntegers(arr []int) bool {
	for _, num := range arr {
		if num < 0 {
			return false
		}
	}
	return true
}

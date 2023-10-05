package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
	"reflect"
	"strings"
)

type MatrixDec struct {
	Lin  int
	Col  int
	Id   string
	Type interface{}
	Def  interfaces.Expression
}

func NewMatrixDec(lin int, col int, id string, tyype interface{}, def interfaces.Expression) MatrixDec {
	NewMatrixDeclaration := MatrixDec{lin, col, id, tyype, def}
	return NewMatrixDeclaration
}

func (p MatrixDec) Execute(ast *environment.AST, env interface{}) interface{} {

	if env.(environment.Environment).VariableExists(p.Id) {
		ast.SetError(p.Lin, p.Col, "Variable ya declarada")
		return nil
	}

	value := p.Def.Execute(ast, env)
	value.Scope =  env.(environment.Environment).Scope
	deepness := GetDepth(value.Value.([]interface{}))
	if _, isString := p.Type.(string); isString {
		if countCharOccurrences(p.Type.(string), ']') == deepness {
			if deepness == 1 {
				var matrixType = getMatrixType(value.Type)
				value.Type = matrixType
				env.(environment.Environment).SaveVariable(p.Id, value)
				ast.SaveSymbol(p.Id,value)
				return nil
			}

			arrayType := getCommonType(value.Value)
			if arrayType != nil {
				if arrayType == reflect.TypeOf(1) && strings.Contains(p.Type.(string), "Int") {
					value.Type = environment.MATRIX_INT
				} else if arrayType == reflect.TypeOf("x") && strings.Contains(p.Type.(string), "String") {
					value.Type = environment.MATRIX_STRING
				} else if arrayType == reflect.TypeOf(int32(0)) && strings.Contains(p.Type.(string), "Character") {
					value.Type = environment.MATRIX_CHAR
				} else if arrayType == reflect.TypeOf(5.121) && strings.Contains(p.Type.(string), "Float") {
					value.Type = environment.MATRIX_FLOAT
				} else if arrayType == reflect.TypeOf(false) && strings.Contains(p.Type.(string), "Bool") {
					value.Type = environment.MATRIX_BOOLEAN
				} else {
					ast.SetError(p.Lin, p.Col, "matriz no coincide con tipo de dato definido")
					return nil
				}
				env.(environment.Environment).SaveVariable(p.Id, value)
				ast.SaveSymbol(p.Id,value)
				return nil
			} else {
				ast.SetError(p.Lin, p.Col, "matriz con varios tipos de dato")
				return nil
			}

		} else {
			ast.SetError(p.Lin, p.Col, "Error: El tamaño con el que se inicializa la matriz no consiste con el tamaño definido")
			return nil
		}
	} else {
		arrayType := getCommonType(value.Value)
		if arrayType != nil {
			if arrayType == reflect.TypeOf(1) {
				value.Type = environment.MATRIX_INT
			} else if arrayType == reflect.TypeOf("x") {
				value.Type = environment.MATRIX_STRING
			} else if arrayType == reflect.TypeOf(int32(0)) {
				value.Type = environment.MATRIX_CHAR
			} else if arrayType == reflect.TypeOf(5.121) {
				value.Type = environment.MATRIX_FLOAT
			} else if arrayType == reflect.TypeOf(false) {
				value.Type = environment.MATRIX_BOOLEAN
			} else {
				ast.SetError(p.Lin, p.Col, "matriz no coincide con tipo de dato definido")
				return nil
			}
			env.(environment.Environment).SaveVariable(p.Id, value)
			ast.SaveSymbol(p.Id,value)
			return nil
		} else {
			ast.SetError(p.Lin, p.Col, "matriz con varios tipos de dato")
			return nil
		}

	}

}

func (p MatrixDec) GetMatrixDec(ast *environment.AST, env interface{}) interface{} {

	value := p.Def.Execute(ast, env)
	value.Scope = env.(environment.Environment).Scope
	deepness := GetDepth(value.Value.([]interface{}))
	if _, isString := p.Type.(string); isString {
		if countCharOccurrences(p.Type.(string), ']') == deepness {
			if deepness == 1 {
				var matrixType = getMatrixType(value.Type)
				value.Type = matrixType
				return value

			}

			arrayType := getCommonType(value.Value)
			if arrayType != nil {
				if arrayType == reflect.TypeOf(1) && strings.Contains(p.Type.(string), "Int") {
					value.Type = environment.MATRIX_INT
				} else if arrayType == reflect.TypeOf("x") && strings.Contains(p.Type.(string), "String") {
					value.Type = environment.MATRIX_STRING
				} else if arrayType == reflect.TypeOf(int32(0)) && strings.Contains(p.Type.(string), "Character") {
					value.Type = environment.MATRIX_CHAR
				} else if arrayType == reflect.TypeOf(5.121) && strings.Contains(p.Type.(string), "Float") {
					value.Type = environment.MATRIX_FLOAT
				} else if arrayType == reflect.TypeOf(false) && strings.Contains(p.Type.(string), "Bool") {
					value.Type = environment.MATRIX_BOOLEAN
				} else {
					ast.SetError(p.Lin, p.Col, "matriz no coincide con tipo de dato definido")
					return nil
				}
				return value

			} else {
				ast.SetError(p.Lin, p.Col, "matriz con varios tipos de dato")
				return nil
			}

		} else {
			ast.SetError(p.Lin, p.Col, "Error: El tamaño con el que se inicializa la matriz no consiste con el tamaño definido")
			return nil
		}
	} else {
		arrayType := getCommonType(value.Value)
		if arrayType != nil {
			if arrayType == reflect.TypeOf(1) {
				value.Type = environment.MATRIX_INT
			} else if arrayType == reflect.TypeOf("x") {
				value.Type = environment.MATRIX_STRING
			} else if arrayType == reflect.TypeOf(int32(0)) {
				value.Type = environment.MATRIX_CHAR
			} else if arrayType == reflect.TypeOf(5.121) {
				value.Type = environment.MATRIX_FLOAT
			} else if arrayType == reflect.TypeOf(false) {
				value.Type = environment.MATRIX_BOOLEAN
			} else {
				ast.SetError(p.Lin, p.Col, "matriz no coincide con tipo de dato definido")
				return nil
			}
			return value

		} else {
			ast.SetError(p.Lin, p.Col, "matriz con varios tipos de dato")
			return nil
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
func countCharOccurrences(input string, char rune) int {
	count := 0
	for _, c := range input {
		if c == char {
			count++
		}
	}
	return count
}

func getMatrixType(typee environment.TipoExpresion) environment.TipoExpresion {
	if typee == environment.VECTOR_INT {
		return environment.MATRIX_INT
	} else if typee == environment.VECTOR_FLOAT {
		return environment.MATRIX_FLOAT
	} else if typee == environment.VECTOR_BOOLEAN {
		return environment.MATRIX_BOOLEAN
	} else if typee == environment.VECTOR_CHAR {
		return environment.MATRIX_CHAR
	} else if typee == environment.VECTOR_STRING {
		return environment.MATRIX_STRING
	}
	return environment.NULL
}

func getCommonType(arr interface{}) reflect.Type {
	switch arr.(type) {
	case []interface{}:
		var commonType reflect.Type
		hasMultipleTypes := false

		for _, item := range arr.([]interface{}) {
			itemType := getCommonType(item)
			if itemType == nil {
				hasMultipleTypes = true
			} else if commonType == nil {
				commonType = itemType
			} else if commonType != itemType {
				hasMultipleTypes = true
			}
		}

		if hasMultipleTypes {
			return nil
		}

		return commonType

	case string:
		if len(arr.(string)) == 1 {
			return reflect.TypeOf('a')
		} else {
			return reflect.TypeOf(arr)
		}

	default:
		return reflect.TypeOf(arr)
	}
}

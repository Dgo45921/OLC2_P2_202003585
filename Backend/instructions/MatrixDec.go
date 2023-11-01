package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
	"fmt"
	"reflect"
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

func (p MatrixDec) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var prueba = p.Def.Execute(ast, env, gen)
	fmt.Println(prueba)
	return environment.Value{}
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

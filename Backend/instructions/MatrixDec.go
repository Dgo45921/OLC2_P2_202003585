package instructions

import (
	"PY1/environment"
	"PY1/expressions"
	"PY1/generator"
	"PY1/interfaces"
	"fmt"
	"reflect"
	"strconv"
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
	if env.(environment.Environment).VariableExists(p.Id) {
		ast.SetError(p.Lin, p.Col, "Error, variable ya declarada!")
		return environment.Value{}
	}

	if _, isBreak := p.Def.(expressions.ManualMatrixDef); isBreak {
		newArray := subtractOneFromElements(p.Def.(expressions.ManualMatrixDef).Value, ast, env, gen)
		response := newArray.([]interface{})[0]
		prueba := flattenArray(response)
		fmt.Println(prueba)
		var result, val environment.Value
		size := len(prueba)

		//generando array
		gen.AddComment("----Generando matriz----")
		newtmp1 := gen.NewTemp()
		newtmp2 := gen.NewTemp()
		gen.AddAssign(newtmp1, "H")
		gen.AddExpression(newtmp2, newtmp1, "1", "+")
		gen.AddSetHeap("(int)H", strconv.Itoa(size))
		gen.AddExpression("H", "H", strconv.Itoa(size+1), "+")
		//recorriendo lista de expressiones
		for _, s := range prueba {
			val = s.(interfaces.Expression).Execute(ast, env, gen)
			gen.AddSetHeap("(int)"+newtmp2, val.Value)
			gen.AddExpression(newtmp2, newtmp2, "1", "+")
		}

		dimentions := getArrayDimensions(response)
		result = environment.Value{
			Value:        newtmp1,
			IsTemp:       true,
			Type:         getTypeMatrix(val.Type),
			TrueLabel:    nil,
			FalseLabel:   nil,
			OutLabel:     nil,
			IntValue:     0,
			FloatValue:   0,
			BreakFlag:    false,
			ContinueFlag: false,
		}

		newVar := env.(environment.Environment).SaveMatrix(p.Id, result.Type, size, p.Def.(expressions.ManualMatrixDef).Value.([]interface{}), dimentions)

		gen.AddSetStack(strconv.Itoa(newVar.Position), result.Value)
		gen.AddBr()

		return result

	} else if _, isBreak := p.Def.(expressions.RepeatingVector); isBreak {

		var arr = p.Def.(expressions.RepeatingVector).Execute2(ast, env, gen).Value
		fmt.Println(arr)

		prueba := flattenArray(arr)
		fmt.Println(prueba)
		var result, val environment.Value
		size := len(prueba)

		//generando array
		gen.AddComment("----Generando matriz----")
		newtmp1 := gen.NewTemp()
		newtmp2 := gen.NewTemp()
		gen.AddAssign(newtmp1, "H")
		gen.AddExpression(newtmp2, newtmp1, "1", "+")
		gen.AddSetHeap("(int)H", strconv.Itoa(size))
		gen.AddExpression("H", "H", strconv.Itoa(size+1), "+")
		//recorriendo lista de expressiones
		for _, s := range prueba {
			val = s.(interfaces.Expression).Execute(ast, env, gen)
			gen.AddSetHeap("(int)"+newtmp2, val.Value)
			gen.AddExpression(newtmp2, newtmp2, "1", "+")
		}

		dimentions := getArrayDimensions(arr)
		if len(dimentions) == 1 {
			dimentions = make([]int, 1)
			dimentions[0] = 1
		}
		result = environment.Value{
			Value:        newtmp1,
			IsTemp:       true,
			Type:         getTypeMatrix(val.Type),
			TrueLabel:    nil,
			FalseLabel:   nil,
			OutLabel:     nil,
			IntValue:     0,
			FloatValue:   0,
			BreakFlag:    false,
			ContinueFlag: false,
		}

		newVar := env.(environment.Environment).SaveMatrix(p.Id, result.Type, size, arr.([]interface{}), dimentions)

		gen.AddSetStack(strconv.Itoa(newVar.Position), result.Value)
		gen.AddBr()

		return result
	}

	var prueba = p.Def.Execute(ast, env, gen)
	fmt.Println(prueba)
	return environment.Value{}
}

func getArrayDimensions(arr interface{}) []int {
	dimensions := []int{}

	var traverse func(interface{})
	traverse = func(a interface{}) {
		if reflect.TypeOf(a).Kind() == reflect.Slice {
			dimensions = append(dimensions, reflect.ValueOf(a).Len())
			if reflect.ValueOf(a).Len() > 0 {
				traverse(reflect.ValueOf(a).Index(0).Interface())
			}
		}
	}

	traverse(arr)
	return dimensions
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

func flattenArray(arr interface{}) []interface{} {
	var result []interface{}

	switch v := arr.(type) {
	case []interface{}:
		// If it's a slice of interfaces, flatten each sub-slice
		for _, subArr := range v {
			result = append(result, flattenArray(subArr)...)
		}
	default:
		// If it's not a slice of interfaces, assume it's an element, and add it to the result
		result = append(result, v)
	}

	return result
}

func subtractOneFromElements(arr interface{}, ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	switch arr.(type) {
	case []interface{}:
		result := make([]interface{}, len(arr.([]interface{})))
		for i, item := range arr.([]interface{}) {
			result[i] = subtractOneFromElements(item, ast, env, gen)
		}
		return result
	case interfaces.Expression:
		return arr.(interfaces.Expression)
	case []interfaces.Expression:
		result := make([]interface{}, len(arr.([]interface{})))
		for i, item := range arr.([]interface{}) {
			result[i] = item.(interfaces.Expression)
		}
		return result
	default:
		return arr
	}
}

func getTypeMatrix(val environment.TipoExpresion) environment.TipoExpresion {
	if val == environment.INTEGER {
		return environment.MATRIX_INT
	} else if val == environment.FLOAT {
		return environment.MATRIX_FLOAT
	} else if val == environment.BOOLEAN {
		return environment.MATRIX_BOOLEAN
	} else if val == environment.CHAR {
		return environment.MATRIX_CHAR
	} else if val == environment.STRING {
		return environment.MATRIX_STRING
	}

	return environment.VECTOR
}

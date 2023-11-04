package expressions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
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

func (p VectorAccess) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {

	var tempArray, tempIndex, result environment.Value

	prueba := VariableAccess{ID: p.Id}

	tempArray = prueba.Execute(ast, env, gen)

	if len(p.Index) == 1 {
		tempIndex = p.Index[0].(interfaces.Expression).Execute(ast, env, gen)
		//llamada
		newTmp := gen.NewTemp()
		lvl1 := gen.NewLabel()
		lvl2 := gen.NewLabel()
		lvl3 := gen.NewLabel()
		gen.AddIf(tempIndex.Value, "0", "<", lvl1)
		tmp := gen.NewTemp()
		gen.AddGetHeap(tmp, "(int)"+tempArray.Value)
		gen.AddComment("checking vector index bounds")
		gen.AddIf(tempIndex.Value, tmp, ">=", lvl1)
		gen.AddGoto(lvl2)
		gen.AddLabel(lvl1)

		gen.AddPrintf("c", "73")
		gen.AddPrintf("c", "78")
		gen.AddPrintf("c", "68")
		gen.AddPrintf("c", "69")
		gen.AddPrintf("c", "88")
		gen.AddPrintf("c", "79")
		gen.AddPrintf("c", "85")
		gen.AddPrintf("c", "84")
		gen.AddPrintf("c", "79")
		gen.AddPrintf("c", "70")
		gen.AddPrintf("c", "66")
		gen.AddPrintf("c", "111")
		gen.AddPrintf("c", "117")
		gen.AddPrintf("c", "110")
		gen.AddPrintf("c", "100")
		gen.AddPrintf("c", "115")
		gen.AddPrintf("c", "69")
		gen.AddPrintf("c", "114")
		gen.AddPrintf("c", "114")
		gen.AddPrintf("c", "111")
		gen.AddPrintf("c", "114")
		gen.AddGoto(lvl3)
		gen.AddLabel(lvl2)
		gen.AddExpression(newTmp, tempArray.Value, tempIndex.Value, "+")
		gen.AddExpression(newTmp, newTmp, "1", "+")
		newTmp2 := gen.NewTemp()
		gen.AddGetHeap(newTmp2, "(int)"+newTmp)
		gen.AddLabel(lvl3)

		result = environment.Value{
			Value:        newTmp2,
			IsTemp:       true,
			Type:         getInsideType(tempArray.Type),
			TrueLabel:    nil,
			FalseLabel:   nil,
			OutLabel:     nil,
			IntValue:     0,
			FloatValue:   0,
			BreakFlag:    false,
			ContinueFlag: false,
		}
		return result

	} else {

		var intSlice []int
		size := len(p.Index)
		intSlice = make([]int, size)
		for index, element := range p.Index {
			intSlice[index] = element.(interfaces.Expression).Execute(ast, env, gen).IntValue
		}

		var realIndex = calculateIndex(intSlice, tempArray.Dimentions)

		var primitive = Primitive{
			Lin:   0,
			Col:   0,
			Valor: realIndex,
			Type:  environment.INTEGER,
		}

		tempIndex = primitive.Execute(ast, env, gen)

		//llamada
		newTmp := gen.NewTemp()
		lvl1 := gen.NewLabel()
		lvl2 := gen.NewLabel()
		lvl3 := gen.NewLabel()
		gen.AddIf(tempIndex.Value, "0", "<", lvl1)
		tmp := gen.NewTemp()
		gen.AddGetHeap(tmp, "(int)"+tempArray.Value)
		gen.AddComment("checking vector index bounds")
		gen.AddIf(tempIndex.Value, tmp, ">=", lvl1)
		gen.AddGoto(lvl2)
		gen.AddLabel(lvl1)
		gen.AddPrintf("c", "66")
		gen.AddPrintf("c", "111")
		gen.AddPrintf("c", "117")
		gen.AddPrintf("c", "110")
		gen.AddPrintf("c", "100")
		gen.AddPrintf("c", "115")
		gen.AddPrintf("c", "69")
		gen.AddPrintf("c", "114")
		gen.AddPrintf("c", "114")
		gen.AddPrintf("c", "111")
		gen.AddPrintf("c", "114")
		gen.AddGoto(lvl3)
		gen.AddLabel(lvl2)
		gen.AddExpression(newTmp, tempArray.Value, tempIndex.Value, "+")
		gen.AddExpression(newTmp, newTmp, "1", "+")
		newTmp2 := gen.NewTemp()
		gen.AddGetHeap(newTmp2, "(int)"+newTmp)
		gen.AddLabel(lvl3)

		result = environment.Value{
			Value:        newTmp2,
			IsTemp:       true,
			Type:         getInsideTypeMatrix(tempArray.Type),
			TrueLabel:    nil,
			FalseLabel:   nil,
			OutLabel:     nil,
			IntValue:     0,
			FloatValue:   0,
			BreakFlag:    false,
			ContinueFlag: false,
		}
		return result

	}

}

func getInsideType(val environment.TipoExpresion) environment.TipoExpresion {
	if val == environment.VECTOR_INT {
		return environment.INTEGER
	} else if val == environment.VECTOR_FLOAT {
		return environment.FLOAT
	} else if val == environment.VECTOR_BOOLEAN {
		return environment.BOOLEAN
	} else if val == environment.VECTOR_CHAR {
		return environment.CHAR
	} else if val == environment.VECTOR_STRING {
		return environment.STRING
	}

	return environment.VECTOR
}

func getInsideTypeMatrix(val environment.TipoExpresion) environment.TipoExpresion {
	if val == environment.MATRIX_INT {
		return environment.INTEGER
	} else if val == environment.MATRIX_FLOAT {
		return environment.FLOAT
	} else if val == environment.MATRIX_BOOLEAN {
		return environment.BOOLEAN
	} else if val == environment.MATRIX_CHAR {
		return environment.CHAR
	} else if val == environment.MATRIX_STRING {
		return environment.STRING
	}

	return environment.VECTOR
}

func calculateIndex(indices []int, dimensions []int) int {
	if len(indices) != len(dimensions) {
		panic("Mismatched dimensions and indices")
	}

	index := 0
	multiplier := 1

	for i := len(indices) - 1; i >= 0; i-- {
		index += indices[i] * multiplier
		multiplier *= dimensions[i]
	}

	return index
}

func convertToIntArray(arr []interface{}) []int {
	intArray := make([]int, len(arr))
	for i, v := range arr {
		if val, ok := v.(int); ok {
			intArray[i] = val
		} else {
			// Handle the case where the element is not an int.
			// You can choose to skip it, set a default value, or handle it as needed.
			intArray[i] = 0 // For example, set to 0.
		}
	}
	return intArray
}

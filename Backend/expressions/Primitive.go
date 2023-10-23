package expressions

import (
	"PY1/environment"
	"PY1/generator"
	"fmt"
	"strconv"
)

type Primitive struct {
	Lin   int
	Col   int
	Valor interface{}
	Type  environment.TipoExpresion
}

func NewPrimitive(lin int, col int, valor interface{}, tipo environment.TipoExpresion) Primitive {
	exp := Primitive{lin, col, valor, tipo}
	return exp
}

func (p Primitive) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	if p.Type == environment.INTEGER {
		result = environment.NewValue(fmt.Sprintf("%v", p.Valor), false, p.Type)
		result.IntValue = p.Valor.(int)
	} else if p.Type == environment.FLOAT {
		result = environment.NewValue(fmt.Sprintf("%v", p.Valor), false, p.Type)
		result.FloatValue = p.Valor.(float64)
	} else if p.Type == environment.NULL {
		result = environment.NewValue("NULLPTR", false, p.Type)
	} else if p.Type == environment.STRING || p.Type == environment.CHAR {
		//nuevo temporal
		newTemp := gen.NewTemp()
		//iguala a heap pointer
		gen.AddAssign(newTemp, "H")
		//recorremos string en ascii
		myString := p.Valor.(string)
		byteArray := []byte(myString)
		for _, asc := range byteArray {
			//se agrega ascii al heap
			gen.AddSetHeap("(int)H", strconv.Itoa(int(asc)))
			//suma heap pointer
			gen.AddExpression("H", "H", "1", "+")
		}
		//caracteres de escape
		gen.AddSetHeap("(int)H", "-1")
		gen.AddExpression("H", "H", "1", "+")
		gen.AddBr()
		result = environment.NewValue(newTemp, true, p.Type)
		//result = environment.Value{Value: newTemp, IsTemp: true, Type: p.Type}
	} else if p.Type == environment.BOOLEAN {
		tempo := gen.NewTemp()
		gen.AddComment("Primitivo bool")

		trueLabel := gen.NewLabel()
		falseLabel := gen.NewLabel()
		if p.Valor.(bool) {
			gen.AddExpression(tempo, "1", "0", "+")
			gen.AddGoto(trueLabel)
		} else {
			gen.AddExpression(tempo, "0", "0", "+")
			gen.AddGoto(falseLabel)
		}

		result = environment.NewValue(tempo, true, environment.BOOLEAN)
		result.TrueLabel = append(result.TrueLabel, trueLabel)
		result.FalseLabel = append(result.FalseLabel, falseLabel)

	}
	return result
}

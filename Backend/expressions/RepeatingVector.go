package expressions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type RepeatingVector struct {
	Lin        int
	Col        int
	MatrixType string
	Repeating  interfaces.Expression
	Count      interfaces.Expression
}

func NewRepeatingVector(lin int, col int, mtype string, repeating interfaces.Expression, count interfaces.Expression) RepeatingVector {
	exp := RepeatingVector{lin, col, mtype, repeating, count}
	return exp
}

func (p RepeatingVector) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
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

func DeepCopyArray(source interface{}) interface{} {
	switch source := source.(type) {
	case []interface{}:
		copyArray := make([]interface{}, len(source))
		for i, val := range source {
			copyArray[i] = DeepCopyArray(val)
		}
		return copyArray
	default:
		return source
	}
}

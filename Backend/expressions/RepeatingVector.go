package expressions

import (
	"PY1/environment"
	"PY1/interfaces"
	"strconv"
	"strings"
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

func (p RepeatingVector) Execute(ast *environment.AST, env interface{}) environment.Symbol {

	// keep executing recursively
	if _, isRepeating := p.Repeating.(RepeatingVector); isRepeating {

		var nextdimension = countCharOccurrences(p.Repeating.(RepeatingVector).MatrixType, '[')
		var currentdimension = countCharOccurrences(p.MatrixType, '[')

		if nextdimension == currentdimension-1 {
			response := p.Repeating.(RepeatingVector).Execute(ast, env)

			// check count, it must be an integer
			var count = p.Count.Execute(ast, env)
			if count.Type != environment.INTEGER {
				ast.SetError(p.Lin, p.Col, "el atributo count debe de ser un entero")
				return environment.Symbol{
					Lin:   p.Lin,
					Col:   p.Col,
					Value: nil,
					Const: true,
				}

			}

			N := count.Value.(int)

			arr := make([]interface{}, N)
			for i := 0; i < N; i++ {
				arr[i] = DeepCopyArray(response.Value)
			}

			if (strings.Contains(p.MatrixType, "Int") && response.Type == environment.VECTOR_INT) || (strings.Contains(p.MatrixType, "String") && response.Type == environment.VECTOR_STRING) || (strings.Contains(p.MatrixType, "Character") && response.Type == environment.VECTOR_CHAR) || (strings.Contains(p.MatrixType, "Float") && response.Type == environment.VECTOR_FLOAT) || (strings.Contains(p.MatrixType, "Bool") && response.Type == environment.VECTOR_BOOLEAN) {
				return environment.Symbol{
					Lin:   p.Lin,
					Col:   p.Col,
					Value: arr,
					Type:  response.Type,
					Const: true,
				}
				// i should be adding the response vector

			} else {
				ast.SetError(p.Lin, p.Col, "no coincide el tipo definido con el 'repeating'")
				return environment.Symbol{
					Lin:   p.Lin,
					Col:   p.Col,
					Value: nil,
				}
			}

		} else {
			ast.SetError(p.Lin, p.Col, "Un array de dimensión: " + strconv.Itoa(currentdimension) + " no puede almacenar uno de dimensión: " + strconv.Itoa(nextdimension))


			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Value: nil,
			}
		}

	} else {
		// check count, it must be an integer
		var count = p.Count.Execute(ast, env)
		if count.Type != environment.INTEGER {

			ast.SetError(p.Lin, p.Col, "el atributo count debe de ser un entero")
			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Value: nil,
			}

		}

		N := count.Value.(int)
		value := p.Repeating.Execute(ast, env)
		// creating the one dimension array
		arr := make([]interface{}, N)
		for i := 0; i < N; i++ {
			arr[i] = value.Value
		}

		// setting up its type

		typesVector := environment.VECTOR_INT
		switch typo := value.Type; typo {
		case environment.INTEGER:
			typesVector = environment.VECTOR_INT
		case environment.FLOAT:
			typesVector = environment.VECTOR_FLOAT
		case environment.CHAR:
			typesVector = environment.VECTOR_CHAR
		case environment.BOOLEAN:
			typesVector = environment.VECTOR_BOOLEAN
		case environment.STRING:
			typesVector = environment.VECTOR_STRING

		}

		if (strings.Contains(p.MatrixType, "Int") && typesVector == environment.VECTOR_INT) || (strings.Contains(p.MatrixType, "String") && typesVector == environment.VECTOR_STRING) || (strings.Contains(p.MatrixType, "Character") && typesVector == environment.VECTOR_CHAR) || (strings.Contains(p.MatrixType, "Float") && typesVector == environment.VECTOR_FLOAT) || (strings.Contains(p.MatrixType, "Bool") && typesVector == environment.VECTOR_BOOLEAN) {
			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Value: arr,
				Type:  typesVector,
			}
		} else {
			ast.SetError(p.Lin, p.Col, "no coincide el tipo definido con el repeating")

			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Value: nil,
			}
		}

	}

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

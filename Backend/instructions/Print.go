package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
	"fmt"
	"strings"
)

type Print struct {
	Lin   int
	Col   int
	Value []interface{}
}

func NewPrint(lin int, col int, val []interface{}) Print {
	return Print{lin, col, val}
}

func (p Print) Execute(ast *environment.AST, env interface{}) interface{} {
	var printedValues []string

	for _, val := range p.Value {
		if expr, ok := val.(interfaces.Expression); ok {

			pivote := expr.Execute(ast, env)
			valueToPrint := pivote.Value
			if pivote.Type == environment.CHAR {
				if _, ischar := valueToPrint.(uint8); ischar {
					valueToPrint = string(valueToPrint.(uint8))
				}
			}

			printedValues = append(printedValues, fmt.Sprintf("%v", valueToPrint))
		} else if _, ok := val.(environment.Symbol); ok {

			valueToPrint := val.(environment.Symbol).Value

			printedValues = append(printedValues, fmt.Sprintf("%v", valueToPrint))
		}
	}

	consoleOut := strings.Join(printedValues, " ")
	ast.SetPrint(consoleOut + "\n")

	return nil
}

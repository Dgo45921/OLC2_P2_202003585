package expressions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type Range struct {
	Lin        int
	Col        int
	FirstIndex interfaces.Expression
	LastIndex  interfaces.Expression
}

func NewRange(lin int, col int, findex interfaces.Expression, lindex interfaces.Expression) Range {
	exp := Range{lin, col, findex, lindex}
	return exp
}

func (p Range) Execute(ast *environment.AST, env interface{}) environment.Symbol {

	var findex = p.FirstIndex.Execute(ast, env)
	var lindex = p.LastIndex.Execute(ast, env)

	if findex.Type == environment.INTEGER && lindex.Type == environment.INTEGER {

		start := findex.Value.(int)
		end := lindex.Value.(int)

		if start <= end {
			size := end - start + 1

			numbers := make([]interface{}, size)
			for i := 0; i < size; i++ {
				numbers[i] = start + i
			}

			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Value: numbers,
				Type:  environment.VECTOR_INT,
			}

		} else {
			ast.SetError(p.Lin, p.Col, "indice inicial es mas grande que el indice final")
			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Value: nil,
			}
		}

	} else {
		ast.SetError(p.Lin, p.Col, "los indices deben de ser enteros")
		return environment.Symbol{
			Lin:   p.Lin,
			Col:   p.Col,
			Value: nil,
		}
	}

}

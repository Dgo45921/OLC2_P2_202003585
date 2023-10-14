package expressions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
	"strconv"
	"strings"
)

type Cast struct {
	Lin         int
	Col         int
	CastingType string
	Val         interfaces.Expression
}

func NewCast(lin int, col int, valor string, tipo interfaces.Expression) Cast {
	cast := Cast{lin, col, valor, tipo}
	return cast
}

func (p Cast) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result, tmp environment.Value
	if p.CastingType == "String" {
		tmp = p.Val.(interfaces.Expression).Execute(ast, env, gen)
		if tmp.Type == environment.STRING || tmp.Type == environment.INTEGER || tmp.Type == environment.FLOAT {
			tmp.Type = environment.STRING
			result = tmp
		} else {
			ast.SetError(p.Lin, p.Col, "No se pudo castear a string")
		}

		return result
	}

	return result
}

func getIntegerValue(input string) (int, error) {
	// Split the input string by the decimal point
	parts := strings.Split(input, ".")

	// Parse the first part as an integer
	intPart, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}

	return intPart, nil
}

func parseFloatFromString(input string) (float64, error) {
	// Attempt to parse the input string into a float64
	result, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0.0, err
	}

	return result, nil
}

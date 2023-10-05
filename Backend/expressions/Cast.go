package expressions

import (
	"PY1/environment"
	"PY1/interfaces"
	"math"
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

func (p Cast) Execute(ast *environment.AST, env interface{}) environment.Symbol {
	val := p.Val.Execute(ast, env)

	if p.CastingType == "Int" {
		if val.Type == environment.STRING || val.Type == environment.CHAR {

			input := val.Value.(string)
			result, err := getIntegerValue(input)
			if err != nil {
				return environment.Symbol{
					Lin:   p.Lin,
					Col:   p.Col,
					Type:  environment.INTEGER,
					Value: nil,
				}
			} else {
				return environment.Symbol{
					Lin:   p.Lin,
					Col:   p.Col,
					Type:  environment.INTEGER,
					Value: result,
				}
			}

		} else if val.Type == environment.FLOAT {
			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Type:  environment.FLOAT,
				Value: int(math.Floor(val.Value.(float64))),
			}

		} else {
			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Type:  environment.NULL,
				Value: nil,
			}

		}

	} else if p.CastingType == "Float" {
		if val.Type == environment.STRING {
			input := val.Value.(string)
			parsedFloat, err := parseFloatFromString(input)
			if err != nil {
				return environment.Symbol{
					Lin:   p.Lin,
					Col:   p.Col,
					Type:  environment.NULL,
					Value: nil,
				}
			} else {
				return environment.Symbol{
					Lin:   p.Lin,
					Col:   p.Col,
					Type:  environment.FLOAT,
					Value: parsedFloat,
				}
			}

		} else if val.Type == environment.INTEGER {
			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Type:  environment.FLOAT,
				Value: float64(val.Value.(int)),
			}

		} else {
			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Type:  environment.NULL,
				Value: nil,
			}
		}

	} else if p.CastingType == "String" {
		if val.Type == environment.BOOLEAN {
			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Type:  environment.STRING,
				Value: strconv.FormatBool(val.Value.(bool)),
			}

		} else if val.Type == environment.FLOAT {
			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Type:  environment.STRING,
				Value: strconv.FormatFloat(val.Value.(float64), 'f', -1, 64),
			}

		} else if val.Type == environment.INTEGER {
			return environment.Symbol{
				Lin:   p.Lin,
				Col:   p.Col,
				Type:  environment.STRING,
				Value: strconv.Itoa(val.Value.(int)),
			}
		} else {

		}

	}

	return environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Type:  environment.NULL,
		Value: nil,
	}
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

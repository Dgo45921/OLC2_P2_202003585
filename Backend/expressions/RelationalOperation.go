package expressions

import (
	"PY1/environment"
	"PY1/interfaces"
	"fmt"
	"strconv"
)

type RelationalOperation struct {
	Lin      int
	Col      int
	LeftExp  interfaces.Expression
	Operator string
	RightExp interfaces.Expression
}

func NewRelationalOperation(lin int, col int, Op1 interfaces.Expression, Operador string, Op2 interfaces.Expression) RelationalOperation {
	exp := RelationalOperation{Lin: lin, Col: col, LeftExp: Op1, Operator: Operador, RightExp: Op2}
	return exp
}

func (o RelationalOperation) Execute(ast *environment.AST, env interface{}) environment.Symbol {

	var op1, op2 environment.Symbol
	op1 = o.LeftExp.Execute(ast, env)
	op2 = o.RightExp.Execute(ast, env)

	switch o.Operator {

	case "<":
		{
			var type0 = op1.Type
			var type1 = op2.Type

			if type0 == environment.INTEGER && type1 == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(int) < op2.Value.(int)}
			} else if type0 == environment.FLOAT && type1 == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: val1 < val2}
			} else if type0 == environment.STRING && type1 == environment.STRING || type0 == environment.CHAR && type1 == environment.CHAR {
				r1 := fmt.Sprintf("%v", op1.Value)
				r2 := fmt.Sprintf("%v", op2.Value)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: r1 < r2}
			} else {
				ast.SetError(o.Lin, o.Col, "No es posible comparar <")
			}
		}
	case ">":
		{
			var type0 = op1.Type
			var type1 = op2.Type

			if type0 == environment.INTEGER && type1 == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(int) > op2.Value.(int)}
			} else if type0 == environment.FLOAT && type1 == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: val1 > val2}
			} else if type0 == environment.STRING && type1 == environment.STRING || type0 == environment.CHAR && type1 == environment.CHAR {
				r1 := fmt.Sprintf("%v", op1.Value)
				r2 := fmt.Sprintf("%v", op2.Value)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: r1 > r2}
			} else {
				ast.SetError(o.Lin, o.Col, "No es posible comparar >")
			}
		}
	case "<=":
		{
			var type0 = op1.Type
			var type1 = op2.Type

			if type0 == environment.INTEGER && type1 == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(int) <= op2.Value.(int)}
			} else if type0 == environment.FLOAT && type1 == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: val1 <= val2}
			} else if type0 == environment.STRING && type1 == environment.STRING || type0 == environment.CHAR && type1 == environment.CHAR {
				r1 := fmt.Sprintf("%v", op1.Value)
				r2 := fmt.Sprintf("%v", op2.Value)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: r1 <= r2}
			} else {
				ast.SetError(o.Lin, o.Col, "No es posible comparar <=")
			}
		}
	case ">=":
		{
			var type0 = op1.Type
			var type1 = op2.Type

			if type0 == environment.INTEGER && type1 == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(int) >= op2.Value.(int)}
			} else if type0 == environment.FLOAT && type1 == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: val1 >= val2}
			} else if type0 == environment.STRING && type1 == environment.STRING || type0 == environment.CHAR && type1 == environment.CHAR {
				r1 := fmt.Sprintf("%v", op1.Value)
				r2 := fmt.Sprintf("%v", op2.Value)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: r1 >= r2}
			} else {
				ast.SetError(o.Lin, o.Col, "No es posible comparar >=")
			}
		}
	case "==":
		{
			var type0 = op1.Type
			var type1 = op2.Type

			if type0 == environment.INTEGER && type1 == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(int) == op2.Value.(int)}
			} else if type0 == environment.FLOAT && type1 == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: val1 == val2}
			} else if type0 == environment.STRING && type1 == environment.STRING || type0 == environment.CHAR && type1 == environment.CHAR {
				r1 := fmt.Sprintf("%v", op1.Value)
				r2 := fmt.Sprintf("%v", op2.Value)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: r1 == r2}
			} else {
				ast.SetError(o.Lin, o.Col, "No es posible comparar ==")
			}
		}

	case "!=":
		{
			var type0 = op1.Type
			var type1 = op2.Type

			if type0 == environment.INTEGER && type1 == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: op1.Value.(int) != op2.Value.(int)}
			} else if type0 == environment.FLOAT && type1 == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Value), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Value), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: val1 != val2}
			} else if type0 == environment.STRING && type1 == environment.STRING || type0 == environment.CHAR && type1 == environment.CHAR {
				r1 := fmt.Sprintf("%v", op1.Value)
				r2 := fmt.Sprintf("%v", op2.Value)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.BOOLEAN, Value: r1 != r2}
			} else {
				ast.SetError(o.Lin, o.Col, "No es posible comparar !=")
			}
		}

	}

	var result interface{}
	return environment.Symbol{Lin: o.Lin, Col: o.Col, Type: environment.NULL, Value: result}
}

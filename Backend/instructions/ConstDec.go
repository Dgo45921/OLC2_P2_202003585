package instructions

import (
	"PY1/environment"
	"PY1/interfaces"
)

type ConstDec struct {
	Lin        int
	Col        int
	Id         string
	Type       interface{}
	Expression interface{}
}

func NewConstDec(lin int, col int, id string, tyype interface{}, val interface{}) ConstDec {
	NewConstDeclaration := ConstDec{lin, col, id, tyype, val}
	return NewConstDeclaration
}

func (p ConstDec) Execute(ast *environment.AST, env interface{}) interface{} {

	if env.(environment.Environment).VariableExists(p.Id) {
		ast.SetError(p.Lin, p.Col, "constante ya declarada")
		return nil
	}
	if p.Type == nil {

		if _, ok := p.Expression.(interfaces.Expression); ok {
			expression := p.Expression.(interfaces.Expression)
			value := expression.Execute(ast, env)
			value.Scope = env.(environment.Environment).Scope
			value.Const = true
			env.(environment.Environment).SaveVariable(p.Id, value)
			ast.SaveSymbol(p.Id,value)
			return nil

		}
	}

	if _, ok := p.Expression.(interfaces.Expression); ok {

		expression := p.Expression.(interfaces.Expression)
		value := expression.Execute(ast, env)
		value.Scope = env.(environment.Environment).Scope
		value.Const = true

		if p.Type == "String" && value.Type == environment.STRING {
			env.(environment.Environment).SaveVariable(p.Id, value)
			ast.SaveSymbol(p.Id,value)
			return nil
		} else if p.Type == "Int" && value.Type == environment.INTEGER {
			env.(environment.Environment).SaveVariable(p.Id, value)
			ast.SaveSymbol(p.Id,value)
			return nil
		} else if p.Type == "Character" && value.Type == environment.CHAR {
			env.(environment.Environment).SaveVariable(p.Id, value)
			ast.SaveSymbol(p.Id,value)
			return nil
		} else if p.Type == "Float" {
			if value.Type == environment.FLOAT {
				env.(environment.Environment).SaveVariable(p.Id, value)
				ast.SaveSymbol(p.Id,value)
				return nil

			} else if value.Type == environment.INTEGER {
				if _, ok := value.Value.(int); ok {
					value.Value = float64(value.Value.(int))
					value.Type = environment.FLOAT
					env.(environment.Environment).SaveVariable(p.Id, value)
					ast.SaveSymbol(p.Id,value)
					return nil

				}

			}

		} else if p.Type == "Bool" && value.Type == environment.BOOLEAN {
			env.(environment.Environment).SaveVariable(p.Id, value)
			ast.SaveSymbol(p.Id,value)
			return nil
		} else {
			ast.SetError(p.Lin, p.Col, "declaración de constante no coincide con el tipo definido")
		}
	}

	return nil
}

func (p ConstDec) GetConstDec(ast *environment.AST, env interface{}) interface{} {
	if p.Type == nil {

		if _, ok := p.Expression.(interfaces.Expression); ok {
			expression := p.Expression.(interfaces.Expression)
			value := expression.Execute(ast, env)
			value.Const = true
			value.Scope = env.(environment.Environment).Scope
			return value

		}
	}

	if _, ok := p.Expression.(interfaces.Expression); ok {

		expression := p.Expression.(interfaces.Expression)
		value := expression.Execute(ast, env)
		value.Scope = env.(environment.Environment).Scope
		value.Const = true
		if p.Type == "String" && value.Type == environment.STRING {
			return value

		} else if p.Type == "Int" && value.Type == environment.INTEGER {
			return value

		} else if p.Type == "Character" && value.Type == environment.CHAR {
			return value

		} else if p.Type == "Float" {
			if value.Type == environment.FLOAT {
				return value

			} else if value.Type == environment.INTEGER {
				if _, ok := value.Value.(int); ok {
					value.Value = float64(value.Value.(int))
					value.Type = environment.FLOAT
					return value

				}
			}

		} else if p.Type == "Bool" && value.Type == environment.BOOLEAN {
			return value

		} else {
			return nil
		}
	}

	return nil
}

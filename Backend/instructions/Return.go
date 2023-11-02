package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type Return struct {
	Lin int
	Col int
	Exp interface{}
}

func NewReturn(lin int, col int, exp interface{}) Return {
	breakInstr := Return{lin, col, exp}
	return breakInstr
}

func (p Return) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	if _, isBreak := p.Exp.(interfaces.Expression); isBreak {
		gen.AddComment("----Return----")
		val := p.Exp.(interfaces.Expression).Execute(ast, env, gen)
		response := environment.Value{
			Value:        val.Value,
			IsTemp:       true,
			Type:         val.Type,
			TrueLabel:    nil,
			FalseLabel:   nil,
			OutLabel:     nil,
			IntValue:     val.IntValue,
			FloatValue:   0,
			BreakFlag:    false,
			ContinueFlag: false,
			Dimentions:   nil,
			Const:        false,
			Scope:        0,
			Lin:          0,
			Col:          0,
			Id:           "",
		}
		gen.LastReturnValue = response
		gen.AddComment("val a retornar")
		gen.AddExpression("stack[(int)P]", "0", response.Value, "+")
		gen.AddGoto(gen.ReturnLabel)
		gen.LastReturnValue = response
		response.ReturnFlag = true
		return response

	}
	return environment.Value{}
}

package instructions

import (
	"PY1/environment"
	"PY1/expressions"
	"PY1/generator"
	"PY1/interfaces"
)

type For struct {
	Lin      int
	Col      int
	Id       string
	Range    interfaces.Expression
	insBlock []interface{}
}

func NewFor(lin int, col int, id string, rangge interfaces.Expression, insBlock []interface{}) For {
	return For{lin, col, id, rangge, insBlock}
}

func (p For) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	gen.AddComment("---FOR---")
	var result environment.Value
	newdecc := VarDec{Lin: p.Lin, Col: p.Col, Id: p.Id, Expression: p.Range.(expressions.Range).FirstIndex, Type: environment.INTEGER}
	newdecc.Execute(ast, env, gen)
	newaccess := expressions.VariableAccess{ID: newdecc.Id}
	expresssion := expressions.RelationalOperation{Lin: p.Lin, Col: p.Col, Operator: "<", LeftExp: newaccess, RightExp: p.Range.(expressions.Range).LastIndex}
	//etiqueta de retorno
	RetLvl := gen.NewLabel()
	gen.AddLabel(RetLvl)
	//ejecutando expresion (if)

	aumento := Asignation{
		Lin: p.Lin,
		Col: p.Col,
		Id:  newdecc.Id,
		Expression: expressions.ArithmeticOperation{
			Lin:      p.Lin,
			Col:      p.Col,
			OpIzq:    newaccess,
			Operator: "+",
			OpDer: expressions.Primitive{
				Lin:   p.Lin,
				Col:   p.Col,
				Valor: 1,
				Type:  environment.INTEGER,
			},
		},
	}

	comodin := append(p.insBlock, aumento)

	newif := If{
		Lin:         p.Lin,
		Col:         p.Col,
		Condition:   expresssion,
		TrueBlock:   comodin,
		ElseIfBlock: nil,
		ElseBlock:   nil,
	}

	ifval := newif.Execute2(ast, env, gen) //!!!!!! este es xd

	//******************** add break & continue lvls
	//add true labels
	for _, lvl := range ifval.TrueLabel {
		gen.AddLabel(lvl.(string))
	}
	//instrucciones while

	//retorno
	gen.AddGoto(RetLvl)
	//add false labels
	for _, lvl := range ifval.FalseLabel {
		gen.AddLabel(lvl.(string))
	}
	return result
}

package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
)

type Switch struct {
	Lin          int
	Col          int
	Condition    interfaces.Expression
	CasesList    []interface{}
	DefaultBlock []interface{}
}

func NewSwitch(lin int, col int, condition interfaces.Expression, casesList []interface{}, defaultstmt []interface{}) Switch {
	return Switch{lin, col, condition, casesList, defaultstmt}
}

func (p Switch) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	return nil
}
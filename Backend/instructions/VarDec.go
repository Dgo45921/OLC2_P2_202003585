package instructions

import (
	"PY1/environment"
	"PY1/generator"
	"PY1/interfaces"
	"strconv"
)

type VarDec struct {
	Lin        int
	Col        int
	Id         string
	Type       interface{}
	Expression interface{}
}

func NewVarDec(lin int, col int, id string, tyype interface{}, val interface{}) VarDec {
	NewVarDeclaration := VarDec{lin, col, id, tyype, val}
	return NewVarDeclaration
}

func (p VarDec) Execute(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	if env.(environment.Environment).VariableExists(p.Id) {
		ast.SetError(p.Lin, p.Col, "Error, variable ya declarada!")
		return environment.Value{}
	}
	var result environment.Value
	var newVar environment.Symbol
	result = p.Expression.(interfaces.Expression).Execute(ast, env, gen)
	if result.Type == environment.STRUCT_IMP {
		env.(environment.Environment).SaveVariableStruct(p.Id, environment.Symbol{Lin: p.Lin, Col: p.Col, Type: environment.STRUCT_IMP, Value: result.StructValues, StructType: result.Id})
		return result
	}

	gen.AddComment("Agregando una declaracion")
	newVar = env.(environment.Environment).SaveVariable(p.Id, result.Type)
	extra := result
	extra.Id = p.Id
	extra.Scope = env.(environment.Environment).Scope
	extra.Lin = p.Lin
	extra.Col = p.Col
	ast.SaveSymbol(p.Id, extra)

	if result.Type == environment.BOOLEAN {
		//si no es temp (boolean)
		newLabel := gen.NewLabel()
		//add labels
		for _, lvl := range result.TrueLabel {
			gen.AddLabel(lvl.(string))
		}
		gen.AddSetStack(strconv.Itoa(newVar.Position), "1")
		gen.AddGoto(newLabel)
		//add labels
		for _, lvl := range result.FalseLabel {
			gen.AddLabel(lvl.(string))
		}
		gen.AddSetStack(strconv.Itoa(newVar.Position), "0")
		gen.AddGoto(newLabel)
		gen.AddLabel(newLabel)
		gen.AddBr()
	} else {
		//si es temp (num,string,etc)
		gen.AddSetStack(strconv.Itoa(newVar.Position), result.Value)
		gen.AddBr()
	}

	return result
}

func (p VarDec) GetVarDec(ast *environment.AST, env interface{}) interface{} {
	return p
}

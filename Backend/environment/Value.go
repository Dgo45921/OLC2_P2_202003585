package environment

type Value struct {
	Value        string
	IsTemp       bool
	Type         TipoExpresion
	TrueLabel    []interface{}
	FalseLabel   []interface{}
	OutLabel     []interface{}
	IntValue     int
	FloatValue   float64
	BreakFlag    bool
	ContinueFlag bool
	ReturnFlag  bool
	Dimentions [] int
	Const 	bool
	Scope EnvType
	Lin int
	Col int
	Id string
	StructValues []KeyValue
}

func NewValue(Val string, tmp bool, typ TipoExpresion) Value {
	result := Value{
		Value:    Val,
		IsTemp:   tmp,
		Type:     typ,
		IntValue: 0,
	}
	return result
}


func (s Value) GetType() string {
	if s.Type == STRING {
		return "STRING"
	} else if s.Type == INTEGER {
		return "INTEGER"
	} else if s.Type == FLOAT {
		return "FLOAT"
	} else if s.Type == CHAR {
		return "CHARACTER"
	} else if s.Type == BOOLEAN {
		return "BOOLEAN"
	} else if s.Type == VECTOR {
		return "VECTOR"
	} else if s.Type == VECTOR_INT {
		return "VECTOR_INT"
	} else if s.Type == VECTOR_FLOAT {
		return "VECTOR_FLOAT"
	} else if s.Type == VECTOR_CHAR {
		return "VECTOR_CHAR"
	} else if s.Type == VECTOR_STRING {
		return "VECTOR_STRING"
	} else if s.Type == VECTOR_BOOLEAN {
		return "VECTOR_BOOLEAN"
	} else if s.Type == VECTOR_STRUCT {
		return "VECTOR_STRUCT"
	} else if s.Type == MATRIX_INT {
		return "MATRIX_INT"
	} else if s.Type == MATRIX_FLOAT {
		return "MATRIX_FLOAT"
	} else if s.Type == MATRIX_CHAR {
		return "MATRIX_CHAR"
	} else if s.Type == MATRIX_STRING {
		return "MATRIX_STRING"
	} else if s.Type == MATRIX_BOOLEAN {
		return "MATRIX_BOOLEAN"
	} else if s.Type == STRUCT_DEF {
		return "STRUCT_DEF"
	} else if s.Type == STRUCT_IMP {
		return "STRUCT_IMP"
	}

	return "NULL"
}

func (s Value) GetScopeType() string {
	if s.Scope == GLOBAL {
		return "GLOBAL"
	} else if s.Scope == IF {
		return "IF"
	} else if s.Scope == FOR {
		return "FOR"
	} else if s.Scope == WHILE {
		return "WHILE"
	} else if s.Scope == FUNC {
		return "FUNC"
	} else if s.Scope == CASE {
		return "CASE"
	} else if s.Scope == DEFAULT {
		return "DEFAULT"
	} else if s.Scope == ElSE {
		return "ELSE"
	} else if s.Scope == ELSEIF {
		return "ELSE-IF"
	} else if s.Scope == GUARD {
		return "GUARD"
	}

	return "NULL"
}
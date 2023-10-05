package environment

type Symbol struct {
	Lin        int
	Col        int
	Type       TipoExpresion
	Value      interface{}
	Const      bool
	StructType string
	Scope      EnvType
}

func (s Symbol) SetnewValue(value interface{}) {
	s.Value = value
}

func (s Symbol) GetType() string {
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

func (s Symbol) GetScopeType() string {
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

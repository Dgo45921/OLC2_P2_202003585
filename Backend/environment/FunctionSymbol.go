package environment

type FunctionSymbol struct {
	Lin        int
	Col        int
	ReturnType TipoExpresion
	Args       []FuncParam
	InsBlock   []interface{}
	StructType string
	Mutating   bool
}

func (s FunctionSymbol) GetType() string {
	if s.ReturnType == STRING {
		return "STRING"
	} else if s.ReturnType == INTEGER {
		return "INTEGER"
	} else if s.ReturnType == FLOAT {
		return "FLOAT"
	} else if s.ReturnType == CHAR {
		return "CHARACTER"
	} else if s.ReturnType == BOOLEAN {
		return "BOOLEAN"
	} else if s.ReturnType == VECTOR {
		return "VECTOR"
	} else if s.ReturnType == VECTOR_INT {
		return "VECTOR_INT"
	} else if s.ReturnType == VECTOR_FLOAT {
		return "VECTOR_FLOAT"
	} else if s.ReturnType == VECTOR_CHAR {
		return "VECTOR_CHAR"
	} else if s.ReturnType == VECTOR_STRING {
		return "VECTOR_STRING"
	} else if s.ReturnType == VECTOR_BOOLEAN {
		return "VECTOR_BOOLEAN"
	} else if s.ReturnType == VECTOR_STRUCT {
		return "VECTOR_STRUCT"
	} else if s.ReturnType == MATRIX_INT {
		return "MATRIX_INT"
	} else if s.ReturnType == MATRIX_FLOAT {
		return "MATRIX_FLOAT"
	} else if s.ReturnType == MATRIX_CHAR {
		return "MATRIX_CHAR"
	} else if s.ReturnType == MATRIX_STRING {
		return "MATRIX_STRING"
	} else if s.ReturnType == MATRIX_BOOLEAN {
		return "MATRIX_BOOLEAN"
	} else if s.ReturnType == STRUCT_DEF {
		return "STRUCT_DEF"
	} else if s.ReturnType == STRUCT_IMP {
		return "STRUCT_IMP"
	}

	return "NULL"
}

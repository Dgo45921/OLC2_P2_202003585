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
	Dimentions [] int
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

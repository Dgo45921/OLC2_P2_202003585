package environment

type EnvType int

const (
	GLOBAL EnvType = iota
	IF
	FOR
	WHILE
	FUNC
	CASE
	DEFAULT
	ElSE
	ELSEIF
	GUARD
)

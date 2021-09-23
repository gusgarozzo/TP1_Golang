package model

type Estructura struct{
	Tipo string
	Largo int
	Valores string
}

func NewEstructura(tipo string, largo int, values string) Estructura{
	return Estructura{tipo, largo, values}
}
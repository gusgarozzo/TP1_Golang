package main

import (
	"fmt"
	"tp1/model"
)

func changeValores(e *model.Estructura, valores string){
	e.Valores = valores
}

func changeTipo(e *model.Estructura, tipo string){
	e.Tipo = tipo
}

func changeLargo(e *model.Estructura, largo int){
	e.Largo = largo
}

func changeStructure(e *model.Estructura, tipo string, largo int, valores string){
	e.Tipo = tipo
	e.Largo = largo
	e.Valores = valores
}

func main() {
	e := model.NewEstructura("AX", 3, "ABC")

	changeValores(&e, "DEF")
	changeTipo(&e, "AZ")
	changeLargo(&e, 5)
	changeStructure(&e, "TX", 4, "XYZ")
	fmt.Println(e)
}

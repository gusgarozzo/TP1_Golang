package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"tp1/model"
)

func TestChangeValores(t *testing.T){
	e := model.NewEstructura("AX", 3, "ABC")

	changeStructure(&e, "XY", 5, "WXY")
	changeLargo(&e, 4)
	changeTipo(&e, "AR")
	changeValores(&e, "FGH")

	
	if assert.NotNil(t, &e){
		assert.Equal(t, e.Valores, "FGH", "Los valores no coinciden")
		assert.NotEqual(t, "FGH", "WXY", "Los valores no deben coincidir")
		assert.Equal(t, e.Tipo, "AR", "El tipo no coincide")
		assert.NotEqual(t, "AR", "XY", "El tipo no debe coincidir")
		assert.Equal(t, e.Largo, 4, "El largo no coincide")
		assert.NotEqual(t, 4, 5, "El largo no debe coincidir")
	}
	

	
}



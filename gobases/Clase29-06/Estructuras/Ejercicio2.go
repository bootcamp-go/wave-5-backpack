package main

import (
	"fmt"
	"math"
)

type Matrix struct {
	Alto       float64
	Ancho      float64
	Max        float64
	Cuadratica bool
}

func (a *Matrix) set(alto, ancho float64) {
	a.Alto = alto
	a.Ancho = ancho
	a.Max = math.Max(alto, ancho)
	if alto == ancho {
		a.Cuadratica = true
	} else {
		a.Cuadratica = false
	}
}

func (a Matrix) print() {
	fmt.Printf("Alto:\t%f\nAncho:\t%f\nMaximo:\t%f\nCuadratica:\t%v\n", a.Alto, a.Ancho, a.Max, a.Cuadratica)
}

func main() {

	matris := Matrix{}

	matris.set(3, 2)

	matris.print()

}

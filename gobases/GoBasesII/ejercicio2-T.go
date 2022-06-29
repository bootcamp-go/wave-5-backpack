package main

func main() {

}

type Matrix struct {
	Valores    float64
	Alto       float64
	Ancho      float64
	Cuadratica bool
	MaxValue   float64
}

type metodos interface {
	set(n ...float64) float64
}

package main

import "fmt"

type Vehículo struct {
	km, tiempo float64
}

func (v Vehículo) detalle() {
	fmt.Printf("km:\t%f\n", v.km)
}

type Auto struct {
	vehiculo Vehículo
}


func (a *Auto) Correr(minutos int) {
	a.vehiculo.tiempo = float64(minutos) / 60
	a.vehiculo.km = a.vehiculo.tiempo * 100
}

func (a *Auto) Detalle() {
	fmt.Println("\nV:\tAuto")
	a.vehiculo.detalle()
}

func main() {

	auto := Auto{}
	// auto.Correr(360)
	auto.Detalle()

}

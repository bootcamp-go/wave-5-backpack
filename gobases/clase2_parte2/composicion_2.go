package main

import "fmt"

type Vehículo struct {
	km, tiempo float64
}

func (v *Vehículo) detalle() {
	fmt.Printf("km:\t%f\n", v.km)
}

type Auto struct {
	v Vehículo		// Se embebe 
}

func (a *Auto) Correr(minutos int) {
	a.v.tiempo = float64(minutos) / 60
	a.v.km = a.v.tiempo * 100
}

func (a *Auto) Detalle() {
	fmt.Println("\nV:\tAuto")
	a.v.detalle()
}

type Moto struct {
	v Vehículo
}

func (m *Moto) Correr(minutos int) {
	m.v.tiempo = float64(minutos) / 60
	m.v.km = m.v.tiempo * 80
}

func (m *Moto) Detalle() {
	fmt.Println("\nV:\tMoto")
	m.v.detalle()
}

func main() {
	auto := Auto{}
	auto.Correr(360)
	auto.Detalle()

	moto := Moto{}
	moto.Correr(360)
	moto.Detalle()
}

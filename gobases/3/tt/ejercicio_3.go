package main

import "fmt"

type Productos struct {
	Nombre   string  `json:"nombre"`
	Precio   float64 `json:"precio"`
	Cantidad int     `json:"cantidad"`
}

type Servicios struct {
	Nombre  string  `json:"nombre"`
	Precio  float64 `json:"precio"`
	Minutos int     `json:"minutos_trabajados"`
}

type Mantenimiento struct {
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
}

var (
	productos      []Productos
	servicios      []Servicios
	mantenimientos []Mantenimiento
)

func main() {
	p1 := Productos{
		Nombre:   "Ladrillo",
		Precio:   1.75,
		Cantidad: 200,
	}
	productos = append(productos, p1)

	s1 := Servicios{
		Nombre:  "Albañileria",
		Precio:  2000,
		Minutos: 29,
	}
	servicios = append(servicios, s1)

	m1 := Mantenimiento{
		Nombre: "Revoque",
		Precio: 500,
	}
	mantenimientos = append(mantenimientos, m1)

	cp := make(chan float64)
	cs := make(chan float64)
	cm := make(chan float64)

	go SumarProductos(productos, cp)
	go SumarServicios(servicios, cs)
	go SumarMantenimiento(mantenimientos, cm)

	total_productos := <-cp
	total_servicios := <-cs
	total_mantenimientos := <-cm

	fmt.Printf(
		"Productos: %.2f\nServicios: %.2f\nMantenimientos: %.2f\nTotal: %.2f\n",
		total_productos,
		total_servicios,
		total_mantenimientos,
		total_productos+total_servicios+total_mantenimientos,
	)
}

func SumarProductos(ps []Productos, c chan float64) {
	var total float64
	for _, p := range ps {
		total += p.Precio
	}
	c <- total
}

func SumarServicios(ss []Servicios, c chan float64) {
	var total float64
	for _, s := range ss {
		if s.Minutos < 30 {
			total += s.Precio / 2
		} else {
			total += s.Precio
		}
	}
	c <- total
}

func SumarMantenimiento(ms []Mantenimiento, c chan float64) {
	var total float64
	for _, m := range ms {
		total += m.Precio
	}
	c <- total
}

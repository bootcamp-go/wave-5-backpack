package main

import "fmt"

type Productos struct {
	nombre   string
	precio   float64
	cantidad int
}

type Servicios struct {
	nombre  string
	precio  float64
	minutos int
}

type Mantenimiento struct {
	nombre string
	precio float64
}

func  sumProductos(p []Productos, c chan float64) {
	var suma float64
	for _, prod := range p {
		suma += prod.precio * float64(prod.cantidad)
	}
	c <- suma
}

func sumServicios(s []Servicios, c chan float64){
	var suma float64
	for _, ser := range s {
		if ser.minutos < 30 {
			ser.minutos = 30
		}
		suma += ser.precio * float64(ser.minutos)
	}
	c <-suma
}

func sumMantenimiento(m []Mantenimiento,c chan float64)   {
	var suma float64
	for _, man := range m {
		suma += man.precio
	}
	c<-suma
} 

func main() {
	productos := []Productos{
		{"Monitor LG", 150000, 5},
		{"Monitor Samsung", 150000, 5},
		{"Monitor", 150000, 5},
	}

	servicios := []Servicios{
		{"Reparacion equipos", 100, 68},
		{"alguito", 254, 2},
	}

	mantenimientos := []Mantenimiento{
		{"PC", 56},
		{"Router", 135},
	}

	res := make(chan float64)

	go sumProductos(productos, res)
	go sumServicios(servicios, res)
	go sumMantenimiento(mantenimientos, res)

	preciofinal := <-res + <-res + <-res

	fmt.Printf("Precio final: $%.f \n",preciofinal)

}

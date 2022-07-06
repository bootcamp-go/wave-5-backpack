package main

import (
	"fmt"
)

type Producto interface {
	CalcularCosto() float64
}

type producto struct {
	nombre       string
	precio       float64
	tipo_de_prod string
}

func (p producto) CalcularCosto() float64 {
	dic_tip_prod := map[string][2]float64{
		"pequeño": {0.0, 0.0},
		"mediana": {1.03, 0.0},
		"grande":  {1.06, 2500},
	}
	return p.precio*(1+dic_tip_prod[p.tipo_de_prod][0]) + dic_tip_prod[p.tipo_de_prod][1]
}

type Ecommerce interface {
	Total() float64
	Agregar(Producto)
}

type tienda struct {
	lista_de_productos []Producto
}

func (t tienda) Total() float64 {
	total := 0.0
	for _, element := range t.lista_de_productos {
		total += element.CalcularCosto()
	}
	return total
}

func (t *tienda) Agregar(p Producto) {
	t.lista_de_productos = append(t.lista_de_productos, p)
}

func nuevoProducto(nombre string, precio float64, tipo string) Producto {
	return producto{nombre, precio, tipo}
}
func nuevaTienda() Ecommerce {
	t := &tienda{}
	return t
}

func detalles(p Producto) {
	fmt.Printf("%f\n", p.CalcularCosto())
}

func imprimir(obj interface{}) {
	fmt.Println(obj.(producto).precio)
}

func main() {
	p1 := producto{"collar", 6500, "pequeño"}
	fmt.Println(p1)
	detalles(p1)
	p2 := producto{"silla", 10000, "pequeño"}
	fmt.Println(p2)
	detalles(p2)
	p3 := nuevoProducto("moto", 2000000, "pequeño")
	fmt.Println(p3)
	detalles(p3)
	var t Ecommerce = nuevaTienda()
	fmt.Println(t)

	t.Agregar(p1)
	fmt.Println(t)
	// t.Agregar(p2)
	// fmt.Println(t)
	// t.Agregar(p3)
	// fmt.Printf("%f\n", t.Total())
}

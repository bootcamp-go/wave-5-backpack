package main

import "fmt"

const (
	PEQUEÑO = "PEQUEÑO"
	MEDIANO = "MEDIANO"
	GRANDE  = "GRANDE"
)

type tienda struct {
	cosas []producto
}

type producto struct {
	tipo   string
	nombre string
	precio float64
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar()
}

func nuevoProducto(tipo, nombre string, precio float64) producto {
	return (producto{tipo, nombre, precio})
}

func nuevaTienda() tienda {
	return (tienda{})
}

func (a producto) CalcularCosto() float64 {
	if a.tipo == PEQUEÑO {
		return a.precio
	} else if a.tipo == MEDIANO {
		return (a.precio + a.precio*0.03)
	} else {
		return (a.precio + a.precio*0.06 + 2500)
	}
}

func (a tienda) Total() float64 {
	var acumulador float64
	for j, _ := range a.cosas {
		acumulador = acumulador + a.cosas[j].CalcularCosto()
	}
	return acumulador
}

func (a *tienda) Agregar(b producto) {
	a.cosas = append(a.cosas, b)
}

func main() {
	t1 := nuevaTienda()
	p1 := nuevoProducto(PEQUEÑO, "jamon", 20)
	p2 := nuevoProducto(PEQUEÑO, "queso", 20)
	t1.Agregar(p1)
	t1.Agregar(p2)

	fmt.Printf("%v\n", t1)
	fmt.Printf("%v\n", t1.Total())
}

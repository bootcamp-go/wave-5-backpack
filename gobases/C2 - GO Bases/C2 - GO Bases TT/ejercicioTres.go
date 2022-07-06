package main

import "fmt"

type producto struct {
	tipoPorducto string
	nombre       string
	precio       float64
}

type tienda struct {
	listaProductos []Producto
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(p Producto)
}

func (p producto) CalcularCosto() float64 {
	costo_adicional := 0.0
	switch p.tipoPorducto {
	case "pequeno":
		costo_adicional = p.precio
	case "mediano":
		costo_adicional = p.precio + (p.precio * 0.03)
	case "grande":
		costo_adicional = p.precio + (p.precio * 0.06)
	}
	return costo_adicional
}

func (t tienda) Total() float64 {
	res := 0.0

	for _, p := range t.listaProductos {
		res += p.CalcularCosto()
	}

	return res
}

func (t *tienda) Agregar(p Producto) {
	t.listaProductos = append(t.listaProductos, p)
}

func nuevoProducto(tipoPorducto string, nombre string, precio float64) Producto {
	return &producto{
		tipoPorducto: tipoPorducto,
		nombre:       nombre,
		precio:       precio,
	}
}

func nuevaTienda(productos ...Producto) Ecommerce {
	return &tienda{
		listaProductos: productos,
	}
}

func main() {
	p1 := nuevoProducto("pequeno", "p1", 100)
	p2 := nuevoProducto("mediano", "p2", 100)
	p3 := nuevoProducto("grande", "p3", 100)

	t1 := nuevaTienda(p1, p2, p3)
	fmt.Println("Costo pequeño", p1.CalcularCosto(),"\nCosto mediano", p2.CalcularCosto(), "\nCosto grande",p3.CalcularCosto())
	//total de costos
	fmt.Println(t1.Total())
	t1.Agregar(nuevoProducto(
		"pequeno", "nuevo", 100,
	))
	//total al agregar 100 de uno nuevo pequeño
	fmt.Println(t1.Total())
}

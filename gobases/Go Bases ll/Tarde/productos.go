package main

import "fmt"

const (
	pequeno = "pequeno"
	mediano = "mediano"
	grande  = "grande"
)

type Producto struct {
	TipProducto string
	Nombre      string
	Precio      float64
}
type Tienda struct {
	Productos []producto
}

type producto interface {
	calcularPrecio() float64
}
type Ecommerce interface {
	Total() float64
	Agregar(unidad producto)
}

func nuevaTienda() Ecommerce {
	return &Tienda{}
}
func nuevoProducto(TipProducto, nombre string, precio float64) producto {
	return &Producto{TipProducto: TipProducto, Nombre: nombre, Precio: precio}
}

func (p Producto) calcularPrecio() float64 {
	switch p.TipProducto {
	case pequeno:
		return p.Precio
	case mediano:
		return p.Precio + ((3 * p.Precio) / 100)
	case grande:
		return p.Precio + ((6 * p.Precio) / 100) + 2500
	}
	return 0
}

func (t *Tienda) Agregar(unidad producto) {
	t.Productos = append(t.Productos, unidad)
}
func (t Tienda) Total() float64 {
	var total float64
	for _, p := range t.Productos {
		subPrecio := p.calcularPrecio()
		total += subPrecio
	}
	return total
}

func main() {
	p1 := nuevoProducto(pequeno, "pan", 100)
	t1 := nuevaTienda()

	t1.Agregar(p1)
	total := t1.Total()

	fmt.Printf("total: %.2f\n", total)
}

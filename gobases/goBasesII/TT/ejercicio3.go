package main

import "fmt"

const (
	PEQUENO = "pequeno"
	MEDIANO = "mediano"
	GRANDE  = "grande"
)

type producto struct {
	tipo   string
	nombre string
	precio float64
}

type tienda struct {
	productos []producto
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(product producto)
}

func nuevoProducto(tipo string, nombre string, precio float64) producto {
	return producto{tipo: tipo, nombre: nombre, precio: precio}
}

func nuevaTienda() tienda {
	return tienda{}
}
func (p *producto) CalcularCosto() float64 {
	pequeno := [2]float64{0, 0}
	mediano := [2]float64{3, 0}
	grande := [2]float64{6, 2500}

	switch p.tipo {
	case PEQUENO:
		return (p.precio*(1+pequeno[0]/100) + pequeno[1])
	case MEDIANO:
		return (p.precio*(1+mediano[0]/100) + mediano[1])
	case GRANDE:
		return (p.precio*(1+grande[0]/100) + grande[1])
	}
	return 0
}

func (t *tienda) Total() (total float64) {
	for _, p := range t.productos {
		total += p.CalcularCosto()
	}
	return
}

func (t *tienda) Agregar(p producto) {
	t.productos = append(t.productos, p)
}
func printProductos(t tienda) {
	for _, p := range t.productos {
		fmt.Printf("Tipo: %s\nNombre: %s\nPrecio: %.2f\n", p.tipo, p.nombre, p.precio)
	}
}
func main() {
	p1 := nuevoProducto(PEQUENO, "licuadora", 2000)
	p2 := nuevoProducto(MEDIANO, "televisor", 10000)
	p3 := nuevoProducto(GRANDE, "nevera", 30000)
	t1 := nuevaTienda()
	t2 := nuevaTienda()
	t1.Agregar(p1)
	t1.Agregar(p2)
	t1.Agregar(p3)
	printProductos(t1)
	fmt.Println(t1.Total())
	t2.Agregar(p1)
	t2.Agregar(p2)
	printProductos(t2)
	fmt.Println(t2.Total())
}

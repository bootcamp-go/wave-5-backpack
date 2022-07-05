package main

import "fmt"

/*Ejercicio 3 - Productos

Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
Las empresas tienen 3 tipos de productos:
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
	Pequeño: El costo del producto (sin costo adicional)
	Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
	Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

Requerimientos:
	Crear una estructura “tienda” que guarde una lista de productos.
	Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
	Crear una interface “Producto” que tenga el método “CalcularCosto”
	Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
	Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
	Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
Interface Producto:
El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
	El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
	El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda*/

const (
	PEQUE string = "PEQUEÑO"
	MED   string = "MEDIANO"
	GRAN  string = "GRANDE"
)

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(producto Producto)
}

type Tienda struct {
	listProd []Producto
}

type producto struct {
	tipo, nombre string
	precio       float64
}

func (p *producto) CalcularCosto() float64 {
	var totalCosto float64

	switch p.tipo {
	case "PEQUEÑO":
		totalCosto = p.precio
	case "MEDIANO":
		totalCosto = p.precio + (p.precio * 0.03)
	case "GRANDE":
		totalCosto = p.precio + (p.precio * 0.06) + 2500
	default:
		fmt.Println("No existe el tipo")
	}
	return totalCosto
}

func (t *Tienda) Total() float64 {
	var total float64

	for _, precio := range t.listProd {
		total += precio.CalcularCosto()
	}

	return total
}

func (t *Tienda) Agregar(prd Producto) {
	t.listProd = append(t.listProd, prd)
}

func nuevoProducto(t, n string, pre float64) Producto {

	return &producto{tipo: t, nombre: n, precio: pre}
}

func nuevaTienda(productos ...Producto) Ecommerce {

	return &Tienda{productos}
}

func main() {

	prd1 := nuevoProducto(PEQUE, "samsung", 400)
	prd2 := nuevoProducto(MED, "gamer", 200)
	prd3 := nuevoProducto(GRAN, "AC", 600)

	tienda := nuevaTienda(prd1, prd2, prd3)

	fmt.Println("Total tienda:", tienda.Total())

	prd4 := nuevoProducto(GRAN, "AC", 700)

	tienda.Agregar(prd4)

	fmt.Println("Total tienda:", tienda.Total())

}

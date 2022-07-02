package main

import (
	"fmt"
	"strings"
)

// Ejercicio 3 - Productos
// Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar
// productos y retornar el valor del precio total.

// Las empresas tienen 3 tipos de productos:
//   - Pequeño, Mediano y Grande. (Se espera que sean muchos más)

// Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

// Sus costos adicionales son:
//   - Pequeño: El costo del producto (sin costo adicional)
//   - Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
//   - Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

//Requerimientos:
//   - Crear una estructura “tienda” que guarde una lista de productos.
//   - Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
//   - Crear una interface “Producto” que tenga el método “CalcularCosto”
//   - Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
//   - Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
//   - Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
//   - Interface Producto:
//       - El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
//   - Interface Ecommerce:
//       - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
//       - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda

// Constantes del tipo de producto
const SMALL = "S"
const MEDIUM = "M"
const BIG = "G"

// Estructura de Tienda
type tienda struct {
	productos []producto
}

// Funcion para obtener el total de los productos de la tienda
func (t tienda) Total() {
	total := 0.0
	for _, v := range t.productos {
		total += v.CalcularCosto()
	}
	fmt.Printf("El total de los productos es de %s pesos\n\n", formatearMoneda(total))
}

// Funcion para agregar productos a la tienda
func (t *tienda) Agregar(p producto) {
	t.productos = append(t.productos, p)
}

// Estructura de Producto
type producto struct {
	tipo   string
	nombre string
	precio float64
}

// Funcion para calcular el costo dependiendo del tipo de producto
func (p producto) CalcularCosto() float64 {
	switch p.tipo {
	case SMALL:
		return p.precio
	case MEDIUM:
		return p.precio + (p.precio * 0.02)
	case BIG:
		return p.precio + (p.precio * 0.06) + 2500
	default:
		return 0.0
	}
}

// Intercace Producto
type Producto interface {
	CalcularCosto() float64
}

// Interface Ecommerce
type Ecommerce interface {
	Total()
	Agregar(producto)
}

// Funcion para crear un nuevo producto
func nuevoProducto(tipo string, nombre string, precio float64) Producto {
	return producto{tipo: tipo, nombre: nombre, precio: precio}
}

// Funcion para crear nueva tienda
func nuevaTienda() Ecommerce {
	return &tienda{productos: make([]producto, 0)}
}

// Función para dar formato a moneda
func formatearMoneda(m float64) string {
	// Formateamos la cantidad a string
	money := fmt.Sprintf("%.2f", m)
	// Separamos la cantidad de su decimal
	moneyElements := strings.Split(money, ".")
	// Invertimos la cantidad
	moneyInverted := ""
	for _, v := range moneyElements[0] {
		moneyInverted = string(v) + moneyInverted
	}
	// Reinvertimos la cantidad y agregamos las comas
	moneyValid := ""
	for i, v := range moneyInverted {
		if (i+1)%3 == 0 && (i+1) != len(moneyInverted) {
			moneyValid = "," + string(v) + moneyValid
		} else {
			moneyValid = string(v) + moneyValid
		}
	}
	// Regresamos el resultado
	return "$" + moneyValid + "." + moneyElements[1]
}

func main() {
	fmt.Println("Ejercicio 3 - Productos")
	fmt.Println("")

	// Creamos 3 productos
	p_cucharas := nuevoProducto(SMALL, "cucharas", 100.0)
	fmt.Println(formatearMoneda(p_cucharas.CalcularCosto()))
	p_tazas := nuevoProducto(MEDIUM, "tazas", 150.0)
	fmt.Println(formatearMoneda(p_tazas.CalcularCosto()))
	p_sillas := nuevoProducto(BIG, "sillas", 400.0)
	fmt.Println(formatearMoneda(p_sillas.CalcularCosto()))

	// Creamos nueva tienda y agregamos los 3 productos creados
	nt := nuevaTienda()
	var i_cucharas interface{} = p_cucharas
	ic := i_cucharas.(producto)
	nt.Agregar(ic)

	var i_tazas interface{} = p_tazas
	it := i_tazas.(producto)
	nt.Agregar(it)

	var i_sillas interface{} = p_sillas
	is := i_sillas.(producto)
	nt.Agregar(is)

	// Mostramos el total de los productos de la tienda
	nt.Total()
}

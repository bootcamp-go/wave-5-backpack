package main

import "fmt"

const (
	PEQUENIO = "PEQUENIO"
	MEDIANO  = "MEDIANO"
	GRANDE   = "GRANDE"
)

type Producto interface {
	costoPorTipo(valor float32) float32
}

type Ecommerce interface {
	total() float32
	agregar(producto producto)
}

type tienda struct {
	productos []producto
}

type producto struct {
	tipo   Producto
	nombre string
	precio float32
}

type pequenio struct {
	tipo string
}

type mediano struct {
	tipo string
}

type grande struct {
	tipo string
}

func (t pequenio) costoPorTipo(valor float32) float32 {
	return valor
}

func (t mediano) costoPorTipo(valor float32) float32 {
	return valor + valor*0.03
}

func (t grande) costoPorTipo(valor float32) float32 {
	return valor + valor*0.06 + 2500
}

func (p producto) calcularcosto() float32 {
	return p.tipo.costoPorTipo(p.precio)
}

func nuevoProducto(tipoProducto Producto, nombre string, precio float32) producto {
	return producto{tipoProducto, nombre, precio}
}

func nuevaTienda() Ecommerce {
	return &tienda{}
}

func (t *tienda) agregar(productoNuevo producto) {
	t.productos = append(t.productos, productoNuevo)
}

func (t *tienda) total() float32 {
	var acum float32 = 0.0
	for _, producto := range t.productos {
		acum += producto.calcularcosto()
	}
	return acum
}

func main() {

	tienda := nuevaTienda()

	lapiz := producto{pequenio{PEQUENIO}, "Lapiz", 2.0}
	licuadora := producto{mediano{MEDIANO}, "Licuadora", 100.0}
	sillon := producto{grande{GRANDE}, "Sillon", 5000.0}

	tienda.agregar(lapiz)
	tienda.agregar(licuadora)
	tienda.agregar(sillon)

	fmt.Println("Total en tienda: ", tienda.total())
}

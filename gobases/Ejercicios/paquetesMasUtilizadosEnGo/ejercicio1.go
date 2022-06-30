package main

import "fmt"

type producto struct {
	ID       int
	Precio   float64
	Cantidad int
}

func (p producto) detalle() string {
	return fmt.Sprintf("ID: %d; Precio: %.2f; Cantidad: %d\n", p.ID, p.Precio, p.Cantidad)
}

func (p producto) imprimir()

type Producto interface {
	detalle() string
}

func newProducto(iD int, precio float64, cantidad int) Producto {
	return &producto{
		ID:       iD,
		Precio:   precio,
		Cantidad: cantidad,
	}

}

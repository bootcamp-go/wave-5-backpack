package main

import "fmt"

type producto struct {
	ID       int64
	Precio   float64
	Cantidad int64
}

func (p producto) detalle() string {
	return fmt.Sprintf("ID: %d; Precio: %.2f; Cantidad: %d\n", p.ID, p.Precio, p.Cantidad)
}

func (p producto) imprimir(i int) string {
	if i == 0 {
		fmt.Println("ID          Precio          Cantidad")
	}
	return fmt.Sprintf("%d %13.f %13.d", p.ID, p.Precio, p.Cantidad)

}

func (p producto) precio() float64 {
	return p.Precio
}

type Producto interface {
	detalle() string
	imprimir(i int) string
	precio() float64
}

func newProducto(iD int64, precio float64, cantidad int64) Producto {
	return &producto{
		ID:       iD,
		Precio:   precio,
		Cantidad: cantidad,
	}

}

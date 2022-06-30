package main

import (
	"fmt"
	"os"
)

type producto struct {
	id       int
	precio   float64
	cantidad int
}

type lista struct {
	productos []producto
}

func (l *lista) Agregar(prod producto) {

	l.productos = append(l.productos, prod)
}

func nuevoProducto(id int, precio float64, cantidad int) producto {
	return producto{id, precio, cantidad}
}

type Productos interface {
	detalle()
	Agregar(prod producto)
	guardarTxt()
}

func (l lista) detalle() {
	for _, prod := range l.productos {
		fmt.Println("ID :", prod.id)
		fmt.Println("Precio :", prod.precio)
		fmt.Println("Cantidad :", prod.cantidad)
		fmt.Println("-----------")
	}
}

func (l lista) guardarTxt() {
	res := ""
	for _, prod := range l.productos {
		res = res + fmt.Sprintf("%d;%f;%d \n", prod.id, prod.precio, prod.cantidad)
	}
	//fmt.Println(res)
	d1 := []byte(res)
	err := os.WriteFile("./productos.txt", d1, 0644)
	if err == nil {
		fmt.Println("Se guardo con exito")
	}

}

func nuevaLista() Productos {
	return &lista{}
}

func main() {

	Productos := nuevaLista()
	Productos.Agregar(nuevoProducto(12, 8000.0, 3))
	Productos.Agregar(nuevoProducto(9, 2000.0, 2))
	Productos.Agregar(nuevoProducto(7, 1000.0, 3))
	//Productos.detalle()
	Productos.guardarTxt()
}

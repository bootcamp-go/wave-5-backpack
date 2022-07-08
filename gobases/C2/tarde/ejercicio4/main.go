package main

import (
	"fmt"
)

/*Ejercicio 4 - Envios

Un Ecommerce necesita realizar una funcionalidad en Go para gestionar el envío y reparto de productos:
La empresa tiene 5 tipos de productos: Chico, Mediano, Grande, Especial, Frágil.
Cada producto tiene el tamaño en centímetros cúbicos. Y además cada tipo de producto requiere un
adicional al momento de ser enviado:

	Chico: Ningún adicional.
	Mediano: Requiere un %5 más de espacio
	Grande: Requiere un %20 más de espacio
	Frágil: Requiere un %75 más de espacio
	Especial: Sólo puede ser enviado con productos especiales

Para ello requerimos que los productos guarden el tamaño y tengan un método Tamaño Total que nos
devuelva el espacio en cm3 que requerimos para ser enviado.

Y una estructura Flete que tenga los métodos:
	1. Agregar Producto: agregar producto al flete
	2. Calcular Envios: calcula la cantidad de envíos que debe realizar sabiendo que solo puede
	cargar un total de 10.000.000 cm3 por envío.
*/

const (
	Chico    = "chico"
	Mediano  = "mediano"
	Grande   = "grande"
	Fragil   = "frágil"
	Especial = "especial"
)

type Ecommerce interface {
	AgregarProducto(prod Producto)
	CalcularEnvio() (float64, error)
}

// Producto
type Producto struct {
	tamanio    string
	tamanioCm3 float64
}

func (t Producto) TamanioTotal() float64 {
	return 10000000 // cm3
}

func New() Ecommerce {
	return &Flete{}
}

type Flete struct {
	producto Producto
}

func (f *Flete) AgregarProducto(prod Producto) {
	f.producto = prod
}

func (f Flete) CalcularEnvio() (float64, error) {
	switch f.producto.tamanio {
	case Chico:
		tamanioProducto := f.producto.tamanioCm3 + ((f.producto.tamanioCm3 * 5) / 100)
		if tamanioProducto > f.producto.TamanioTotal() {
			return 0, fmt.Errorf("Tamaño %.0f es superior al permitido %.0f", tamanioProducto, f.producto.TamanioTotal())
		}
		return tamanioProducto, nil
	case Mediano:
		tamanioProducto := f.producto.tamanioCm3 + ((f.producto.tamanioCm3 * 20) / 100)
		if tamanioProducto > f.producto.TamanioTotal() {
			return 0, fmt.Errorf("Tamaño %.0f es superior al permitido %.0f", tamanioProducto, f.producto.TamanioTotal())
		}
		return tamanioProducto, nil
	case Grande:
		tamanioProducto := f.producto.tamanioCm3 + ((f.producto.tamanioCm3 * 75) / 100)
		if tamanioProducto > f.producto.TamanioTotal() {
			return 0, fmt.Errorf("Tamaño %.0f es superior al permitido %.0f", tamanioProducto, f.producto.TamanioTotal())
		}
		return tamanioProducto, nil
	case Especial:
		return 0, fmt.Errorf("Solo puede ser enviado con productos especiales")
	default:
		return 0, fmt.Errorf("Tipo de tamaño no especificado.")
	}
}

func main() {

	flete := New()
	producto := Producto{
		tamanio:    Chico,
		tamanioCm3: 100000,
	}

	flete.AgregarProducto(producto)
	calculo, err := flete.CalcularEnvio()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Calculo de envío", calculo)
}

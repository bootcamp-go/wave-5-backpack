package main

import (
	"errors"
	"fmt"
)

// -------- Productos ---------

type Producto interface {
	CalcularCosto() int
}

type producto struct {
	tipoProducto string
	nombre       string
	precio       int
}

func (p producto) CalcularCosto() int {
	costeo, err := newCosteo(p.tipoProducto)
	if err != nil {
		return p.precio
	}
	return costeo.calcularCosto(p.precio)
}

func nuevoProducto(tipo string, nombre string, precio int) Producto {
	producto := producto{nombre: nombre, precio: precio, tipoProducto: tipo}
	return &producto
}

// ------ Tienda ----------

type Ecommerce interface {
	Total() int
	Agregar(Producto)
}

type tienda struct {
	Productos []Producto
}

func (t *tienda) Total() int {
	total := 0
	for _, prod := range t.Productos {
		total += prod.CalcularCosto()
	}
	return total
}

func (t *tienda) Agregar(producto Producto) {
	t.Productos = append(t.Productos, producto)
}

func nuevaTienda() Ecommerce {
	return &tienda{}
}

// -------------- Costeos -----------------

type costeo interface {
	calcularCosto(costo int) int
}

type costeoPequeño struct{}
type costeoGrande struct{}
type costeoMediano struct{}

func (cos costeoPequeño) calcularCosto(costo int) int {
	return costo
}
func (cos costeoGrande) calcularCosto(costo int) int {
	return costo + calcularPorcentage(costo, 6) + 2500
}

func (cos costeoMediano) calcularCosto(costo int) int {
	return costo + calcularPorcentage(costo, 3)
}

func newCosteoPequeño() costeo { return &costeoPequeño{} }
func newCosteoGrande() costeo  { return &costeoGrande{} }
func newCosteoMediano() costeo { return &costeoMediano{} }
func newCosteo(tipoProducto string) (costeo, error) {
	costosRegistrados := map[string]costeo{
		PEQUEÑO: newCosteoPequeño(),
		MEDIANO: newCosteoMediano(),
		GRANDE:  newCosteoGrande(),
	}
	costo := costosRegistrados[tipoProducto]
	if costo == nil {
		return nil, errors.New("Error: Tipo de Producto sin costo Implementado")
	}
	return costo, nil
}

// --------- TIPO CONSTANTS ----------

const (
	PEQUEÑO = "pequeño"
	MEDIANO = "mediano"
	GRANDE  = "grande"
)

// --------- utils ------------------

func calcularPorcentage(numero int, porcentage int) int {
	return int(float32(numero) * (float32(porcentage) / float32(100)))
}

// Comando
func main() {
	tienda := nuevaTienda()
	tienda.Agregar(nuevoProducto(PEQUEÑO, "Télefono s9", 200000))
	tienda.Agregar(nuevoProducto(MEDIANO, "Microondas", 100000))
	tienda.Agregar(nuevoProducto(GRANDE, "Lavadora", 300000))
	fmt.Println(tienda.Total())
}

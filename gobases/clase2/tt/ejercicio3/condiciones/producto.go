package condiciones

import "errors"

func NuevoProducto(tipo, nombre string, precio float64) (Producto, error) {
	// Switch para cada producto
	switch tipo {
	case "pequeno":
		return &productoPequeno{
			producto{
				tipo:   tipo,
				nombre: nombre,
				precio: precio,
			},
		}, nil
	case "mediano":
		return &productoMediano{
			producto{
				tipo:   tipo,
				nombre: nombre,
				precio: precio,
			},
		}, nil
	case "grande":
		return &productoGrande{
			producto{
				tipo:   tipo,
				nombre: nombre,
				precio: precio,
			},
		}, nil
	}

	return nil, errors.New("no tenemos ese tipo de producto")
}

type Producto interface {
	CalcularCosto() float64
}

type producto struct {
	tipo   string
	nombre string
	precio float64
}

// Productos concretos, deben implementar la interfaz
type productoPequeno struct {
	producto
}

func (p productoPequeno) CalcularCosto() float64 {
	return p.precio
}

type productoMediano struct {
	producto
}

func (p productoMediano) CalcularCosto() float64 {
	extra := p.precio * 0.03

	return p.precio + extra
}

type productoGrande struct {
	producto
}

func (p productoGrande) CalcularCosto() float64 {
	extra := p.precio * 0.06
	envio := float64(2500)

	return p.precio + extra + envio
}

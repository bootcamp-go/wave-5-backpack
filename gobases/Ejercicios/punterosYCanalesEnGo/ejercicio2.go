package main

type producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func sumarProductos(p []producto, prod chan float64) {
	total := 0.0

	for _, producto := range p {
		total += (producto.Precio * float64(producto.Cantidad))
	}

	prod <- total
}

func newProducto(nombre string, precio float64, cantidad int) producto {
	return producto{
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: cantidad,
	}
}

func addProducto(u *usuario, p *producto, cantidad int) {
	p.Cantidad = cantidad

	u.Productos = append(u.Productos, *p)
}

func deleteProdUsu(u *usuario) {
	var arrayVacio []producto
	u.Productos = arrayVacio
}

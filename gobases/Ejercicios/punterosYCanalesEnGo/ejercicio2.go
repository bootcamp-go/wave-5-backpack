package main

type producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func newProducto(nombre string, precio float64) producto {
	return producto{
		Nombre: nombre,
		Precio: precio,
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

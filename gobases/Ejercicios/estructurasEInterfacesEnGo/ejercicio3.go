package main

type tienda struct {
	Productos []producto
}

func (tienda *tienda) agregar(producto producto) {
	tienda.Productos = append(tienda.Productos, producto)
}

func (tienda tienda) total() float64 {
	var precioTotal float64 = 0
	for _, producto := range tienda.Productos {
		precioTotal += producto.Precio
	}

	return precioTotal
}

type producto struct {
	Tipo   string
	Nombre string
	Precio float64
}

func (producto *producto) calcularCosto() producto {
	switch producto.Tipo {
	case "mediano":
		producto.Precio += (producto.Precio * 3) / 100
	case "grande":
		producto.Precio += ((producto.Precio * 6) / 100) + 2500
	}
	return *producto
}

type Producto interface {
	calcularCosto() producto
}

type Ecommerce interface {
	total() float64
	agregar(producto)
}

func newProducto(tipo string, nombre string, precio float64) Producto {
	return &producto{
		Tipo:   tipo,
		Nombre: nombre,
		Precio: precio,
	}
}

func newTienda(productos []producto) Ecommerce {
	return &tienda{
		Productos: productos,
	}
}

package condiciones

func NuevaTienda(productos ...Producto) Ecommerce {
	return &tienda{
		productos: productos,
	}
}

type Ecommerce interface {
	Total() float64
	Agregar(Producto)
}

type tienda struct {
	productos []Producto
}

func (t *tienda) Total() float64 {
	var total float64

	for _, v := range t.productos {
		total += v.CalcularCosto()
	}

	return total
}

func (t *tienda) Agregar(p Producto) {
	t.productos = append(t.productos, p)
}

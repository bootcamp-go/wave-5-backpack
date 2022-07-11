package main

type tienda struct {
	products map[string]producto
}

type producto struct {
	name    string
	price   float64
	product string
}

type Producto interface {
}

type Ecomerce interface {
}

func nuevoProducto(productType string, productName string, price float64) producto {
	product := producto{
		name:    productName,
		price:   price,
		product: productType,
	}
	return product
}

func nuevaTienda() Ecomerce {
	var ecomerce Ecomerce
	return ecomerce
}

func main() {

}

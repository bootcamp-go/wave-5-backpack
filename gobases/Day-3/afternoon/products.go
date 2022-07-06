package main

import "fmt"

type User struct {
	Name     string
	LastName string
	Email    string
	Products []Product
}

type Product struct {
	Name   string
	Price  float64
	Amount int
}

func NewProduct(nombre *string, precio *float64) *Product {
	return &Product{Name: *nombre, Price: *precio}
}

func (this *User) AddProduct(product *Product, cantidad *int) {
	product.Amount = *cantidad
	this.Products = append(this.Products, *product)
}

func (this *User) DeleteProducts() {
	this.Products = []Product{}
}

func main() {
	var (
		nombre   = "Milo"
		precio   = 3.0
		cantidad = 4
	)

	producto := NewProduct(&nombre, &precio)
	usuario := &User{
		Name:     "Marcelo",
		LastName: "Gonzalez",
		Email:    "test@gmail.com",
	}
	usuario.AddProduct(producto, &cantidad)

	for _, value := range usuario.Products {
		fmt.Printf("Producto - %s - Cantidad %d - Precio %.2f - Total $%.2f\n", value.Name, value.Amount, value.Price, (value.Price * float64(value.Amount)))
	}
	fmt.Println(usuario.Products)
	usuario.DeleteProducts()
	fmt.Println(usuario.Products)
}

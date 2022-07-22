package main

import "fmt"

type User struct {
	Name     string
	Surname  string
	Email    string
	Products []Product
}

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func NewProduct(name string, price float64) *Product {
	return &Product{Name: name, Price: price}
}

func (u *User) AddProduct(product Product, quantity int) {
	product.Quantity = quantity
	u.Products = append(u.Products, product)
}

func (u *User) DeleteProduct() {
	u.Products = []Product{}
}

func main() {
	product := NewProduct("La serenisima", 1085.2)
	user := &User{
		Name:    "Marcelo",
		Surname: "Gonzalez",
		Email:   "test@gmail.com",
	}

	user.AddProduct(*product, 20)

	fmt.Println("Usuario - ", user.Name, user.Surname)
	fmt.Println("Correo - ", user.Email)
	for _, value := range user.Products {
		fmt.Println(value)
	}

	user.DeleteProduct()

	fmt.Println("PRODUCTOS ELIMINADOS\n", user.Products)
}

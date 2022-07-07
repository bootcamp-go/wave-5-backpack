package main

import "fmt"

type Usuario struct {
	name     string
	lastName string
	email    string
	products []Producto
}

type Producto struct {
	name  string
	price float64
	stock int
}

func main() {
	var (
		pUser *Usuario
		pItem *Producto
		user  = Usuario{"Andy", "Esquivel", "a@gmail.com", []Producto{}}
		item1 = newItem("Lavadora", 16832.94)
	)
	pUser = &user
	pItem = &item1
	fmt.Println("User created: ", user)
	fmt.Println("Item created: ", item1)
	addItem(pUser, pItem, 4)
	fmt.Println("User products: ", *&pUser.products)

	deleteItems(pUser)
	fmt.Println("User products: ", *&pUser.products)

}

func newItem(name string, price float64) Producto {
	return Producto{name, price, 0}
}

func deleteItems(user *Usuario) {
	*&user.products = []Producto{}
}

func addItem(user *Usuario, item *Producto, stock int) {
	*&item.stock += stock
	*&user.products = append(*&user.products, *item)
}

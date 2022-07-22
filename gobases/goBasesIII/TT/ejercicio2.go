package main

import "fmt"

type Product struct {
	Name   string
	Price  float64
	Amount int
}

type User struct {
	Name     string
	LastName string
	Mail     string
	Products []Product
}

func newProduct(name string, price float64) Product {
	return Product{Name: name, Price: price}
}

func addProduct(u *User, p Product, amount int) {
	p.Amount = amount
	u.Products = append(u.Products, p)
}

func deleteProduct(u *User) {
	u.Products = []Product{}
}

func main() {
	p1 := newProduct("lavadora", 12000)
	p2 := newProduct("nevera", 30000)
	user1 := User{Name: "camilo", LastName: "perez", Mail: "mymail@mail.com"}
	addProduct(&user1, p1, 6)
	addProduct(&user1, p2, 2)
	fmt.Println(user1)
	fmt.Println(p1)
	fmt.Println(p2)
	deleteProduct(&user1)
	fmt.Println(user1)
}

package main

import "fmt"

type user struct{
	name 		string
	lastName 	string
	email 		string
	userProduct product
}
type product struct{
	name 		string
	price 		float64
	quantity 	int
}

func main(){
//	users 		:= []user
	products 	:= []product


	newProduct(products,"cafe",4.3)
	fmt.Println(products)
}

// func addProduct(){

// }

func newProduct(p *product, name string, price float64){
		prod:= product{name,price,nil}
		p=append(p,prod)
}


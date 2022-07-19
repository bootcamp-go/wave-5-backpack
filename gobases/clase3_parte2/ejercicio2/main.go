package main

import "fmt"
/*
	*************** STRUCTURES ***************
*/
type Usser struct {
	name     string
	l_name   string
	email    string
	products []Product
}

type Product struct {
	name     string
	price    int
	quantity int
}

/*
*************** METHODS ***************
 */
func (p *Product) newProduct(name string, price int) Product {
	(*p).name = name
	(*p).price = price
	return *p
}
func (u *Usser) addProduct(product Product, quantity int) {
	product.quantity = quantity
	(*u).products = append((*u).products, product)
}

func (u *Usser) deleteProduct() {
	if len(u.products) > 0{
		u.products = nil
	}
}

/*
*************** MAIN ***************
 */
func main() {
	p1 := Product{}
	p2 := Product{}

	p1.newProduct("Jabon de losa x500", 12000)
	p2.newProduct("Arroz", 1000)

	var usser = Usser{
		name: "Michael",
		l_name: "Torres",
		email: "michaelstiven.torres@mercadolibre.com.co",
	}
	var usser2 = Usser{
		name: "Ander",
		l_name: "Torres",
		email: "ander.torres@mercadolibre.com.co",
	}

	usser.addProduct(p1, 5)
	usser2.addProduct(p2, 5)

	usser.deleteProduct()

	fmt.Printf("%v\n", usser)
}

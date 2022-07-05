package main

import "fmt"

type product struct {
	name     string
	price    int
	quantity int
}

type service struct {
	name    string
	price   int
	minutes int
}

type mantainence struct {
	name  string
	price int
}

func sumProducts(products *[]product, c chan int) {
	sum := 0
	for _, p := range *products {
		sum += p.price * p.quantity
	}
	c <- sum
}

func sumServices(services *[]service, c chan int) {
	sum := 0
	for _, s := range *services {
		sum += s.price / 2 * ((s.minutes + 29) / 30)
	}
	c <- sum
}

func sumMantainence(mantainences *[]mantainence, c chan int) {
	sum := 0
	for _, m := range *mantainences {
		sum += m.price
	}
	c <- sum
}

func main() {
	products := []product{
		{
			name:     "p1",
			price:    100,
			quantity: 10,
		},
		{
			name:     "p2",
			price:    200,
			quantity: 10,
		},
		{
			name:     "p3",
			price:    3,
			quantity: 10,
		},
	}

	mantainences := []mantainence{
		{
			name:  "m1",
			price: 1000,
		},
		{
			name:  "m2",
			price: 3000,
		},
		{
			name:  "m3",
			price: 6000,
		},
	}

	services := []service{
		{
			name:    "s1",
			price:   100,
			minutes: 120,
		},
		{
			name:    "s1",
			price:   300,
			minutes: 15,
		},
	}
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	go sumProducts(&products, c1)
	go sumMantainence(&mantainences, c2)
	go sumServices(&services, c3)
	for i := 0; i < 3; i++ {
		select {
		case ans1 := <-c1:
			fmt.Println("Products sum: ", ans1)
		case ans2 := <-c2:
			fmt.Println("Mantainences sum: ", ans2)
		case ans3 := <-c3:
			fmt.Println("Services sum: ", ans3)
		}
	}

}

package main

import "fmt"

func main() {
	var (
		p *int
		v int = 19
	)
	fmt.Printf("&p: %v, p: %v\n", &p, p)

	var p2 = new(int)

	fmt.Printf("p2: %v\n", p2)
	p3 := &v
	fmt.Printf("p3: %p\n", p3)

	fmt.Printf("La dirección de memoria de v es: %p\n", &p3)
	fmt.Printf("El valor almacenado en %p es: %d\n", p3, *p3)

}

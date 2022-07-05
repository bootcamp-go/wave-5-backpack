package main

import "fmt"

func Incrementar(v *int) {
	*v = 50
}

func main() {
	var v int = 19
	fmt.Println("El valor de v antes vale:", v)
	Incrementar(&v)
	fmt.Println("El valor de v ahora vale:", v)
}

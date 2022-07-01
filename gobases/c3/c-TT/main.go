package main

import "fmt"

func Incrementar(v *int) {
	*v++
}

func main() {
	var v int = 19

	Incrementar(&v)

	fmt.Println("El valor de v ahora vale:", v)
}

// func Incrementar(v int) {
// 	v++

// }

// func main() {
// 	var v int = 19
// 	fmt.Println("El valor de v ahora vale:", v)

// 	Incrementar(v)

// 	fmt.Println("El valor de v ahora vale:", v)
// 	//fmt.Println("El valor de c ahora vale:", c)
// }

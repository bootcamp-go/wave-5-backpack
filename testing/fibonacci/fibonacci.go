package main

import "fmt"

func main() {
	fmt.Println(fibonacci(2))
}

func fibonacci(numero int) int {
	if numero == 0 {
		return 0
	}
	if numero == 1 || numero == 2 {
		return 1
	}

	return fibonacci(numero-1) + fibonacci(numero-2)
}

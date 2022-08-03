package main

import (
	"fmt"
	"log"

	"github.com/bootcamp-go/go-testing/fibonacci"
)

func main() {
	log.Println("TDD - Fibonacci")
	sucesion, total := fibonacci.Fibonnaci(5)

	fmt.Println("Sucesión ", sucesion)
	fmt.Println("Total ", total)
}

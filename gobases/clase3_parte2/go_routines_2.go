package main

import (
	"fmt"
	"time"
)

type Persona struct {
	Data string
}

func procesar(i int, c chan int) {
	fmt.Println(i, "- Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Termina")
	c <- i
}

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go procesar(i, c)
	}
	for i := 0; i < 10; i++ {
		a := <-c
		fmt.Println(a, "-Termina")
	}

	fmt.Println("Termino el programa")
}

package main

import (
	"fmt"
	"time"
)

func procesar(i int, c chan int) {
	fmt.Println(i, "-Inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-Termina")
	c <- i
}

func main() {
	c := make(chan int)
	go procesar(1, c)
	fmt.Println("Termino el programa")
}

package main

import (
	"fmt"
	"time"
)

func procesar(i int, c chan int) {
	fmt.Println(i, "-inicia")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, "-final")
	c <- i
}
func main() {
	// for i := 0; i < 10; i++ {
	// 	go procesar(i)
	// }
	//time.Sleep(5000*time.Millisecond)
	//fmt.Println("termina programa")
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go procesar(i, c)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("termino el program", <-c)
		// fmt.Println("hola")
	}

}

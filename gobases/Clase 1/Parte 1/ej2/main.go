package main

import (
	"fmt"
)

func main() {
	var temp int = 100
	var presion int = 100
	var humedad float32 = 0.32

	fmt.Printf("La temperatura registrada es: %d C \n", temp)
	fmt.Printf("La presion registrada es: %d bar \n", presion)
	fmt.Printf("La humedad registrada es: %d%%\n", int(humedad*100))
}

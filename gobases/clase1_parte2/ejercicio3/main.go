package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Digite el número del mes: ")
	var mes string
	fmt.Scanln(&mes)

	mesInt, _ := strconv.Atoi(mes)

	if mesInt < 13 {
		fmt.Println(time.Month(mesInt))
	} else {
		println("Número no válido!")
	}
}

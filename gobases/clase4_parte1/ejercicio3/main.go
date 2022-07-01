package main

import (
	"fmt"
	"time"
)

func calculoSueldo(s int) {
	if s < 150000 {
		err := fmt.Errorf("error: el minimo imponible es de 150.000 y el salario ingresado es de: %d %v", s, time.Now())
		fmt.Println(err)
		return
	} else if s >= 150000 {
		err := fmt.Errorf("error: El minimo imponible es de 150.000, Debe pagar impuestos %d, en un tiempo de: %v", s, time.Now())
		fmt.Println(err)
		return
	}
	fmt.Println("El programa finalizo OK")
}

func main() {
	var salary int = 140000
	calculoSueldo(salary)

}

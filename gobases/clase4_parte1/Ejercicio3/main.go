package main

import (
	"fmt"
	"time"
)

func main() {
	salary := 140000
	if salary < 150000 {
		err := fmt.Errorf("Error: El minimo imposible es de 150000 y el salario ingresado es de: %d,  El error se presento el: %v", salary, time.Now())
		fmt.Println("error ocurrido: ", err)
		return
	}
	fmt.Println("La persona debe pagar impuesto")
}

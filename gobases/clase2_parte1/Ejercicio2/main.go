package main

import (
	"errors"
	"fmt"
)

func main(){
	response, err := calcularPromedio(4,4,-8)

	if err != nil {
		fmt.Println("Hubo un error: ", err)
	}else{
		fmt.Println("El promedio de notas es: ", response)
	}
}

func calcularPromedio(notas ...int)(int, error){
	suma :=0
	count := 0

	for _, value := range notas {
		if value < 0 {
			return 0, errors.New("No se permiten numeros negativos")
		}
		suma += value
		count++
	}
	return suma/count, nil
}
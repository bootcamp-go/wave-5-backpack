package main

import (
	"fmt"
	"errors"
)

func main(){

	prueba:= calcularPromedio(10,10,10,10,9)
	fmt.Println(prueba)
}



func calcularPromedio(valores... int)int{
	var resultado int	
	for _, value := range valores {
		if value < 0 {
			errors.New("Error")
		}
		resultado += value
	}
	resultado = resultado / len(valores)
	return resultado
} 
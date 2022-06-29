package main

import (
	"fmt"
	"errors"
)

func calcularPromedio(valores ...float64)  (float64,error){
	suma := 0.0
	err := false

	for _, valor := range valores{

		if valor >= 0{
			suma = suma+valor
		} else {
			err = true
			break
		}
	}

	if err == true {

		return 0, errors.New ("No puede usar numeros negativos")

	} else {

		return suma/float64(len(valores)), nil
	}
	
	
}
func main() {
	promedio, err := calcularPromedio(5.0,5.0,5.0,4.0,-1.0,5.0)

	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("El promedio es: ", promedio)
	}

}
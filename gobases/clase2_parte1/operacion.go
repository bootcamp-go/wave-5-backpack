package main

import (
	"errors"
)

func main(){

}


const (
	min = "minimum"
	ave = "average"
	max = "maximum"
)

func operation(operacion string){
	switch operacion {
	case min:
		return minimum(valores)
	case ave:
		return average(valores)
	case max:
		return maximum(valores)
	}
	errors.New("Debes ingresar una operación valida")
} 

func minimum(valores ...float64)float64{
	minimum:= valores[0]
	for _, value := range valores {
		if value < minimum {
			minimum = value
		}
	}
	return minimum
}

func average(valores ...float64)float64{
	var average float64

	for _, value := range valores {
		average += value
	}
	average = average/len(valores)
	return average
}

func maximum(valores ...float64)float64{
	maximum:= valores[0]
	for _, value := range valores {
		if value > maximum {
			maximum = value
		}
	}
	return maximum
}



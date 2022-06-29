package main

import (
	"errors"
	"fmt"
)

func main() {

	resultado, error := operation("ss")
	if error != nil {
		fmt.Println("Se ingreso un operador inexistente")
	} else {
		finalR := resultado(3.0, 4.2, 1.0)
		fmt.Println(finalR)
	}

}

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func opMinimo(valores ...float32) float32 {
	var aux float32 = 200000000
	for _, valor := range valores {
		if aux >= valor {
			aux = valor
		}
	}
	return aux
}

func opMaximo(valores ...float32) float32 {
	var aux float32 = 0
	for _, valor := range valores {
		if aux <= valor {
			aux = valor
		}
	}
	return aux
}

func opPromedio(valores ...float32) float32 {
	var aux float32
	var total float32
	for _, valor := range valores {
		aux += 1
		total += valor
	}
	return (total / aux)
}

func operation(oper string) (func(valores ...float32) float32, error) {
	switch oper {
	case "minimum":
		return opMinimo, nil
	case "maximum":
		return opMaximo, nil
	case "average":
		return opPromedio, nil
	}
	return nil, errors.New("No existe la operacion")
}

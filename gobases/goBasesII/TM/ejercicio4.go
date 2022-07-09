package main

import (
	"errors"
	"fmt"
)

var (
	minimum string = "minimum"
	average string = "average"
	maximum string = "maximum"
)

func minFunc(values []int) (resultado float32) {
	for i, value := range values {
		if i == 0 {
			resultado = float32(value)
		} else if float32(value) < resultado {
			resultado = float32(value)
		}
	}
	return
}

func aveFunc(values []int) (resultado float32) {
	for i, value := range values {
		if i == 0 {
			resultado = float32(value)
		} else {
			resultado += float32(value)
		}
	}
	resultado /= float32(len(values))
	return
}

func maxFunc(values []int) (resultado float32) {
	for i, value := range values {
		if i == 0 {
			resultado = float32(value)
		} else if float32(value) > resultado {
			resultado = float32(value)
		}
	}
	return
}

func operadorMath(operacion string) (operacionFunc func(values []int) float32, err error) {
	switch operacion {
	case minimum:
		return minFunc, nil
	case average:
		return aveFunc, nil
	case maximum:
		return maxFunc, nil
	default:
		return nil, errors.New("No hay una operacion definida para resolver la solicitud")
	}
}
func operar(operador string, values ...int) (resultado float32, err error) {
	fun, err := operadorMath(operador)

	if err != nil {
		return 0, err
	} else {
		resultado = fun(values)
	}
	return
}

func main() {

	result, err := operar(minimum, 1, 2, 3, 4, 5, 6, 7)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Resultado: %.2f\n", result)
	}

}

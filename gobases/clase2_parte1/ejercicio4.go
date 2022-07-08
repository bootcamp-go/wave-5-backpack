package main

import "fmt"
import "errors"

const(
	MINIMUM = "minimum"
	MAXIMUM = "maximum"
	AVERAGE = "average"
)

func opMin(values ...int) float32 {
	var minimo int
	for key, value := range values {
		if key == 0 {
			minimo = value
		} else if value < minimo {
			minimo = value
		}
	}

	return float32(minimo)
}

func opMax(values ...int) float32 {
	var maximo int
	for key, value := range values {
		if key == 0 {
			maximo = value
		} else if value > maximo {
			maximo = value
		}
	}

	return float32(maximo)
}

func opAvg(values ...int) float32 {
	var suma int
	for _, value := range values {
		suma += value
	}

	division := float32(suma) / float32(len(values))
	return division
}

func operacionEstadistica(operacion string) (func(valores ...int) float32, error) {
	switch operacion {
		case MINIMUM:
			return opMin, nil
		case MAXIMUM:
			return opMax, nil
		case AVERAGE:
			return opAvg, nil
		default:
			return nil, errors.New("No se reconoce la operación")
	}

	return nil, nil
}

func main() {
	minFunc, errmin := operacionEstadistica(MINIMUM)
	maxFunc, errmax := operacionEstadistica(MAXIMUM)
	avgFunc, erravg := operacionEstadistica(AVERAGE)
	if errmin == nil {
		minValue := minFunc(3,4,5,7,6,4,2,6,6)
		fmt.Printf("Valor mínimo: %d\n", int(minValue))
	} else {
		fmt.Println(errmin)
	}

	if errmax == nil {
		maxValue := maxFunc(3,6,3,2,6,6,5,3,9,8,0)
		fmt.Printf("Valor máximo: %d\n", int(maxValue))
	} else {
		fmt.Println(errmax)
	}
	
	
	if erravg == nil {
		avgValue := avgFunc(4,7,6,3,9)
		fmt.Printf("Promedio: %f\n", avgValue)
	} else {
		fmt.Println(erravg)
	}
}
package main

import (
	"errors"
	"fmt"
)

const (
	MINIMUN = "minimun"
	MAXIMUN = "maximun"
	AVERAGE = "average"
)

func average(notas ...float32) float32 {
	total := float32(0)
	for _, nota := range notas {
		total += nota
	}
	return total / float32(len(notas))
}
func maximun(notas ...float32) float32 {
	max := float32(0)
	for i, nota := range notas {
		if i == 0 {
			max = nota
		} else if max < nota {
			max = nota
		}
	}
	return max
}
func minimun(notas ...float32) float32 {
	min := float32(0)
	for i, nota := range notas {
		if i == 0 {
			min = nota
		} else if min > nota {
			min = nota
		}
	}
	return min
}

func operation(oper string) (func(notas ...float32) float32, error) {
	switch oper {
	case MINIMUN:
		return minimun, nil
	case MAXIMUN:
		return maximun, nil
	case AVERAGE:
		return average, nil
	}

	return nil, errors.New("Error: Calculo no encontrado")
}

func execOperation(oper string, data []float32) {
	opFunc, opError := operation(oper)
	if opError != nil {
		fmt.Println(opError)
	} else {
		fmt.Println("Resultado para operacion", oper, opFunc(data...))
	}
}

func main() {
	notas := []float32{3, 2.4, 1.5, 7, 6.5, 3.4}
	execOperation(MINIMUN, notas)
	execOperation(MAXIMUN, notas)
	execOperation(AVERAGE, notas)

}

package main

import (
	"errors"
	"fmt"
)

func main() {
	promedio, err := calcularPromedio(5, 3, 2)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("El promedio es: ", promedio)

}

func calcularPromedio(notas ...float32) (float32, error) {
	var cantidadNotas int = len(notas)
	var acum float32 = 0.0
	var err error = nil

	for _, nota := range notas {
		if nota == 0 {
			err = errors.New("Alguno de los datos ingresados no es válido")
			break
		}
		acum += nota
	}

	return acum / float32(cantidadNotas), err
}

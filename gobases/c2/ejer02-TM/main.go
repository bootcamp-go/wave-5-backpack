package main

import (
	"errors"
	"fmt"
)

// Un colegio necesita calcular el promedio (por alumno) de sus calificaciones.
// Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros
// y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo

func promedioNotas(notas ...int) (float32, error) {
	var promedio float32
	total := 0

	for _, nota := range notas {
		if nota < 1 {
			return promedio, errors.New("No puede ingresar notas negativas ni cero")
		} else {
			total += nota

		}
	}

	promedio = float32(total) / float32(len(notas))
	return promedio, nil

}

func main() {

	fmt.Println(promedioNotas(1, 2, 3, 4))

	fmt.Println(promedioNotas(0, 2, 3, 4, 5))

	fmt.Println(promedioNotas(-1, 2, 3, 4, 5))
}

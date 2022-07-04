package main

import "fmt"

func promedioCalificaciones(calificaciones ...int) float64 {

	var sumaCalificaciones float64
	numeroCalificaciones := len(calificaciones)

	for _, calificacion := range calificaciones {
		sumaCalificaciones += float64(calificacion)
	}

	promedio := sumaCalificaciones / float64(numeroCalificaciones)

	return promedio

}

func main() {

	fmt.Printf("El promedio de calificaicones es: %v \n", promedioCalificaciones(50, 40, 60, 70, 70, 60))
}

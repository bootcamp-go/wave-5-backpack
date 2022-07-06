package main
import "fmt"
import "errors"

func main() {
	fmt.Printf("El promedio del alumno es: %v \n", calcPromedio(10, 8, 9, 10, 8, 8, 9, 10, 7, 10))
}

func calcPromedio(valores ...int) int{
	var total int = 0
	var unidades int = 0
	for _, calificacion := range valores {
		total += calificacion
		unidades++
	}
	return total / unidades
}

//Un colegio necesita calcular el promedio (por alumno) de sus calificaciones.
//Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio
//y un error en caso que uno de los números ingresados sea negativo
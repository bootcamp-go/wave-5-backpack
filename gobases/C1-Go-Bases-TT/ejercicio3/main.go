package main

import "fmt"

// Ejercicio 3 - A qué mes corresponde

// Realizar una aplicación que contenga una variable con el número del mes.
// Según el número, imprimir el mes que corresponda en texto.
// /¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?

// Con switch porque los meses son definidos

func main() {
	var mes int = 8
	var mesStr string = ""

	switch mes {
	case 1:
		mesStr = "Enero"
	case 2:
		mesStr = "Febrero"
	case 3:
		mesStr = "Marzo"
	case 4:
		mesStr = "Abril"
	case 5:
		mesStr = "Mayo"
	case 6:
		mesStr = "Junio"
	case 7:
		mesStr = "Julio"
	case 8:
		mesStr = "Agosto"
	case 9:
		mesStr = "Septiembre"
	case 10:
		mesStr = "Octubre"
	case 11:
		mesStr = "Noviembre"
	case 12:
		mesStr = "Diciembre"
	default:
		mesStr = "ohh"
	}

	fmt.Printf("%s\n", mesStr)
}

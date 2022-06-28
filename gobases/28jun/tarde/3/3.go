// Ejercicio 3 - A qué mes corresponde

// Realizar una aplicación que contenga una variable con el número del mes.
// Según el número, imprimir el mes que corresponda en texto.
// // ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?
// // Ej: 7, Julio
package main

import "fmt"

var (
	mes int = 5
)

func main() {
	meses := []string{"Enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "sep", "oct", "nov", "dic"}
	if mes >= 1 && mes <= 12 {
		fmt.Printf("el mes es %v\n", meses[mes-1])
	} else {
		fmt.Printf("error")
	}
}

package main

import "fmt"

/*Ejercicio 1 - Imprimí tu nombre

Crea una aplicación donde tengas como variable tu nombre y dirección.
Imprime en consola el valor de cada una de las variables.
*/

func main() {

	var (
		nombre    = "Juan Gabriel"
		direccion = "Independencia 1232"
	)

	// Ejemplos para representar distinto tipos de datos en un string:
	// %s string
	// %d interger
	// %f float

	fmt.Printf("✅ Mi nombre es %s y vivo en %s\n", nombre, direccion) // \n es solamente un salto de linea cuando imprimimos
	fmt.Printf("✅ Mi nombre es %v y vivo en %v\n", nombre, direccion)
}

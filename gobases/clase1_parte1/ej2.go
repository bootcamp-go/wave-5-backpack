// Ejercicio 2 - Clima

// Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura y humedad y presión atmosférica de distintos lugares.
// Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres.
// Imprime los valores de las variables en consola.
// ¿Qué tipo de dato le asignarías a las variables?

package main

import "fmt"

func main() {
	var temperatura int = 36
	var humedad int = 80
	var presion int = 40

	fmt.Printf("La temperatura es %d, la humedad es: %d %% y la presion es: %d \n", temperatura, humedad, presion)
}

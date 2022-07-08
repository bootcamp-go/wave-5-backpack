package main

import "fmt"

/* Ejercicio 1 - Clima

Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura y humedad y
presión atmosférica de distintos lugares.
	1. Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura,
	humedad y presión de donde te encuentres.
	2. Imprime los valores de las variables en consola.
	3. ¿Qué tipo de dato le asignarías a las variables? */

func main() {
	var temperatura, humedad, pression float32 = 16, 80, 1000.2

	fmt.Println("Reporte del clima: 🌥️")
	fmt.Printf("Temperatura: %f C°\n", temperatura)
	fmt.Printf("Humedad: %.0f %%\n", humedad)                // %% para escapar el porcentaje
	fmt.Printf("Presión Atmosferica: %.2f hPa°\n", pression) //o hectopascal (hPa).
}

/*
REFERENCIAS:

La presión atmosférica es el peso del aire sobre la superficie de la Tierra. La capa de aire
que envuelve la Tierra es la atmósfera. Esta capa ejerce un peso sobre la superficie
terrestre: es esto lo que llamamos presión atmosférica.
*/

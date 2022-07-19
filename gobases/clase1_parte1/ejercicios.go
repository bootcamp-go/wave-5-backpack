package main

import "fmt"

func main() {
	//Ejercicio 1
	ejercicio1()
	//Ejercicio 2
	ejercicio2()
	//Ejercicio 3
}

func ejercicio1() {
	fmt.Println("******************** INICIO EJERCICIO 1 ********************")
	var (
		nombre    = "Michael Torres"
		direccion = "Cra 13C #14B - 04"
	)
	fmt.Println("Mi nombre es: " + nombre + " y vivo en la direccion: " + direccion)
	fmt.Println("******************** FIN EJERCICIO 1 ********************\n\n\n")
}
func ejercicio2() {
	fmt.Println("******************** INICIO EJERCICIO 2 ********************")
	var temperatura float32 = 29
	var humedad float32 = 0.7
	var presion float32 = 30.3
	fmt.Println("La temperatura en Bogota es: ", temperatura, ", la humedad es: ", humedad, " y su presion es: ", presion)
	fmt.Println("******************** FIN EJERCICIO 2 ********************\n\n\n")
}
func ejercicio3() {
	//var 1nombre string - incorrecta - NO se puede empezar variables por numero
	//var nombre1 string
	//var apellido string //Correcta
	//var int edad        //Correcta
	//1apellido := 6 - incorrecta - NO se puede declarar una variable sin empezar por numero, sin embargo esta bien declarada
	//var licencia_de_conducir = true
	//var estatura de la persona int - incorrecta - No se puede declarar varaibles con espacios
	//cantidadDeHijos := 2
}

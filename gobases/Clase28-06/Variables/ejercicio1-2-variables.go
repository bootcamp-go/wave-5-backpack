package main

import (
	"fmt"

	"github.com/ddadumitrescu/hellomod"
)

var (
	name   = "Naim Rossetti"
	addres = "Angel J Carranza 1234"
)

func main() {
	fmt.Println("EJERCICIO 1:\n ")
	fmt.Printf("Mi nombre es: %s y mi direccion es: %s\n ", name, addres)

	fmt.Println("\n EJERCICIO 2:\n ")
	var temperatura int = 14
	var humedad int = 55
	var presion float64 = 1015.3

	fmt.Println("DATOS CLIMATICOS:\n ")
	fmt.Printf("Temperatura: %d \n Humedad: %d\n Presion: %v", temperatura, humedad, presion)

	hellomod.SayHello()

}

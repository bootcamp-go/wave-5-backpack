package main

import "fmt"

/*Ejercicio 3 - Declaración de variables

Un profesor de programación está corrigiendo los exámenes de sus estudiantes de la materia Programación I para poder brindarles
las correspondientes devoluciones. Uno de los puntos del examen consiste en declarar distintas variables.
Necesita ayuda para:
Detectar cuáles de estas variables que declaró el alumno son correctas.
Corregir las incorrectas:
	var 1nombre string
	var apellido string
	var int edad
	1apellido := 6
	var licencia_de_conducir = true
	var estatura de la persona int
	cantidadDeHijos := 2
*/

func main() {
	// No se pueden crear variables si empiezan por número
	// var 1nombre string
	var nombre1 string

	var apellido string

	// Se necesita declarar el nombre de la variable y después el tipo de dato
	//var int edad
	var edad int

	// No se pueden crear variables si empiezan por número
	// 1apellido := 6
	apellido1 := 6
	var licencia_de_conducir = true

	// No podemos definir una variable con espacios
	// var estatura de la persona int
	var estaturaDeLaPersona int
	cantidadDeHijos := 2

	fmt.Println(nombre1, apellido, edad, apellido1, licencia_de_conducir, estaturaDeLaPersona, cantidadDeHijos)
}

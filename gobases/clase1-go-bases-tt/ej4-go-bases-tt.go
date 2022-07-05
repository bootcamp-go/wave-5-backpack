/*---------------------------------------------------------------------------------*

     Assignment:	Ejercicio #4:  Que edad tiene
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Bases

	Description:
		An employee of a company wants to know the name and age of one
		of his employees. According to the map below, he helps print
		Benjamin's age.

		var  employees  =  map  [  string  ]  int  {  "Benjamin"  :  20,
		"Nahuel"  :  26,  "Brenda"  :  19,  "Darío"  :  44,  "Pedro"  :  30  }

		On the other hand it is also necessary:
			• Know how many of your employees are over the age of 21.
			• Add a new employee to the list, named Federico,
			  who is 25 years old.
			• Eliminate Pedro from the map.

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

//	PACKAGE & LIBRARY
package main

import "fmt"

// VARIABLE : employees
var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19,
	"Darío": 44, "Pedro": 30}

//	MAIN PROGRAM
func main() {
	fmt.Printf("\n|| Que edad tiene ||\n LISTA => ")
	fmt.Println(employees)

	// Saber cuántos de sus empleados son mayores de 21 años.
	fmt.Printf("\n > Empleados mayores a 21: \n")
	for name, age := range employees {
		if age >= 21 {
			fmt.Println("	Name: ", name, " Age:", age)
		}
	}

	// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	fmt.Printf("\n > Agregando al empleado 'Federico': \n")
	employees["Federico"] = 25
	fmt.Println(employees)

	// Eliminar a Pedro del mapa.
	fmt.Printf("\n > Eliminando al empleado 'Pedro': \n")
	delete(employees, "Pedro")
	fmt.Println(employees)
}

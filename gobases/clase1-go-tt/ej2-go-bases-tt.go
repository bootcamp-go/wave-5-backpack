/*Ejercicio 2 - Préstamo*/
// Debe imprimir 3 mjs

package main

import "fmt"

type Data struct {
	Name     string
	Edad     int
	Empleado bool
	Tiempo   int
}

func main() {
	employee := Data{"Samuel", 25, true, 2}
	fmt.Printf("\n|| Préstamo ||")
	fmt.Printf("\nPersona: ", employee.Name)
	if employee.Edad > 22 {
		fmt.Printf("\n>> Posee mayor a 22 años :: Tiene %d años\n", employee.Edad)
	} else {
		fmt.Printf("\n>> No posse mayor a 22 años ")
	}
	if employee.Empleado == true {
		fmt.Printf(">> Si es empleado\n")
	} else {
		fmt.Printf("\n>> No es empleado ")
	}
	if employee.Tiempo > 1 {
		fmt.Printf(">> Posee mayor a 1 año :: Tiene %d (años)\n", employee.Tiempo)
	} else {
		fmt.Printf("\n>> Tiene menor a un año ")
	}
}

package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {
	//Empleado mayor a 21

	contador := 0
	for _, edad := range employees {
		if edad > 21 {
			contador++
		}
	}
	fmt.Println("La cantidad de empleados mayores de 21 años es: ", contador, "\n")

	//Edad Benjamin
	fmt.Println("La edad de Benjamin es: ", (employees["Benjamin"]), "\n")
	// Agregar Info
	employees["Federico"] = 25
	fmt.Println("Se agrego a Federico ==> ", employees, "\n")
	//Eliminar Info
	delete(employees, "Pedro")
	fmt.Println("Se Elimino a Pedro ==> ", employees, "\n")
}

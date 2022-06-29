package main

import (
	"fmt"
)

func letrasDeUnaPalabra() {
	//ejercicio 1
	palabra := "Hola"

	fmt.Printf("Cantidad de letras: %d \n", len(palabra))

	for i := range palabra {
		fmt.Printf("%c ", palabra[i])
	}

	fmt.Println()
}

func prestamo() {
	//ejercicio 2

	edad := 23
	empleado := true
	antiguedad := 3
	sueldo := 200000

	if edad > 22 && empleado && antiguedad > 1 {
		if sueldo > 100000 {
			fmt.Println("Prestamo aprobado! Le informamos tambien que el mismo no tiene intereses.")
		} else {
			fmt.Println("Prestamo aprobado!")
		}
	} else {
		fmt.Println("Lo sentimos, no cumple los requisitos para acceder al prestamo")
	}

}

func aQueMesCorresponde() {
	//ejercicio 3
	num := 5

	switch num {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 6:
		fmt.Println("Junio")
	case 7:
		fmt.Println("Julio")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Septiembre")
	case 10:
		fmt.Println("Octubre")
	case 11:
		fmt.Println("Noviembre")
	case 12:
		fmt.Println("Diciembre")
	default:
		fmt.Println("Opcion invalida, ingrese un numero del 1 al 12")
	}
}

func queEdadTiene() {
	//ejercicio 4
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println("La edad de Benjamin es:", employees["Benjamin"])

	count := 0

	for key := range employees {
		if employees[key] > 21 {
			count++
		}
	}

	fmt.Println("Empleados mayores de 21 años: ", count)

	employees["Federico"] = 25

	delete(employees, "Pedro")

}

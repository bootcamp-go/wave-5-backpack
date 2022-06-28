package main

import "fmt"

func main() {

	const EDAD_MINIMA = 22
	const ANTIGUEDAD_MINIMA = 1
	const SUELDO_NO_INTERES = 100000

	edad := 5
	empleado := true
	antiguedad := 2
	sueldo := 500000

	if edad < EDAD_MINIMA {
		fmt.Println("No se otorga crédito por ser menor de la edad establecida")
		return
	}

	if antiguedad < ANTIGUEDAD_MINIMA {
		fmt.Println("No se otorga crédito por no tener antigüedad suficiente")
		return
	}

	if !empleado {
		fmt.Println("No se otorga crédito por no estar empleado")
		return
	}

	if sueldo < SUELDO_NO_INTERES {
		fmt.Println("Se otorga crédito con intereses")
		return
	}

	fmt.Println("Se otorga crédito sin intereses")

	//Nota: en vez de usar if-else anidados, se utiliza este patrón llamado 'early return' para facilitar la lectura del código
}

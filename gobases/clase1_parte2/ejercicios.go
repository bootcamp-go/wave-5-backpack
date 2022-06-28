package main

import "fmt"

func main() {

	//EJERCICIO 1
	var word = "Un balon"
	fmt.Println("Tu palabra tiene: ", len(word))
	for _, letter := range word {
		fmt.Println("Tu apalabra elegida es: ", string(letter))
	}

	//EJERCICIO 2
	var edad int = 25
	var antiguedad int = 2
	var empleado bool = true
	var sueldo float32 = 100500

	if edad > 22 && empleado && antiguedad > 1 {
		fmt.Println("Cumple con los requisitos basicos para aplicar al prestamos")
		if sueldo > 100000 {
			fmt.Println("Usted esta excento de interes, por tu alto sueldo.")
		} else {
			fmt.Println("Usted tendra que pagar una tasa de interes, por su bajo salario.")
		}
	} else {
		fmt.Println("Su prestamo ha sido rechazado, no cumple con los requsitos basicos")
	}

	//EJERCICIO 3
	var meses = map[int]string{1: "ENERO", 2: "FEBRERO", 3: "MARZO", 4: "ABRIL", 5: "MAYO", 6: "JUNIO", 7: "JULIO", 8: "AGOSTO", 9: "SEPTIEMBRE", 10: "OCTUBRE", 11: "NOVIEMBRE", 12: "DICIEMBRE"}
	var indice int = 2
	fmt.Println(meses[indice])

	//EJERCICIO 4
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var counter = 0
	fmt.Println(employees["Benjamin"])

	for _, edad := range employees {
		if edad > 21 {
			counter++
		}
	}
	fmt.Println("La cantidad de empleados que son mayores de 21", counter)

	employees["Federico"] = 25

	fmt.Println(employees)

	delete(employees, "Pedro")

	fmt.Println(employees)
}

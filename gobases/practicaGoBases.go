package main

import "fmt"

type Employee struct {
	cliente    bool
	edad       int
	antiguedad int
	sueldo     float32
}

func main() {

	//ejercicio 1

	//Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
	//Luego imprimí cada una de las letras.

	var palabra string = "hola"

	fmt.Printf("ejercicio01\ntamaño: %d\n", len(palabra))

	for _, letra := range palabra {
		fmt.Printf("%s\n", string(letra))
	}

	//ejercicio 2

	//Otorgar prestamos a sus clientes , solo le otorga a mayores de 22 , empleados y mas de un año de antiguedad.
	//No se cobra interes a sueldos > 100.000

	fmt.Println("ingrese la edad del Cliente:")
	var edad int
	fmt.Scanln(&edad)
	fmt.Println("ingrese la antiguedad del Cliente:")
	var antiguedad int
	fmt.Scanln(&antiguedad)
	fmt.Println("ingrese la sueldo del Cliente:")
	var sueldo float32
	fmt.Scanln(&sueldo)

	e1 := Employee{true, edad, antiguedad, sueldo}

	if e1.cliente != true {
		fmt.Println("Tenes que ser cliente para pedir un prestamo")
	} else if e1.edad < 22 {
		fmt.Println("No tenes edad para hacer un prestamo")
	} else if e1.antiguedad < 1 {
		fmt.Println("No tenes antiguedad para hacer un prestamo")
	} else if e1.sueldo > 100000 {
		fmt.Println("No se te cobra interes y podes acceder al prestamo")
	} else {
		fmt.Println("Toma tu dinero")
	}

	//ejercicio 3

	fmt.Println("Ingrese el numero del mes")
	var mes int
	fmt.Scanln(&mes)
	switch mes {
	case 1:
		fmt.Printf("Enero\n")
		break
	case 2:
		fmt.Printf("Febrero\n")
		break
	case 3:
		fmt.Printf("Marzo\n")
		break
	case 4:
		fmt.Printf("Abril\n")
		break
	case 5:
		fmt.Printf("Mayo\n")
		break
	case 6:
		fmt.Printf("Junio\n")
		break
	case 7:
		fmt.Printf("Julio\n")
		break
	case 8:
		fmt.Printf("Agosto\n")
		break
	case 9:
		fmt.Printf("Septiembre\n")
		break
	case 10:
		fmt.Printf("Octubre\n")
		break
	case 11:
		fmt.Printf("Noviembre\n")
		break
	case 12:
		fmt.Printf("Diciembre\n")
		break
	}

	//Ejercicio 4

	var empleados = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var count int = 0
	for key, element := range empleados {
		if element > 21 {
			count++
		}

		if key == "Benjamin" {
			fmt.Println("Edad de Benjamin:", element)
		}
	}

	fmt.Println("Empleados mayores a 21:", count)

	empleados["Facundo"] = 25
	delete(empleados, "Pedro")

}

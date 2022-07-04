package main

import "fmt"

const (
	edadMin       int     = 22
	empleado   bool    = true
	antiguedadMin int     = 1
	sueldoMin float32 = 100000.00
)

var (
	edad int = 23
	empleo bool = true
	antiguedad int = 3
	sueldo float32 = 70000.00
)

func check(edad int, empleo bool, antiguedad int, sueldo float32){
	if edad > edadMin {
		fmt.Println("Cumple con la edad mínima")
	} else {
		fmt.Println("No cumple con la edad mínima")
	}

	if empleo == empleado {
		fmt.Println("Cumple con la empleabilidad")
	} else {
		fmt.Println("No cumple con la empleabilidad")
	}

	if antiguedad > antiguedadMin {
		fmt.Println("Cumple con la antigüedad")
	} else {
		fmt.Println("No cumple con la antigüedad")
	}

	if sueldo > sueldoMin {
		fmt.Println("Otorga crédito sin interés")
	} else {
		fmt.Println("Otorga crédito con interés")
	}
}

func checkV2(edad int, empleo bool, antiguedad int, sueldo float32){
	switch {
	case edad <= edadMin:
		fmt.Println("Debes ser mayor de 22 años de edad.")
	case empleo != empleado:
		fmt.Println("Debes estar empleado.")
	case antiguedad < antiguedadMin:
		fmt.Println("Debes tener una antiguedad mayor a 1 año")
	case sueldo > sueldoMin:
		fmt.Println("Otorgar credito sin interés ✅")
	default:
		fmt.Println("Otorgar credito con interés ✔️")
	}
}

func main(){
	//check(edad, empleo, antiguedad, sueldo)
	checkV2(edad, empleo, antiguedad, sueldo)
}
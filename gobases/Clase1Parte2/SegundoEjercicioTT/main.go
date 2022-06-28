package main

import "fmt"

func main() {
	clienteEdad := 23
	clienteEmpleo := true
	clienteAntiguedad := 2
	var cumpleEdad bool
	var cumpleEmpleo bool
	var cumpleAntiguedad bool
	if clienteEdad < 22 {
		fmt.Printf("Debe ser mayor de 22 años \n")
		cumpleEdad = false
	} else {
		fmt.Printf("Cumple con la edad\n")
		cumpleEdad = true
	}
	if clienteEmpleo == false {
		fmt.Printf("Debe tener empleo \n")
		cumpleEmpleo = false
	} else {
		fmt.Printf("Cumple con la condicion de estar empleado\n")
		cumpleEmpleo = true
	}
	if clienteAntiguedad < 1 {
		fmt.Printf("Debe tener un año o más de antiguedad \n")
		cumpleAntiguedad = false
	} else {
		fmt.Printf("Cumple con la antiguedad \n")
		cumpleAntiguedad = true
	}
	if cumpleAntiguedad && cumpleEdad && cumpleEmpleo {
		fmt.Printf("Credito Otorgado\n")
	}
}

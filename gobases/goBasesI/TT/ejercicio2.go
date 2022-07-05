package main

import "fmt"

const (
	edad       int     = 22
	empleado   bool    = true
	antiguedad int     = 12
	sueldo     float64 = 100000
)

func main() {
	var (
		edadE       int
		empleadoE   bool
		antiguedadE int
		suedoE      float64
		prestamo    bool
	)
	fmt.Print("Ingrese la edad en anos: ")
	fmt.Scanln(&edadE)
	fmt.Print("Es empleado: ")
	fmt.Scanln(&empleadoE)
	fmt.Print("Ingrese la antiguedad en meses: ")
	fmt.Scanln(&antiguedadE)
	fmt.Print("Ingrese sueldo:")
	fmt.Scanln(&suedoE)
	if edadE > edad && (empleadoE == empleado) && antiguedadE > antiguedad {
		prestamo = true
		fmt.Print("Es posible otorgar prestamo")
	} else {
		fmt.Print("No es posible otorgar un prestamo")
	}
	if prestamo && suedoE > sueldo {
		fmt.Print(" y esta excento de intereses")
	}
	fmt.Print(".\n")
}

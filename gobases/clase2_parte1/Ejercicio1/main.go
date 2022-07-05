package main

import "fmt"

const (
	Mult = "*"
)

func impuestoSalario(sueldo1, porcentaje float64, operador string) float64 {
	switch operador {
	case Mult:
		return sueldo1 * porcentaje
	}
	return 0
}

func main() {
	fmt.Println("Sueldo total empleado 1: ", (impuestoSalario(50000.00, 0.83, Mult)))
	r2 := impuestoSalario(150000.00, 0.83, Mult)
	fmt.Println("Sueldo total empleado 2: ", (impuestoSalario(r2, 0.90, Mult)))

}

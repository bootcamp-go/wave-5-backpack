package main

import "fmt"

func main() {
	var sueldo int = 23
	impuestos(sueldo)
}

func impuestos(sueldo int) {
	var descuento float32 = float32(sueldo) * (0.17)
	if sueldo < 50000 {
		fmt.Println("No se le descuenta nada")
	} else {
		if sueldo > 50000 && sueldo < 150000 {
			descuento = float32(sueldo) * (0.17)
			fmt.Println("El impuesto correspondiente es: ", descuento)
		}
		if sueldo > 150000 {
			descuento = float32(sueldo) * (0.27)
			fmt.Println("Por ganar mas de $150000 se le descuenta un total de: ", descuento)
		}
	}

}

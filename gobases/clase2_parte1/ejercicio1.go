package main

import "fmt"


func main()  {
	var impuesto1, impuesto2, sueldo_empleado float64

	impuesto1 = 0.17
	impuesto2 = 0.27

	sueldo_empleado = 200000
	resultado := calculo_impuestos(impuesto1, impuesto2, sueldo_empleado)
	fmt.Printf("Este es el impuesto que se le descontara al empleado: %.2f\n", resultado)
}


func calculo_impuestos(impuesto1 float64, impuesto2 float64, sueldo float64) float64  {
	if sueldo > 50000 && sueldo < 150000 {
		return sueldo * impuesto1
	} else if sueldo > 150000 {
		return sueldo * impuesto2
	}
	return 0
}

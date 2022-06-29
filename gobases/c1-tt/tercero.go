package main

import "fmt"

func main() {
	var numberMes int = 1
	var meses [12]string
	meses[0] = "Enero"
	meses[1] = "Febrero"
	meses[2] = "Marzo"
	meses[3] = "Abril"
	meses[4] = "Mayo"
	meses[5] = "Junio"
	meses[6] = "Julio"
	meses[7] = "Agosto"
	meses[8] = "Septiembre"
	meses[9] = "Octubre"
	meses[10] = "Noviembre"
	meses[11] = "Diciembre"

	fmt.Printf("El mes indicado es: %s\n", meses[numberMes-1])
}

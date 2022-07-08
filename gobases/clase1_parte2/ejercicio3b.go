package main

import "fmt"

func main() {
	meses := [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	mes := 2

	fmt.Printf("El mes es %s\n", meses[mes - 1])
}
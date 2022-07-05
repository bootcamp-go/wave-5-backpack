package main

import "fmt"

var meses = [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

func main() {
	var mes int
	fmt.Print("Ingrese un numero: ")
	fmt.Scanln(&mes)
	if mes >= 12 {
		fmt.Println("No hay un mes para el valor ingresado")
	} else {
		fmt.Printf("%d, %s\n", mes, meses[mes-1])
	}
}

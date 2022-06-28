package main

import "fmt"

func main() {
	numero_mes := 6
	getMes(numero_mes)
	numero_mes = 13
	getMes(numero_mes)

}

func getMes(numero_mes int) {
	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	if numero_mes > 0 && numero_mes <= 12 {
		fmt.Printf("El mes %d corresponde a %s \n", numero_mes, meses[numero_mes-1])
	} else {
		fmt.Println("El numero del mes debe estar entre 1 y 12!")
	}
}

package main

import "fmt"

func imprimirMes(mes int, meses [12]string) {
	if mes > 12 {
		fmt.Println("No existe ese mes")
		return
	}

	fmt.Printf("Mes: %s\n", meses[mes])
}

func main() {
	meses := [12]string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"}

	imprimirMes(7, meses)

	imprimirMes(14, meses)
}

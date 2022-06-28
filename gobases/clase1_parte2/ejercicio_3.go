package main

import "fmt"

func main() {
	mes := 1

	switch mes { // Opción 1: switch (muy largo)
	case 1:
		fmt.Println("Enero")
	}

	if mes == 1 { // Opción 2: if (muy largo)
		fmt.Println("Enero")
	}

	// Opción 3: map (más corto)
	var meses = map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio", 8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}
	fmt.Println(meses[8])

}

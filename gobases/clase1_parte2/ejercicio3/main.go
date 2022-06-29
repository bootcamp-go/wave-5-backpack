package main

import "fmt"

func main() {

	mes := map[int]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo", 6: "Junio", 7: "Julio",
		8: "Agosto", 9: "Septiembre", 10: "Octubre", 11: "Noviembre", 12: "Diciembre"}

	numero := 5

	fmt.Println("mes seleccionado: ", mes[numero])
}

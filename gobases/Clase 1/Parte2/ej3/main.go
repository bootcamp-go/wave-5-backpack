package main

import "fmt"

func main() {
	mes := 12
	listaMeses := make([]string, 12)
	listaMeses = []string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"}

	fmt.Printf("el mes %d es %s\n", mes, listaMeses[mes-1])
}

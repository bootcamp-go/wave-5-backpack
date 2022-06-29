package main

import "fmt"

func main() {
	var month int = 1
	var month_name = [12]string{
		"Enero",
		"Febrero",
		"Marzo",
		"Abril",
		"Mayo",
		"Junio",
		"Julio",
		"Agosto",
		"Septiembre",
		"Octubre",
		"Noviembre",
		"Diciembre",
	}
	
	fmt.Printf("%d, %s\n", month, month_name[month - 1])
}

// Tambien se podria usar un slice o un map[int]string
// Elegí array porque en este caso particular, se que solo hay 12 elementos fijos(uno por cada mes) y se puede vincular facilmente cada indice con un mes,
// no hay necesidad de que el tamaño sea variable como en un slice, 
// un map[int]string aumentaria la complejidad de la estructura innecesariamente.
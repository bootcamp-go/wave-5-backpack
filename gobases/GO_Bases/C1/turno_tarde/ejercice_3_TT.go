package main

import "fmt"

func main() {
	mounth := 2

	switch mounth {
	case 1:
		fmt.Println("MES : Enero")
	case 2:
		fmt.Println("MES : Febrero")
	case 3:
		fmt.Println("MES : Marzo")
	case 4:
		fmt.Println("MES : Abril")
	case 5:
		fmt.Println("MES : Mayo")
	case 6:
		fmt.Println("MES : Junio")
	case 7:
		fmt.Println("MES : Julio")
	case 8:
		fmt.Println("MES : Agosto")
	case 9:
		fmt.Println("MES : Septiembre")
	case 10:
		fmt.Println("MES : Octubre")
	case 11:
		fmt.Println("MES : Noviembre")
	case 12:
		fmt.Println("MES : Diciembre")
	}

	//Se podria utilizar una estructua map, agregando como key el numero del mes y como valor el
	//nombre del mismo
}

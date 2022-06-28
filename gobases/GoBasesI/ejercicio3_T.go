package main

import "fmt"

func main() {
	var nmes int = 10

	switch {
	case nmes == 1:
		fmt.Println("Enero")
	case nmes == 2:
		fmt.Println("Febrero")
	case nmes == 3:
		fmt.Println("Marzo")
	case nmes == 4:
		fmt.Println("Abril")
	case nmes == 5:
		fmt.Println("Mayo")
	case nmes == 6:
		fmt.Println("Junio")
	case nmes == 7:
		fmt.Println("Julio")
	case nmes == 8:
		fmt.Println("Agosto")
	case nmes == 9:
		fmt.Println("Septiembre")
	case nmes == 10:
		fmt.Println("Octubre")
	case nmes == 11:
		fmt.Println("Noviembre")
	case nmes == 12:
		fmt.Println("Diciembre")
	default:
		fmt.Println("No representa un numero de mes")
	}

}

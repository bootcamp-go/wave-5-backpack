package main

import "fmt"

func main() {
	mes := 6

	switch mes {
	case 1:
		fmt.Printf("Enero \n")
	case 2:
		fmt.Printf("Febrero \n")
	case 3:
		fmt.Printf("Marzo \n")
	case 4:
		fmt.Printf("Abril \n")
	case 5:
		fmt.Printf("Mayo \n")
	case 6:
		fmt.Printf("Junio \n")
	case 7:
		fmt.Printf("Julio \n")
	case 8:
		fmt.Printf("Agosto \n")
	case 9:
		fmt.Printf("Septiembre \n")
	case 10:
		fmt.Printf("Octubre \n")
	case 11:
		fmt.Printf("Noviembre \n")
	case 12:
		fmt.Printf("Diciembre \n")
	default:
		fmt.Printf("No es un mes\n")
	}
}

//Esto podría resolverse con if - else, pero sería tedioso y repetitivo y no se estaría aprovechando una estructura que nos facilita y hace mas pulcro el código

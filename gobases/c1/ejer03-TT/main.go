package main

import "fmt"

// Realizar una aplicación que contenga una variable con el número del mes.
// Según el número, imprimir el mes que corresponda en texto.
// ¿Se te ocurre si se puede resolver de más de una manera? ¿Cuál elegirías y por qué?

func main() {

	var mes int

	mes = 1

	switch mes {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 6:
		fmt.Println("Junio")
	case 7:
		fmt.Println("Julio")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Septiembre")
	case 10:
		fmt.Println("Octubre")
	case 11:
		fmt.Println("Noviembre")
	case 12:
		fmt.Println("Diciembre")
	default:
		fmt.Println("yo soy un default")
	}

	var meses = [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	fmt.Println(meses[mes-1])
}

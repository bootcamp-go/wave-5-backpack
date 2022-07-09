package main

import "fmt"

func main() {

	var MONTH int = 3      //Marzo 1-12
	var LASTMONTH int = 12 //Diciembre 1-12

	getMonth(MONTH)           //Forma 1, utilizando Arrays
	getMonthSwitch(LASTMONTH) //Forma 2, utilizando Switchs

}

func getMonth(m int) {
	calendar := [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

	fmt.Println(calendar[m-1])
}

func getMonthSwitch(m int) {
	switch m {
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
	}

}

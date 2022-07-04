package main

import "fmt"

var meses = [12] string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

func main(){
	month := 1
	if month >=1 && month <=12 {
		fmt.Println(meses[month-1])
	} else {		
		fmt.Println("Inserte un numero valido")
	}

}
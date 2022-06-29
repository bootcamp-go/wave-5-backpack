package main

import "fmt"

func main(){

	var number = 3
	// first option declare a slice , and print with negative offset (-1)
	var monthy =[12]string {"Enero",
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
							"Diciembre"} 
	fmt.Println(monthy[number-1])

	// second option , create a map , an associate the number of month like key and name month like value
	var months = map[int] string{	
							1: "Enero",
							2: "Febrero",	
							3: "Marzo", 
							4: "Abril", 
							5: "Mayo", 
							6: "Junio", 
							7: "Julio", 
							8: "Agosto", 
							9: "Septiembre", 
							10: "Octubre", 
							11: "Noviembre", 
							12: "Diciembre"}
	fmt.Println(months[number])
}
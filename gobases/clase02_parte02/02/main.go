package main

import "fmt"

var edad int = 20
var empleado bool = true
var antiguedad int = 3

func check(edad int, empleado bool, antiguedad int){
	if edad > 22 {
		fmt.Println("Cumple con la edad mínima")
	} else {
		fmt.Println("No cumple con la edad mínima")
	}

	if empleado == true {
		fmt.Println("Cumple con la empleabilidad")
	} else {
		fmt.Println("No cumple con la empleabilidad")
	}

	if antiguedad > 1 {
		fmt.Println("Cumple con la antigüedad")
	} else {
		fmt.Println("No cumple con la antigüedad")
	} 
}

func main(){
	check(edad, empleado, antiguedad)
}
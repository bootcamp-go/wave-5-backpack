package main

import "fmt"

var apellidos string = "Gomez"
var edades int = 35
var sueldo float64 = 45857.90
var nombres string = "Julián"

func main() {
	boolean := false
	fmt.Printf("Las variables coregidas son:\n apellidos: %s\n edades: %d\n sueldo: %v\n nombres: %s\n boolean: %v\n", apellidos, edades, sueldo, nombres, boolean)
}

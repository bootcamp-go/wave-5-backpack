package main

import "fmt"

var nombre, direccion = "Andy Esquivel", "Jiutepec, Morelos, México"
func main(){
	fmt.Println(`Hola, yo soy ` + nombre + " y vivo en " + direccion) 
}
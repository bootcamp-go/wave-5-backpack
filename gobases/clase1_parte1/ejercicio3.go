package main

import (
	"fmt"
	"strconv"
)

func main() {
	var lnombre string
	var apellido string
	var edad int
	lapellido := 8
	var licencia_de_conducir = true
	var estatura_de_la_persona int
	cantidadDeHijos := 2

	lnombre = "Camilo"
	apellido = "Calderón"
	edad = 21
	estatura_de_la_persona = 175

	fmt.Println("Nombre: ", lnombre, "Apellido: ", apellido, "Edad: ", edad)
	fmt.Println(strconv.Itoa(lapellido), licencia_de_conducir, estatura_de_la_persona, cantidadDeHijos)

}

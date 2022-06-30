package main

import "fmt"

//Definiendo una estrucutra
type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

//Generando un METODO de la estrucutura de Alumno
func (a Alumno) mostrarDetalle() {
	fmt.Printf("Nombre: %s\nApellido: %s\nDNI: %d\nFecha: %s\n", a.Nombre, a.Apellido, a.DNI, a.Fecha)
}

func main() {
	estudiante1 := Alumno{"Karla", "Estrada", 10073452123, "1998-01-25"} //Una manera de instanciar una estrucutra
	estudiante2 := Alumno{                                               //Otra manera de instanciar la estrucutura
		Nombre:   "Maicol",
		Apellido: "Torres",
		DNI:      1073251375,
		Fecha:    "1998-01-01",
	}

	estudiante1.mostrarDetalle()
	estudiante2.mostrarDetalle()
}

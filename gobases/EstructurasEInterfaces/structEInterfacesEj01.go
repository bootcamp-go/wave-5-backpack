package gobases

import (
	"fmt"
	"time"
)

type Alumnos struct {
	nombre   string
	apellido string
	dni      int
	fecha    time.Time
}

func (a *Alumnos) mostrarAlumno() {
	fmt.Println("Nombre:", a.nombre, "\n", "Apellido:", a.apellido, "\n", "DNI:", a.dni, "\n", "Fecha de ingreso:", a.fecha.Day(), "/", a.fecha.Month(), "/", a.fecha.Year())
}

func main() {

	fmt.Println("Ingrese nombre del alumno")
	var nombre string
	fmt.Scanf("%s", &nombre)

	fmt.Println("Ingrese apellido del alumno")
	var apellido string
	fmt.Scanf("%s", &apellido)

	fmt.Println("Ingrese dni del alumno")
	var dni int
	fmt.Scanf("%d", &dni)

	alumno := Alumnos{
		nombre:   nombre,
		apellido: apellido,
		dni:      dni,
		fecha:    time.Now(),
	}

	alumno.mostrarAlumno()
}

/*
Ejercicio 1 - Registro de estudiantes

Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle

*/

package main
import "fmt"

type Estudiante struct {
	Nombre string
	Apellido string
	DNI int
	Fecha string
}

func (estudiante Estudiante) detalle() {
	fmt.Printf("Detalles:\nNombre: %v, \nApellido: %v, \nDNI: %v, \nFecha: %v", estudiante.Nombre, estudiante.Apellido, estudiante.DNI, estudiante.Fecha)
}

func main(){
	
	e1 := Estudiante {"Juan", "Perez", 33444555, "22/03/2001"}

	e1.detalle()

}

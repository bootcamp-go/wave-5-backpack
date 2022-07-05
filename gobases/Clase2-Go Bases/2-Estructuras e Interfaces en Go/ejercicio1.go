package main

import "fmt"

type Estudiante struct {
	Nombre string
	Apellido string
	DNI string
	FechaIngreso string
}
func (individuo Estudiante) detalle(){
	salida:="Nombre: "+individuo.Nombre+"\nApellido: "+individuo.Apellido+"\nDNI: "+individuo.DNI+"\nFecha: "+individuo.FechaIngreso+"\n"
	fmt.Printf(salida) 
}
func main(){
	alumno:=Estudiante{
		Nombre: "German",
		Apellido: "Rodriguez",
		DNI: "4.705.340-4",
		FechaIngreso:"29/06/2022",
	}
	alumno.detalle()
}






/*Ejercicio 1 - Registro de estudiantes
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle


*/
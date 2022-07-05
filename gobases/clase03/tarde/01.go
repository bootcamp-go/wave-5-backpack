/*
Ejercicio 1 - Red social
Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones:
La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
Y deben implementarse las funciones:
cambiar nombre: me permite cambiar el nombre y apellido.
cambiar edad: me permite cambiar la edad.
cambiar correo: me permite cambiar el correo.
cambiar contraseña: me permite cambiar la contraseña.

*/

package main
import "fmt"

type usuarios struct {
	Nombre string
	Apellido string
	Edad int
	Mail string
	Contrasena string
}

func (u *usuarios) cambiarNombre(nuevoNombre *string) {
	u.Nombre = *nuevoNombre
}

func (u *usuarios) cambiarEdad(nuevaEdad *int){
	u.Edad = *nuevaEdad
}

func (u *usuarios) cambiarMail(nuevoMail *string){
	u.Mail = *nuevoMail
}

func (u *usuarios) cambiarContrasena(nuevaContrasena *string){
	u.Contrasena = *nuevaContrasena
}

func main() {
	usuario1 := usuarios{
		Nombre: "Roberto",
		Apellido: "Gomez Bolaños",
		Edad: 100,
		Mail: "elChavo@mail.com",
		Contrasena: "EsoEsoEso",

	}

	fmt.Print(usuario1, "\n")

	var (
		Nombre string = "Robertoooo"
		Edad int = 99
		Mail string = "elChavito@mail.com"
		Contrasena string = "EsQueNoMeTienenPaciencia"
	)

	usuario1.cambiarNombre(&Nombre)
	usuario1.cambiarEdad(&Edad)
	usuario1.cambiarMail(&Mail)
	usuario1.cambiarContrasena(&Contrasena)

	fmt.Print(usuario1)
}

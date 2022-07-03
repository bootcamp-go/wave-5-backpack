/*Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan
 agregando información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo
 lugar en memoria para el main del programa y para las funciones:
La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
Y deben implementarse las funciones:
cambiar nombre: me permite cambiar el nombre y apellido.
cambiar edad: me permite cambiar la edad.
cambiar correo: me permite cambiar el correo.
cambiar contraseña: me permite cambiar la contraseña.
*/
package main

import "fmt"

func main() {
	mati := Usuario{
		Nombre:      "Matias",
		Apellido:    "Fante",
		Edad:        33,
		Correo:      "mati@mati.com",
		Contrasenia: "123456",
	}
	fmt.Println("Los valores antes del cambio son:", mati)

	cambiarNombre(&mati, "Pedro")
	cambiarEdad(&mati, 40)
	cambiarCorreo(&mati, "pedrito@pedro.com")
	cambiarContrasenia(&mati, "678")

	fmt.Println("Los valores luego de los cambios son:", mati)
}

type Usuario struct {
	Nombre      string
	Apellido    string
	Edad        int
	Correo      string
	Contrasenia string
}

func cambiarNombre(u *Usuario, nombre string) {
	u.Nombre = nombre
}

func cambiarEdad(u *Usuario, edad int) {
	u.Edad = edad
}

func cambiarCorreo(u *Usuario, mail string) {
	u.Correo = mail
}

func cambiarContrasenia(u *Usuario, clave string) {
	u.Contrasenia = clave
}

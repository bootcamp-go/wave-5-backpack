package main

import "fmt"

// Ejercicio 1 - Red social

// Una empresa de redes sociales requiere implementar una estructura usuario con funciones
// que vayan agregando información a la estructura.
// Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo lugar en memoria
// para el main del programa y para las funciones:

// La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
// y deben implementarse las funciones:
//  - cambiar nombre: me permite cambiar el nombre y apellido.
//  - cambiar edad: me permite cambiar la edad.
//  - cambiar correo: me permite cambiar el correo.
//  - cambiar contraseña: me permite cambiar la contraseña.

// Estructura Usuario
type usuario struct {
	nombre     string
	apellido   string
	edad       int
	correo     string
	contrasena string
}

// Función para actualizar el nombre y apellido
func actualizarNombre(u *usuario, n string, a string) {
	u.nombre = n
	u.apellido = a
}

// Función para actualizar la edad
func actualizarEdad(u *usuario, e int) {
	u.edad = e
}

// Función para actualizar el correo
func actualizarCorreo(u *usuario, c string) {
	u.correo = c
}

// Función para actualizar la contraseña
func actualizarPass(u *usuario, p string) {
	u.contrasena = p
}

func main() {
	fmt.Println("Ejercicio 1 - Red social")
	fmt.Println("")

	// Creamos un usuario con valores por default
	user := usuario{nombre: "", apellido: "", edad: 0, correo: "", contrasena: ""}
	fmt.Println(user)

	// Se actualiza el nombre y apellido
	actualizarNombre(&user, "nombre", "apellido")
	fmt.Println(user)

	// Se actualiza la edad
	actualizarEdad(&user, 26)
	fmt.Println(user)

	// Se actualiza el correo
	actualizarCorreo(&user, "correo@dominio.com")
	fmt.Println(user)

	// Se actualiza la contraseña
	actualizarPass(&user, "******")
	fmt.Println(user)
}

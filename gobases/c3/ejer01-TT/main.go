package main

import "fmt"

// Una empresa de redes sociales requiere implementar una estructura usuario
// con funciones que vayan agregando información a la estructura.
// Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe
// el mismo lugar en memoria para el main del programa y para las funciones:
// La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
// Y deben implementarse las funciones:
// cambiar nombre: me permite cambiar el nombre y apellido.
// cambiar edad: me permite cambiar la edad.
// cambiar correo: me permite cambiar el correo.
// cambiar contraseña: me permite cambiar la contraseña.

type Usuario struct {
	Nombre   string
	Apellido string
	Edad     int
	Correo   string
	Password string
}

func cambiarNombre(u *Usuario, nombre, apellido string) {
	u.Nombre = nombre
	u.Apellido = apellido
}

func cambiarEdad(u *Usuario, edad int) {
	u.Edad = edad
}

func cambiarCorreo(u *Usuario, correo string) {
	u.Correo = correo
}

func cambiarPassword(u *Usuario, password string) {
	u.Password = password
}

func main() {
	user := Usuario{
		Nombre:   "Juan",
		Apellido: "Perez",
		Edad:     30,
		Correo:   "juan@perez.com",
		Password: "Pass1234",
	}

	fmt.Println(user)

	cambiarNombre(&user, "Jose", "Lopez")
	cambiarEdad(&user, 35)
	cambiarCorreo(&user, "jose@lopez.com")
	cambiarPassword(&user, "4567Pass")

	fmt.Println(user)

}

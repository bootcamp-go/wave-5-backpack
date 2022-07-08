package main

import (
	"fmt"
)

/*Ejercicio 1 - Red social
Una empresa de redes sociales requiere implementar una estructura usuario con funciones
que vayan agregando información a la estructura. Para optimizar y ahorrar memoria
requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del
programa y para las funciones:

La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
Y deben implementarse las funciones:

    1. cambiar nombre: me permite cambiar el nombre y apellido.
    2. cambiar edad: me permite cambiar la edad.
    3. cambiar correo: me permite cambiar el correo.
    4. cambiar contraseña: me permite cambiar la contraseña. */

type Usuario struct {
	Nombre   string
	Apellido string
	Edad     int
	Correo   string
	Password string
}

func (u *Usuario) UpdateNombre(nombre *string) {
	u.Nombre = *nombre
}

func (u *Usuario) UpdateEdad(edad *int) {
	u.Edad = *edad
}

func (u *Usuario) UpdateCorreo(email *string) {
	u.Correo = *email
}

func (u *Usuario) UpdatePassword(password *string) {
	u.Password = *password
}

func main() {
	usuario := &Usuario{
		Nombre:   "María",
		Apellido: "Martinez",
		Edad:     18,
		Correo:   "maria@gmail.com",
		Password: "1234123",
	}

	var (
		nombre   string = "María Elena"
		edad     int    = 22
		correo   string = "mariaelena@gmail.com"
		password string = "1234123***"
	)

	usuario.UpdateNombre(&nombre)
	usuario.UpdateEdad(&edad)
	usuario.UpdateCorreo(&correo)
	usuario.UpdatePassword(&password)

	fmt.Println("👧", usuario.Nombre, usuario.Edad, usuario.Correo, usuario.Password)
}

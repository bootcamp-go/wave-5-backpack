package main

import "fmt"

type Usuario struct {
	Nombre, Apellido string
	Edad             int
	Correo, Password string
}

func cambiarNombre(u *Usuario, nombre string, apellido string) {
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
		Nombre:   "Cristobal",
		Apellido: "Monsalve",
		Edad:     27,
		Correo:   "cmonsalve@gmail.com",
		Password: "12345",
	}

	fmt.Println(user)

	cambiarNombre(&user, "Juan", "Perez")
	cambiarEdad(&user, 30)
	cambiarCorreo(&user, "cris@gmail.cl")
	cambiarPassword(&user, "67890")

	fmt.Println(user)

}

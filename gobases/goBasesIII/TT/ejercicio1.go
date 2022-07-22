package main

import "fmt"

type usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contrasena string
}

func (u *usuario) cambioNombre(name, lastName string) {
	u.Nombre = name
	u.Apellido = lastName
}

func (u *usuario) cambioEdad(age int) {
	u.Edad = age
}

func (u *usuario) cambioCorreo(mail string) {
	u.Correo = mail
}

func (u *usuario) cambioContrasena(password string) {
	u.Contrasena = password
}

func main() {
	usuario := usuario{Nombre: "Raul", Apellido: "Gonzalez", Edad: 25, Correo: "mymail@mail.com", Contrasena: "abc123"}
	fmt.Println(usuario)
	usuario.cambioNombre("Mathias", "Porta")
	usuario.cambioEdad(36)
	usuario.cambioCorreo("newmail@mail.com")
	usuario.cambioContrasena("4545ADc")
	fmt.Println(usuario)

}

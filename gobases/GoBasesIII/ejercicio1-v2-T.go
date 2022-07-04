package main

import "fmt"

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contrasena string
}

type Methods interface {
	cambiarNombre(name string) string
	cambiarApellido(lastname string) string
	cambiarEdad(age int) int
	cambiarCorreo(mail string) string
	cambiarContrasena(pass string) string
}

func (u *Usuario) cambiarNombre(name string) string {
	u.Nombre = name
	return u.Nombre
}
func (u *Usuario) cambiarApellido(lastname string) string {
	u.Apellido = lastname
	return u.Apellido
}
func (u *Usuario) cambiarEdad(age int) int {
	u.Edad = age
	return u.Edad
}
func (u *Usuario) cambiarCorreo(mail string) string {
	u.Correo = mail
	return u.Correo
}
func (u *Usuario) cambiarContrasena(pass string) string {
	u.Contrasena = pass
	return u.Contrasena
}

func newName(str string) Methods {
	return &Usuario{
		Nombre: str,
	}
}
func newLastname(str string) Methods {
	return &Usuario{
		Apellido: str,
	}
}
func newAge(str int) Methods {
	return &Usuario{
		Edad: str,
	}
}
func newMail(str string) Methods {
	return &Usuario{
		Correo: str,
	}
}
func newContrasena(str string) Methods {
	return &Usuario{
		Contrasena: str,
	}
}

func main() {
	a := newName("Juan")
	b := newLastname("Serna")
	c := newAge(23)
	d := newMail("juan.serna")
	e := newContrasena("12345")

	fmt.Println("Los datos actuales son los siguientes ")
	fmt.Printf("Nombre: %s, Apellido: %s, Edad: %d, Mail: %s y el password: %s \n", a, b, c, d, e)

	a = newName("David")
	c = newAge(20)
	d = newMail("serna.juan")
	fmt.Println("Los datos modificados quedan asi: ")
	fmt.Printf("Nombre: %s, Apellido: %s, Edad: %d, Mail: %s y el password: %s\n", a, b, c, d, e)

}

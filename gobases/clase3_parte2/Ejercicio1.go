package main

import (
	"fmt"
)

type Usuario struct {
	Nombre   string
	Apellido string
	Edad     int
	Correo   string
	Pwd      string
}

func changePwd(u *Usuario, new string) {
	u.Pwd = new
}

func changeName(u *Usuario, name string, lastname string) {
	u.Nombre = name
	u.Apellido = lastname
}
func changeAge(u *Usuario, new int) {
	u.Edad = new
}
func changeEmail(u *Usuario, new string) {
	u.Correo = new
}
func main() {
	u := Usuario{"Sebastian", "Olivera", 25, "lseba@correo.com", "contra"}

	fmt.Println("Usuario al inicio :", u)
	changePwd(&u, "nueva")
	changeAge(&u, 99)
	changeEmail(&u, "correo@correo.com")
	changeName(&u, "Luis", "Oliveraa")
	fmt.Println("Usuario al final :", u)
}

package main

import "fmt"

//Ejercicio 1 - Red social
type Usuario struct {
	Nombre string
	Apellido string
	Edad int
	Correo string
	Contrasena string
}

func (u *Usuario) cambiarNombre(nombre *string, apellido *string)  {
	u.Nombre = *nombre
	u.Apellido = *apellido

}

func (u *Usuario) cambiarEdad(edad *int)  {
	u.Edad = *edad
}
func (u *Usuario) cambiarCorreo(correo *string)  {
	u.Correo = *correo
}
func (u *Usuario) cambiarContrasena(contrasena *string)  {
	u.Contrasena = *contrasena
}
func main()  {
	user := &Usuario{"Vanessa", "Sotomayor", 26, "vane@mail.com", "123"}
	nombre := "Nathaly"
	apellido := "Mera"
	edad := 25
	correo := "nath@mail.com"
	contrasena := "000"
	
	fmt.Println(user)

	user.cambiarNombre(&nombre, &apellido)
	user.cambiarEdad(&edad)
	user.cambiarCorreo(&correo)
	user.cambiarContrasena(&contrasena)

	fmt.Println(user)
}
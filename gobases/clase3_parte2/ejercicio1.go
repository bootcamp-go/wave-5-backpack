package main

import "fmt"

type usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contrasena string
}

type Usuario interface {
	cambiarNombre(Nombre, Apellido string)
	cambiarEdad(Edad int)
	cambiarCorreo(Correo string)
	cambiarContrasena(contrasenaAntigua, contrasenaNueva string)
}

func nuevoUsuario(Nombre, Apellido, Correo, Contrasena string, Edad int) Usuario {
	return &usuario{Nombre, Apellido, Edad, Correo, Contrasena}
}

func (u *usuario) cambiarNombre(Nombre, Apellido string) {
	u.Nombre = Nombre
	u.Apellido = Apellido
}

func (u *usuario) cambiarEdad(Edad int) {
	u.Edad = Edad
}

func (u *usuario) cambiarCorreo(Correo string) {
	u.Correo = Correo
}

func (u *usuario) cambiarContrasena(contrasenaAntigua, contrasenaNueva string) {
	if u.Contrasena == contrasenaAntigua {
		u.Contrasena = contrasenaNueva
	}
}

func main() {
	usuario1 := nuevoUsuario("Camilo", "Calderon", "camilo@correo.com", "12345", 21)
	usuario1.cambiarNombre("Carlos", "Amaya")
	usuario1.cambiarEdad(22)
	usuario1.cambiarCorreo("carlos@correo.com")
	usuario1.cambiarContrasena("123", "321")
	fmt.Println(usuario1)
	usuario1.cambiarContrasena("12345", "321")
	fmt.Println(usuario1)
}

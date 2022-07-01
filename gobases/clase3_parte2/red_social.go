package main

import "fmt"

type Usuario struct {
	Nombre string
	Apellido string
	edad int
	correo string
	contraseña string
}

func (u *Usuario) cambiarNombre(nombre, apellido string){
	u.Nombre = nombre
	u.Apellido = apellido
}

func (u *Usuario) cambiarEdad(edad int){
	u.edad = edad
}

func (u *Usuario) cambiarCorreo(correo string){
	u.correo = correo
}

func (u *Usuario) cambiarContraseña(contraseña string){
	u.contraseña = contraseña
}

func main(){
	u := Usuario{"nahu", "monserrat", 13, "nahuel123", "123"}
	u.cambiarNombre("j", "alvarez")
	fmt.Println(u)
}
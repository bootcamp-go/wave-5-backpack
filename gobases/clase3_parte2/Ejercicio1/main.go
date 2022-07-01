package main

import "fmt"

type Usuario struct{
	Nombre, Apellido string
	edad int
	correo, contraseña string
	
}

func (u *Usuario) imprimir()  {
	fmt.Printf("Nombre completo: %s %s \nEdad: %d\nCorreo: %s\nContraseña: %s\n", u.Nombre, u.Apellido,u.edad,u.correo,u.contraseña)
}

func (u *Usuario) editNombre(nombre, apellido string){
	u.Nombre = nombre
	u.Apellido = apellido
}

func (u *Usuario) editEdad(edad int){
	u.edad = edad
}

func (u *Usuario) editCorreo(correo string){
	u.correo = correo
}

func (u *Usuario) editContraseña(contraseña string){
	u.Nombre = contraseña
}

func main(){
	user1 := Usuario{"Daniela", "Bedoya", 20, "email@gmail.com", "123"}

	user1.imprimir()
}
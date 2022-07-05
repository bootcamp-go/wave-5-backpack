package main

import "fmt"

type usuario struct{
	Nombre string
	Apellido string
	edad int
	correo string
	contraseña string
}

func cambiarNombre(user *usuario,nomb string, apell string) {
	(*user).Nombre=nomb
	(*user).Apellido=apell
}
func cambiarEdad(user *usuario,ed int) {
	(*user).edad=ed
}
func cambiarCorreo(user *usuario,mail string) {
	(*user).correo=mail
}
func cambiarContraseña(user *usuario,pass string) {
	(*user).contraseña=pass
}

func main(){
	user:=usuario{
	Nombre:"Diego",
	Apellido:"Jota",
	edad: 25,
	correo: "dj@gmail.com",
	contraseña: "kkkkk",
	}	
	fmt.Println(user)
	cambiarEdad(&user,35)
	fmt.Println(user)
}


/*
Ejercicio 1 - Red social
Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones:
La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
Y deben implementarse las funciones:
cambiar nombre: me permite cambiar el nombre y apellido.
cambiar edad: me permite cambiar la edad.
cambiar correo: me permite cambiar el correo.
cambiar contraseña: me permite cambiar la contraseña.

*/
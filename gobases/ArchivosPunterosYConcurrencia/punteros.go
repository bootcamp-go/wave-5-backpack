package main

import (
	"fmt"
	"net/mail"
)

const (
	n    = "nombre"
	c    = "correo"
	pass = "contraseña"
)

type Usuario struct {
	nombre     string
	apellido   string
	edad       int
	correo     string
	contraseña string
}

func ImprimirUsuario(user *Usuario) string {

	return fmt.Sprintf("Nombre:%s,Edad:%d", user.nombre, user.edad)
}

func Operaciones(operador string) func(user *Usuario, valor string) string {

	switch operador {
	case "nombre":
		return modificarNombre
	case "correo":
		return modificarCorreo
	case "contraseña":
		return modificarContraseña
	}

	return nil
}

func modificarEdad(user *Usuario, nuevaEdad int) string {

	user.edad = nuevaEdad
	return fmt.Sprintf("La nueva edad del usuario es: %d\n", user.edad)

}

func modificarNombre(user *Usuario, nuevoNombre string) string {

	user.nombre = nuevoNombre
	return fmt.Sprintf("El nuevo nombre del usuario es: %s\n", user.nombre)

}

func modificarCorreo(user *Usuario, nuevoCorreo string) string {

	if validarCorreo(nuevoCorreo) {
		user.correo = nuevoCorreo
		return fmt.Sprintf("El nuevo correo del usuario es: %s\n", user.correo)
	}

	return fmt.Sprintln(fmt.Errorf(mail.ErrHeaderNotPresent.Error()))

}

func validarCorreo(correo string) bool {
	_, err := mail.ParseAddress(correo)
	return err == nil
}

func modificarContraseña(user *Usuario, nuevaContraseña string) string {

	user.contraseña = nuevaContraseña
	return fmt.Sprintf("El nuevo correo del usuario es: %s\n", user.correo)

}

func main() {

	user := Usuario{nombre: "Denis", apellido: "Roldan", edad: 27, correo: "denis.roldan@mercadolibre.com", contraseña: "holamundo"}
	var punteroAUser *Usuario = &user

	//fmt.Println(*punteroAUser)
	//ImprimirUsuario(punteroAUser)

	modificaNombre := Operaciones(n)

	modificaNombre(punteroAUser, "Michel")
	// ImprimirUsuario(*punteroAUser)
	fmt.Println(*punteroAUser)

}

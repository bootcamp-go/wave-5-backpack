package main

import "fmt"

type user struct {
	nombre string
	edad   int
	correo string
	pass   string
}

func (u *user) cambiarNombre(name *string) {
	u.nombre = *name
}

func (u *user) cambiarEdad(edad *int) {
	u.edad = *edad
}

func (u *user) cambiarCorreo(correo *string) {
	u.correo = *correo
}

func (u *user) cambiarPass(pass *string) {
	u.pass = *pass
}

func main() {
	user := user{
		nombre: "Gonza Vernackt",
		edad:   22,
		correo: "padrej04@gmail.com",
		pass:   "87654321",
	}

	fmt.Println(user)

	var (
		nombre string = "Gonzalo Estaban Vernackt Vince"
		edad   int    = 28
		correo string = "gonzalovernackt@gmail.com"
		pass   string = "12345678"
	)

	user.cambiarNombre(&nombre)
	user.cambiarEdad(&edad)
	user.cambiarCorreo(&correo)
	user.cambiarPass(&pass)

	fmt.Println(user)
}

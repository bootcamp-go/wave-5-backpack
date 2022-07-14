package main

import (
	"fmt"
)

type User struct {
	Nombre string
	Apellido string
	Edad int
	Correo string
	Password string
}

func main()  {
	usuario := &User{
		Nombre:   "María",
		Apellido: "Martinez",
		Edad:     18,
		Correo:   "maria@gmail.com",
		Password: "1234123",
	}

	fmt.Println("---------------------")
	fmt.Println("Primer usuario")
	fmt.Println(usuario.Nombre, usuario.Apellido, usuario.Edad, usuario.Correo, usuario.Password)
	fmt.Println("---------------------")

	var (
		nombre string = "Nicolas"
		apellido string = "Herrera"
		edad int = 19
		correo string = "nicolas.29@gmail.com"
		password string = "123456"
	)

	usuario.UpdateNombre(&nombre)
	usuario.UpdateApellido(&apellido)
	usuario.UpdateEdad(&edad)
	usuario.UpdateCorreo(&correo)
	usuario.UpdatePassword(&password)

	fmt.Println("---------------------")
	fmt.Println("Update Usuario")
	fmt.Println(usuario.Nombre, usuario.Apellido, usuario.Edad, usuario.Correo, usuario.Password)
	fmt.Println("---------------------")
}

func (u *User) UpdateNombre(nombre *string) {
	u.Nombre = *nombre
}

func (u *User) UpdateApellido(apellido *string) {
	u.Apellido = *apellido
}

func (u *User) UpdateEdad(edad *int) {
	u.Edad = *edad
}

func (u *User) UpdateCorreo(correo *string) {
	u.Correo = *correo
}

func (u *User) UpdatePassword(password *string) {
	u.Password = *password
}

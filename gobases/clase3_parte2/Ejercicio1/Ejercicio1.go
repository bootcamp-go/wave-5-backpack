package main

type Usuario struct {
	Nombre   string
	Apellido string
	Edad     int
	Correo   string
	Password string
}

func (user *Usuario) cambiarNombre(nombre *string, apellido *string) {
	user.Nombre = *nombre
	user.Apellido = *apellido
}

func (user *Usuario) cambiarEdad(edad *int) {
	user.Edad = *edad
}

func (user *Usuario) cambiarCorreo(correo *string) {
	user.Correo = *correo
}

func (user *Usuario) cambiarPassword(pass *string) {
	user.Password = *pass
}

func main() {
	user1 := Usuario{
		Nombre:   "pepe",
		Apellido: "user",
		Edad:     46,
		Correo:   "dummy",
		Password: "dummy",
	}
	var nombre, apellido, correo, password string = "carlitos", "perez", "a@a.com", "hola123"
	var edad int = 35
	user1.cambiarNombre(&nombre, &apellido)
	user1.cambiarCorreo(&correo)
	user1.cambiarPassword(&password)
	user1.cambiarEdad(&edad)
	println(user1.Nombre, user1.Apellido, user1.Edad, user1.Correo, user1.Password)
}

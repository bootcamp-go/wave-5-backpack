// Ejercicio 1 - Red social
// Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura. Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones:
// La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
// Y deben implementarse las funciones:
// cambiar nombre: me permite cambiar el nombre y apellido.
// cambiar edad: me permite cambiar la edad.
// cambiar correo: me permite cambiar el correo.
// cambiar contraseña: me permite cambiar la contraseña.


package main

import "fmt"

type Usuario struct {
	Name   string
	LastName string
	Age     int
	Email   string
	Pass string
}

func (u *Usuario) UpdateNombre(nombre *string) {
	u.Name = *nombre
}

func (u *Usuario) UpdateEdad(edad *int) {
	u.Age = *edad
}

func (u *Usuario) UpdateCorreo(email *string) {
	u.Email = *email
}

func (u *Usuario) UpdatePassword(password *string) {
	u.Pass = *password
}

func main() {
	user := &Usuario{
		Name:   "Eimi",
		LastName: "Galvan",
		Email:   "eimi@mail.com",
		Age:     33,
		Pass: "lalalala",
	}

	var (
		name   string = "fran"
		age     int    = 36
		email   string = "fran@mail.com"
		pass string = "lelele"
	)

	user.UpdateNombre(&name)
	user.UpdateEdad(&age)
	user.UpdateCorreo(&email)
	user.UpdatePassword(&pass)

	fmt.Println("user:", user.Name, user.Age, user.Email, user.Pass)
}

package main

import "fmt"

//Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información a la estructura.
//Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa
//y para las funciones:
//La estructura debe tener los campos: Nombre, Apellido, edad, correo y contraseña
//Y deben implementarse las funciones:
//cambiar nombre: me permite cambiar el nombre y apellido.
//cambiar correo: me permite cambiar el correo.
//cambiar contraseña: me permite cambiar la contraseña.

type User struct {
	Name     string
	LastName string
	Age      int
	Email    string
	Password string
}

func (u *User) ChangeUser(name, lastname, email, password *string, edad *int) {
	u.Name = *name
	u.LastName = *lastname
	u.Email = *email
	u.Password = *password
	u.Age = *edad
}

func main() {
	user := &User{
		Name:     "José Luis",
		LastName: "Riverón",
		Age:      28,
		Email:    "jriveronrodriguez@gmail.com",
		Password: "123456",
	}

	var (
		name     = "Carlos Manuel"
		lastname = "Campos"
		email    = "cmcampos0@gmail.com"
		password = "123456**"
		edad     = 34
	)

	user.ChangeUser(&name, &lastname, &email, &password, &edad)
	fmt.Println("Los datos del usuario son:", user.Name, user.LastName, user.Age, user.Email, user.Password, user.Age)
}

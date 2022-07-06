package main

import "fmt"

type user struct {
	name     string
	lastName string
	age      int
	email    string
	password string
}

func (this *user) changeNameAndLastname(newName *string, newLastName *string) {
	this.name = *newName
	this.lastName = *newLastName
}

func (this *user) changeAge(newAge *int) {
	this.age = *newAge
}
func (this *user) changeEmail(newEmail *string) {
	this.email = *newEmail
}
func (this *user) changePassword(newPassword *string) {
	fmt.Println(&this.email)
	this.password = *newPassword
}

func main() {
	u1 := &user{
		name:     "Santiago",
		lastName: "Salcedo",
		age:      18,
		email:    "santiago.salcedo@mercadolibre.com.co",
		password: "m.4,SK+k,jRp82D5",
	}
	var (
		nombre      = "Santiago Rafael"
		apellido    = "Salcedo Camacho"
		edad        = 19
		correo      = "xantiago.salcedo@gmail.com"
		contrasenia = "santi12345"
	)
	fmt.Println(&u1.password)
	u1.changeNameAndLastname(&nombre, &apellido)
	u1.changeEmail(&correo)
	u1.changeAge(&edad)
	u1.changePassword(&contrasenia)
}

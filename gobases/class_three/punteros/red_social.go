package main

import "fmt"

type Usuario struct {
	name     string
	lastName string
	age      int
	email    string
	password string
}

func main() {
	var (
		p     *Usuario
		user1 = Usuario{"Andrea", "Esquivel", 23, "andy@gmail.com", "jKF$6Fe"}
	)

	p = &user1
	//Name
	fmt.Printf("El nombre con el que se inicalizó el usuario es: %s\n", *&p.name)
	changeName(p, "Gabriela")
	fmt.Printf("El nuevo nombre del usuario es: %s\n", *&p.name)

	//Age
	fmt.Printf("La edad con la que se inicalizó el usuario es: %d\n", *&p.age)
	changeAge(p, 24)
	fmt.Printf("La nueva edad del usuario es: %d\n", *&p.age)

	// Email
	fmt.Printf("El email con el que se inicalizó el usuario es: %s\n", *&p.email)
	changeEmail(p, "gaby@gmail.com")
	fmt.Printf("El nuevo email del usuario es: %s\n", *&p.email)

	// Passwd
	fmt.Printf("La password con la que se inicalizó el usuario es: %s\n", *&p.password)
	changePaswd(p, "JJkd37&fd")
	fmt.Printf("La nueva password del usuario es: %s\n", *&p.password)
}

func changeName(user *Usuario, newName string) {
	*&user.name = newName
}

func changeAge(user *Usuario, newAge int) {
	*&user.age = newAge
}

func changeEmail(user *Usuario, newEmail string) {
	*&user.email = newEmail

}

func changePaswd(user *Usuario, newPassword string) {
	*&user.password = newPassword
}

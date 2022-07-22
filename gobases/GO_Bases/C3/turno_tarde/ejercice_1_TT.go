package main

import (
	"fmt"
)

type User struct {
	Name     string
	Surname  string
	Age      int
	Email    string
	Password string
}

func (u *User) UpdateName(name string) {
	u.Name = name
}

func (u *User) UpdateAge(age int) {
	u.Age = age
}

func (u *User) UpdateEmail(email string) {
	u.Email = email
}

func (u *User) UpdatePassword(password string) {
	u.Password = password
}

func main() {
	user := &User{
		Name:     "Franco",
		Surname:  "Pergolini",
		Age:      28,
		Email:    "franco@gmail.com",
		Password: "1234",
	}

	user.UpdateName("Franco")
	user.UpdateAge(27)
	user.UpdateEmail("franco@gmail.com")
	user.UpdatePassword("1234")

	fmt.Println(user)
}

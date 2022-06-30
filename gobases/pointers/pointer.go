package main

import "fmt"

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"document"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (stu *Person) CreatePerson() {
	*stu = Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       26,
		Email:     "john.doe@meli.com",
		Password:  "john.doe@meli01",
	}
}
func UpdateName(fName, lName string, stu *Person) {
	stu.FirstName = fName
	stu.LastName = lName
}
func UpdateAge(age int, stu *Person) {
	stu.Age = age
}
func UpdateEmail(email string, stu *Person) {
	stu.Email = email
}
func UpdatePassword(pass string, stu *Person) {
	stu.Password = pass
}
func main() {
	p1 := Person{
		FirstName: "J J",
		LastName:  "Lopez",
		Age:       12,
		Email:     "asd@asd",
		Password:  "asd",
	}
	fmt.Println("Antes de los cambios")
	fmt.Println(p1.Age)
	fmt.Println(p1.FirstName)
	fmt.Println(p1.LastName)
	fmt.Println(p1.Email)
	UpdateAge(29, &p1)
	UpdateEmail("johnd@gmail.com", &p1)
	UpdatePassword("asd123", &p1)
	UpdateName("John", "Maxwell", &p1)
	fmt.Println("Despues de los cambios")
	fmt.Println(p1.Age)
	fmt.Println(p1.FirstName)
	fmt.Println(p1.LastName)
	fmt.Println(p1.Email)
}

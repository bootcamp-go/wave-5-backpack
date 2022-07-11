package main

import (
	"fmt"
)

type Studient struct{
	Name string
	LastName string
	ID int
	Date string
}

func main()  {
	student := Studient{
		Name: 		"Yoshua",
		LastName: 	"Cary",
		ID: 		44111726,
		Date: 		"19/12/1988",
	}

	fmt.Println("Estudiante: ", student)
}
package main

import "fmt"

type Student struct {
	Name     string
	LastName string
	DNI      int
	Date     string
}

func main() {
	s1 := Student{
		Name:     "Santiago Rafael",
		LastName: "Salcedo Camacho",
		DNI:      1118528414,
		Date:     "Junio 21 del 2022",
	}
	fmt.Println(s1)
}

package main

import "fmt"

type Student struct {
	Name     string
	LastName string
	DNI      int
	Date     string
}

func (s Student) Details() string {
	return fmt.Sprintf(" Name: %s \n LastName: %s \n DNI: %d \n Date: %s", s.Name, s.LastName, s.DNI, s.Date)
}

func main() {
	s1 := Student{
		Name:     "Santiago Rafael",
		LastName: "Salcedo Camacho",
		DNI:      1118528414,
		Date:     "Junio 21 del 2022",
	}
	fmt.Println(s1.Details())
}

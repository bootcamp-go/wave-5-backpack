package main

import (
	"errors"
	"fmt"
	"time"
)

func TaxSalary(salary float64) float64 {
	switch {
	case salary >= 150000:
		return salary - (salary * 10 / 100)
	case salary >= 50000:
		return salary - (salary * 17 / 100)
	default:
		return salary
	}
}

func School(calif ...float64) (float64, error) {
	var subtotal float64
	var i float64
	for _, calific := range calif {
		if calific < 0 {
			return 0, errors.New("No puede haber una calificacion negativa")
		}
		subtotal = subtotal + calific
		i++
	}
	return subtotal / i, nil

}

func SalaryCalc(min float32, cat string) float32 {
	var result float32 = 0.0
	switch cat {
	case "A":
		{
			var salMin float32 = 3000.0 / 60.0
			result = salMin * min
			result = result * 1.5
		}
	case "B":
		{
			var salMin float32 = 1500.0 / 60.0
			result = salMin * min
			result = result * 1.2
		}
	case "C":
		{
			var salMin float32 = 1000.0 / 60.0
			result = salMin * min
			fmt.Println(salMin)

		}
	}
	return result
}

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func MinInt(vn ...float64) (m float64) {
	m = vn[0]
	for i := 1; i < len(vn); i++ {
		if vn[i] < m {
			m = vn[i]
		}
	}
	return
}

func MaxInt(vn ...float64) (m float64) {
	m = vn[0]
	for i := 1; i < len(vn); i++ {
		if vn[i] > m {
			m = vn[i]
		}
	}
	return
}

func Average(num ...float64) float64 {
	var subtotal float64
	var i float64
	for _, n := range num {
		subtotal = subtotal + n
		i++
	}
	return subtotal / i
}

func operation(op string, vals ...float64) (float64, error) {
	switch op {
	case "Promedio":
		return Average(vals...), nil
	case "Maximo":
		return MaxInt(vals...), nil
	case "Minimo":
		return MinInt(vals...), nil
	default:
		return 0.0, errors.New("Operacion Invalida")
	}

}
func Animal(op string) (float64, error) {
	switch op {
	case "Perro":
		return 10, nil
	case "Gato":
		return 5, nil
	case "Tarantula":
		return 0.15, nil
	case "Hamster":
		return 0.25, nil
	default:
		return 0, errors.New("No existe ese animal")
	}
}

type Student struct {
	Id        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Document  Document  `json:"document"`
	EntryDate time.Time `json:"entryDate"`
}

type Document struct {
	DocumentNumber int    `json:"documentNumber"`
	DocumentType   string `json:"documentType"`
}

func CreateStudent() (stu Student) {
	stu = Student{
		Id:        1,
		FirstName: "John",
		LastName:  "Doe",
		Document: Document{
			DocumentNumber: 123456,
			DocumentType:   "DNI",
		},
		EntryDate: time.Now(),
	}
	return stu
}

func Detail(stu Student) {
	fmt.Println(stu)
}

func main() {
	fmt.Println("Hello Mundo!")
	//fmt.Println(School(5.0, 5.0, 5.0))
	//fmt.Println(SalaryCalc(60, "C"))
	//fmt.Println(operation("Minsßßimo", 9, 10, 8))
	//fmt.Println(Animal("Perro"))
	var stu = CreateStudent()
	Detail(stu)
}

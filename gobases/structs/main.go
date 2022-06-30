package main

import (
	"fmt"
	"time"
)

type Student struct {
	Id        int      `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Document  Document `json:"document"`
	EntryDate string   `json:"entryDate"`
}
type Document struct {
	DocumentNumber int    `json:"documentNumber"`
	DocumentType   string `json:"documentType"`
}

func (stu *Student) CreateStudent() {
	stu = &Student{
		Id:        1,
		FirstName: "John",
		LastName:  "Doe",
		Document: Document{
			DocumentNumber: 123456,
			DocumentType:   "DNI",
		},
		EntryDate: time.Now().Format("01-02-2006"),
	}
}
func Detail(stu Student) {
	fmt.Println(stu)
}

func main() {
	fmt.Println("Hello Mundo!")
	stu := Student{Id: 1,
		FirstName: "John",
		LastName:  "Doe",
		Document: Document{
			DocumentNumber: 123456,
			DocumentType:   "DNI",
		},
		EntryDate: time.Now().Format("01-02-2006")}
	//stu.CreateStudent()
	Detail(stu)
}

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var tickets []Ticket
	// Funcion para obtener tickets del archivo csv
	NewBookings(tickets)

	file := File{}
	tickets, err := file.ReadFile() // slice de tipo byte con al info , se debe separar la info

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("no error reading file")
	//fmt.Printf("%v \n", tickets)
	fmt.Println("end of reading file")

}

// services
type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	return Ticket{}, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	return Ticket{}, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	return Ticket{}, nil
}

func (b *bookings) Delete(id int) (int, error) {
	return 0, nil
}

// file
type File struct {
	path string
}

//
func (f *File) ReadFile() ([]Ticket, error) {

	var fileSliceFormat []Ticket
	file, err := os.ReadFile("./tickets.csv")

	//fmt.Printf("%v", file)
	if err != nil {
		return []Ticket{}, errors.New("cannot read file")
	}
	// type Ticket struct {
	// 	Id                              int
	// 	Names, Email, Destination, Date string
	// 	Price                           int
	// }
	for _, v := range file {
		var position int // position on slice type ticket
		var temp Ticket
		switch position {
		case 0:
			temp.Id += int(v)
		case 1:
			temp.Names += string(v)
		case 2:
			temp.Email += string(v)
		case 3:
			temp.Destination += string(v)
		case 4:
			temp.Date += string(v)
		case 5:
			temp.Price += int(v)
		}

		if v == 44 { //separamos por commas ASCII 44d
			var count, index int
			count++
			position++
			if v == 13 { // encontramos el fin de la fila
				fileSliceFormat[index] = temp
				index++
				position = 0
			}

			fmt.Printf("%v \n", temp.Names)
		}

	}
	return fileSliceFormat, nil
}

func (f *File) Write(Ticket) error {
	return nil
}

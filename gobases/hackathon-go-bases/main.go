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
	fmt.Printf("%v /n", tickets)

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

	file, err := os.ReadFile("./tickets.csv")

	//fmt.Printf("%v", file)
	if err != nil {
		return []Ticket{}, errors.New("cannot read file")
	}

	for _, v := range file {
		if v == 110 {
			fmt.Println(v)
		}
	}

	return nil, nil
}

func (f *File) Write(Ticket) error {
	return nil
}

package service

import (
	"errors"
	"fmt"
)

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

	for _, value := range b.Tickets {
		if value.Id == t.Id {
			return t, errors.New("EL id ya esta repetido")
		}
	}

	b.Tickets = append(b.Tickets, t)

	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	for _, value := range b.Tickets {
		if value.Id == id {
			return value, nil
		}
	}
	return Ticket{}, errors.New("No hay vuelo con el ID ingresado")
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	for key, _ := range b.Tickets {
		if b.Tickets[key].Id == id {
			fmt.Println(key)
			fmt.Println(b.Tickets[key])
			b.Tickets[key].Names = t.Names
			b.Tickets[key].Email = t.Email
			b.Tickets[key].Price = t.Price
			b.Tickets[key].Destination = t.Destination
			b.Tickets[key].Date = t.Date

			return b.Tickets[key], nil
		}
	}
	return Ticket{}, errors.New("No hay vuelo con el ID ingresado")
}

func (b *bookings) Delete(id int) (int, error) {
	for _, value := range b.Tickets {
		if value.Id == id {
			b.Tickets = append(b.Tickets[:id-1], b.Tickets[id:]...)
			return id, nil
		}
	}
	return id, errors.New("EL ID INGRESADO NO EXISTE")
}

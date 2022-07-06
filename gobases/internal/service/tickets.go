package service

import (
	"errors"
	"fmt"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (*Ticket, error)
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

func (t Ticket) Detalle() string {
	return fmt.Sprintf("\n%d,%s,%s,%s,%s,%d", t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price)
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	for _, ticket := range b.Tickets {
		if t.Id == ticket.Id {
			return ticket, errors.New("El id ya existe")
		}
	}

	return t, nil
}

func (b *bookings) Read(id int) (*Ticket, error) {
	for _, ticket := range b.Tickets {
		if id == ticket.Id {
			return &ticket, nil
		}
	}

	return &Ticket{}, errors.New("No se encontro ningun ticket con el ID ingresado")
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	var ticket *Ticket
	ticket, err := b.Read(id)
	if err != nil {
		return Ticket{}, err
	} else {
		ticket.Id = t.Id
		ticket.Names = t.Names
		ticket.Email = t.Email
		ticket.Destination = t.Destination
		ticket.Date = t.Date
		ticket.Price = t.Price
	}

	return *ticket, nil
}

func (b *bookings) Delete(id int) (int, error) {
	for i := range b.Tickets {
		if id == b.Tickets[i].Id {
			b.Tickets = append(b.Tickets[:i], b.Tickets[i+1:]...)
			return id, nil
		}
	}
	return 0, errors.New("No se encontro el registo a borrar")
}

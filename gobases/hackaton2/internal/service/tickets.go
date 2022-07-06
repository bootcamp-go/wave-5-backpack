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
	CreateID() int
	GetTickets() []Ticket
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
	t.Id = b.CreateID()
	if t.Names == "" || t.Email == "" || t.Destination == "" || t.Date == "" || t.Price == 0 {
		return Ticket{}, errors.New("error: algún dato es inválido")
	} else {
		b.Tickets = append(b.Tickets, t)
		return t, nil
	}
}

func (b *bookings) Read(id int) (Ticket, error) {
	for _, value := range b.Tickets {
		if id == value.Id {
			return b.Tickets[id-1], nil
		}
	}
	return Ticket{}, errors.New("error: no existe el id")
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	if t.Names == "" || t.Email == "" || t.Destination == "" || t.Date == "" || t.Price == 0 {
		return Ticket{}, errors.New("error: algún dato es inválido")
	}
	for _, value := range b.Tickets {
		if id == value.Id {
			value.Names = t.Names
			value.Email = t.Email
			value.Destination = t.Destination
			value.Date = t.Date
			value.Price = t.Price
			b.Tickets[id-1] = value
			return b.Tickets[id-1], nil
		}
	}
	return Ticket{}, errors.New("error: no se pueden editar los datos porque el id no existe")
}

func (b *bookings) Delete(id int) (int, error) {
	if id <= 0 || id > b.Tickets[len(b.Tickets)-1].Id {
		return 0, errors.New("error: el id es inválido")
	}
	for idx, value := range b.Tickets {
		if id == value.Id {
			b.Tickets = append(b.Tickets[:idx], b.Tickets[idx+1:]...)
			fmt.Println("Registro eliminado satisfactoriamente!")
			return id, nil
		}
	}
	return id, errors.New("error: no se pueden eliminar los datos porque el id no existe")
}

func (b bookings) CreateID() int {
	id := b.Tickets[len(b.Tickets)-1].Id + 1
	return id
}

func (b *bookings) GetTickets() []Ticket {
	return b.Tickets
}

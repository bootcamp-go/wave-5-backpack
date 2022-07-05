package service

import (
	"errors"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
	Detalle() []Ticket
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

func (b *bookings) Create(t Ticket) {
	b.Tickets = append(b.Tickets, t)
}

func (b *bookings) Read(id int) (Ticket, error) {
	for _, ticket := range b.Tickets {
		if ticket.Id == id {
			return ticket, nil
		}
	}
	return Ticket{}, errors.New("no se encontró el ticket que se busca")
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	for i := 0; i < len(b.Tickets); i++ {
		ticket := &b.Tickets[i]
		if id == ticket.Id {
			ticket.Id = t.Id
			ticket.Names = t.Names
			ticket.Email = t.Email
			ticket.Destination = t.Destination
			ticket.Date = t.Date
			ticket.Price = t.Price
			return *ticket, nil
		}
	}
	return Ticket{}, errors.New("no se encontro el ticket ingresado con id")
}

func (b *bookings) Delete(id int) (int, error) {
	for i, ticket := range b.Tickets {
		if id == ticket.Id {
			b.Tickets = append(b.Tickets[:i], b.Tickets[i+1:]...)
			return id, nil
		}
	}

	return 0, errors.New("no se encontro el ticket con id indicado para eliminarlo")
}

func (b *bookings) Detalle() []Ticket {
	return b.Tickets
}

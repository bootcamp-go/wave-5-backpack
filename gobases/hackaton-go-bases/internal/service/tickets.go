package service

import (
	"errors"
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
			return Ticket{}, errors.New("Id de ticket ya existente")
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
	return Ticket{}, errors.New("Id de ticket no existente")
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	for i, value := range b.Tickets {
		if value.Id == id {
			b.Tickets[i] = t
			return value, nil
		}
	}
	return Ticket{}, errors.New("Id de ticket no existente")
}

func (b *bookings) Delete(id int) (int, error) {
	for i, value := range b.Tickets {
		if value.Id == id {
			b.Tickets = append(b.Tickets[:i], b.Tickets[i+1:]...)
			return id, nil
		}
	}
	return 0, errors.New("Id de ticket no existente")
}

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
func NewBookings(Tickets []Ticket) bookings {
	return bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	b.Tickets = append(b.Tickets, t)
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	t, err := b.getById(id)
	return t, err
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	index, err := b.getIndexById(id)
	if err != nil {
		return Ticket{}, err
	}
	b.Tickets[index] = t
	return b.Tickets[index], nil
}

func (b *bookings) Delete(id int) (int, error) {
	index, err := b.getIndexById(id)
	if err != nil {
		return 0, err
	}
	b.Tickets = append(b.Tickets[:index], b.Tickets[index+1:]...)
	return index, nil
}

func (b *bookings) getById(id int) (Ticket, error) {
	var ticket Ticket

	for _, t := range b.Tickets {
		if t.Id == id {
			return t, nil
		}
	}

	return ticket, errors.New(fmt.Sprint("Could not find item with given Id :", id))
}

func (b *bookings) getIndexById(id int) (int, error) {
	for i, t := range b.Tickets {
		if t.Id == id {
			return i, nil
		}
	}

	return 0, errors.New(fmt.Sprint("Could not find item with given Id :", id))
}

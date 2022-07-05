package service

import (
	"errors"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Read all Tickets
	ReadAll() ([]Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets map[int]Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	TicketsMap := make(map[int]Ticket)
	for _, ticket := range Tickets {
		TicketsMap[ticket.Id] = ticket
	}
	return &bookings{Tickets: TicketsMap}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	b.Tickets[t.Id] = t
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	ticket, ok := b.Tickets[id]
	if ok {
		return ticket, nil
	} else {
		return Ticket{}, errors.New("el ticket no existe")
	}
}

func (b *bookings) ReadAll() ([]Ticket, error) {
	tickets := []Ticket{}
	for _, ticket := range b.Tickets {
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	_, ok := b.Tickets[id]
	if ok {
		b.Tickets[id] = t
		return t, nil
	} else {
		return Ticket{}, errors.New("el ticket no existe")
	}
}

func (b *bookings) Delete(id int) (int, error) {
	delete(b.Tickets, id)
	return 0, nil
}

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
	Get() []Ticket
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int64
	Names, Email, Destination, Date string
	Price                           int64
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	previousLen := len(b.Tickets)
	t.Id = b.Tickets[len(b.Tickets)-1].Id + 1
	b.Tickets = append(b.Tickets, t)
	postLen := len(b.Tickets)
	if previousLen >= postLen {
		err := errors.New("error: create ticket failed")
		return Ticket{}, err
	}
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	if b.Tickets[id].Id-1 != 0 {
		return b.Tickets[id], nil
	} else {
		return Ticket{}, errors.New("error: Ticket does not exists")
	}
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	_, err := b.Read(id - 1)
	if err != nil {
		return Ticket{}, err
	}
	b.Tickets[id-1] = t
	b.Tickets[id-1].Id = int64(id)
	return t, nil
}

func (b *bookings) Delete(id int) (int, error) {
	_, err := b.Read(id)
	if err != nil {
		return 0, err
	}
	b.Tickets[id-1] = Ticket{}
	return 0, nil
}

func (b *bookings) Get() []Ticket {
	return b.Tickets
}

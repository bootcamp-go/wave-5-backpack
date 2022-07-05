package service

import (
	"errors"
	"strconv"
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

	CountAllTickets() []Ticket
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
	for _, tiquete := range b.Tickets {
		if tiquete.Id == t.Id {
			return Ticket{}, errors.New("Ya existe el ID: " + strconv.Itoa(t.Id))
		}
	}
	b.Tickets = append(b.Tickets, t)
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	for _, t := range b.Tickets {
		if t.Id == id {
			return t, nil
		}
	}
	return Ticket{}, errors.New("ID Ticket" + strconv.Itoa(id) + " no encontrado")
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	for i, ticket := range b.Tickets {
		if t.Id == id {
			b.Tickets[i] = ticket
			return ticket, nil
		}
	}
	return Ticket{}, errors.New("ID Ticket" + strconv.Itoa(id) + " no encontrado")
}

func (b *bookings) Delete(id int) (int, error) {
	ticket := -1
	for i, t := range b.Tickets {
		if t.Id == id {
			ticket = i
			break
		}
	}

	if ticket == -1 {
		return -1, errors.New("ID Ticket" + strconv.Itoa(id) + " no encontrado")
	}

	b.Tickets = append(b.Tickets[:ticket], b.Tickets[ticket+1:]...)
	return ticket, nil
}

func AutoIncrement(tickets []Ticket) int {
	max := 0
	for _, ticket := range tickets {
		if ticket.Id > max {
			max = ticket.Id
		}
	}
	return max + 1
}

func (b *bookings) CountAllTickets() []Ticket {
	return b.Tickets
}

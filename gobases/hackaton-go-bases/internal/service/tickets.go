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
	// ReadAll read all Tickets
	ReadAll() []Ticket
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
			return Ticket{}, errors.New("Ticket con id " + strconv.Itoa(t.Id) + " ya existe")
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
	return Ticket{}, errors.New("Ticket con id " + strconv.Itoa(id) + " no encontrado")
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	for i, t := range b.Tickets {
		if t.Id == id {
			b.Tickets[i] = t
			return t, nil
		}
	}
	return Ticket{}, errors.New("Ticket con id " + strconv.Itoa(id) + " no encontrado, por lo tanto no se puede actualizar")
}

func (b *bookings) Delete(id int) (int, error) {
	idTicket := -1
	for i, t := range b.Tickets {
		if t.Id == id {
			idTicket = i
			break
		}
	}

	if idTicket == -1 {
		return -1, errors.New("Ticket con id " + strconv.Itoa(id) + " no encontrado")
	}

	b.Tickets = append(b.Tickets[:idTicket], b.Tickets[idTicket+1:]...)
	return idTicket, nil
}

func (b *bookings) ReadAll() []Ticket {
	return b.Tickets
}

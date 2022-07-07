package service

import (
	"errors"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// ReadAll  read al Tickets
	ReadAll() ([]Ticket)
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
	err := checkEmpty(t)
	if err != nil {
		return Ticket{}, err
	}

	// toma el ultimo id y le asigna el siguiente num al nuevo ticket
	last := len(b.Tickets) - 1

	t.Id = b.Tickets[last].Id + 1

	b.Tickets =  append(b.Tickets, t)

	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {	
	for _, t := range b.Tickets {
		if t.Id == id {
			return t, nil
		}
	}

	return Ticket{}, errors.New("error: no existe el ticket")
}

func (b *bookings) ReadAll() ([]Ticket) {
	return b.Tickets
}

// t Ticket contiene una instancia de ticket con los datos a actualizar
func (b *bookings) Update(id int, t Ticket) (Ticket, error) {	
	for i, tt := range b.Tickets {
		if tt.Id == id {
			tt = update(t, &tt)

			b.Tickets[i] = tt

			return tt, nil
		}
	}

	return Ticket{}, errors.New("error: no existe el id a actualizar")
}

// t update info, tt *Ticket to update
func update(t Ticket, tt *Ticket) Ticket {
	if t.Names != "" {
		tt.Names = t.Names
	}

	if t.Email != "" {
		tt.Email = t.Email
	}

	if t.Destination != "" {
		tt.Destination = t.Destination
	}

	if t.Date != "" {
		tt.Date = t.Date
	}

	if t.Price != 0 {
		tt.Price = t.Price
	}

	return *tt
}

func (b *bookings) Delete(id int) (int, error) {
	// Obtener el index del Ticket a eliminar
	for i, ticket := range b.Tickets {
		if ticket.Id == id {
			// actualiza slice
			b.Tickets = append(b.Tickets[:i], b.Tickets[i+1:]...)
			return id, nil
		}
	}

	return 0, errors.New("error: no existe el id a eliminar")
}

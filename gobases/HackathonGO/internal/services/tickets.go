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
	// ReadAll Ticket
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

//Leer todo el archivo
func (b *bookings) ReadAll() []Ticket {
	return b.Tickets
}

// Crear nuevo ticket
func (b *bookings) Create(t Ticket) (Ticket, error) {
	for _, value := range b.Tickets {
		if value.Id == t.Id {
			return Ticket{}, errors.New("Este ticket ya existe")

		}
	}
	b.Tickets = append(b.Tickets, t)
	return t, nil
}

// Leer Ticket
func (b *bookings) Read(id int) (Ticket, error) {
	for _, value := range b.Tickets {
		if value.Id == id {
			return value, nil
		}
		return Ticket{}, errors.New("El ticket no existe")
	}
	return Ticket{}, nil
}

// Actualizar Ticket
func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	for i, value := range b.Tickets {
		if value.Id == id {
			b.Tickets[i] = value
			return value, nil
		}
	}
	return Ticket{}, errors.New("No se encontro el ticket, no se puede actualizar")
}

// Borrar Ticket
func (b *bookings) Delete(id int) (int, error) {
	eliminarT := -1
	for i, t := range b.Tickets {
		if t.Id == id {
			eliminarT = i
			break
		}
	}

	if eliminarT == -1 {
		return -1, errors.New("Ticket no encontrado")
	}

	b.Tickets = append(b.Tickets[:eliminarT], b.Tickets[eliminarT+1:]...)
	return eliminarT, nil
}

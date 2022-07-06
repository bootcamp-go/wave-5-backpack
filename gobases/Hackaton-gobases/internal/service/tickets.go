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
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

//Create functions
func (b *bookings) Create(t Ticket) (Ticket, error) {
	err := Validation(t)
	if err != nil {
		return Ticket{}, nil
	}
	id := len(b.Tickets) + 2
	t.Id = id
	b.Tickets = append(b.Tickets, t)
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	var find bool = false
	var ticket Ticket
	for _, values := range b.Tickets {
		if values.Id == id {
			find = true
			ticket = values
		}
	}
	if !find {
		return Ticket{}, fmt.Errorf("El ticket con el id ", id, "no existe")
	}
	return ticket, nil
}

//Update Function
func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	err := Validation(t)
	if err != nil {
		return Ticket{}, nil
	}
	var find bool = false
	var ticket Ticket
	for _, values := range b.Tickets {
		if values.Id == id {
			find = true
			ticket = t
			ticket.Id = id
		}
	}
	if !find {
		return Ticket{}, fmt.Errorf("el ticket con el id ", id, "no existe.")
	}
	return ticket, nil
}

func (b *bookings) Delete(id int) (int, error) {
	var find bool = false
	for _, value := range b.Tickets {
		if value.Id == id {
			find = true
		}
	}
	if !find {
		return 0, fmt.Errorf("el ticket id %d no existe", id)
	}
	return id, nil
}

//Funcion para validar los campos del ticket
func Validation(t Ticket) error {
	if t.Names == "" {
		return errors.New("El campo Nombres es obligatorio")
	}
	if t.Email == "" {
		return errors.New("El campo Email es obligatorio")
	}
	if t.Destination == "" {
		return errors.New("El campo destino es obligatorio")
	}
	if t.Date == "" {
		return errors.New("El campo fecha es obligatorio")
	}
	if t.Price == 0 {
		return errors.New("El campo price es obligatorio")
	}
	return nil
}

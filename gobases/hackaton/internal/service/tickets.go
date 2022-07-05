package service

import (
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

func (b *bookings) Create(t Ticket) (Ticket, error) {

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("Se detectaron varios errores al crear el ticket")
		}
	}()

	_, err := b.Read(t.Id)

	if err != nil {
		panic("Id already exists")
	}

	if t.Id == 0 {
		panic("Id no puede ser 0")
	}

	if t.Names == "" {
		panic("Nombre debe completarse")
	}

	if t.Email == "" {
		panic("Email requerido")
	}

	if t.Destination == "" {
		panic("Destination requerido")
	}

	if t.Date == "" {
		panic("Fecha requerido")
	}

	if t.Price <= 0 {
		panic("Precio debe ser mayor a 0")
	}

	b.Tickets = append(b.Tickets, t)

	return Ticket{t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price}, nil
}

func (b *bookings) Read(id int) (Ticket, error) {

	for i, tickets := range b.Tickets {
		if b.Tickets[i].Id == id {
			return Ticket{tickets.Id, tickets.Names, tickets.Email, tickets.Destination, tickets.Date, tickets.Price}, nil
		}
	}
	return Ticket{}, fmt.Errorf("Ticket not found")
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {

	for i := range b.Tickets {
		if b.Tickets[i].Id == id {
			return Ticket{t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price}, nil
		}
	}

	return Ticket{}, fmt.Errorf("Ticket not found")
}

func (b *bookings) Delete(id int) (int, error) {

	for i := range b.Tickets {
		if b.Tickets[i].Id == id {
			b.Tickets = []Ticket{}
			return id, nil
		}
	}

	return 0, fmt.Errorf("Ticket not found")
}

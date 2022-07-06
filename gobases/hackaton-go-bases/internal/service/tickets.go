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
	// Delete delete a Ticket by id
	GetTickets(cantidad int) []Ticket
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
	// Buscar el ultimo ID en el archivo
	actualID := 0
	for _, v := range b.Tickets {
		if v.Id > actualID {
			actualID = v.Id
		}
	}

	t.Id = actualID + 1

	b.Tickets = append(b.Tickets, t)

	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	found := false
	ticket := Ticket{}
	for _, p := range b.Tickets {
		if p.Id == id {
			found = true
			ticket = p
			break
		}
	}

	if found {
		return ticket, nil
	}
	return Ticket{}, errors.New("no se encontró el ticket")
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	temp := []Ticket{}
	found := false
	for _, v := range b.Tickets {
		if v.Id == id {
			v.Names = t.Names
			v.Email = t.Email
			v.Destination = t.Destination
			v.Date = t.Date
			v.Price = t.Price
			temp = append(temp, v)
			found = true
		} else {
			temp = append(temp, v)
		}
	}

	b.Tickets = temp

	if found {
		return t, nil
	}
	return Ticket{}, errors.New("no existe el ticket a actualizar")
}

func (b *bookings) Delete(id int) (int, error) {
	temp := []Ticket{}
	found := false
	for _, v := range b.Tickets {
		if v.Id != id {
			temp = append(temp, v)
		} else {
			found = true
		}
	}

	b.Tickets = temp

	if found {
		return id, nil
	}
	return 0, errors.New("no existe el ticket a borrar")
}

func (b *bookings) GetTickets(cantidad int) []Ticket {
	if cantidad == 0 {
		return b.Tickets
	} else {
		tickets := make([]Ticket, cantidad)
		index := len(b.Tickets) - cantidad
		fmt.Println(index)
		for i := index; i < len(b.Tickets); i++ {
			tickets[i-index] = b.Tickets[i]
		}
		return tickets
	}
}

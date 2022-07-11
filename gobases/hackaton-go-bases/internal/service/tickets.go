package service

import (
	"fmt"
	"os"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, []Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) ([]Ticket, error)
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

	index := len(b.Tickets) - 1
	new_id := b.Tickets[index].Id + 1

	Ticket := Ticket{
		Id:          new_id,
		Names:       t.Names,
		Email:       t.Email,
		Destination: t.Destination,
		Date:        t.Date,
		Price:       t.Price,
	}

	b.Tickets = append(b.Tickets, Ticket)

	return Ticket, nil
}

func (b *bookings) Read(id int) (Ticket, error) {

	Ticket := Ticket{}
	for _, ticketSearched := range b.Tickets {
		if ticketSearched.Id == id {
			Ticket.Id = ticketSearched.Id
			Ticket.Names = ticketSearched.Names
			Ticket.Email = ticketSearched.Email
			Ticket.Destination = ticketSearched.Destination
			Ticket.Date = ticketSearched.Date
			Ticket.Price = ticketSearched.Price
		}
	}

	if Ticket.Id == 0 {
		fmt.Println("error: no se ha encontralado el emento con id", id)
		os.Exit(1)
	}
	return Ticket, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, []Ticket, error) {

	//ticketUpdated, _ := b.Read(id)
	var ticketUpdated Ticket

	for i, value := range b.Tickets {
		if value.Id == id {
			b.Tickets[i].Names = t.Names
			b.Tickets[i].Email = t.Email
			b.Tickets[i].Destination = t.Destination
			b.Tickets[i].Date = t.Date
			b.Tickets[i].Price = t.Price
			ticketUpdated = b.Tickets[i]
		}
	}

	return ticketUpdated, b.Tickets, nil
}

func (b *bookings) Delete(id int) ([]Ticket, error) {
	var index int
	for i, ticketDeleted := range b.Tickets {
		if ticketDeleted.Id == id {
			index = i
		}
	}
	b.Tickets = append(b.Tickets[:index], b.Tickets[index+1:]...)
	return b.Tickets, nil
}

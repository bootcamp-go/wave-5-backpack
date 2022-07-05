package bookings

import (
	"errors"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/tickets"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t tickets.Ticket) (tickets.Ticket, error)
	// Read read a Ticket by id
	Read(id int) (tickets.Ticket, error)
	// Update update values of a Ticket
	Update(id int, t tickets.Ticket) (tickets.Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []tickets.Ticket
	csv     file.File
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []tickets.Ticket, csv file.File) Bookings {
	return &bookings{Tickets: Tickets, csv: csv}
}

func (b *bookings) Create(t tickets.Ticket) (tickets.Ticket, error) {
	b.Tickets = append(b.Tickets, t)
	b.csv.Write(b.Tickets)
	return t, nil
}

func (b *bookings) Read(id int) (tickets.Ticket, error) {
	var ticketFound tickets.Ticket
	for _, ticket := range b.Tickets {
		if ticket.Id == id {
			ticketFound = ticket
		}
	}

	if ticketFound.Names == "" {
		return ticketFound, errors.New("ticket not found")
	}

	return ticketFound, nil
}

func (b *bookings) Update(id int, t tickets.Ticket) (tickets.Ticket, error) {
	for i, ticket := range b.Tickets {
		if ticket.Id == id {
			b.Tickets[i] = t
		}
	}
	b.csv.Write(b.Tickets)
	return t, nil
}

func (b *bookings) Delete(id int) (int, error) {
	var pos int
	for i, ticket := range b.Tickets {
		if ticket.Id == id {
			pos = i
		}
	}

	if pos == 0 {
		return 0, errors.New("no se encuentra el registro")
	}
	b.Tickets = append(b.Tickets[:pos-1], b.Tickets[pos+1:]...)
	b.csv.Write(b.Tickets)
	return id, nil
}

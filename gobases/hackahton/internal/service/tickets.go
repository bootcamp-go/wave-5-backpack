package service

import (
	"errors"
	"fmt"
	"math/rand"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	getTicketById(id int64) (Ticket, int)
	// Get all Tickets
	ReadAll() ([]Ticket, error)
	// Update update values of a Ticket
	Update(id int64, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int64) (int, error)
}

type bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id, Price                       int64
	Names, Email, Destination, Date string
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {

	return &bookings{Tickets: Tickets}
}
func (b *bookings) Create(t Ticket) (Ticket, error) {
	defer func() {
		errRecovered := recover()
		if errRecovered != nil {
			fmt.Println(errRecovered)
		}
	}()

	id, err := generateId(b)
	if err != nil {
		panic(err)
	}
	t.Id = id
	b.Tickets = append(b.Tickets, t)
	return t, nil
}

func CreateTicket(id, price int64, names, email, destination, date string) (Ticket, error) {
	return Ticket{Id: id, Names: names,
		Email: email, Destination: destination,
		Date: date}, nil
}

func (b *bookings) getTicketById(id int64) (Ticket, int) {
	var ticketToFind Ticket
	var pos int
	for key, ticket := range b.Tickets {
		if ticket.Id == id {
			ticketToFind = ticket
			pos = key
		}
	}

	return ticketToFind, pos

}
func (b *bookings) ReadAll() ([]Ticket, error) {
	if len(b.Tickets) == 0 {
		return nil, errors.New("No hay tickets registrados.")
	} else {
		return b.Tickets, nil
	}
}

func (b *bookings) Update(id int64, t Ticket) (Ticket, error) {
	ticketToEdit, pos := b.getTicketById(id)

	if (ticketToEdit == Ticket{}) {
		return Ticket{}, errors.New("No existe ningún ticket con el ID especificado.")
	}

	ticketToEdit = t
	ticketToEdit.Id = id
	b.Tickets[pos] = ticketToEdit
	return ticketToEdit, nil
}

func (b *bookings) Delete(id int64) (int, error) {
	ticketToEdit, pos := b.getTicketById(id)

	if (ticketToEdit == Ticket{}) {
		return 1, errors.New("No existe ningún ticket con el ID especificado.")
	}
	newTickets := []Ticket{}

	for key, ticket := range b.Tickets {
		if key != pos {
			newTickets = append(newTickets, ticket)
		}
	}
	b.Tickets = newTickets
	return 0, nil
}

func generateId(b *bookings) (int64, error) {
	id := int64(rand.Intn(100000))
	previousTicket, _ := b.getTicketById(id)

	if (Ticket{} == previousTicket) {
		return id, nil
	} else {
		return 0, errors.New("Conflicto con ID, por favor, intente de nuevo.")
	}
}

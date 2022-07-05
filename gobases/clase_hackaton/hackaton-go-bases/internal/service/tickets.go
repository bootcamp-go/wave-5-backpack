package service

import (
	"errors"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/interfaces"
)

const PATH = "./tickets.csv"

type Bookings interface {
	// Create create a new Ticket
	Create(t interfaces.Ticket) (interfaces.Ticket, error)
	// Read read a Ticket by id
	Read(id int) (interfaces.Ticket, error)
	// Update update values of a Ticket
	Update(t interfaces.Ticket) (interfaces.Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
}

type bookings struct {
	Tickets []interfaces.Ticket
}

// NewBookings creates a new bookings service
func NewBookings() Bookings {
	file := file.File{Path: PATH}
	tickets, err := file.Read()

	if err != nil {
		panic("Error al leer el archivo")
	}
	return &bookings{Tickets: tickets}
}

func (b *bookings) Create(t interfaces.Ticket) (interfaces.Ticket, error) {
	nuevoTicket := interfaces.Ticket{
		Id:          generateId(b),
		Names:       t.Names,
		Email:       t.Email,
		Destination: t.Destination,
		Date:        t.Date,
		Price:       t.Price,
	}
	file := file.File{Path: PATH}

	err := file.Write(nuevoTicket)
	return nuevoTicket, err
}

func (b *bookings) Read(id int) (interfaces.Ticket, error) {
	ticket, err := b.findById(id)

	if err != nil {
		return ticket, err
	}
	return ticket, nil
}

func (b *bookings) Update(t interfaces.Ticket) (interfaces.Ticket, error) {
	file := file.File{Path: PATH}

	err := file.Update(t)
	if err != nil {
		return t, err
	}
	return t, nil
}

func (b *bookings) Delete(id int) (int, error) {
	file := file.File{Path: PATH}

	err := file.Remove(id)
	if err != nil {
		return 0, err
	}
	return 0, err
}

func (b *bookings) findById(id int) (interfaces.Ticket, error) {
	for _, ticket := range b.Tickets {
		if ticket.Id == id {
			return ticket, nil
		}
	}
	return interfaces.Ticket{}, errors.New("No se encontró el id")
}

func generateId(b *bookings) int {
	lastId := b.Tickets[len(b.Tickets)-1].Id
	return lastId + 1
}

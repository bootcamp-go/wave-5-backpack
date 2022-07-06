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

// Este punto hay que corregirlo para que devuelva Bokings y no bookings
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	err := Validation(t)
	if err != nil {
		return Ticket{}, err
	}

	t.Id = len(b.Tickets) + 1
	b.Tickets = append(b.Tickets, t)
	return t, nil
}


func (b *bookings) Read(id int) (Ticket, error) {
	for _, ticket := range b.Tickets{
		if ticket.Id == id{
			return ticket,nil
		}
	}
	err := fmt.Errorf("error: no existe el id: %d", id)
	return Ticket{}, err
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	err := Validation(t)
	if err != nil {
		return Ticket{}, err
	}
	
	var find bool = false
	var ticket Ticket
	for _, values :=range b.Tickets{ 
		if values.Id == id{
			ticket = t
			find = true
			break
		}
	}

	if !find {
		return Ticket{}, fmt.Errorf("No existe el id: %d", id)
	}

	return ticket, nil
}

func (b *bookings) Delete(id int) (int, error) {

	var find bool = false
	for i, values :=range b.Tickets{ 
		if values.Id == id{
			find = true
			b.Tickets = append(b.Tickets[:i], b.Tickets[i+1:]...) // REVISAR DONDE LO PONE BOOTCAMP
			break
		}
	}

	if !find {
		return 0, fmt.Errorf("el ticket id %d no existe", id)
	}

	return id, nil

}

func Validation(t Ticket) error {
	if t.Date == "" {
		return errors.New("el campo data es requerido")
	}

	if t.Destination == "" {
		return errors.New("el campo destino es requerido")
	}

	if t.Email == "" {
		return errors.New("el campo email es requerido")
	}

	if t.Names == "" {
		return errors.New("el campo nombres es requerido")
	}

	if t.Price == 0 {
		return errors.New("el campo precio es requerido")
	}

	return nil
}

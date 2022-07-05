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

	ReadAll() []Ticket
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

func validateTicket(t Ticket) (bool, error) {
	if t.Id == 0 {
		return false, errors.New("Ticket id cant be 0")
	}
	if t.Names == "" {
		return false, errors.New("Ticket name cant be empty")
	}
	if t.Email == "" {
		return false, errors.New("Ticket email cant be empty")
	}
	if t.Destination == "" {
		return false, errors.New("Ticket destination cant be empty")
	}
	if t.Date == "" {
		return false, errors.New("Ticket date cant be empty")
	}
	if t.Price == 0 {
		return false, errors.New("Ticket price cant be 0")
	}
	return true, nil
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	valid, err := validateTicket(t)
	if !valid {
		return Ticket{}, err
	}
	_, err = b.Read(t.Id)
	if err == nil {
		return t, fmt.Errorf("Ticket id: %d already exists", t.Id)
	}
	b.Tickets = append(b.Tickets, t)
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	for i := 0; i < len(b.Tickets); i++ {
		if b.Tickets[i].Id == id {
			return b.Tickets[i], nil
		}
	}
	return Ticket{}, fmt.Errorf("Ticket id: %d doesnt exists", id)

}

func (b *bookings) ReadAll() []Ticket {
	return b.Tickets
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	valid, err := validateTicket(t)
	if !valid {
		return Ticket{}, err
	}
	for i := 0; i < len(b.Tickets); i++ {
		if b.Tickets[i].Id == id {
			b.Tickets[i] = t
			return t, nil
		}
	}
	return Ticket{}, fmt.Errorf("Ticket id: %d doesnt exists", id)
}

func (b *bookings) Delete(id int) (int, error) {
	for i := 0; i < len(b.Tickets); i++ {
		if b.Tickets[i].Id == id {
			b.Tickets = append(b.Tickets[:i], b.Tickets[i+1:]...)
			return id, nil
		}
	}
	return 0, fmt.Errorf("Ticket id: %d doesnt exists", id)
}

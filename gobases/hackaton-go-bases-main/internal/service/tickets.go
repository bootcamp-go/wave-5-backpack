package service

import "errors"

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
	// Traer todos los Tickets
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

func (b *bookings) Create(t Ticket) (Ticket, error) {
	b.Tickets = append(b.Tickets, t)
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	for _, t := range b.Tickets {
		if t.Id == id {
			return t, nil
		}
	}
	return Ticket{}, errors.New("el ticket no fue en contrado")
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	for i, tem := range b.Tickets {
		if tem.Id == id {
			b.Tickets[i] = t
			return t, nil
		}
	}
	return Ticket{}, errors.New("el ticket no fue en contrado")
}

func (b *bookings) Delete(id int) (int, error) {
	for i, t := range b.Tickets {
		if t.Id == id {
			b.Tickets = append(b.Tickets[:i], b.Tickets[i+1:]...)
			return 0, nil
		}
	}
	return -1, errors.New("el ticket no fue en contrado")
}

func (b *bookings) ReadAll() []Ticket {
	return b.Tickets
}

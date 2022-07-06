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
	return Ticket{}, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	t := Ticket{}
	found := false
	for _, item := range b.Tickets {
		if item.Id == id {
			t = item
			found = true
		}
	}
	if !found {
		return Ticket{}, errors.New("no existe el dato")
	}
	return t, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	found := false
	for _, item := range b.Tickets {
		if item.Id == id {
			item = t
			found = true
		}
	}
	if !found {
		return Ticket{}, errors.New("no existe el dato")
	}
	return t, nil
}

func (b *bookings) Delete(id int) (int, error) {
	found := false
	for _, item := range b.Tickets {
		if item.Id == id {

			found = true
		}
	}
	if !found {
		return 0, errors.New("no existe el dato")
	}

	return 0, nil
}

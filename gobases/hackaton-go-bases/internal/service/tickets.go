package service

import (
	"errors"
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
	Id          int    `json:"id"`
	Names       string `json:"names"`
	Email       string `json:"email"`
	Destination string `json:"destination"`
	Date        string `json:"date"`
	Price       int    `json:"price"`
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	if t.Names == "" {
		return Ticket{}, errors.New("el campo names no puede estar vacio")
	}
	if t.Email == "" {
		return Ticket{}, errors.New("el campo email no puede estar vacio")
	}
	if t.Destination == "" {
		return Ticket{}, errors.New("el campo destination no puede estar vacio")
	}
	if t.Date == "" {
		return Ticket{}, errors.New("el campo date no puede estar vacio")
	}
	if t.Price == 0 {
		return Ticket{}, errors.New("el campo price no puede estar vacio")
	}

	new_id := b.Tickets[len(b.Tickets)-1].Id + 1
	t.Id = new_id

	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	for _, v := range b.Tickets {
		if v.Id == id {
			return v, nil
		}
	}

	return Ticket{}, errors.New("no se encontro un ticket con el id solicitado")
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	var to_update Ticket

	for _, v := range b.Tickets {
		if v.Id == id {
			to_update.Id = t.Id

			if t.Names != "" && v.Names != t.Names {
				to_update.Names = t.Names
			} else {
				to_update.Names = v.Names
			}
			if t.Email != "" && v.Email != t.Email {
				to_update.Email = t.Email
			} else {
				to_update.Email = v.Email
			}
			if t.Destination != "" && v.Destination != t.Destination {
				to_update.Destination = t.Destination
			} else {
				to_update.Destination = v.Destination
			}
			if t.Date != "" && v.Date != t.Date {
				to_update.Date = t.Date
			} else {
				to_update.Date = v.Date
			}
			if t.Price != 0 && v.Price != t.Price {
				to_update.Price = t.Price
			} else {
				to_update.Price = v.Price
			}

			return to_update, nil
		}
	}

	return Ticket{}, errors.New("no se encontro un ticket con el id solicitado")
}

func (b *bookings) Delete(id int) (int, error) {
	for _, v := range b.Tickets {
		if v.Id == id {
			return v.Id, nil
		}
	}

	return 0, errors.New("no se encontro un ticket con el id solicitado")
}

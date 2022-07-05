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
	Deleted                         bool
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	res := Ticket{Id: b.Tickets[len(b.Tickets)-1].Id + 1, Names: t.Names, Email: t.Email, Destination: t.Destination, Date: t.Date, Price: t.Price, Deleted: false}
	b.Tickets = append((*b).Tickets, res)
	return res, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	if len(b.Tickets) < id {
		err := errors.New("no hay ticket para ese codigo ID")
		panic("No hay ticket para ese codigo ID")
		return Ticket{}, err
	}
	if b.Tickets[id-1].Deleted {
		err := errors.New("ese ticket fue eliminado")
		panic("Ese ticket fue eliminado")
		return Ticket{}, err
	}
	res := b.Tickets[id-1]
	return res, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	tkt, _ := b.Read(id)
	b.Tickets[id-1] = tkt
	return tkt, nil
}

func (b *bookings) Delete(id int) (int, error) {
	_, err := b.Read(id)
	if err != nil {
		return 404, nil
	}
	b.Tickets[id-1].Deleted = true
	return 0, nil
}

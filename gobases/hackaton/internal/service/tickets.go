package service

import "errors"

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) error
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)
	Details() []Ticket
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

func (b *bookings) Create(t Ticket) (err error) {
	canAdd := true
	for _, value := range b.Tickets {
		if value.Id == t.Id {
			canAdd = false
			err = errors.New("error: ya existe un ticket con dicho id")
		}
	}
	if canAdd {
		b.Tickets = append(b.Tickets, t)
	}
	return err
}

func (b *bookings) Read(id int) (Ticket, error) {
	for _, value := range b.Tickets {
		if value.Id == id {
			return value, nil
		}
	}
	return Ticket{}, errors.New("error: no se ha encontrado el id especificado")
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	// no usar range porque hace copia del valor del array
	for i := 0; i < len(b.Tickets)-1; i++ {
		ticket := &b.Tickets[i]
		if ticket.Id == id {
			ticket.Names = t.Names
			ticket.Email = t.Email
			ticket.Destination = t.Destination
			ticket.Date = t.Date
			ticket.Price = t.Price
			return *ticket, nil
		}
	}
	return Ticket{}, errors.New("error: no se encontro el ticket con el id especificado para realizar update")
}

func (b *bookings) Delete(id int) (int, error) {
	for i, value := range b.Tickets {
		if value.Id == id {
			b.Tickets = append(b.Tickets[:i], b.Tickets[i+1:]...)
			return id, nil
		}
	}
	return 0, errors.New("error: no se encontro el ticket con el id especificado para realizar el borrado")
}

func (b *bookings) Details() []Ticket {
	return b.Tickets
}

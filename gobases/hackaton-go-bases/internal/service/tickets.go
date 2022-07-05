package service

import (
	"errors"
	"strconv"
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
	// Get entire list of Tickets
	Get() []Ticket
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

// Agrega un ticket generando el ID automáticamente, para esto
// comprueba que los datos no sean vacíos y asigna un ID tomando
// el ultimo ID registrado y sumándole 1. Si bookings está vacío
// el ID se inicializa en 1
func (b *bookings) Create(t Ticket) (Ticket, error) {
	if t.Names == "" || t.Email == "" || t.Destination == "" || t.Date == "" || t.Price <= 0 {
		return Ticket{}, errors.New("error: debe llenar todos los datos del ticket")
	} else {
		id := 1
		if len(b.Tickets) > 0 {
			id = b.Tickets[len(b.Tickets)-1].Id + 1
		}
		t.Id = id
		b.Tickets = append(b.Tickets, t)
		return t, nil
	}
}

// Encuentra un ticket por su ID, para esto el ID no debe ser
// negativo ni 0, si no encuentra el id solicitado arroja un error
func (b *bookings) Read(id int) (Ticket, error) {
	if id > 0 {
		for _, ticket := range b.Tickets {
			if ticket.Id == id {
				return b.Tickets[id-1], nil
			}
		}
	}

	return Ticket{}, errors.New("Ticket con id " + strconv.Itoa(id) + " no encontrado")
}

// Actualiza un ticket de la lista, para esto ninguno de los datos
// del ticket debe estar vacío y el id debe existir en la lista
func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	if t.Names == "" || t.Email == "" || t.Destination == "" || t.Date == "" || t.Price <= 0 {
		return Ticket{}, errors.New("error: debe llenar todos los datos del ticket")
	}

	for _, ticket := range b.Tickets {
		if ticket.Id == id {
			newTicket := Ticket{
				Id:          id,
				Names:       t.Names,
				Email:       t.Email,
				Destination: t.Destination,
				Date:        t.Date,
				Price:       t.Price,
			}
			b.Tickets[id-1] = newTicket
			return newTicket, nil
		}
	}
	return Ticket{}, errors.New("Ticket con id " + strconv.Itoa(id) + " no encontrado, no fue posible actualizar")
}

// Elimina un ticket de la lista, para esto el id no debe ser
// negativo o 0 ni ser mayor al ultimo id de la lista
func (b *bookings) Delete(id int) (int, error) {
	if id > 0 && id <= b.Tickets[len(b.Tickets)-1].Id {
		for idx, ticket := range b.Tickets {
			if ticket.Id == id {
				b.Tickets = append(b.Tickets[:idx], b.Tickets[idx+1:]...)
				return id, nil
			}
		}
	}
	return id, errors.New("error: El id ingresado no existe o no es valido")

}

// Devuelve la lista de Tickets
func (b *bookings) Get() []Ticket {
	return b.Tickets
}

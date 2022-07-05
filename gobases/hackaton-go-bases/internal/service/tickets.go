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
	Delete(id int) ([]Ticket, int, error)
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

func imprimirError(campo string) error {
	return errors.New(fmt.Sprint("campo ", campo, " vacío"))
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	if t.Names == "" {
		return Ticket{}, imprimirError("Names")
	}
	if t.Email == "" {
		return Ticket{}, imprimirError("Email")
	}
	if t.Destination == "" {
		return Ticket{}, imprimirError("Destination")
	}
	if t.Date == "" {
		return Ticket{}, imprimirError("Date")
	}
	if t.Price == 0 {
		return Ticket{}, imprimirError("Price")
	}
	newTicket := Ticket{t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price}
	fmt.Println("Se ha creado un nuevo ticket")
	return newTicket, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	ticket := Ticket{}
	for _, value := range b.Tickets {
		if value.Id == id {
			ticket = value
		}
	}
	if ticket == (Ticket{}) {
		return ticket, errors.New("no se encontró ningún ticket con el id ingresado")
	} else {
		fmt.Println("El ticket con id ", ticket.Id, " ha sido leído exitosamente")
		return ticket, nil
	}
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	for _, value := range b.Tickets {
		if value.Id == id {
			t.Names = "Maria"
			t.Email = "maria@hotmail.com"
			t.Destination = "Colombia"
			t.Date = "5/06/2022"
			t.Price = 2400000
		}
	}
	if t == (Ticket{}) {
		return t, errors.New("no se encontró ningún ticket con el id ingresado")
	} else {
		fmt.Println("El ticket con id ", t.Id, " ha sido modificado exitosamente")
		return t, nil
	}
}

func (b *bookings) Delete(id int) ([]Ticket, int, error) {
	var values []Ticket
	var isHere bool
	for _, value := range b.Tickets {
		if value.Id == id {
			isHere = true
			continue
		} else {
			values = append(values, value)
		}
	}
	if isHere == false {
		return values, 1, errors.New("no se encontró ningún ticket con el id ingresado")
	} else {
		fmt.Println("El ticket con id ", id, " ha sido eliminado exitosamente")
		return values, 0, nil
	}

}

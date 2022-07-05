package service

import (
	"errors"
	"fmt"
)

type iBookings interface {
	// Create create a new Ticket
	Create(t Ticket) error
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) error
	// Delete delete a Ticket by id
	Delete(id int) error
}

type Bookings struct {
	Tickets []Ticket
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

// NewBookings creates a new bookings service
func (b *Bookings) NewBookings(ts []Ticket) error {
	if len(ts) == 0 {
		return errors.New("No hay elementos para agregar.")
	} else {
		b.Tickets = ts
		return nil
	}
}

func (b *Bookings) Create(t *Ticket) error {
	mayor := b.Tickets[0].Id

	for _, val := range b.Tickets {
		if val.Id > mayor {
			mayor = val.Id
		}
	}
	t.Id = mayor + 1

	b.Tickets = append(b.Tickets, *t)

	return nil
}

func (b *Bookings) Read(id int) (Ticket, error) {
	var t Ticket

	for _, val := range b.Tickets {
		if val.Id == id {
			t = val
			return t, nil
		}
	}

	return Ticket{}, errors.New("No se encontró ningun ticket con ese ID.")
}

func (b *Bookings) Update(id int, t Ticket) error {
	for i, val := range b.Tickets {
		if val.Id == id {
			b.Tickets[i] = Ticket{Id: val.Id, Names: t.Names, Email: t.Email, Destination: t.Destination, Date: t.Date, Price: t.Price}
			return nil
		}
	}

	return errors.New("No se encontró el ticket o hubo algun problema.")
}

func (b *Bookings) Delete(id int) error {
	for i, val := range b.Tickets {
		if val.Id == id {
			b.Tickets = append(b.Tickets[:i], b.Tickets[i+1:]...)

			return nil
		}
	}

	return errors.New("No se encontró el ticket o hubo algun problema.")
}

func (b *Bookings) Imprimir(err error, operacion string) {
	fmt.Println("\n------------------------------------------------------------------------------")
	fmt.Println("\t\t\t\t", operacion)
	fmt.Println("------------------------------------------------------------------------------")

	if err != nil {
		fmt.Println(err)
	} else {
		b.print()
	}
}

func (b *Bookings) print() {
	for _, val := range b.Tickets {
		fmt.Printf("Id: %d | Names: %s | Email: %s | Date: %s | Destination: %s | Price: %d\n",
			val.Id, val.Names, val.Email, val.Date, val.Destination, val.Price)
	}
}

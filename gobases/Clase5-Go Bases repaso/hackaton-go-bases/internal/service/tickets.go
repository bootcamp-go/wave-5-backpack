package service

import (
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

func NewTicket(id int, names string, email string, destination string, date string, price int) Ticket {
	tick := Ticket{
		Id:          id,
		Names:       names,
		Email:       email,
		Destination: destination,
		Date:        date,
		Price:       price,
	}
	return tick
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	_, err := (*b).Read(t.Id)
	if err == nil {
		return Ticket{}, fmt.Errorf("No se pudo crear el ticket porque ya existe uno con el mismo Id: %v", t.Id)
	}
	(*b).Tickets = append((*b).Tickets, t)
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	for _, tick := range *&b.Tickets {
		if tick.Id == id {
			return tick, nil
		}
	}
	return Ticket{}, fmt.Errorf("No se encontró ningun ticket con Id:%v", id)
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	_, err := (*b).Read(t.Id)
	if err != nil {
		return Ticket{}, err
	}
	listaTickets := (*b).Tickets
	index := buscarIndice(listaTickets, id)
	listaTickets[index] = t
	return t, nil
}

func (b *bookings) Delete(id int) (int, error) {
	_, err := (*b).Read(id)
	if err != nil {
		return 0, err
	}
	listaTickets := (*b).Tickets
	index := buscarIndice(listaTickets, id)
	listaTickets = append(listaTickets[:index], listaTickets[index+1:]...)
	//listaTickets[index] = listaTickets[len(listaTickets)-1]
	//prueba := listaTickets[:len(listaTickets)-2]
	//(*b).Tickets = prueba
	fmt.Println(listaTickets)
	return id, nil
}

func buscarIndice(lista []Ticket, id int) (index int) {
	var cont int

	for _, tick := range lista {
		if tick.Id != id {
			cont++
		} else {
			break
		}
		//fmt.Println(cont)
	}
	return cont
}

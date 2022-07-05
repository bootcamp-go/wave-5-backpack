package service

import (
	"errors"
	"fmt"
	"net/mail"
	"strconv"
	"strings"
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
	// Data
	ReadAll() ([]Ticket, error)
}

type bookings struct {
	Tickets []Ticket
}

type ErrorTicketNotFound struct {
	idTicket int
}

func (err ErrorTicketNotFound) Error() string {
	return fmt.Sprintf("error: ticket no encontrado con id %d", err.idTicket)
}

type Ticket struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

func MinStringLength(data string, length int) error {
	if len(data) <= length {
		return fmt.Errorf("error: nombre muy corto por lo menos %d caracteres", length)
	}
	return nil
}

func ValidateEmailAddress(address string) error {
	_, err := mail.ParseAddress(address)
	return err
}

func ValidateNotNegative(num int) error {
	if num < 0 {
		return errors.New("error: negative value")
	}
	return nil
}

func removeNilErrors(errs []error) []error {
	errsFiltered := make([]error, 0)
	for _, err := range errs {
		if err != nil {
			errsFiltered = append(errsFiltered, err)
		}
	}
	return errsFiltered
}

func ValidateTime(da string) error {
	result := strings.Split(da, ":")
	if len(result) != 2 {
		return errors.New("error: la fecha no esta formateada correctamente ejemplo 23:59")
	}

	hours, err := strconv.Atoi(result[0])
	if err != nil {
		return errors.New("error: no se pudo obtener la hora")
	}
	if hours > 23 || hours < 0 {
		return errors.New("error: la hora esta fuera de rango 00:00 a 23:00")
	}
	mins, err := strconv.Atoi(result[1])
	if err != nil {
		return errors.New("error: no se pudo obtener la hora")
	}
	if mins > 60 || hours < 0 {
		return errors.New("error: los minutos estan fuera de rango entre 0 a 60")
	}

	return nil

}

func mergeErrors(errs []error) error {
	errsMessages := ""
	if len(errs) > 1 {
		errsMessages = "error: multiple validation error \n"
	}
	for _, err := range errs {
		errsMessages += err.Error() + "\n"
	}
	return errors.New(errsMessages)
}

func ValidateTicket(t Ticket) error {
	errors := make([]error, 0)
	errors = append(errors,
		MinStringLength(t.Names, 3),
		MinStringLength(t.Destination, 3),
		ValidateNotNegative(t.Price),
		ValidateEmailAddress(t.Email),
		ValidateTime(t.Date),
	)
	errors = removeNilErrors(errors)
	if len(errors) > 0 {
		err := mergeErrors(errors)
		return err
	}
	return nil
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	t.Id = b.Tickets[len(b.Tickets)-1].Id + 1
	err := ValidateTicket(t)
	if err != nil {
		return Ticket{}, err
	}
	b.Tickets = append(b.Tickets, t)
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	for _, ticket := range b.Tickets {
		if ticket.Id == id {
			return ticket, nil
		}
	}
	return Ticket{}, &ErrorTicketNotFound{idTicket: id}
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	for index, ticket := range b.Tickets {
		if ticket.Id == id {
			t.Id = id
			err := ValidateTicket(t)
			if err != nil {
				return Ticket{}, err
			}
			b.Tickets[index] = t
			return t, nil
		}
	}
	return Ticket{}, &ErrorTicketNotFound{idTicket: id}
}

func (b *bookings) Delete(id int) (int, error) {
	for index, ticket := range b.Tickets {
		if ticket.Id == id {
			if len(b.Tickets) == index+1 {
				b.Tickets = b.Tickets[:index]
				return id, nil
			}
			b.Tickets = append(b.Tickets[:index], b.Tickets[index+1:]...)
			return id, nil
		}
	}
	return 0, &ErrorTicketNotFound{idTicket: id}
}

func (b *bookings) ReadAll() ([]Ticket, error) {
	return b.Tickets, nil
}

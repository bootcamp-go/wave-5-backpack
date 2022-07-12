package service

import (
	// "github.com/bootcamp-go/wave-5-backpack/tree/tapia_luis/gobases/Hackaton/internal/file"
	"fmt"
	"os"
	// "strings"
	// "strconv"
	"errors"
)

const(
	create = "create" 
	read = "read"
	update = "update"
	delete = "delete"
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

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	return &bookings{Tickets: Tickets}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	i:=len(b.Tickets) - 1 
	t.Id = b.Tickets[i].Id + 1
	b.Tickets = append(b.Tickets, t)
	//write
	e := Write(t,"create")
	if e != nil{
		return Ticket{},fmt.Errorf("Ah ocurrido un error: %v",e)
	}
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	result := Ticket{}
	for _,tic := range b.Tickets{
		if tic.Id == id {
			result = tic
		}else{
			return result, fmt.Errorf("No existe un ticket con el Id %d",id)
		}
	}
	return result, nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	result := Ticket{}
	for _,tic := range b.Tickets{
		if tic.Id == id {
			//write
		}else{
			return result, fmt.Errorf("No existe un ticket con el Id %d",id)
		}
	}
	return result, nil
}

func (b *bookings) Delete(id int) (int, error) {
	return 0, nil
}
func Write(t Ticket,wType string) error {
	path := "./tickets.csv"
	data, err := os.ReadFile(path)
	if err != nil{
		return fmt.Errorf("Ah ocurrido un error al leer el archivo: %v",err)
	}
	strData := string(data)
	switch wType{
		case create:
			newTicker := fmt.Sprintf("\n%d,%s,%s,%s,%s,%d",t.Id,t.Names,t.Email,t.Destination,t.Date,t.Price)
			strData += newTicker
			newBytes := []byte(strData)
			e := os.WriteFile(path,newBytes,0644)
			if e != nil{
				return fmt.Errorf("Ah ocurrido un error al escribir el archivo: %v",err)
			}
			return nil
		case update:
			return nil
		case delete:
			return nil
	}
	
	return errors.New("No se a seleccionado un tipo de escritura valido")
}

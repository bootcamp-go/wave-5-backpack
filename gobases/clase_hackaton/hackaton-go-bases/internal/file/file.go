package file

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bootcamp-go/hackaton-go-bases/internal/interfaces"
)

const PATH = "./tickets.csv"

type File struct {
	Path string
}

func (f *File) Read() ([]interfaces.Ticket, error) {
	tickets, err := leerYParsear(f.Path)
	return tickets, err
}

func (f *File) Write(ticket interfaces.Ticket) error {
	tickets, err := leerYParsear(f.Path)
	tickets = append(tickets, ticket)
	guardarArchivo(tickets)
	return err
}

func (f *File) Update(ticket interfaces.Ticket) error {
	tickets, err := leerYParsear(f.Path)
	if err != nil {
		return err
	}
	var indexUpdate int = 0
	for index, t := range tickets {
		if t.Id == ticket.Id {
			indexUpdate = index
			break
		}
	}
	tickets[indexUpdate].Names = ticket.Names
	tickets[indexUpdate].Email = ticket.Email
	tickets[indexUpdate].Date = ticket.Date
	tickets[indexUpdate].Destination = ticket.Destination
	tickets[indexUpdate].Price = ticket.Price

	guardarArchivo(tickets)

	return nil
}

func (f *File) Remove(id int) error {
	tickets, err := leerYParsear(f.Path)
	if err != nil {
		return err
	}
	var indexAEliminar int = 0

	for index, ticket := range tickets {
		if ticket.Id == id {
			indexAEliminar = index
		}
	}
	tickets = remove(tickets, indexAEliminar)

	guardarArchivo(tickets)

	return nil
}

func leerYParsear(path string) ([]interfaces.Ticket, error) {
	file, err := openFile(path)
	if err != nil {
		return nil, err
	}

	tickets, err := parseBytesToTickets(file)
	if err != nil {
		return nil, err
	}

	return tickets, err
}

func openFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}
	return file, nil
}

func parseBytesToTickets(bytes []byte) ([]interfaces.Ticket, error) {
	registros := strings.Split(string(bytes), "\n")

	var tickets []interfaces.Ticket = []interfaces.Ticket{}

	for _, registro := range registros {
		columnas := strings.Split(registro, ",")
		if len(columnas) < 6 {
			continue
		}

		id, err1 := strconv.Atoi(columnas[0])
		price, err2 := strconv.Atoi(columnas[5])

		if err1 != nil || err2 != nil {
			return nil, errors.New("Error al parsear el archivo")
		}

		var ticket interfaces.Ticket = interfaces.Ticket{
			Id:          id,
			Names:       columnas[1],
			Email:       columnas[2],
			Destination: columnas[3],
			Date:        columnas[4],
			Price:       price,
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func guardarArchivo(tickets []interfaces.Ticket) error {
	var ticketsConcatenados string

	for _, ticket := range tickets {
		ticketsConcatenados = ticketsConcatenados + ticketToString(ticket)
	}
	registrosBytes := []byte(ticketsConcatenados)
	err := os.WriteFile(PATH, registrosBytes, 0644)
	return err
}

func ticketToString(ticket interfaces.Ticket) string {
	registro := fmt.Sprintf("%v,%s,%s,%s,%s,%v\n", ticket.Id, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticket.Price)
	return registro
}

func remove(slice []interfaces.Ticket, s int) []interfaces.Ticket {
	return append(slice[:s], slice[s+1:]...)
}

package file

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

const SEPARADOR = ","

type File struct {
	Path string
}

func TicketCsvToTicket(csvValues []string) (service.Ticket, error) {
	ticket := service.Ticket{}
	if len(csvValues) == 6 {
		id, err := strconv.Atoi(csvValues[0])
		if err != nil {
			return ticket, err
		}
		ticket.Id = id
		ticket.Names, ticket.Email, ticket.Destination, ticket.Date = csvValues[1], csvValues[2], csvValues[3], csvValues[4]
		precio, err := strconv.Atoi(csvValues[5])
		ticket.Price = precio
		if err != nil {
			return ticket, err
		}
		return ticket, nil
	}
	return ticket, errors.New("error: la lista de valores debe ser igual a 6")
}

func TicketToCSV(ticket service.Ticket) string {
	return fmt.Sprintf("%d,%s,%s,%s,%s,%d", ticket.Id, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticket.Price)
}

func (f *File) Read() ([]service.Ticket, error) {

	data, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}
	csvProduct := strings.Split(string(data), "\n")

	tickets := []service.Ticket{}
	for _, line := range csvProduct {
		valores := strings.Split(line, SEPARADOR)
		ticket, err := TicketCsvToTicket(valores)
		if err != nil {
			fmt.Println("error en convertir", err)
			continue
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil

}

func (f *File) WriteAll(tickets []service.Ticket) error {
	fileData := ""
	for index, ticket := range tickets {
		if index+1 == len(tickets) {
			fileData += TicketToCSV(ticket)
			continue
		}
		fileData += fmt.Sprint(TicketToCSV(ticket), "\n")
	}
	return os.WriteFile(f.Path, []byte(fileData), 0644)
}

func (f *File) Write(ticket service.Ticket) error {
	csvFile, err := os.OpenFile(f.Path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer csvFile.Close()
	_, err = csvFile.WriteString("\n" + TicketToCSV(ticket))
	return err
}

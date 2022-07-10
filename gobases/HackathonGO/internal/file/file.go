package file

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {
	tickets := []service.Ticket{}

	data, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}

	csvFile := csv.NewReader(bytes.NewBuffer(data))

	for {
		line, err := csvFile.Read()
		if err != nil {
			break
		}

		id, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, err
		}
		price, err := strconv.Atoi(line[5])
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, service.Ticket{
			Id:          id,
			Names:       line[1],
			Email:       line[2],
			Destination: line[3],
			Date:        line[4],
			Price:       price,
		})
	}
	return tickets, nil
}

func (f *File) Write(tickets []service.Ticket) error {
	data, err := os.ReadFile(f.Path)

	if err != nil {
		return errors.New("Error")
	}

	// Copia del archivo con los datos nuevos

	newCsv := strings.TrimSuffix(f.Path, ".csv")
	newCsv += "new-tickets.csv"

	if _, err := os.Stat(newCsv); err == nil {
		os.Remove(newCsv)
	}

	file, err := os.Create(newCsv)

	if err != nil {
		return errors.New("Error")
	}

	file.Write(data)

	os.Remove(f.Path)
	newTicket := ""
	for _, ticket := range tickets {
		newTicket += fmt.Sprintf("%d,%s,%s,%s,%s,%d\n", ticket.Id, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticket.Price)
	}

	err = os.WriteFile(f.Path, []byte(newTicket), 0644)
	if err != nil {
		return errors.New("Error")
	}

	return nil
}

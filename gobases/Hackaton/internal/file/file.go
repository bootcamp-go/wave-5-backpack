package file

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"github.com/bootcamp-go/hackaton-go-bases/internal/tickets"
)

type File struct {
	Path string
}

func (f *File) Read() ([]tickets.Ticket, error) {
	readTickets := []tickets.Ticket{}
	file, err := os.Open(f.Path)
	if err != nil {
		return nil, errors.New("ha ocurrido un error al intentar abrir el archivo")
	}
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New("ha ocurrido un error al intentar leer el archivo")
	}

	for _, column := range data {
		if len(column) == 6 {
			id, _ := strconv.Atoi(column[0])
			price, _ := strconv.Atoi(column[5])
			ticket := tickets.Ticket{
				Id:          id,
				Names:       column[1],
				Email:       column[2],
				Destination: column[3],
				Date:        column[4],
				Price:       price,
			}
			readTickets = append(readTickets, ticket)

		}

	}
	file.Close()

	return readTickets, nil
}

func (f *File) Write(tickets []tickets.Ticket) error {
	file, err := os.Create(f.Path)
	if err != nil {
		return errors.New("ha ocurrido un error al intentar abrir el archivo")
	}
	writer := csv.NewWriter(file)

	var data [][]string

	for _, ticket := range tickets {
		id := strconv.Itoa(ticket.Id)
		price := strconv.Itoa(ticket.Price)
		data = append(data, []string{
			id,
			ticket.Names,
			ticket.Email,
			ticket.Destination,
			ticket.Date,
			price,
		})
	}

	writer.WriteAll(data)
	file.Close()

	return nil
}

package file

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/anesquivel/wave-5-backpack/tree/esquivel_andrea/gobases/hackathon/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {
	file, err := os.Open(f.Path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','

	var data []service.Ticket
	var errorReading error
	for {
		record, e := reader.Read()
		if e != nil && e.Error() != "EOF" {
			return nil, e
		} else if e == nil && record != nil {

			id, errId := strconv.ParseInt(record[0], 0, 64)
			price, errprice := strconv.ParseInt(record[5], 0, 64)
			if errId != nil || errprice != nil {
				errorReading = errors.New("Ocurrió un error al parsear data")

			}

			ticket, err := service.CreateTicket(id, price, record[1], record[2], record[3], record[4])

			if err != nil {
				errorReading = err
				break
			} else {
				data = append(data, ticket)
			}

		} else {
			break
		}

	}

	if errorReading != nil {
		return nil, errorReading
	} else {
		return data, nil
	}
}

func (f *File) Write(tickets []service.Ticket) error {
	var allData string = ""
	for _, ticket := range tickets {
		allData += formatData(ticket)
	}
	ticketBytes := []byte(allData)
	err := os.WriteFile(f.Path, ticketBytes, 0644)

	if err != nil {
		return errors.New("Ocurrió un problema al registrar el nuevo ticket en el archivo.")
	}
	return nil
}

func formatData(ticket service.Ticket) string {
	return fmt.Sprintf("%v,%s,%s,%s,%s,%v\n", ticket.Id, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticket.Price)
}

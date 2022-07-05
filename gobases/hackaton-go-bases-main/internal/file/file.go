package file

import (
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
	file, err := os.ReadFile(f.Path)

	if err != nil {
		return nil, errors.New("error abriendo el archivo")
	}

	datos := string(file)
	filas := strings.Split(datos, "\n")
	tickets := []service.Ticket{}
	for _, row := range filas {
		if row != "" {
			fields := strings.Split(row, ",")

			id, errid := strconv.Atoi(fields[0])
			price, err := strconv.Atoi(fields[5])

			if err != nil || errid != nil {
				return nil, errors.New("Error")
			}

			tickets = append(tickets, service.Ticket{
				Id:          id,
				Names:       fields[1],
				Email:       fields[2],
				Destination: fields[3],
				Date:        fields[4],
				Price:       price,
			})
		}
	}
	return tickets, nil
}

func (f *File) Write(tickets []service.Ticket) error {
	_, err := os.ReadFile(f.Path)
	if err != nil {
		return errors.New("error abriendo el archivo")
	}

	strTickets := ""
	for _, ticket := range tickets {
		strTickets += fmt.Sprintf("%d,%s,%s,%s,%s,%d\n", ticket.Id, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticket.Price)
	}

	strTickets = strings.TrimSuffix(strTickets, "\n")

	err = os.WriteFile(f.Path, []byte(strTickets), 0644)
	if err != nil {
		return errors.New("error guardando los tickets")
	}

	return nil
}

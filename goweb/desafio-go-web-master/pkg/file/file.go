package file

import (
	"desafio-go-web/internal/domain"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Path string
}

func (f *File) Read() ([]domain.Ticket, error) {
	data, err := os.ReadFile(f.Path)
	datos := strings.Split(string(data), "\n")

	var tickets []domain.Ticket

	for _, val := range datos {
		var ticket domain.Ticket

		a := strings.Split(val, ",")

		ticket.Id = a[0]
		ticket.Name = a[1]
		ticket.Email = a[2]
		ticket.Country = a[3]
		ticket.Time = a[4]
		ticket.Price, _ = strconv.ParseFloat(a[5], 64)

		tickets = append(tickets, ticket)
	}

	return tickets, err
}

func (f *File) Write(ticket domain.Ticket) error {
	tickets, _ := f.Read()
	var tkt string

	for _, val := range tickets {
		tkt += fmt.Sprintf("%s,%s,%s,%s,%s,%f\n", val.Id, val.Name, val.Email, val.Country, val.Time, val.Price)
	}

	tkt += fmt.Sprintf("%s,%s,%s,%s,%s,%f", ticket.Id, ticket.Name, ticket.Email, ticket.Country, ticket.Time, ticket.Price)

	d1 := []byte(tkt)
	err := os.WriteFile(f.Path, d1, 0644)

	return err
}

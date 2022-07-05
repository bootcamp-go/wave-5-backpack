package file

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"hackaton/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {
	data, err := os.ReadFile(f.Path)
	datos := strings.Split(string(data), "\n")

	var tickets []service.Ticket

	for _, val := range datos {
		var ticket service.Ticket

		a := strings.Split(val, ",")

		ticket.Id, _ = strconv.Atoi(a[0])
		ticket.Names = a[1]
		ticket.Email = a[2]
		ticket.Destination = a[3]
		ticket.Date = a[4]
		ticket.Price, _ = strconv.Atoi(a[5])

		tickets = append(tickets, ticket)
	}

	return tickets, err
}

func (f *File) Write(ticket service.Ticket) error {
	tickets, _ := f.Read()
	var tkt string

	for _, val := range tickets {
		tkt += fmt.Sprintf("%d,%s,%s,%s,%s,%d\n", val.Id, val.Names, val.Email, val.Destination, val.Date, val.Price)
	}

	tkt += fmt.Sprintf("%d,%s,%s,%s,%s,%d", ticket.Id, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticket.Price)

	d1 := []byte(tkt)
	err := os.WriteFile(f.Path, d1, 0644)

	return err
}

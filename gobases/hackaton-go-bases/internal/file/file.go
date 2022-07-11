package file

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {
	data, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}

	var tickets []service.Ticket

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		fields := strings.Split(line, ",")

		if ok := checkEmpty(fields); ok {
			log.Println("campo vacio, no creamos el ticket")
			continue
		}

		id, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, err
		}

		names, email, destination, date := fields[1], fields[2], fields[3], fields[4]

		price, err := strconv.Atoi(fields[5])
		if err != nil {
			return nil, err
		}

		ticket := service.Ticket{
			Id:          id,
			Names:       names,
			Email:       email,
			Destination: destination,
			Date:        date,
			Price:       price,
		}

		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func (f *File) Write(tickets []service.Ticket) error {
	var data string

	for i, t := range tickets {
		if i == len(tickets)-1 {
			data += fmt.Sprintf("%d,%s,%s,%s,%s,%d", t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price)

			if err := os.WriteFile(f.Path, []byte(data), 664); err != nil {
				return err
			}
			break
		}

		data += fmt.Sprintf("%d,%s,%s,%s,%s,%d\n", t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price)
	}

	return nil
}

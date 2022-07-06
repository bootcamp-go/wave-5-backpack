package file

import (
	"os"
	"strconv"
	"strings"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	path string
}

func (f *File) Read() ([]service.Ticket, error) {
	read, err := os.ReadFile(f.path)
	if err != nil {
		return nil, err
	}

	var tickets []service.Ticket
	for _, linea := range read {
		lect := strings.Split(string(linea), ",")
		price, err := strconv.Atoi(lect[5])
		if err != nil {
			return nil, err
		}
		id, err := strconv.Atoi(lect[0])
		if err != nil {
			return nil, err
		}

		t := service.Ticket{
			Id:          int(id),
			Names:       lect[1],
			Email:       lect[2],
			Destination: lect[3],
			Date:        lect[4],
			Price:       int(price),
		}
		tickets = append(tickets, t)
	}
	return tickets, nil
}

func (f *File) Write(service.Ticket) error {

	return nil
}

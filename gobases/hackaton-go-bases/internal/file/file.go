package file

import (
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
	data, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}
	rows := strings.Split(string(data), "\n")
	tickets := []service.Ticket{}
	for _, row := range rows {
		values := strings.Split(row, ",")
		Id, err := strconv.Atoi(values[0])
		if err != nil {
			return nil, err
		}
		Names, Email, Destination, Date := values[1], values[2], values[3], values[4]
		Price, err := strconv.Atoi(values[5])
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, service.Ticket{Id, Names, Email, Destination, Date, Price})
	}
	return tickets, nil
}

func (f *File) Write(tickets []service.Ticket) error {
	csv := []byte{}

	for i, t := range tickets {
		row := fmt.Sprintf("%d,%s,%s,%s,%s,%d", t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price)
		if i < len(tickets)-1 {
			row += "\n"
		}
		csv = append(csv, []byte(row)...)
	}

	return os.WriteFile(f.Path, csv, 0644)
}

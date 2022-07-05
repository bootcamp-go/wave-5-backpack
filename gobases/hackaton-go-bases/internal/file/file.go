package file

import (
	"bytes"
	"encoding/csv"
	"os"
	"strconv"

	"hackaton-go-bases/internal/service"
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
		ticket := service.Ticket{
			Id:          id,
			Names:       line[1],
			Email:       line[2],
			Destination: line[3],
			Date:        line[4],
			Price:       price,
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func (f *File) Write(t service.Ticket) error {
	data, err := os.OpenFile(f.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	d := []string{strconv.Itoa(t.Id), t.Names, t.Email, t.Destination, t.Date, strconv.Itoa(t.Price)}
	csvFile := csv.NewWriter(data)
	err = csvFile.Write(d)
	csvFile.Flush()
	data.Close()
	if err != nil {
		return err
	}
	return nil
}

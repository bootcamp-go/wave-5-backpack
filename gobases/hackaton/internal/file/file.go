package file

import (
	"encoding/csv"
	"fmt"
	"hackaton-go-bases/internal/service"
	"os"
	"strconv"
	"strings"
)

type File struct {
	path string
}

func (f *File) Read() ([]service.Ticket, error) {

	var data []service.Ticket
	read, err := os.ReadFile(f.path)
	file := string(read)
	rows := strings.Split(file, "\n")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, row := range rows {
		ticketS := strings.Split(row, ";")
		id, err := strconv.Atoi(ticketS[0])
		if err != nil {
			fmt.Print(err)
		}

		price, err := strconv.Atoi(ticketS[5])
		if err != nil {
			fmt.Print(err)
		}
		ticket := service.Ticket{Id: id, Names: ticketS[1], Email: ticketS[2], Destination: ticketS[3], Date: ticketS[4], Price: price}
		data = append(data, ticket)
	}

	return data, nil
}

func (f *File) Write(t service.Ticket) error {

	data, err := os.OpenFile(f.path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
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

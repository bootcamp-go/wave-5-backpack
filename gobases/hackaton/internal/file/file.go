package file

import (
	"fmt"
	"hackaton-go-bases/internal/service"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {

	var data []service.Ticket
	read, err := os.ReadFile(f.Path)
	file := string(read)
	rows := strings.Split(file, "\n")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, row := range rows {
		ticketS := strings.Split(row, ",")
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

func (f *File) Write(t []service.Ticket) error {

	resultado := ""
	for _, value := range t {
		resultado += fmt.Sprintf("%d,%s,%s,%s,%s,%d\n", value.Id, value.Names, value.Email, value.Destination, value.Date, value.Price)
	}
	err := os.WriteFile(f.Path, []byte(resultado), 0644)
	return err
}

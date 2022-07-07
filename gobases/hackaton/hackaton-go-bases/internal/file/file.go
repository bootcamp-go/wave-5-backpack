package file

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	path string
}

func New(path string) File {
	return File{
		path: path,
	}
}

func (f *File) Read() ([]service.Ticket, error) {
	file, err := os.ReadFile(f.path)
	if err != nil {
		return []service.Ticket{}, err
	}
	return byteToTickets(file)
}

func (f *File) Write(t service.Ticket) error {
	file, err := os.OpenFile(f.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.WriteString(ticketToCSVLine(t)); err != nil {
		return err
	}
	return nil
}

func (f *File) WriteAll(t []service.Ticket) error {
	if err := os.Truncate(f.path, 0); err != nil {
		return err
	}
	file, err := os.OpenFile(f.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, ticket := range t {
		file.WriteString(ticketToCSVLine((ticket)))
	}

	return nil
}

func ticketToCSVLine(t service.Ticket) string {
	return fmt.Sprintf("%d,%s,%s,%s,%s,%d\n", t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price)
}

func byteToTickets(file []byte) ([]service.Ticket, error) {
	tickets := []service.Ticket{}
	fileStr := string(file)
	fields := strings.Split(fileStr, "\n")
	for _, line := range fields {
		attrs := strings.Split(line, ",")
		if len(attrs) == 6 {
			if _, err := strconv.Atoi(attrs[0]); err != nil {
				return []service.Ticket{}, err
			}
			if _, err := strconv.Atoi(attrs[5]); err != nil {
				return []service.Ticket{}, err
			}
			id, _ := strconv.Atoi(attrs[0])
			names := attrs[1]
			email := attrs[2]
			destination := attrs[3]
			date := attrs[4]
			price, _ := strconv.Atoi(attrs[5])
			t := service.Ticket{
				Id:          id,
				Names:       names,
				Email:       email,
				Destination: destination,
				Date:        date,
				Price:       price,
			}
			tickets = append(tickets, t)
		}
	}
	return tickets, nil
}

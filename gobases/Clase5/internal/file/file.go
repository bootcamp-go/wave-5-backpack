package file

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {
	var sT service.Ticket
	var service []service.Ticket

	read, err := os.Open(f.Path)
	if err != nil {
		fmt.Println(err)
	}
	data := csv.NewReader(read)
	content, err2 := data.ReadAll()
	if err != nil {
		fmt.Println(err2)
	}
	for _, row := range content {
		sT.Id, _ = strconv.ParseInt(row[0], 10, 64)
		sT.Names = row[1]
		sT.Email = row[2]
		sT.Destination = row[3]
		sT.Date = row[4]
		sT.Price, _ = strconv.ParseInt(row[5], 10, 64)
		service = append(service, sT)
	}
	return service, nil
}

func (f *File) Write(t []service.Ticket) error {
	var w string
	for i, value := range t {
		if i < len(t)-1 {
			w += fmt.Sprintf("%v,%v,%v,%v,%v,%v\n", value.Id, value.Names, value.Email, value.Destination, value.Date, value.Price)
		} else {
			w += fmt.Sprintf("%v,%v,%v,%v,%v,%v", value.Id, value.Names, value.Email, value.Destination, value.Date, value.Price)
		}
	}
	err := os.WriteFile("./tickets.csv", []byte(w), 0644)
	if err != nil {
		return err
	}
	return nil
}

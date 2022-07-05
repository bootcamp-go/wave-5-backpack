package file

import (
	"errors"
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
		return nil, errors.New("ha ocurrido un error al leer el archivo")
	} else {
		var tickets []service.Ticket
		slit := strings.Split(string(data), "\n")
		for _, value := range slit {
			line := strings.Split(value, ",")
			id, _ := strconv.Atoi(line[0])
			name := line[1]
			email := line[2]
			destination := line[3]
			date := line[4]
			price, _ := strconv.Atoi(line[5])
			t := service.Ticket{Id: id, Names: name, Email: email, Destination: destination, Date: date, Price: price}
			tickets = append(tickets, t)
		}
		return tickets, nil
	}
}

func (f *File) Write(t []service.Ticket) error {
	var file_str string
	for _, value := range t {
		str := fmt.Sprintf("%d,%s,%s,%s,%s,%d", value.Id, value.Names, value.Email, value.Destination, value.Date, value.Price)
		file_str = fmt.Sprintf("%s\n%s", file_str, str)
	}
	file := []byte(file_str[1:])
	err := os.WriteFile("./tickets.csv", file, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

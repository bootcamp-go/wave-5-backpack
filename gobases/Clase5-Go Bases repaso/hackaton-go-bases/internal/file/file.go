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

func NewFile(ruta string) File {
	f := File{
		path: "./tickets.csv",
	}
	return f
}

func (f *File) Read() ([]service.Ticket, error) {
	var listaTickets []service.Ticket
	data, err := os.ReadFile((*f).path)
	if err == nil {
		for _, linea := range strings.Split(string(data), "\n") {
			campos := strings.Split(linea, ",")
			if true {
				var id, price int
				id, err = strconv.Atoi(campos[0])
				var names = campos[1]
				var email = campos[2]
				var destination = campos[3]
				var date = campos[4]
				price, err = strconv.Atoi(campos[5])
				listaTickets = append(listaTickets, service.NewTicket(id, names, email, destination, date, price))
			}

		}
		return listaTickets, err
	}
	return listaTickets, err

}

func (f *File) Write(t service.Ticket) error {
	_, err := os.Stat(f.path)
	linea := fmt.Sprintf("%v,%v,%v,%v,%v,%v", t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price)
	var texto string
	if err != nil {
		os.Create(f.path)
		texto = linea
	} else {
		data, _ := os.ReadFile((*f).path)
		texto = string(data) + "\n" + linea
	}
	os.WriteFile(f.path, []byte(texto), 0644)
	return err
}

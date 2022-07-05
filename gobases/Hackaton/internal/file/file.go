package file

import (
	"errors"
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
		return nil, errors.New("Error al leer el archivo")
	} else {
		var tickets []service.Ticket

		datos := string(data)

		filas := strings.Split(datos, "\n")

		for _, fila := range filas {

			campos := strings.Split(fila, ",")

			id, _ := strconv.Atoi(campos[0])
			name := campos[1]
			mail := campos[2]
			destiny := campos[3]
			date := campos[4]
			price, _ := strconv.Atoi(filas[5])

			tickets = append(tickets, service.Ticket{
				Id:          id,
				Names:       name,
				Email:       mail,
				Destination: destiny,
				Date:        date,
				Price:       price,
			})
		}
		return tickets, nil
	}
}

func (f *File) Write(a service.Ticket) error {

	return nil
}

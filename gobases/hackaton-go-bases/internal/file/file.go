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

// Lee un archivo csv y lo organiza en una lista de Tickets
func (f *File) Read() ([]service.Ticket, error) {
	data, err := os.ReadFile(f.Path)

	if err != nil {
		panic("Error: No se pudo leer el archivo")
	} else {
		ticketList := []service.Ticket{}
		fileLines := strings.Split(string(data), "\n")
		for _, line := range fileLines {
			data := strings.Split(line, ",")
			id, _ := strconv.Atoi(data[0])
			price, _ := strconv.Atoi(data[5])
			ticket := service.Ticket{
				Id:          id,
				Names:       data[1],
				Email:       data[2],
				Destination: data[3],
				Date:        data[4],
				Price:       price,
			}
			ticketList = append(ticketList, ticket)
		}
		return ticketList, nil
	}

}

// Guardar todos los tickets, recibe el booking para generar el archivo
func (f *File) Write(tickets []service.Ticket) error {
	csvFormat := ""
	for _, t := range tickets {
		newline := fmt.Sprintf("%d,%s,%s,%s,%s,%d", t.Id, t.Names, t.Email, t.Destination, t.Date, t.Price)
		csvFormat = fmt.Sprintf("%s\n%s", csvFormat, newline)
	}
	fmt.Println(csvFormat)
	err := os.WriteFile("tickets.csv", []byte(csvFormat[1:]), 0644)
	if err != nil {
		return errors.New("el archivo no pudo crearse")
	}
	return nil
}

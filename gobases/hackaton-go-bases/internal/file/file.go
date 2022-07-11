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

func (f *File) Read() ([]service.Ticket, error) {

	f.path = "./tickets.csv"
	var Tickets []service.Ticket
	data, err := os.ReadFile(f.path)
	if err != nil {
		fmt.Println("error: la ruta del archivo csv no existe o se encuentra errónea")
		os.Exit(1)
	}

	ticketsOrdenados := strings.Split(string(data), "\n")

	for i := 0; i < len(ticketsOrdenados); i++ {
		archivoSinComas := strings.Split(ticketsOrdenados[i], ",")

		id, _ := strconv.Atoi(archivoSinComas[0])
		nombre := archivoSinComas[1]
		email := archivoSinComas[2]
		paisDestino := archivoSinComas[3]
		horaVuelo := archivoSinComas[4]
		precio, _ := strconv.Atoi(archivoSinComas[5])

		Tickets = append(Tickets, service.Ticket{Id: int(id), Names: nombre, Email: email, Destination: paisDestino, Date: horaVuelo, Price: int(precio)})
	}

	return Tickets, nil
}

func (f *File) Write(nuevoTicket service.Ticket, tickets []service.Ticket) error {

	var textAdd string
	f.path = "./tickets.csv"
	for _, nT := range tickets {
		textAdd += fmt.Sprintf("%d,%s,%s,%s,%s,%d\n",
			nT.Id,
			nT.Names,
			nT.Email,
			nT.Destination,
			nT.Date,
			nT.Price,
		)
	}

	if nuevoTicket.Id != 0 {
		textAdd += fmt.Sprintf("%d,%s,%s,%s,%s,%d\n",
			nuevoTicket.Id,
			nuevoTicket.Names,
			nuevoTicket.Email,
			nuevoTicket.Destination,
			nuevoTicket.Date,
			nuevoTicket.Price,
		)
	}

	newFile := []byte(textAdd)
	error := os.WriteFile(f.path, newFile, 0644)
	return error
}

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
	file, err := os.ReadFile(f.Path)

	if err != nil {
		return nil, errors.New("Error abriendo el archivo tickets.csv")
	}

	datos := string(file)
	filas := strings.Split(datos, "\n")
	tickets := []service.Ticket{}
	for _, fila := range filas {
		if fila != "" {
			fields := strings.Split(fila, ",")

			id, errId := strconv.Atoi(fields[0])
			names := fields[1]
			email := fields[2]
			destination := fields[3]
			date := fields[4]
			price, errPrice := strconv.Atoi(fields[5])

			if errId != nil || errPrice != nil {
				return nil, errors.New("Error al convertir los datos")
			}

			tickets = append(tickets, service.Ticket{
				Id:          id,
				Names:       names,
				Email:       email,
				Destination: destination,
				Date:        date,
				Price:       price,
			})
		}
	}
	return tickets, nil
}

func (f *File) Write(tickets []service.Ticket) error {
	data, err := os.ReadFile(f.Path)
	if err != nil {
		return errors.New("Error abriendo el archivo tickets.csv")
	}

	// --------- Crear archivo de copia del csv ----------

	// Hacer una copia del archivo
	// eliminar el .csv del string path
	pathSinCSV := strings.TrimSuffix(f.Path, ".csv")
	pathSinCSV += "-copy.csv"

	// Borrar copia si ya existia una
	if _, err := os.Stat(pathSinCSV); err == nil {
		os.Remove(pathSinCSV)
	}

	file, err := os.Create(pathSinCSV)
	if err != nil {
		return errors.New("Error creando el archivo copia de tickets.csv")
	}

	// Escribir el archivo
	file.Write(data)

	// --------- Guardar tickets en el csv ----------

	// Borrar el archivo si existe
	os.Remove(f.Path)

	strTickets := ""
	for _, ticket := range tickets {
		strTickets += fmt.Sprintf("%d,%s,%s,%s,%s,%d\n", ticket.Id, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticket.Price)
	}

	// Quitar el último \n
	strTickets = strings.TrimSuffix(strTickets, "\n")

	err = os.WriteFile(f.Path, []byte(strTickets), 0644)
	if err != nil {
		return errors.New("Error guardando los tickets en tickets.csv")
	}

	return nil
}

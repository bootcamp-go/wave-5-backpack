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

func (f *File) Read() (tickets []service.Ticket, myerror error) {
	bytesArchivo, _ := os.ReadFile(f.Path)
	stringArchivo := string(bytesArchivo)
	linesArchivo := strings.Split(stringArchivo, "\n")
	for _, line := range linesArchivo {
		valueLine := strings.Split(line, ",")
		id, _ := strconv.ParseInt(valueLine[0], 10, 32)
		nombre := valueLine[1]
		correo := valueLine[2]
		destino := valueLine[3]
		fechaH := valueLine[4]
		precio, _ := strconv.ParseInt(valueLine[5], 10, 32)
		tickets = append(tickets, service.Ticket{Id: int(id), Names: nombre, Email: correo, Destination: destino, Date: fechaH, Price: int(precio)})

	}
	return tickets, nil
}

func (f *File) Write(ticketColeccion []service.Ticket) error {
	var result string
	for _, lineas := range ticketColeccion {
		result += fmt.Sprintln(lineas.Id, lineas.Names, lineas.Email, lineas.Destination, lineas.Date, lineas.Price)
	}
	bytesArray := []byte(result)
	error := os.WriteFile(f.Path, bytesArray, 0644)
	if error != nil {
		return errors.New("se produjo un error al escribir el archivo")
	}
	return nil
}

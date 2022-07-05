package file

import (
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

	var tickets []service.Ticket
	//Lectura de archivo
	read, err := os.ReadFile(f.Path)
	if err != nil {
		fmt.Println(err)
		//return tickets, err
	}

	//Imprimiendo en pantalla la información del csv
	data := string(read)
	fmt.Println("----------\nInformación del csv\n----------")
	fmt.Println(strings.ReplaceAll(data, ",", "\t\t\t"))

	// Recorro para asignar a service.Ticket
	splitJump := strings.Split(string(read), "\n")
	for _, valueLine := range splitJump {

		splitComma := strings.Split(string(valueLine), ",")
		var (
			Id, Price                        int
			Nombre, Email, Destination, Date string
		)
		for j, data := range splitComma {
			if j == 0 {
				Id, err = strconv.Atoi(data)
				if err != nil {
					return tickets, err
				}
			}
			if j == 1 {
				Nombre = data
			}
			if j == 2 {
				Email = data
			}
			if j == 3 {
				Destination = data
			}
			if j == 4 {
				Date = data
			}
			if j == 5 {
				Price, err = strconv.Atoi(data)
				if err != nil {
					return tickets, err
				}
			}
		}
		tickets = append(tickets, service.Ticket{Id, Nombre, Email, Destination, Date, Price})
	}
	return tickets, nil
}

func (f *File) Write(s string) error {
	//Traigo todo lo del csv para añadir los nuevos
	//elementos
	data, _ := os.ReadFile(f.Path)

	dataWrite := string(data)
	dataWrite += s
	//Obtener la informacion que ingresó
	//almacenarla en data

	//WriteFile para escribirlo en el csv info
	dataTicket := []byte(dataWrite)
	err := os.WriteFile("./writingTicket.csv", dataTicket, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (f *File) WriteId(s string) error {
	dataWrite := s
	dataTicket := []byte(dataWrite)
	err := os.WriteFile("./writingTicket.csv", dataTicket, 0644)
	if err != nil {
		return err
	}
	return nil
}

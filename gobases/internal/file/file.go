package file

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bootcamp-go/wave-5-backpack/gobases/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {
	var t []service.Ticket
	var registers []string
	var datos []string

	fileB, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, errors.New("error: no se pudo leer el archivo")
	} else {
		fileS := string(fileB)
		registers = strings.Split(fileS, "\n")
		for _, dato := range registers {
			datos = strings.Split(dato, ",")

			id, err1 := strconv.ParseInt(datos[0], 10, 64)
			if err1 != nil {
				return nil, errors.New("error: fallo ParseInt en \"ID\"")
			}

			name := datos[1]
			email := datos[2]
			destination := datos[3]
			date := datos[4]
			price, err2 := strconv.ParseInt(datos[5], 10, 64)

			if err2 != nil {
				return nil, errors.New("error: fallo ParseInt en \"Price\"")
			}

			tickets := service.Ticket{Id: int(id), Names: name, Email: email, Destination: destination, Date: date, Price: int(price)}

			t = append(t, tickets)

		}
	}

	return t, nil
}

func (f *File) Write(t service.Ticket, u bool) error {
	fileB, err := os.ReadFile(f.Path)
	if err != nil {
		return errors.New("error: no se pudo leer el archivo para escribirlo")
	} else {
		if u {
			wArchivo := []byte(fmt.Sprint(string(fileB), t.Detalle()))
			err := os.WriteFile("./tickets.csv", wArchivo, 0644)
			return err
		} else {
			wArchivo := []byte(fmt.Sprint(string(fileB)))
			err := os.WriteFile("./tickets.csv", wArchivo, 0644)
			return err
		}
	}
}

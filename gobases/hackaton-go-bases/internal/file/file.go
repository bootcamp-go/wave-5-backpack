package file

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/bootcamp-go/wave-5-backpack/gobases/hackaton-go-bases/internal/service"
)

type File struct {
	path string
}

func (f *File) Read() ([]service.Ticket, error) {
	// Se lee el archivo CSV
	data, err := os.ReadFile(f.path)

	if err != nil {
		return nil, errors.New("no se puede abrir el archivo csv")
	}

	return readData(data), nil
}

func (f *File) Write(service.Ticket) error {
	return nil
}

func (f *File) SetPath(p string) {
	f.path = p
}

func readData(bytes []byte) []service.Ticket {
	data := []service.Ticket{}
	valid := true
	id, price := 0, 0
	names, email, destination, date := "", "", "", ""

	// Separamos la información por cada línea leída
	lineas := strings.Split(string(bytes), "\n")

	// Verificamos la información por cada línea leída
	for _, p := range lineas {
		// Separamos cada línea por comas
		linea := strings.Split(p, ",")

		// Se verifica que cada línea tenga la misma cantidad de datos
		if len(linea) == 6 {
			for i, l := range linea {
				switch i {
				case 0:
					idInt, err := strconv.Atoi(l)
					if err != nil {
						valid = false
					} else {
						id = idInt
					}
				case 1:
					names = l
				case 2:
					email = l
				case 3:
					destination = l
				case 4:
					date = l
				case 5:
					priceInt, err := strconv.Atoi(l)
					if err != nil {
						valid = false
					} else {
						price = priceInt
					}
				}
			}

			// Si la información de la línea es válida se agrega al arreglo
			if valid {
				custom := service.Ticket{Id: id, Names: names, Email: email, Destination: destination, Date: date, Price: price}
				data = append(data, custom)
			}
		}
	}

	return data
}

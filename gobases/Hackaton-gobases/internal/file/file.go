package file

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() ([][]string, error) {
	file, err := os.Open(f.Path)
	if err != nil {
		return nil, fmt.Errorf("Unable to read input file "+f.Path, err)
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}
	return lines, nil
}

func (f *File) Write(t service.Ticket, operacion string) error {
	file, err := os.OpenFile(f.Path, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("No se pudo leer el archivo "+f.Path, err)
	}
	defer file.Close()

	writter := csv.NewWriter(file)
	records, _ := f.Read()
	id := strconv.Itoa(t.Id)

	switch operacion {
	case "CREATE":
		records = append(records, []string{id, t.Names, t.Email, t.Destination, t.Date, strconv.Itoa(t.Price)})
		for _, record := range records {
			if err := writter.Write(record); err != nil {
				log.Fatalln("Error al escribir el archivo", err)
			}
		}
		writter.Flush()
	case "UPDATE":
		for _, record := range records {
			if record[0] == id {
				record = []string{id, t.Names, t.Email, t.Destination, t.Date, strconv.Itoa(t.Price)}
			}
			if err := writter.Write(record); err != nil {
				log.Fatalln("Error al escribir el archivo")
			}
		}
		writter.Flush()
	case "DELETE":
		csvData := [][]string{}
		for _, record := range records {
			if record[0] != id {
				csvData = append(csvData, record)
			}
		}
		_ = os.Truncate(f.Path, 0)

		err = writter.WriteAll(csvData)
		if err != nil {
			fmt.Println("Error al escribir en el archivo", err)
			return err
		}
		writter.Flush()
	default:
		return errors.New("operación inválida")
	}
	return nil
}

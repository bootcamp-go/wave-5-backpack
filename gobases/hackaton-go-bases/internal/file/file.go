package file

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type FileFuncs interface {
	Read() ([]service.Ticket, error)
	Write([]service.Ticket) error
}

type File struct {
	path string
}

func NewFile(path string) FileFuncs {
	return &File{path: path}
}

func (f *File) Read() ([]service.Ticket, error) {
	tickets := []service.Ticket{}

	file, err := os.Open(f.path)
	if err != nil {
		return nil, errors.New("no se ha podido abrir el archivo")
	}
	defer file.Close()

	r := csv.NewReader(file)
	tickets_slice, _ := r.ReadAll()

	for _, row := range tickets_slice {
		id, err := strconv.Atoi(row[0])
		if err != nil {
			fmt.Println(err)
		}

		price, err := strconv.Atoi(row[5])
		if err != nil {
			fmt.Println(err)
		}

		tickets = append(tickets, service.Ticket{
			Id:          id,
			Names:       row[1],
			Email:       row[2],
			Destination: row[3],
			Date:        row[4],
			Price:       price,
		})
	}

	return tickets, nil
}

func (f *File) Write(ts []service.Ticket) error {
	file, err := os.Create(f.path)
	if err != nil {
		return errors.New("no se ha podido abrir el archivo")
	}
	defer file.Close()

	var to_write [][]string
	for _, v := range ts {
		to_append := []string{
			strconv.Itoa(v.Id),
			v.Names,
			v.Email,
			v.Destination,
			v.Date,
			strconv.Itoa(v.Price),
		}
		to_write = append(to_write, to_append)
	}

	w := csv.NewWriter(file)
	w.WriteAll(to_write)
	w.Flush()

	return nil
}

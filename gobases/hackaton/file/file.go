package file

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"hackaton/service"
)

type File struct {
	path string
}

func ProccessFile() []service.Ticket {
	r := File{}
	res, err := r.Read()
	if err != nil {
		fmt.Println(err)
	}
	return res
}

func (f *File) Read() ([]service.Ticket, error) {
	dataSlice := []service.Ticket{}
	f.path = "./tickets.csv"
	file, err := os.Open(f.path)
	if err != nil {
		fmt.Printf("Ocurrio un error: %v \n", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)

	ticketSlice, err := reader.ReadAll()

	for _, r := range ticketSlice {
		id, err := strconv.Atoi(r[0])
		if err != nil {
			fmt.Println(err)
		}

		price, err := strconv.Atoi(r[5])
		if err != nil {
			fmt.Println(err)
		}

		dataSlice = append(dataSlice, service.Ticket{
			Id:          id,
			Names:       r[1],
			Email:       r[2],
			Destination: r[3],
			Date:        r[4],
			Price:       price,
		})
	}

	return dataSlice, nil
}

func (f *File) Write(service.Ticket) error {
	return nil
}

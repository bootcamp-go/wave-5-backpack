package store

import (
	"desafio-go-web/internal/domain"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Store interface {
	ReadTickets() ([]domain.Ticket, error)
}

type fileStore struct {
	filePath string
}

func NewStore(filePath string) Store {
	return &fileStore{
		filePath: filePath,
	}
}

func (fs *fileStore) ReadTickets() ([]domain.Ticket, error) {

	var ticketList []domain.Ticket

	file, err := os.Open(fs.filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}

	return ticketList, nil
}

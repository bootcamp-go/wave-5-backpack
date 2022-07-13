package store

import (
	"desafio-go-web/internal/domain"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Store interface {
	Read(data interface{}) ([]domain.Ticket, error)
	Write(data interface{}) error
	Ping() error
}

type fileStore struct {
	FilePath string
}

func NewStore(fileName string) Store {
	return &fileStore{fileName}
}

func (fs *fileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}

func (fs *fileStore) Read(data interface{}) ([]domain.Ticket, error) {
	var ticketList []domain.Ticket

	file, err := os.Open(fs.FilePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	readData, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not parse file: %w", err)
	}

	for _, row := range readData {
		price, err := strconv.ParseFloat(row[5], 64)
		if err == nil {
			ticketList = append(ticketList, domain.Ticket{
				Id:      row[0],
				Name:    row[1],
				Email:   row[2],
				Country: row[3],
				Time:    row[4],
				Price:   price,
			})
		}
	}

	return ticketList, nil
}

func (fs *fileStore) Ping() error {
	if _, err := os.ReadFile(fs.FilePath); err != nil {
		return err
	}
	return nil
}

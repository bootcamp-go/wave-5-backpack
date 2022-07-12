package storage

import (
	// "encoding/json"
	"desafio_web/internal/domain"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Store interface {
	Read(p *[]domain.Ticket) error
	Write(data interface{}) error
}

type fileStore struct {
	FilePath string
}

func NewStore(fileName string) Store {
	return &fileStore{fileName}
}

func (f *fileStore) Read(p *[]domain.Ticket) error {
	file, err := os.Open(f.FilePath)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return err
		}
		*p = append(*p, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}
	return nil
}
func (f *fileStore) Write(data interface{}) error {
	// file, err := json.Marshal(data)
	// if err != nil {
	// 	return err
	// }
	// return os.WriteFile(f.FilePath, file, 0644)
	return fmt.Errorf("")
}

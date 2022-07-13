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
	ReadJSON(data interface{}) error
	WriteJSON(data interface{}) error
	ReadCSV() ([]domain.Ticket, error)
}

type fileStore struct {
	FilePath string
}

func NewStore(fileName string) Store {
	return &fileStore{
		fileName,
	}
}

func (fs *fileStore) WriteJSON(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", " ")

	//fmt.Println(fileData, err)
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}
func (fs *fileStore) ReadJSON(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	//fmt.Println(file, err)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

func (fs *fileStore) ReadCSV() ([]domain.Ticket, error) {
	var ticketList []domain.Ticket

	file, err := os.Open(fs.FilePath)
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

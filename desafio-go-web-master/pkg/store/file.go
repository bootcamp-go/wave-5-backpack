package store

import (
	"desafio-go-web/internal/domain"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type fileStore struct {
	FilePath string
}



type Store interface {
	Read(data interface{}) error
}

func NewStore(filePath string) Store {
	return &fileStore{FilePath: filePath}

}

// Read implements Store
func (fs *fileStore) Read(data interface{}) error {

	var ticketList []domain.Ticket

	file, err := os.Open(fs.FilePath)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}

	csvR := csv.NewReader(file)
	data1, err1 := csvR.ReadAll()
	if err1 != nil {
		return fmt.Errorf("could not open file: %w", err)
	}

	for _, row := range data1 {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return err
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

	pl, err3:= json.MarshalIndent(ticketList, "", " ")

	if err3 != nil {
		return fmt.Errorf("could not open file: %w", err)
	}

	return json.Unmarshal(pl, &data)
}
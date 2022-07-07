package file

import (
	"encoding/json"
	"os"

	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/domain"
)

type DB struct {
	Users  []domain.User
	LastId int
}

func ReadJSONFile(fileName string) (DB, error) {
	data, err := os.ReadFile(fileName)
	db := DB{}
	if err == nil {
		err = json.Unmarshal(data, &db)
	}
	return db, err
}

func WriteJSONFile(fileName string, db DB) error {
	data, err := json.Marshal(db)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

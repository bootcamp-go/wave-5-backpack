package services

import (
	"encoding/json"
	"errors"
	"goweb/models"
	"io/ioutil"
	"log"
)

func Read() ([]models.User, error) {
	file, err := ioutil.ReadFile("./users.json")
	if err != nil {
		log.Println(err)
		return nil, errors.New("error al leer el archivo users.json")
	}

	users := []models.User{}
	if err := json.Unmarshal([]byte(file), &users); err != nil {
		log.Println(err)
		return nil, errors.New("error al deserializar el archivo users.json")
	}

	return users, nil
}

package services

import (
	"encoding/json"
	"errors"
	"goweb/models"
	"log"
	"os"
)

var PATH string = "./resources/products.json"

func ReadAllProducts() ([]models.Product, error) {
	data, err := os.ReadFile(PATH)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error al leer el archivo products.json")
	}

	productos := []models.Product{}

	if err := json.Unmarshal([]byte(data), &productos); err != nil {
		log.Println(err)
		return nil, errors.New("error al deserializar el archivo products.json")
	}

	return productos, nil
}

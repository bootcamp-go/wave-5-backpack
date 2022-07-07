package services

import (
	"encoding/json"
	"errors"
	"goweb/models"
	"log"
	"os"
	"strconv"
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

func GetById(id string) (models.Product, error) {
	productos, err := ReadAllProducts()
	if err != nil {
		log.Println(err)
		return models.Product{}, err
	}

	strId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return models.Product{}, err
	}

	for _, producto := range productos {
		if producto.Id == strId {
			return producto, nil
		}
	}

	log.Println("no se encontro el producto con id: ", id)
	return models.Product{}, errors.New("producto no encontrado")
}

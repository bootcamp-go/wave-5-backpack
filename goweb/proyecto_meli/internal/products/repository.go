package products

import (
	"encoding/json"
	"os"
	"proyecto_meli/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	//Store(id int, nombre, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	//LastID() (int, error)
}

type repository struct{}

var ps []domain.Product
var lastID int

func NewRepository() Repository {
	p, _ := readFile()
	ps = p
	return &repository{}
}

func readFile() ([]domain.Product, error) {
	var productos []domain.Product
	dataBit, err := os.ReadFile("products.json")
	if err != nil {
		return productos, err
	} else {
		err := json.Unmarshal(dataBit, &productos)
		if err != nil {
			return productos, err
		}
	}

	return productos, nil
}

func (r *repository) GetAll() ([]domain.Product, error) {
	return ps, nil
}

package products

import (
	"encoding/json"
	"fmt"
	"goweb/internal/domain"
	"log"
	"os"
	"time"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	Create(name string, color string, price float64, stock int, code string, publisher bool) (domain.Product, error)
}

var ps []domain.Product
var lastId int

type repository struct{}

func (*repository) Create(name string, color string, price float64, stock int, code string, publisher bool) (domain.Product, error) {

	var t time.Time = time.Now()
	created := t.Format("2006-01-02 15:04:05")

	pst, err := NewRepository().GetAll()
	if err != nil {
		
	}
	lastId = pst[len(pst) - 1].Id + 1
	p := domain.Product{Id:lastId, Name: name, Color: color, Price: price, Stock: stock, Code: code, Publisher: publisher, CreatedAt:  created}
	
	ps = append(ps, p)

	toJson, err := toJSON(ps)

	if err != nil {
		fmt.Println("Error:", err.Error())
		
	}

	err1 := os.WriteFile("/Users/gtorrealba/go-base-TT/wave-5-backpack/goweb/internal/productos.json", toJson, 0644)

	if err1 != nil {
		fmt.Println("Error:", err1.Error())
		
	}
	
	return p, nil
}

func toJSON(p []domain.Product)  ([]byte, error){
	jsonData, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]domain.Product, error) {
	file, err := os.ReadFile("/Users/gtorrealba/go-base-TT/wave-5-backpack/goweb/internal/productos.json")

	if err != nil {
		fmt.Println("Error abriendo el archivo productos.json")
	}

	if err1 := json.Unmarshal([]byte(file), &ps); err1 != nil {
		log.Fatal(err1)
	}

	return ps, nil
}

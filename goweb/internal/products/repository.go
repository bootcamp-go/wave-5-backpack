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
	Update(id int, name string, color string, price float64, stock int, code string, publisher bool) (domain.Product, error)
	Delete(id int) error
	ParcialUpdate(id int, name string, price float64) (domain.Product, error)
}

var ps []domain.Product
var lastId int

type repository struct{}

// ParcialUpdate implements Repository
func (r *repository) ParcialUpdate(id int, name string, price float64) (domain.Product, error) {

	updated := false
	ps, err := r.GetAll()
	var p domain.Product
	if err != nil {
		fmt.Println("Error:", err)
	}

	for i := range ps {
		if ps[i].Id == id {
		
			if name != "" && price > 0 {
				ps[i].Name = name
				ps[i].Price = price
				p = ps[i]
				updated = true
			}
			if name == "" && price > 0{
				ps[i].Price = price
				p = ps[i]
				updated = true
			}
			if name != "" && price == 0 {
				ps[i].Name = name
				p = ps[i]
				updated = true
			}
		}
	}
	if !updated {
		return domain.Product{}, fmt.Errorf("producto %d no encontrado", id)
	}

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

// Delete implements Repository
func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	ps, err := r.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
	}

	for i := range ps {
		if ps[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("producto %d no encontrado", id)
	}

	ps = append(ps[:index], ps[index+1:]...)

	toJson, err := toJSON(ps)

	if err != nil {
		fmt.Println("Error:", err.Error())

	}

	err1 := os.WriteFile("/Users/gtorrealba/go-base-TT/wave-5-backpack/goweb/internal/productos.json", toJson, 0644)

	if err1 != nil {
		fmt.Println("Error:", err1.Error())

	}
	return nil
}

// Update implements Repository
func (r *repository) Update(id int, name string, color string, price float64, stock int, code string, publisher bool) (domain.Product, error) {
	p := domain.Product{Name: name, Color: color, Price: price, Stock: stock, Code: code, Publisher: publisher}
	update := false
	ps, err := r.GetAll()
	if err != nil {
		fmt.Println("Error:", err)
	}

	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			p.CreatedAt = ps[i].CreatedAt
			ps[i] = p
			update = true
		}
	}

	toJson, err := toJSON(ps)

	if err != nil {
		fmt.Println("Error:", err.Error())

	}

	err1 := os.WriteFile("/Users/gtorrealba/go-base-TT/wave-5-backpack/goweb/internal/productos.json", toJson, 0644)

	if err1 != nil {
		fmt.Println("Error:", err1.Error())

	}

	if !update {
		return domain.Product{}, fmt.Errorf("producto %d no encontrado", id)
	}

	return p, nil
}

func (*repository) Create(name string, color string, price float64, stock int, code string, publisher bool) (domain.Product, error) {

	var t time.Time = time.Now()
	created := t.Format("2006-01-02 15:04:05")

	pst, err := NewRepository().GetAll()
	if err != nil {
		fmt.Println("Error:", err)
	}
	lastId = pst[len(pst)-1].Id + 1
	p := domain.Product{Id: lastId, Name: name, Color: color, Price: price, Stock: stock, Code: code, Publisher: publisher, CreatedAt: created}

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

func toJSON(p []domain.Product) ([]byte, error) {
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

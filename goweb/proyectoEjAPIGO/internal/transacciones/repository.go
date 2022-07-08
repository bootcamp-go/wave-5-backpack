package transacciones

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Transaccion struct {
	ID                 int     `json:"id"`
	Codigo_transaccion string  `json:"codigo_transaccion" binding:"required"`
	Moneda             string  `json:"moneda"`
	Monto              float64 `json:"monto"`
	Emisor             string  `json:"emisor"`
	Receptor           string  `json:"receptor"`
	Fecha_transaccion  string  `json:"fecha_transaccion"`
}

type Repository interface {
	getAll() ([]Transaccion, error)
	Store(id int, codigo_transaccion, moneda, emisor, receptor, fecha_transaccion string, monto float64) (Transaccion, error)
	LastID() (int, error)
	Update(id int, codigo_transaccion, moneda, emisor, receptor, fecha_transaccion string, monto float64) (Transaccion, error)
	UpdateCTandMonto(id int, codigo_transaccion string, monto float64) (Transaccion, error)
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

var ts []Transaccion
var lastID int = 3

func (r *repository) getAll() ([]Transaccion, error) {

	data, err := os.ReadFile("./transacciones.json")

	if err != nil {
		fmt.Println("Error en la lectura: %v", err)
	}

	if err := json.Unmarshal(data, &ts); err != nil {
		log.Fatal(err)
	}
	return ts, nil

}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, codigo_transaccion, moneda, emisor, receptor, fecha_transaccion string, monto float64) (domain.Transaccion, error) {
	t := Transaccion{ID: id, Codigo_transaccion: codigo_transaccion, Moneda: moneda, Emisor: emisor, Receptor: receptor, Fecha_transaccion: fecha_transaccion, Monto: monto}
	ts = append(ts, t)
	lastID = t.ID
	return t, nil
}

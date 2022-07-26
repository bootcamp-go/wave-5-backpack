package transactions

import (
	"errors"
	"fmt"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/models"
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/pkg/storage"
)

type Repository interface {
	Store(monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
	Update(id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
	Patch(id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error)
	GetAll() ([]models.Transaction, error)
	GetByID(id int) (models.Transaction, error)
	GetLastID() (int, error)
	Delete(id int) (int, error)
}

func NewRepository(storage storage.Storage) Repository {
	return &repository{storage}
}

type repository struct {
	storage storage.Storage
}

func (r repository) Store(monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	var tr []models.Transaction
	if err := r.storage.Read(&tr); err != nil {
		return models.Transaction{}, fmt.Errorf("error: al leer el archivo %v", err)
	}

	newID := (tr[len(tr)-1].ID) + 1
	t := models.Transaction{
		ID:       newID,
		Monto:    monto,
		Cod:      cod,
		Moneda:   moneda,
		Emisor:   emisor,
		Receptor: receptor,
		Fecha:    time.Now().Local().String(),
	}

	// Actualiza memoria
	tr = append(tr, t)

	// Escribe archivo
	err := r.storage.Write(tr)
	if err != nil {
		return models.Transaction{}, fmt.Errorf("error: al escribir el archivo %v", err)
	}

	return t, nil
}

func (r repository) GetAll() ([]models.Transaction, error) {
	var tr []models.Transaction
	if err := r.storage.Read(&tr); err != nil {
		return nil, err
	}

	if len(tr) == 0 {
		return nil, errors.New("no hay registros")
	}

	return tr, nil
}

func (r repository) GetByID(id int) (models.Transaction, error) {
	var tr []models.Transaction
	if err := r.storage.Read(&tr); err != nil {
		return models.Transaction{}, err
	}

	for _, t := range tr {
		if t.ID == id {
			return t, nil
		}
	}

	return models.Transaction{}, fmt.Errorf("trasaction con ID: %v no encontrado", id)
}

func (r repository) GetLastID() (int, error) {
	var tr []models.Transaction
	if err := r.storage.Read(&tr); err != nil {
		return 0, err
	}

	if len(tr) == 0 {
		return 0, errors.New("error: no hay transactiones")
	}

	id := tr[len(tr)-1].ID

	return id, nil
}

func (r repository) Update(id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	var tr []models.Transaction
	if err := r.storage.Read(&tr); err != nil {
		return models.Transaction{}, fmt.Errorf("error: al leer el archivo %v", err)
	}

	for i, tt := range tr {
		if tt.ID == id {
			t := models.Transaction{
				ID:       id,
				Monto:    monto,
				Cod:      cod,
				Moneda:   moneda,
				Emisor:   emisor,
				Receptor: receptor,
				Fecha:    tt.Fecha,
			}

			// Actualiza la memoria
			tr[i] = t

			if err := r.storage.Write(tr); err != nil {
				return models.Transaction{}, fmt.Errorf("error: al escribir el archivo %v\n", err)
			}

			return t, nil
		}
	}

	return models.Transaction{}, fmt.Errorf("error: no existe el ID: %v", id)
}

func (r repository) Patch(id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	var tr []models.Transaction
	if err := r.storage.Read(&tr); err != nil {
		return models.Transaction{}, fmt.Errorf("error: al leer el archivo %v", err)
	}

	var t models.Transaction
	for i, tt := range tr {
		if tt.ID == id {
			patch(monto, cod, moneda, emisor, receptor, &t, tt)

			tr[i] = t

			if err := r.storage.Write(tr); err != nil {
				return models.Transaction{}, fmt.Errorf("error: al escribir el archivo %v", err)
			}

			return t, nil
		}
	}

	return models.Transaction{}, fmt.Errorf("error: no existe el ID %v", id)
}

func patch(monto float64, cod, moneda, emisor, receptor string, t *models.Transaction, tt models.Transaction) {
	t.ID = tt.ID
	t.Fecha = tt.Fecha

	if monto != 0 {
		t.Monto = monto
	} else {
		t.Monto = tt.Monto
	}

	if cod != "" {
		t.Cod = cod
	} else {
		t.Cod = tt.Cod
	}

	if moneda != "" {
		t.Moneda = moneda
	} else {
		t.Moneda = tt.Moneda
	}

	if emisor != "" {
		t.Emisor = emisor
	} else {
		t.Emisor = tt.Emisor
	}

	if receptor != "" {
		t.Receptor = receptor
	} else {
		t.Receptor = tt.Receptor
	}
}

func (r repository) Delete(id int) (int, error) {
	var tr []models.Transaction
	if err := r.storage.Read(&tr); err != nil {
		return 0, err
	}

	for i, t := range tr {
		if t.ID == id {
			tr = append(tr[:i], tr[i+1:]...)
			if err := r.storage.Write(tr); err != nil {
				return 0, fmt.Errorf("error: no se pudo escribir el archivo %v\n", err)
			}

			return id, nil
		}
	}

	return 0, fmt.Errorf("error: ID %v no existe\n", id)
}

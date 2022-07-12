package transacciones

import (
	"fmt"
	"goweb/internal/domain"
	"goweb/pkg/store"
)

const (
	TransaccionNotFound = "transaccion %d not found"
	FailReading         = "cant read database"
	FailWriting         = "cant write database, error: %w"
)

type Repository interface {
	GetAll() ([]domain.Transaccion, error)
	Store(id int, codigo_transaccion, moneda, emisor, receptor, fecha_transaccion string, monto float64) (domain.Transaccion, error)
	LastID() (int, error)
	Update(id int, codigo_transaccion, moneda, emisor, receptor, fecha_transaccion string, monto float64) (domain.Transaccion, error)
	UpdateCTandMonto(id int, codigo_transaccion string, monto float64) (domain.Transaccion, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

var ts []domain.Transaccion
var lastID int = 3

func (r *repository) GetAll() ([]domain.Transaccion, error) {

	var ts []domain.Transaccion
	if err := r.db.Read(&ts); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return ts, nil

}

func (r *repository) LastID() (int, error) {
	var ts []domain.Transaccion
	if err := r.db.Read(&ts); err != nil {
		return 0, fmt.Errorf(FailReading)
	}
	if len(ts) == 0 {
		return 0, nil
	}
	return ts[len(ts)-1].ID, nil
}

func (r *repository) Store(id int, codigo_transaccion, moneda, emisor, receptor, fecha_transaccion string, monto float64) (domain.Transaccion, error) {
	var ts []domain.Transaccion

	if err := r.db.Read(&ts); err != nil {
		return domain.Transaccion{}, fmt.Errorf(FailReading)
	}

	t := domain.Transaccion{ID: id, Codigo_transaccion: codigo_transaccion, Moneda: moneda, Emisor: emisor, Receptor: receptor, Fecha_transaccion: fecha_transaccion, Monto: monto}

	ts = append(ts, t)

	if err := r.db.Write(ts); err != nil {
		return domain.Transaccion{}, fmt.Errorf(FailWriting, err)
	}
	return t, nil

	//----------------------------------------------------------------------------------------------------------------------
	// t := domain.Transaccion{ID: id, Codigo_transaccion: codigo_transaccion, Moneda: moneda, Emisor: emisor, Receptor: receptor, Fecha_transaccion: fecha_transaccion, Monto: monto}
	// ts = append(ts, t)
	// lastID = t.ID
	// return t, nil
}

func (r *repository) Update(id int, codigo_transaccion, moneda, emisor, receptor, fecha_transaccion string, monto float64) (domain.Transaccion, error) {
	var ts []domain.Transaccion

	if err := r.db.Read(&ts); err != nil {
		return domain.Transaccion{}, fmt.Errorf(FailReading)
	}
	t := domain.Transaccion{Codigo_transaccion: codigo_transaccion, Moneda: moneda, Emisor: emisor, Receptor: receptor, Fecha_transaccion: fecha_transaccion, Monto: monto}
	updated := false

	for i := range ts {
		if ts[i].ID == id {
			t.ID = id
			ts[i] = t
			updated = true
		}
	}
	if !updated {
		return domain.Transaccion{}, fmt.Errorf(TransaccionNotFound, id)

	}
	if err := r.db.Write(ts); err != nil {
		return domain.Transaccion{}, fmt.Errorf(FailWriting, err)
	}
	return t, nil
	// t := domain.Transaccion{ID: id, Codigo_transaccion: codigo_transaccion, Moneda: moneda, Emisor: emisor, Receptor: receptor, Fecha_transaccion: fecha_transaccion, Monto: monto}
	// updated := false
	// for i := range ts {
	// 	if ts[i].ID == id {
	// 		t.ID = id
	// 		ts[i] = t
	// 		updated = true
	// 	}
	// }
	// if !updated {
	// 	return domain.Transaccion{}, fmt.Errorf("producto %d no encontrado", id)
	// }
	// return t, nil
}

func (r *repository) UpdateCTandMonto(id int, codigo_transaccion string, monto float64) (domain.Transaccion, error) {
	updated := false
	var ts []domain.Transaccion

	if err := r.db.Read(&ts); err != nil {
		return domain.Transaccion{}, fmt.Errorf(FailReading)
	}
	var t domain.Transaccion
	for i := range ts {
		if ts[i].ID == id {
			ts[i].Codigo_transaccion = codigo_transaccion
			ts[i].Monto = monto
			t = ts[i]
			updated = true
		}
	}
	if !updated {

		return domain.Transaccion{}, fmt.Errorf(TransaccionNotFound, id)
	}

	if err := r.db.Write(ts); err != nil {
		return domain.Transaccion{}, fmt.Errorf(FailWriting, err)
	}
	return t, nil
}

func (r *repository) Delete(id int) error {
	var ts []domain.Transaccion

	if err := r.db.Read(&ts); err != nil {
		return fmt.Errorf(FailReading)
	}

	deleted := false
	var index int
	for i := range ts {
		if ts[i].ID == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf(TransaccionNotFound, id)
	}

	ts = append(ts[:index], ts[index+1:]...)

	if err := r.db.Write(ts); err != nil {
		return fmt.Errorf(FailWriting, err)
	}

	return nil
}

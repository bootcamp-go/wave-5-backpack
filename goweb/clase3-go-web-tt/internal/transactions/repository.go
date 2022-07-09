package transactions

import (
	"fmt"

	"goweb/clase3-go-web-tt/internal/domain"
	"goweb/clase3-go-web-tt/pkg/bank"
)

//	CONSTANTS
const ( //	ERRORS = messages
	TransactionNotFound = "transaction with id: %d, not found 😵‍💫"
	FailReading         = "cant read database 🫠"
	FailWriting         = "cant write database 😱, error: %w"
)

//	INTERFACES
type Repository interface {
	GetAll() ([]domain.Transaction, error)
	Ecommerce(id int, codeTra string, coin string, monto float64, emisor string,
		receptor string, fecha string) (domain.Transaction, error)
	LastID() (int, error)
	GetOne(id int) (domain.Transaction, error)
	Update(id int, codeTra string, coin string, monto float64, emisor string,
		receptor string, fecha string) (domain.Transaction, error)
	UpdateOne(id int, codeTra string, monto float64) (domain.Transaction, error)
	Delete(id int) error
}

type repository struct {
	db bank.Bank
}

//	FUNCTIONS
func NewRepository(db bank.Bank) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Transaction, error) {
	var ts []domain.Transaction
	if err := r.db.Read(&ts); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return ts, nil
}

func (r *repository) Ecommerce(id int, codeTra string, coin string, monto float64,
	emisor string, receptor string, fecha string) (domain.Transaction, error) {

	var ts []domain.Transaction

	if err := r.db.Read(&ts); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailReading)
	}

	t := domain.Transaction{Id: id, CodigoTransaccion: codeTra, Moneda: coin, Monto: monto,
		Emisor: emisor, Receptor: receptor, Fecha: fecha}
	ts = append(ts, t)

	if err := r.db.Write(ts); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailWriting, err)
	}

	return t, nil
}

func (r *repository) LastID() (int, error) {
	var ps []domain.Transaction
	if err := r.db.Read(&ps); err != nil {
		return 0, fmt.Errorf(FailReading)
	}
	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].Id, nil
}

func (r *repository) GetOne(id int) (domain.Transaction, error) {

	var ts []domain.Transaction
	if err := r.db.Read(&ts); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailReading)
	}

	for _, transaction := range ts {
		if id == transaction.Id {
			return transaction, nil
		}
	}
	return domain.Transaction{}, fmt.Errorf(FailReading)
}

func (r *repository) Update(id int, codeTra string, coin string, monto float64, emisor string,
	receptor string, fecha string) (domain.Transaction, error) {

	var ts []domain.Transaction

	if err := r.db.Read(&ts); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailReading)
	}

	update := false
	tmpT := domain.Transaction{
		CodigoTransaccion: codeTra,
		Moneda:            coin,
		Monto:             monto,
		Emisor:            emisor,
		Receptor:          receptor,
		Fecha:             fecha,
	}

	for i := range ts {
		if ts[i].Id == id {
			tmpT.Id = id
			ts[i] = tmpT
			update = true
			break
		}
	}

	if !update {
		return domain.Transaction{}, fmt.Errorf(TransactionNotFound, id)
	}

	if err := r.db.Write(ts); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailWriting, err)
	}

	return tmpT, nil
}

func (r *repository) UpdateOne(id int, codeTra string, monto float64) (domain.Transaction, error) {
	var ts []domain.Transaction
	if err := r.db.Read(&ts); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailReading)
	}

	update := false
	var tmpT domain.Transaction
	for i := range ts {
		if ts[i].Id == id {
			ts[i].CodigoTransaccion = codeTra
			ts[i].Monto = monto
			update = true
			tmpT = ts[i]
			break
		}
	}

	if !update {
		return domain.Transaction{}, fmt.Errorf(TransactionNotFound, id)
	}
	if err := r.db.Write(ts); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailWriting, err)
	}

	return tmpT, nil
}

func (r *repository) Delete(id int) error {
	var ts []domain.Transaction

	if err := r.db.Read(&ts); err != nil {
		return fmt.Errorf(FailReading)
	}

	deleted := false
	var index int

	for i := range ts {
		if ts[i].Id == id {
			index = i
			deleted = true
			break
		}
	}
	if !deleted {
		return fmt.Errorf(TransactionNotFound, id)
	}

	ts = append(ts[:index], ts[index+1:]...)

	if err := r.db.Write(ts); err != nil {
		return fmt.Errorf(FailWriting, err)
	}

	return nil
}

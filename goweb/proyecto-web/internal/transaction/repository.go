package transaction

import (
	"fmt"
	"proyecto-web/internal/domain"
	"proyecto-web/internal/transaction/interfaces"
	"proyecto-web/pkg/store"
)

const (
	TransactionNotFound = "transaction %d not found"
	FailReading         = "cant read database"
	FailWriting         = "cant write database"
)

type repository struct {
	bd store.Store
}

func NewRepository(store store.Store) interfaces.IRepository {
	return &repository{bd: store}
}

func (r *repository) GetAll() ([]domain.Transaction, error) {
	var registros []domain.Transaction

	err := r.bd.Read(&registros)
	if err != nil {
		return nil, fmt.Errorf(FailReading)
	}

	return registros, err
}

func (r *repository) GetById(id int) (domain.Transaction, error) {
	transactions, err := r.GetAll()
	if err != nil {
		return domain.Transaction{}, err
	}

	transaccionBuscada, encontrada := findById(id, transactions)
	if encontrada {
		return *transaccionBuscada, nil
	}

	return domain.Transaction{}, fmt.Errorf(TransactionNotFound, id)
}

func (r *repository) Create(codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fecha string) (domain.Transaction, error) {

	var nuevaTransaccion = domain.Transaction{
		Id:                r.generateId(),
		CodigoTransaccion: codigoTransaccion,
		Moneda:            moneda,
		Monto:             monto,
		Emisor:            emisor,
		Receptor:          receptor,
		FechaTransaccion:  fecha,
	}
	transacciones, _ := r.GetAll()

	var nuevasTransaciones = append(transacciones, nuevaTransaccion)

	if err := r.bd.Write(nuevasTransaciones); err != nil {
		return domain.Transaction{}, fmt.Errorf(FailWriting)
	}
	return nuevaTransaccion, nil
}

func (r *repository) Update(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fecha string) (domain.Transaction, error) {
	transacciones, err := r.GetAll()
	if err != nil {
		return domain.Transaction{}, err
	}
	transaccionAActualizar, encontrada := findById(id, transacciones)

	if !encontrada {
		return domain.Transaction{}, fmt.Errorf(TransactionNotFound, id)
	}

	transaccionAActualizar.CodigoTransaccion = codigoTransaccion
	transaccionAActualizar.Moneda = moneda
	transaccionAActualizar.Monto = monto
	transaccionAActualizar.Emisor = emisor
	transaccionAActualizar.Receptor = receptor
	transaccionAActualizar.FechaTransaccion = fecha

	if err := r.bd.Write(transacciones); err != nil {
		return *transaccionAActualizar, fmt.Errorf(FailWriting)
	}
	return *transaccionAActualizar, nil
}

func (r *repository) UpdateParcial(id int, codigoTransaccion string, monto float64) (domain.Transaction, error) {
	transacciones, err := r.GetAll()
	if err != nil {
		return domain.Transaction{}, err
	}
	transaccionAActualizar, encontrada := findById(id, transacciones)

	if !encontrada {
		return domain.Transaction{}, fmt.Errorf(TransactionNotFound, id)
	}

	transaccionAActualizar.CodigoTransaccion = codigoTransaccion
	transaccionAActualizar.Monto = monto

	if err := r.bd.Write(transacciones); err != nil {
		return *transaccionAActualizar, fmt.Errorf(FailWriting)
	}
	return *transaccionAActualizar, nil
}

func (r *repository) Delete(id int) error {
	var indexBuscado int = -1
	transacciones, err := r.GetAll()
	if err != nil {
		return err
	}

	for index, transaccion := range transacciones {
		if transaccion.Id == id {
			indexBuscado = index
		}
	}

	if indexBuscado < 0 {
		return fmt.Errorf(TransactionNotFound, id)
	}
	nuevasTransaciones := remove(transacciones, indexBuscado)

	if err := r.bd.Write(nuevasTransaciones); err != nil {
		return fmt.Errorf(FailWriting)
	}
	return nil
}

func (r *repository) GetByCodigoTransaccion(codigo string) (domain.Transaction, error) {
	//not implemented. Only for bd repository
	return domain.Transaction{}, nil
}

func (r *repository) generateId() int {
	transacciones, _ := r.GetAll()
	if len(transacciones) == 0 {
		return 1
	}
	lastId := transacciones[(len(transacciones) - 1)]
	return lastId.Id + 1
}

func findById(id int, transacciones []domain.Transaction) (*domain.Transaction, bool) {
	var transaccionBuscada *domain.Transaction
	var encontrada bool
	for i, transaccion := range transacciones {
		if transaccion.Id == id {
			transaccionBuscada = &transacciones[i]
			encontrada = true
			break
		}
	}
	return transaccionBuscada, encontrada
}

func remove(slice []domain.Transaction, s int) []domain.Transaction {
	return append(slice[:s], slice[s+1:]...)
}

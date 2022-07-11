package transaction

import (
	"errors"
	"fmt"
	"proyecto-web/internal/domain"
	"proyecto-web/pkg/store"
)

type IRepository interface {
	GetAll() []domain.Transaction
	Create(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fecha string) domain.Transaction
	GetById(id int) (domain.Transaction, error)
	Update(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fecha string) (domain.Transaction, error)
	UpdateParcial(id int, codigoTransaccion string, monto float64) (domain.Transaction, error)
	Delete(id int) error
}

type repository struct {
	bd store.Store
}

func NewRepository(store store.Store) IRepository {
	return &repository{bd: store}
}

func (r *repository) GetAll() []domain.Transaction {
	var registros []domain.Transaction
	r.bd.Read(&registros)
	return registros
}

func (r *repository) GetById(id int) (domain.Transaction, error) {
	transaccionBuscada, encontrada := findById(id, r.GetAll())

	if encontrada {
		return *transaccionBuscada, nil
	}
	return domain.Transaction{}, errors.New("Recurso no encontrado")
}

func (r *repository) Create(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fecha string) domain.Transaction {

	var nuevaTransaccion = domain.Transaction{
		Id:                r.generateId(),
		CodigoTransaccion: codigoTransaccion,
		Moneda:            moneda,
		Monto:             monto,
		Emisor:            emisor,
		Receptor:          receptor,
		FechaTransaccion:  fecha,
	}
	var nuevasTransaciones = append(r.GetAll(), nuevaTransaccion)
	r.bd.Write(nuevasTransaciones)
	return nuevaTransaccion
}

func (r *repository) Update(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fecha string) (domain.Transaction, error) {
	transacciones := r.GetAll()
	transaccionAActualizar, encontrada := findById(id, transacciones)

	if !encontrada {
		return domain.Transaction{}, errors.New("no se encontro el recurso a actualizar")
	}

	transaccionAActualizar.CodigoTransaccion = codigoTransaccion
	transaccionAActualizar.Moneda = moneda
	transaccionAActualizar.Monto = monto
	transaccionAActualizar.Emisor = emisor
	transaccionAActualizar.Receptor = receptor
	transaccionAActualizar.FechaTransaccion = fecha

	r.bd.Write(transacciones)
	return *transaccionAActualizar, nil
}

func (r *repository) UpdateParcial(id int, codigoTransaccion string, monto float64) (domain.Transaction, error) {
	transacciones := r.GetAll()
	transaccionAActualizar, encontrada := findById(id, transacciones)

	if !encontrada {
		return domain.Transaction{}, errors.New("No se encontro el recurso a actualizar")
	}

	transaccionAActualizar.CodigoTransaccion = codigoTransaccion
	transaccionAActualizar.Monto = monto

	r.bd.Write(transacciones)
	return *transaccionAActualizar, nil
}

func (r *repository) Delete(id int) error {
	var indexBuscado int = -1
	transacciones := r.GetAll()

	for index, transaccion := range transacciones {
		if transaccion.Id == id {
			indexBuscado = index
		}
	}
	if indexBuscado < 0 {
		fmt.Println("no encontrado")
		return errors.New("no se encontró el recurso a eliminar")
	}
	nuevasTransaciones := remove(transacciones, indexBuscado)

	return r.bd.Write(nuevasTransaciones)
}

func (r *repository) generateId() int {
	transacciones := r.GetAll()
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

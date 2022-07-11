package transaction

import (
	"errors"
	"proyecto-web/internal/domain"
	"proyecto-web/pkg/store"
)

type IRepository interface {
	GetAll() []domain.Transaction
	Create(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fecha string) domain.Transaction
	GetById(id int) (domain.Transaction, error)
	Update(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fecha string) (domain.Transaction, error)
	UpdateParcial(id int, codigoTransaccion string, monto float64) (domain.Transaction, error)
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
	transaccionBuscada, encontrada := r.findById(id)

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
	transaccionAActualizar, encontrada := r.findById(id)

	if !encontrada {
		return domain.Transaction{}, errors.New("No se encontro el recurso a actualizar")
	}

	transaccionAActualizar.CodigoTransaccion = codigoTransaccion
	transaccionAActualizar.Moneda = moneda
	transaccionAActualizar.Monto = monto
	transaccionAActualizar.Emisor = emisor
	transaccionAActualizar.Receptor = receptor
	transaccionAActualizar.FechaTransaccion = fecha

	return *transaccionAActualizar, nil
}

func (r *repository) UpdateParcial(id int, codigoTransaccion string, monto float64) (domain.Transaction, error) {
	transaccionAActualizar, encontrada := r.findById(id)

	if !encontrada {
		return domain.Transaction{}, errors.New("No se encontro el recurso a actualizar")
	}

	transaccionAActualizar.CodigoTransaccion = codigoTransaccion
	transaccionAActualizar.Monto = monto

	return *transaccionAActualizar, nil
}

func (r *repository) generateId() int {
	transacciones := r.GetAll()
	if len(transacciones) == 0 {
		return 1
	}
	lastId := transacciones[(len(transacciones) - 1)]
	return lastId.Id + 1
}

func (r *repository) findById(id int) (*domain.Transaction, bool) {
	transacciones := r.GetAll()
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

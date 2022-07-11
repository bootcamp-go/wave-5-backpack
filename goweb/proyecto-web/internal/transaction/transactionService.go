package transaction

import "proyecto-web/internal/domain"

type ITransactionService interface {
	GetAll() []domain.Transaction
	Create(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fecha string) domain.Transaction
	GetById(id int) (domain.Transaction, error)
	Update(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fecha string) (domain.Transaction, error)
	UpdateParcial(id int, codigoTransaccion string, monto float64) (domain.Transaction, error)
	Delete(id int) error
}

type transactionService struct {
	repository IRepository
}

func NewService(r IRepository) ITransactionService {
	return &transactionService{
		repository: r,
	}
}

func (s *transactionService) GetAll() []domain.Transaction {
	return s.repository.GetAll()
}

func (s *transactionService) Create(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fecha string) domain.Transaction {
	return s.repository.Create(id, codigoTransaccion, moneda, monto, emisor, receptor, fecha)
}

func (s *transactionService) GetById(id int) (domain.Transaction, error) {
	return s.repository.GetById(id)
}

func (s *transactionService) Update(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fecha string) (domain.Transaction, error) {
	return s.repository.Update(id, codigoTransaccion, moneda, monto, emisor, receptor, fecha)
}

func (s *transactionService) UpdateParcial(id int, codigoTransaccion string, monto float64) (domain.Transaction, error) {
	return s.repository.UpdateParcial(id, codigoTransaccion, monto)
}

func (s *transactionService) Delete(id int) error {
	return s.repository.Delete(id)
}

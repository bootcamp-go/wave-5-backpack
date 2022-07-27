package transactions

import (
	"goweb/internal/domain"
	"goweb/internal/dto/responsedto"
)

type IService interface {
	Delete(int) error
	GetAll() ([]domain.Transaction, error)
	Store(string, string, int, string, string, string) (responsedto.TransactionResponse, error)
	Update(int, string, string, int, string, string, string) (domain.Transaction, error)
	UpdateAmount(int, int) (domain.Transaction, error)
}

type Service struct {
	repository IRepository
}

func NewService(repository IRepository) IService {
	return &Service{
		repository: repository,
	}
}

func (service *Service) Delete(id int) error {
	return service.repository.Delete(id)
}

func (service *Service) GetAll() ([]domain.Transaction, error) {
	transactions, err := service.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (service *Service) Store(codTransaction, currency string, amount int, sender, receiver, dateOrder string) (responsedto.TransactionResponse, error) {
	lastID, err := service.repository.LastID()
	if err != nil {
		return responsedto.TransactionResponse{}, err
	}
	lastID++

	transaction, err := service.repository.Store(lastID, codTransaction, currency, amount, sender, receiver, dateOrder)
	if err != nil {
		return responsedto.TransactionResponse{}, err
	}

	var transactionResponse responsedto.TransactionResponse
	transactionResponse.CodTransaction = transaction.CodTransaction
	transactionResponse.Currency = transaction.Currency
	transactionResponse.Amount = transaction.Amount
	transactionResponse.Sender = transaction.Sender
	transactionResponse.Receiver = transaction.Receiver
	transactionResponse.DateOrder = transaction.DateOrder

	return transactionResponse, nil
}

func (service *Service) Update(id int, codTransaction, currency string, amount int, sender, receiver, dateOrder string) (domain.Transaction, error) {
	return service.repository.Update(id, codTransaction, currency, amount, sender, receiver, dateOrder)
}

func (service *Service) UpdateAmount(id, amount int) (domain.Transaction, error) {
	return service.repository.UpdateAmount(id, amount)
}
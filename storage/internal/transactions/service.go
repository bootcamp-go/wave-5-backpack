package transactions

import (
	"context"
	"goweb/internal/domain"
	"goweb/internal/dto/responsedto"
)

type IService interface {
	Delete(context.Context, int) error
	GetAll(context.Context) ([]domain.Transaction, error)
	GetBySender(context.Context, string) (domain.Transaction, error)
	Store(context.Context, string, string, int, string, string, string) (responsedto.TransactionResponse, error)
	Update(context.Context, int, string, string, int, string, string, string) (domain.Transaction, error)
	UpdateAmount(context.Context, int, int) (domain.Transaction, error)
}

type Service struct {
	repository IRepository
}

func NewService(repository IRepository) IService {
	return &Service{
		repository: repository,
	}
}

func (service *Service) Delete(ctx context.Context, id int) error {
	return service.repository.Delete(ctx, id)
}

func (service *Service) GetAll(ctx context.Context) ([]domain.Transaction, error) {
	transactions, err := service.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func(service *Service) GetBySender(ctx context.Context, sender string) (domain.Transaction, error) {
	return service.repository.GetBySender(ctx, sender)
}

func (service *Service) Store(ctx context.Context, codTransaction, currency string, amount int, sender, receiver, dateOrder string) (responsedto.TransactionResponse, error) {
	lastID, err := service.repository.LastID(ctx)
	if err != nil {
		return responsedto.TransactionResponse{}, err
	}
	lastID++

	transaction, err := service.repository.Store(ctx, lastID, codTransaction, currency, amount, sender, receiver, dateOrder)
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

func (service *Service) Update(ctx context.Context, id int, codTransaction, currency string, amount int, sender, receiver, dateOrder string) (domain.Transaction, error) {
	return service.repository.Update(ctx, id, codTransaction, currency, amount, sender, receiver, dateOrder)
}

func (service *Service) UpdateAmount(ctx context.Context, id, amount int) (domain.Transaction, error) {
	return service.repository.UpdateAmount(ctx, id, amount)
}
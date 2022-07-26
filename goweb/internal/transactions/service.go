package transactions

import (
	"errors"
	"fmt"
	"goweb/internal/domain"
	"goweb/pkg/utils"
	"time"
)

type Service interface {
	GetAll() ([]domain.Transaction, error)
	Store(Currency string, Amount float64, Sender string, Reciever string) (domain.Transaction, error)
	GetById(Id int) (domain.Transaction, error)
	Update(id int, Currency string, Amount float64, Sender string, Reciever string) (domain.Transaction, error)
	UpdateCurrencyAndAmount(id int, Currency string, Amount float64) (domain.Transaction, error)
	Delete(id int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAll() ([]domain.Transaction, error) {
	transactions, err := s.repo.GetAll()
	if err != nil {
		return nil, errors.New("error: transactions cannot be served")
	}
	return transactions, nil
}

func isAmountZeroOrNegative(amount float64) bool {
	return amount <= 0
}

func (s *service) Store(Currency string, Amount float64, Sender string, Reciever string) (domain.Transaction, error) {
	lastId, err := s.repo.lastId()
	if err != nil {
		return domain.Transaction{}, errors.New("error: cannot generate ID")
	}

	Id := lastId + 1
	TransactionDate := time.Now()
	TransactionCode := utils.RandomString(30)
	if isAmountZeroOrNegative(Amount) {
		return domain.Transaction{}, &NotAllowedAmountZeroOrNegative{}
	}
	transaction, err := s.repo.Store(
		Id,
		TransactionCode,
		Currency,
		Amount,
		Sender,
		Reciever,
		TransactionDate,
	)
	if err != nil {
		return domain.Transaction{}, errors.New("error: cannot generate transaction")
	}
	return transaction, nil
}

func (s *service) GetById(Id int) (domain.Transaction, error) {
	transaction, err := s.repo.GetById(Id)
	if err != nil {
		return domain.Transaction{}, errors.New("error: transaction not found")
	}
	return transaction, nil
}

func (s *service) Update(id int, Currency string, Amount float64, Sender string, Reciever string) (domain.Transaction, error) {
	if isAmountZeroOrNegative(Amount) {
		return domain.Transaction{}, &NotAllowedAmountZeroOrNegative{}
	}
	transaction, err := s.repo.Update(
		id,
		Currency,
		Amount,
		Sender,
		Reciever,
	)
	if err != nil {
		return domain.Transaction{}, errors.New("error: cannot update transaction")
	}
	return transaction, nil
}

func (s *service) Delete(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("error: cannot be deleted id %d %w", id, err)
	}
	return nil
}

func (s *service) UpdateCurrencyAndAmount(id int, Currency string, Amount float64) (domain.Transaction, error) {
	if isAmountZeroOrNegative(Amount) {
		return domain.Transaction{}, &NotAllowedAmountZeroOrNegative{}
	}
	transaction, err := s.repo.UpdateCurrencyAndAmount(id, Currency, Amount)
	if err != nil {
		return domain.Transaction{}, errors.New("error: transaction not found")
	}
	return transaction, nil
}

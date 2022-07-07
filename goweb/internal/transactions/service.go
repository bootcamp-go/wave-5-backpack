package transactions

import (
	"errors"
	"fmt"
	"goweb/internal/domain"
	"math/rand"
	"time"
)

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}

type Service interface {
	GetAll() ([]domain.Transaction, error)
	Store(Currency string, Amount float64, Sender string, Reciever string) (domain.Transaction, error)
	GetById(Id int) (domain.Transaction, error)
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
	TransactionCode := randomString(30)
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

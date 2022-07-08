package transactions

import "GoWeb/internals/domain"

type Service interface {
	GetAll() ([]domain.Transanction, error)
	Store(code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]domain.Transanction, error) {
	tran, err := s.repository.GetAll()

	if err != nil {
		return []domain.Transanction{}, err
	}

	return tran, nil
}

func (s *service) Store(code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error) {
	lastID, err := s.repository.lastID()

	if err != nil {
		return domain.Transanction{}, nil
	}
	lastID++

	transaccion, err := s.repository.Store(lastID, code, coin, amount, emisor, receptor, date)

	if err != nil {
		return domain.Transanction{}, err
	}

	return transaccion, nil
}

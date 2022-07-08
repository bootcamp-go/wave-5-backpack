package transactions

import "GoWeb/internals/domain"

type Service interface {
	GetAll() ([]domain.Transanction, error)
	Store(code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error)
	Update(id int, code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error)
	Delete(id int) error
	UpdateCode(id int, code string, amount float64) (domain.Transanction, error)
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

func (s *service) Update(id int, code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error) {
	tran, err := s.repository.Update(id, code, coin, amount, emisor, receptor, date)

	if err != nil {
		return domain.Transanction{}, err
	}

	return tran, nil
}
func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateCode(id int, code string, amount float64) (domain.Transanction, error) {
	return s.repository.UpdateCode(id, code, amount)
}

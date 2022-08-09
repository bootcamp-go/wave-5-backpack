package transactions

import (
	"GoWeb/internals/domain"
	"context"
)

type Service interface {
	GetAll() ([]domain.Transanction, error)
	Store(code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error)
	Update(id int, code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error)
	Delete(id int) (domain.Transanction, error)
	UpdateCode(id int, code string, amount float64) (domain.Transanction, error)
	GetById(id int) (domain.Transanction, error)
	GetByName(name string) ([]domain.Transanction, error)
	GetByIdCtx(ctx context.Context, id int) (domain.Transanction, error)
	UpdateCtx(ctx context.Context, id int, code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error)
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

	transaccion, err := s.repository.Store(code, coin, amount, emisor, receptor, date)

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
func (s *service) Delete(id int) (domain.Transanction, error) {

	tran, err := s.repository.Delete(id)
	if err != nil {
		return domain.Transanction{}, err
	}

	return tran, nil
}

func (s *service) UpdateCode(id int, code string, amount float64) (domain.Transanction, error) {

	tran, err := s.repository.UpdateCode(id, code, amount)
	if err != nil {
		return domain.Transanction{}, err
	}
	return tran, nil
}

func (s *service) GetById(id int) (domain.Transanction, error) {
	tran, err := s.repository.GetById(id)
	if err != nil {
		return domain.Transanction{}, err
	}

	return tran, nil
}

func (s *service) GetByName(name string) ([]domain.Transanction, error) {

	tran, err := s.repository.GetByName(name)
	if err != nil {
		return []domain.Transanction{}, err
	}
	return tran, nil
}

func (s *service) GetByIdCtx(ctx context.Context, id int) (domain.Transanction, error) {
	tran, err := s.repository.GetById(id)
	if err != nil {
		return domain.Transanction{}, err
	}

	return tran, nil
}

func (s *service) UpdateCtx(ctx context.Context, id int, code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error) {
	tran, err := s.repository.UpdateCtx(ctx, id, code, coin, amount, emisor, receptor, date)

	if err != nil {
		return domain.Transanction{}, err
	}

	return tran, nil
}

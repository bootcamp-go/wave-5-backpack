package users

import (
	"context"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Users, error)
	GetByName(ctx context.Context, name string) ([]domain.Users, error)
	Store(ctx context.Context, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error)
	Update(ctx context.Context, id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error)
	UpdateLastNameAndAge(ctx context.Context, id, age int, lastName string) (domain.Users, error)
	Delete(ctx context.Context, id int) error
}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Users, error) {
	ps, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) GetByName(ctx context.Context, name string) ([]domain.Users, error) {
	return s.repository.GetByName(ctx, name)
}

func (s *service) Store(ctx context.Context, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	//TODO: arreglar con repository sin LastID, 0  dummy value
	user, err := s.repository.Store(ctx, 0, age, name, lastName, email, creationDate, height, active)
	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
}

func (s *service) Update(ctx context.Context, id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	return s.repository.Update(ctx, id, age, name, lastName, email, creationDate, height, active)
}

func (s *service) UpdateLastNameAndAge(ctx context.Context, id, age int, lastName string) (domain.Users, error) {
	return s.repository.UpdateLastNameAndAge(ctx, id, age, lastName)
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

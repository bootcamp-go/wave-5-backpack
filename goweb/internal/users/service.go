package users

import (
	"context"
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.User, error)
	StoreUser(name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error)
	GetById(ctx context.Context,id int) (domain.User, error)
	UpdateUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error)
	DeleteUser(id int) error
	UpdateLastnameAndAge(id int, lastname string, age int) (*domain.User, error)
	GetByName(name string) ([]domain.User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(ctx context.Context) ([]domain.User, error) {
	us, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (s *service) GetById (ctx context.Context, id int) (domain.User, error) {
	user, err := s.repository.GetById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *service) StoreUser(name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error) {
	lastID, err := s.repository.LastId()
	if err != nil {
		return domain.User{}, err
	}

	lastID++

	newUser, err := s.repository.StoreUser(lastID, name, lastname, email, age, height, active, doCreation)

	if err != nil {
		return domain.User{}, err
	}

	return newUser, nil
}
func (s *service) UpdateUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error) {
	updatedUser, err := s.repository.UpdateUser(id, name, lastname, email, age, height, active, doCreation)

	if err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}

func (s *service) DeleteUser(id int) error {
	err := s.repository.DeleteUser(id)
	if err != nil {
		return fmt.Errorf("error deleting user %w", err)
	}
	return nil
}


func(s *service) UpdateLastnameAndAge(id int, lastname string, age int) (*domain.User, error){
	updatedUser, err := s.repository.UpdateLastnameAndAge(id, lastname, age)
	if err != nil {
		return nil, fmt.Errorf("error deleting user %w", err)
	}
	return updatedUser, nil
}

func (s *service) GetByName(name string) ([]domain.User, error) {
	return s.repository.GetByName(name)
}
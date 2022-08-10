package users

import (
	"clase3-storage-implementation-tm/internal/domain"
	"context"

	"github.com/google/uuid"
)

// Service ...
type Service interface {
	GetOne(id string) (domain.User, error)
	Store(firstname, lastname, username, email string) (domain.User, error)
	Update(id, firstname, lastname, username, email string) (domain.User, error)
	Delete(id string) error
}

type service struct {
	repository Repository
}

// NewService ...
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetOne(id string) (domain.User, error) {
	user, err := s.repository.GetOne(context.Background(), id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *service) Store(firstname, lastname, username, email string) (domain.User, error) {
	id := uuid.New().String()
	user := domain.User{
		ID:        id,
		Firstname: firstname,
		Lastname:  lastname,
		Username:  username,
		Email:     email,
	}
	err := s.repository.Store(context.Background(), &user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *service) Update(id, firstname, lastname, username, email string) (domain.User, error) {
	return s.repository.Update(id, firstname, lastname, username, email)
}

func (s *service) Delete(id string) error {
	return s.repository.Delete(context.Background(), id)
}

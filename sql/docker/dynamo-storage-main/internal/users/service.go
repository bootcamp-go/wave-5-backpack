package users

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	GetOne(id string) (User, error)
	Store(firstname, lastname, username, email string) (User, error)
	Delete(id string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetOne(id string) (User, error) {
	user, err := s.repository.GetOne(context.Background(), id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *service) Store(firstname, lastname, username, email string) (User, error) {
	id := uuid.New().String()
	user := User{
		Id:        id,
		Firstname: firstname,
		Lastname:  lastname,
		Username:  username,
		Email:     email,
	}
	err := s.repository.Store(context.Background(), &user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s *service) Delete(id string) error {
	return s.repository.Delete(context.Background(), id)
}

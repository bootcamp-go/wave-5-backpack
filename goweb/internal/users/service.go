package users

import (
	"errors"
	"fmt"
	"goweb/internal/domain"
	"time"
)

type Service interface {
	GetAll() ([]domain.User, error)
	Store(Name string, LastName string, Email string, Age int, Height float64, Active bool) (domain.User, error)
	Update(Id int, Name string, LastName string, Email string, Age int, Height float64, Active bool) (domain.User, error)
	Delete(Id int) error
	GetById(Id int) (domain.User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAll() ([]domain.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, errors.New("error loading users")
	}

	return users, nil
}

func (s *service) Store(Name string, LastName string, Email string, Age int, Height float64, Active bool) (domain.User, error) {
	lastId, err := s.repo.LastId()
	if err != nil {
		return domain.User{}, errors.New("error: can´t generate ID")
	}

	Id := lastId + 1
	creationDate := time.Now()

	user, err := s.repo.Store(
		Id,
		Name,
		LastName,
		Email,
		Age,
		Height,
		Active,
		creationDate,
	)
	if err != nil {
		return domain.User{}, errors.New("error: can't generate user")
	}

	return user, nil
}

func (s *service) Update(Id int, Name string, LastName string, Email string, Age int, Height float64, Active bool) (domain.User, error) {
	user, err := s.repo.Update(
		Id,
		Name,
		LastName,
		Email,
		Age,
		Height,
		Active,
	)

	if err != nil {
		return domain.User{}, errors.New(("error: can't update user"))
	}

	return user, nil
}

func (s *service) Delete(Id int) error {
	if err := s.repo.Delete(Id); err != nil {
		return fmt.Errorf("error: can´t be deleted Id %d %w", Id, err)
	}

	return nil
}

func (s *service) GetById(Id int) (domain.User, error) {
	user, err := s.repo.GetById(Id)
	if err != nil {
		return domain.User{}, errors.New("error: user not found")
	}

	return user, nil
}

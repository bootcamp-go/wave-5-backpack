package users

import (
	"fmt"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/domain"
)

type Service interface {
	GetById(Id int) (domain.User, error)
	GetAll(filters map[string]interface{}) ([]domain.User, error)
	Store(Age int, FirstName, LastName, Email string, Height float64, Active bool) (domain.User, error)
	Update(Id, Age int, FirstName, LastName, Email, CreatedAt string, Height float64, Active bool) (domain.User, error)
	UpdateAgeLastName(Id, Age int, LastName string) (domain.User, error)
	Delete(Id int) error
}

type service struct {
	repository Repository
}

func (s *service) GetById(Id int) (domain.User, error) {
	return s.repository.GetById(Id)
}

func (s *service) GetAll(filters map[string]interface{}) ([]domain.User, error) {
	return s.repository.GetAll(filters)
}

func (s *service) Store(Age int, FirstName, LastName, Email string, Height float64, Active bool) (domain.User, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.User{}, err
	}

	Id := lastID + 1

	t := time.Now()
	CreatedAt := fmt.Sprintf("%02d/%02d/%d", t.Day(), t.Month(), t.Year())

	user, err := s.repository.Store(Id, Age, FirstName, LastName, Email, CreatedAt, Height, Active)
	return user, err
}

func (s *service) Update(Id, Age int, FirstName, LastName, Email, CreatedAt string, Height float64, Active bool) (domain.User, error) {
	return s.repository.Update(Id, Age, FirstName, LastName, Email, CreatedAt, Height, Active)
}

func (s *service) UpdateAgeLastName(Id, Age int, LastName string) (domain.User, error) {
	return s.repository.UpdateAgeLastName(Id, Age, LastName)
}

func (s *service) Delete(Id int) error {
	return s.repository.Delete(Id)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

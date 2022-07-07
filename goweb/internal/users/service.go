package users

import (
	"fmt"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/domain"
)

type Service interface {
	GetById(Id int) (domain.User, error)
	GetAll() ([]domain.User, error)
	Store(Age int, FirstName, LastName, Email string, Height float64, Active bool) (domain.User, error)
	FilterUsers(filters map[string]interface{}, users []domain.User) (*[]domain.User, error)
}

type service struct {
	repository Repository
}

func (s *service) GetById(Id int) (domain.User, error) {
	return s.repository.GetById(Id)
}

func (s *service) GetAll() ([]domain.User, error) {
	return s.repository.GetAll()
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

func (s *service) FilterUsers(filters map[string]interface{}, users []domain.User) (*[]domain.User, error) {
	resultUsers := []domain.User{}

	for _, user := range users {
		if Id, ok := filters["Id"]; ok && Id != user.Id {
			continue
		}
		if Age, ok := filters["Age"]; ok && Age != user.Age {
			continue
		}
		if FirstName, ok := filters["FirstName"]; ok && FirstName != user.FirstName {
			continue
		}
		if LastName, ok := filters["LastName"]; ok && LastName != user.LastName {
			continue
		}
		if Email, ok := filters["Email"]; ok && Email != user.Email {
			continue
		}
		if CreatedAt, ok := filters["CreatedAt"]; ok && CreatedAt != user.CreatedAt {
			continue
		}
		if Height, ok := filters["Height"]; ok && Height != user.Height {
			continue
		}
		if Active, ok := filters["Active"]; ok && Active != user.Active {
			continue
		}

		resultUsers = append(resultUsers, user)
	}

	return &resultUsers, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

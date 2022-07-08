package users

import (
	"goweb/internal/domain"
)

type Service interface{
	GetAllUsers() ([]domain.User, error)
	GetUserById(id int) (domain.User, error)
	StoreUser(name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error)
	UpdateTotal(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAllUsers() ([]domain.User, error) {
	users, err := s.repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *service) GetUserById(id int) (domain.User, error) {
	user, err := s.repository.GetUserById(id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *service) StoreUser(name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error){
	//lo primero que hago es generar un ID
	lastID, err := s.repository.LastID()
	if err!= nil{
		return domain.User{}, err
	}

	lastID++

	newUser, err := s.repository.StoreUser(lastID, name, lastname, email, age, height, active, createdat)
	if err!= nil{
		return domain.User{}, err
	}

	return newUser, nil
}

func (s *service) UpdateTotal(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error) {
	return s.repository.UpdateTotal(id, name, lastname, email, age, height, active, createdat)
}



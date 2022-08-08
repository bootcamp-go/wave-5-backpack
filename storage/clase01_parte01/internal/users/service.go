package users

import (
	"goweb/internal/domain"
	"fmt"
)

type Service interface{
	GetAllUsers() ([]domain.User, error)
	GetUserById(id int) (domain.User, error)
	StoreUser(name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error)
	UpdateTotal(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error)
	UpdatePartial(id int, lastname string, age int) (domain.User, error)
	Delete(id int) error
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
	//lo primero que hago es generar un ID. Lo unico que cambió respecto a la versión anterior es agregar un msj para identificar mejor el error
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.User{}, fmt.Errorf("error getting product last id: %w", err)
	}

	lastID++

	newUser, err := s.repository.StoreUser(lastID, name, lastname, email, age, height, active, createdat)
	if err!= nil{
		return domain.User{}, fmt.Errorf("error creating user: %w", err)
	}

	return newUser, nil
}

func (s *service) UpdateTotal(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error) {
	user, err := s.repository.UpdateTotal(id, name, lastname, email, age, height, active, createdat)
	if err != nil {
		return domain.User{}, fmt.Errorf("error updating user: %w", err)
	}
	return user, nil
}

func (s *service) UpdatePartial(id int, lastname string, age int) (domain.User, error){
	user, err := s.repository.UpdatePartial(id, lastname, age)
	if err != nil {
		return domain.User{}, fmt.Errorf("error updating user: %w", err)
	}
	return user, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}




package users

import (
	"goweb/internal/domain"
	"fmt"
)

type Service interface{
	// metodos viejos
	GetAllUsers() ([]domain.User, error)
	GetUserById(id int) (domain.User, error)
	UpdateTotal(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error)
	UpdatePartial(id int, lastname string, age int) (domain.User, error)
	Delete(id int) error
	
	//métodos nuevos o actualizados
	GetUserByName(name string) (domain.User, error)
	StoreUser(name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetUserByName(name string) (domain.User, error){
	user, err := s.repository.GetUserByName(name)
	if err!= nil{
		return domain.User{}, fmt.Errorf("error searching user: %w", err)
	}
	// esta modificación la hago porque si la DB no encuentra el usuario, devuelve un vacío pero no un error
	if user.Id == 0{
		return domain.User{}, fmt.Errorf("user not found")
	}
	
	return user, nil
	
}

func (s *service) GetUserById(id int) (domain.User, error) {

	user, err := s.repository.GetUserById(id)
	if err!= nil{
		return domain.User{}, fmt.Errorf("error searching user: %w", err)
	}
	// esta modificación la hago porque si la DB no encuentra el usuario, devuelve un vacío pero no un error
	if user.Id == 0{
		return domain.User{}, fmt.Errorf("user not found")
	}
	
	return user, nil
}

func (s *service) StoreUser(name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error){
	
	newUser, err := s.repository.StoreUser(name, lastname, email, age, height, active, createdat)
	if err!= nil{
		return domain.User{}, fmt.Errorf("error creating user: %w", err)
	}
	
	return newUser, nil
}

func (s *service) UpdateTotal(id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error) {
	userUpdated, err := s.repository.UpdateTotal(id, name, lastname, email, age, height, active, createdat)
	if err!= nil{
		return domain.User{}, fmt.Errorf("error updating user: %w", err)
	}
	
	return userUpdated, nil
}






func (s *service) GetAllUsers() ([]domain.User, error) {
	return []domain.User{}, nil
}



func (s *service) UpdatePartial(id int, lastname string, age int) (domain.User, error){
	return domain.User{}, nil
}

func (s *service) Delete(id int) error {
	return nil
}




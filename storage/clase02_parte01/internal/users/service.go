package users

import (
	"fmt"
	"goweb/internal/domain"

	"golang.org/x/net/context"
)

type Service interface{
	// metodos viejos
	UpdatePartial(id int, lastname string, age int) (domain.User, error)
	Delete(id int) error
	
	//métodos nuevos o actualizados
	GetAllUsers(ctx context.Context) ([]domain.User, error)
	GetUserById(ctx context.Context, id int) (domain.User, error)
	UpdateTotal(ctx context.Context, id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error)
	GetUserByName(ctx context.Context, name string) (domain.User, error)
	StoreUser(ctx context.Context, name, lastname, email string, age int, height float32, active bool) (domain.User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	users, err := s.repository.GetAllUsers(ctx)
	if err!= nil{
		return nil, fmt.Errorf("error searching user: %w", err)
	}
	// esta modificación la hago porque si la DB no encuentra el usuario, devuelve un vacío pero no un error
	if len(users) == 0{
		return nil, fmt.Errorf("user not found")
	}
	
	return users, nil
}

func (s *service) GetUserByName(ctx context.Context, name string) (domain.User, error){
	user, err := s.repository.GetUserByName(ctx, name)
	if err!= nil{
		return domain.User{}, fmt.Errorf("error searching user: %w", err)
	}
	// esta modificación la hago porque si la DB no encuentra el usuario, devuelve un vacío pero no un error
	if user.Id == 0{
		return domain.User{}, fmt.Errorf("user not found")
	}
	
	return user, nil
	
}

func (s *service) GetUserById(ctx context.Context, id int) (domain.User, error) {

	user, err := s.repository.GetUserById(ctx, id)
	if err!= nil{
		return domain.User{}, fmt.Errorf("error searching user: %w", err)
	}
	// esta modificación la hago porque si la DB no encuentra el usuario, devuelve un vacío pero no un error
	if user.Id == 0{
		return domain.User{}, fmt.Errorf("user not found")
	}
	
	return user, nil
}

func (s *service) StoreUser(ctx context.Context, name, lastname, email string, age int, height float32, active bool) (domain.User, error){
	
	newUser, err := s.repository.StoreUser(ctx, name, lastname, email, age, height, active)
	if err!= nil{
		return domain.User{}, fmt.Errorf("error creating user: %w", err)
	}
	
	return newUser, nil
}

func (s *service) UpdateTotal(ctx context.Context, id int, name, lastname, email string, age int, height float32, active bool, createdat string) (domain.User, error) {
	userUpdated, err := s.repository.UpdateTotal(ctx, id, name, lastname, email, age, height, active, createdat)
	if err!= nil{
		return domain.User{}, fmt.Errorf("error updating user: %w", err)
	}
	
	return userUpdated, nil
}


func (s *service) UpdatePartial(id int, lastname string, age int) (domain.User, error){
	return domain.User{}, nil
}

func (s *service) Delete(id int) error {
	return nil
}




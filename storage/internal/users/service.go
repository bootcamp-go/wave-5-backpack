package users

import (
	"context"
	"errors"
	"fmt"
	"goweb/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.User, error)
	Store(ctx context.Context, Nombre string, Apellido string, Email string, Edad int, Altura float64, Activo bool, FechaCreacion string) (domain.User, error)
	GetById(ctx context.Context, id int) (domain.User, error)
	GetByName(ctx context.Context, nombre string) (domain.User, error)
	Update(ctx context.Context, id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, id int, apellido string, edad int) (domain.User, error)
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
	users, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *service) Store(ctx context.Context, nombre string, apellido string, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error) {

	usr := domain.User{
		Nombre:        nombre,
		Apellido:      apellido,
		Email:         email,
		Edad:          edad,
		Altura:        altura,
		Activo:        activo,
		FechaCreacion: fechaCreacion,
	}

	lastID, err := s.repository.Store(ctx, usr)
	if err != nil {
		return domain.User{}, errors.New("no se pudo guardar el usuario")
	}
	usr.Id = lastID

	return usr, nil
}

func (s *service) GetById(ctx context.Context, id int) (domain.User, error) {
	user, err := s.repository.GetById(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("no se pudo encontrar el usuario con el id: %d", id)
	}
	return user, nil
}

func (s *service) GetByName(ctx context.Context, nombre string) (domain.User, error) {
	user, err := s.repository.GetByName(ctx, nombre)
	if err != nil {
		return domain.User{}, fmt.Errorf("no se pudo encontrar el usuario con el nombre: %s", nombre)
	}
	return user, nil
}

func (s *service) Update(ctx context.Context, id int, nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error) {
	usr := domain.User{
		Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, FechaCreacion: fechaCreacion,
	}
	user, err := s.repository.Update(ctx, usr)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	if err := s.repository.Delete(ctx, id); err != nil {
		return fmt.Errorf("error al eliminar el usuario %d", id)
	}
	return nil
}

func (s *service) Patch(ctx context.Context, id int, apellido string, edad int) (domain.User, error) {
	user, err := s.repository.Patch(ctx, id, apellido, edad)

	if err != nil {
		return domain.User{}, fmt.Errorf("error al actualizar datos del usuario %d", id)
	}

	return user, nil
}

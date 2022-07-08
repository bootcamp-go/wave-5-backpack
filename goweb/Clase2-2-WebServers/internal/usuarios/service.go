package usuarios

import "goweb/Clase2-2-WebServers/internal/domain"

type Service interface {
	GetAll() ([]domain.Usuario, error)
	Store(nombre, apellido, email string, edad, altura int, activo bool, fecha string) (domain.Usuario, error)
}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Usuario, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(nombre, apellido, email string, edad, altura int, activo bool, fecha string) (domain.Usuario, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Usuario{}, err
	}

	lastID++

	user, err := s.repository.Store(lastID, nombre, apellido, email, edad, altura, activo, fecha)
	if err != nil {
		return domain.Usuario{}, err
	}

	return user, nil
}

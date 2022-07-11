package users

import "goweb/clase2_parte2/internal/domain"

// Se debe generar la interface Service con todos sus métodos.
type Service interface {
	GetAll() ([]domain.User, error)
	Store(nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error)
}

// Se debe generar la estructura service que contenga el repositorio.
type service struct {
	repository Repository
}

// Se debe generar una función que devuelva el Servicio.
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..).
func (s *service) GetAll() ([]domain.User, error) {
	us, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return us, nil
}

func (s *service) Store(nombre, apellido, email string, edad int, altura float64, activo bool, fechaCreacion string) (domain.User, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.User{}, err
	}

	lastID++
	
	user, err := s.repository.Store(lastID, nombre, apellido, email, edad, altura, activo, fechaCreacion)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
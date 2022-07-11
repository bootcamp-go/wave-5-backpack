package users

import "goweb/clase3_parte1/internal/domain"

// Se genera la interface Service con todos sus métodos
type Service interface {
	GetAll() ([]domain.User, error)
	Store(nombre, apellido, email string, edad int, altura float64, activo *bool, fechaCreacion string) (domain.User, error)
	Update(id int, nombre, apellido, email string, edad int, altura float64, activo *bool, fechaCreacion string) (domain.User, error)
	UpdateLastNameAndAge(id int, apellido string, edad int) (domain.User, error)
	Delete(id int) error
}

// Se genera la estructura service que contiene el repositorio
type service struct {
	repository Repository
}

// Se genera una función que devuelve el Servicio
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

/* Se implementan todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..) */

func (s *service) GetAll() ([]domain.User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) Store(nombre, apellido, email string, edad int, altura float64, activo *bool, fechaCreacion string) (domain.User, error) {
	user, err := s.repository.Store(nombre, apellido, email, edad, altura, activo, fechaCreacion)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// Se llama al repositorio para que proceda a actualizar el usuario
func (s *service) Update(id int, nombre, apellido, email string, edad int, altura float64, activo *bool, fechaCreacion string) (domain.User, error) {
	return s.repository.Update(id, nombre, apellido, email, edad, altura, activo, fechaCreacion)
}

// Se llama al repositorio para que proceda a actualizar el nombre y la edad del usuario
func (s *service) UpdateLastNameAndAge(id int, apellido string, edad int) (domain.User, error) {
	return s.repository.UpdateLastNameAndAge(id, apellido, edad)
}

// Se llama al repositorio para que proceda a elimiinar el usuario
func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

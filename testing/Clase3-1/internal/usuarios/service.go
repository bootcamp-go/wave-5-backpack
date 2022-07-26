package usuarios

import "Clase3-1/internal/domain"

type Service interface {
	GetAll() ([]domain.Usuario, error)
	Store(nombre, apellido, email string, edad, altura int, activo bool, fecha string) (domain.Usuario, error)
	Update(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (domain.Usuario, error)
	UpdateSurnameAndAge(id int, surname string, age int) (domain.Usuario, error)
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

func (s *service) GetAll() ([]domain.Usuario, error) {
	us, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return us, nil
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
func (s *service) Update(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha string) (domain.Usuario, error) {
	return s.repository.Update(id, nombre, apellido, email, edad, altura, activo, fecha)
}

func (s *service) UpdateSurnameAndAge(id int, surname string, age int) (domain.Usuario, error) {
	return s.repository.UpdateSurnameAndAge(id, surname, age)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

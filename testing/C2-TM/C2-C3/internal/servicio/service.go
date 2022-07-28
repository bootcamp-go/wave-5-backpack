package servicio

import (
	"C2-C3/internal/domain"
	"C2-C3/internal/repositorio"
)

type Service interface {
	GetAll() ([]*domain.User, error)
	Update(id int, firstName string, lastName string, email string, age int, height float64, activo bool, createdAt string) (domain.User, error)
	UpdateLastNameAge(id int, lastName string, age int) (domain.User, error)
	Delete(id int) error
	Store(firstName string, lastName string, email string, age int, height float64, activo bool, createdAt string) (domain.User, error)
}

type service struct {
	repository repositorio.Repository
}

func NewService(r repositorio.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]*domain.User, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) Store(firstName string, lastName string, email string, age int, height float64, activo bool, createdAt string) (domain.User, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.User{}, err
	}
	lastID++

	user, err := s.repository.Store(lastID, firstName, lastName, email, age, height, activo, createdAt)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (s *service) Update(id int, firstName string, lastName string, email string, age int, height float64, activo bool, createdAt string) (domain.User, error) {

	return s.repository.Update(id, firstName, lastName, email, age, height, activo, createdAt)
}
func (s *service) UpdateLastNameAge(id int, lastName string, age int) (domain.User, error) {

	return s.repository.UpdateLastNameAge(id, lastName, age)
}
func (s *service) Delete(id int) error {

	return s.repository.Delete(id)
}

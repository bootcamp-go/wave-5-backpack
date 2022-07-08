package users

import "github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"

type Service interface {
	GetAll() ([]domain.Users, error)
	Store(age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error)
	Update(id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error)
	UpdateLastNameAndAge(id, age int, lastName string) (domain.Users, error)
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

func (s *service) GetAll() ([]domain.Users, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Users{}, err
	}

	lastID++

	user, err := s.repository.Store(lastID, age, name, lastName, email, creationDate, height, active)
	if err != nil {
		return domain.Users{}, err
	}

	return user, nil
}

func (s *service) Update(id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	return s.repository.Update(id, age, name, lastName, email, creationDate, height, active)
}

func (s *service) UpdateLastNameAndAge(id, age int, lastName string) (domain.Users, error) {
	return s.repository.UpdateLastNameAndAge(id, age, lastName)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

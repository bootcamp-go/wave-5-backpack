package users

import "goweb/internal/domain"

type Service interface {
	GetAll() ([]domain.User, error)
	NewUser(name, lastname, email string, age int, height float64, active bool, creationDate string) (domain.User, error)
	Update(id int, name, lastname, email string, age int, height float64, active bool, creationDate string) (domain.User, error)
	UpdateName(id int, name string) (domain.User, error)
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

func (s *service) GetAll() ([]domain.User, error) {
	us, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return us, nil
}

func (s *service) NewUser(name, lastname, email string, age int, height float64, active bool, creationDate string) (domain.User, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.User{}, err
	}

	lastID++

	user, err := s.repository.NewUser(lastID, name, lastname, email, age, height, active, creationDate)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *service) Update(id int, name, lastname, email string, age int, height float64, active bool, creationDate string) (domain.User, error) {
	return s.repository.Update(id, name, lastname, email, age, height, active, creationDate)
}

func (s *service) UpdateName(id int, name string) (domain.User, error) {
	return s.repository.UpdateName(id, name)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

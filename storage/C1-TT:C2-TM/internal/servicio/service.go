package servicio

import (
	"C1-TT/internal/domain"
	"C1-TT/internal/repositorio"
)

type Service interface {
	//GetAll() ([]domain.User, error)
	//Update(id int, firstName string, lastName string, email string, age int, height float64, activo bool, createdAt string) (domain.User, error)
	//UpdateLastNameAge(id int, lastName string, age int) (domain.User, error)
	//Delete(id int) error
	Store(user domain.User) (domain.User, error)
	GetOne(id int) (domain.User, error)
	GetByName(firstName string) (domain.User, error)
}

type service struct {
	repository repositorio.Repository
}

func NewService(r repositorio.Repository) Service {
	return &service{
		repository: r,
	}
}

/*
func (s *service) GetAll() ([]domain.User, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}
*/
func (s *service) Store(user domain.User) (domain.User, error) {

	user, err := s.repository.Store(user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (s *service) GetOne(id int) (domain.User, error) {

	user, err := s.repository.GetOne(id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (s *service) GetByName(firstName string) (domain.User, error) {

	user, err := s.repository.GetByName(firstName)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

/*
func (s *service) Update(id int, firstName string, lastName string, email string, age int, height float64, activo bool, createdAt string) (domain.User, error) {

	return s.repository.Update(id, firstName, lastName, email, age, height, activo, createdAt)
}
func (s *service) UpdateLastNameAge(id int, lastName string, age int) (domain.User, error) {

	return s.repository.UpdateLastNameAge(id, lastName, age)
}
func (s *service) Delete(id int) error {

	return s.repository.Delete(id)
}
*/

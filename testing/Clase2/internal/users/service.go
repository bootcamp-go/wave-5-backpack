package users

import "clase2_2/internal/domain"

type Service interface {
	GetAll() ([]domain.User, error)
	AddUser(name, lastName, mail, createDate string, year int, tall float64, enable bool) (domain.User, error)
	UpdateUser(name, lastName, mail, createDate string, year, id int, tall float64, enable bool) (domain.User, error)
	UpdateUserName(name string, id int) (domain.User, error)
	Delete(id int) error
}

type service struct {
	rep Repository
}

func NewService(r Repository) Service {
	return &service{rep: r}
}

func (s *service) GetAll() ([]domain.User, error) {
	users, err := s.rep.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *service) AddUser(name, lastName, mail, createDate string, year int, tall float64, enable bool) (domain.User, error) {
	lastId, err := s.rep.LastId()
	if err != nil {
		return domain.User{}, err
	}

	lastId++
	user, err := s.rep.AddUser(name, lastName, mail, createDate, year, lastId, tall, enable)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

//clase 3_1
func (s *service) UpdateUser(name, lastName, mail, createDate string, year, id int, tall float64, enable bool) (domain.User, error) {

	user, err := s.rep.UpdateUser(name, lastName, mail, createDate, year, id, tall, enable)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
func (s *service) Delete(id int) error {
	return s.rep.Delete(id)
}
func (s *service) UpdateUserName(name string, id int) (domain.User, error) {
	user, err := s.rep.UpdateUserName(name, id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

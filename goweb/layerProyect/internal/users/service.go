package users

type Service interface {
	GetAll() ([]*User, error)
	AddUser(name, lastName, mail, createDate string, year int, tall float64, enable bool) (*User, error)
	UpdateUser(name, lastName, mail, createDate string, year, id int, tall float64, enable bool) (*User, error)
	UpdateUserName(name string, id int) (*User, error)
	Delete(id int) error
}

type service struct {
	rep Repository
}

func NewService(r Repository) Service {
	return &service{rep: r}
}

func (s *service) GetAll() ([]*User, error) {
	users, err := s.rep.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *service) AddUser(name, lastName, mail, createDate string, year int, tall float64, enable bool) (*User, error) {
	lastId, err := s.rep.LastId()
	if err != nil {
		return nil, err
	}

	lastId++
	user, err := s.rep.AddUser(name, lastName, mail, createDate, year, lastId, tall, enable)
	if err != nil {
		return nil, err
	}

	return user, nil
}

//clase 3_1
func (s *service) UpdateUser(name, lastName, mail, createDate string, year, id int, tall float64, enable bool) (*User, error) {

	user, err := s.rep.UpdateUser(name, lastName, mail, createDate, year, id, tall, enable)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (s *service) Delete(id int) error {
	return s.rep.Delete(id)
}
func (s *service) UpdateUserName(name string, id int) (*User, error) {
	user, err := s.rep.UpdateUserName(name, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

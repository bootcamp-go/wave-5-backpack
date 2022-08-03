package products

type Service interface {
	GetAll() ([]*Product, error)
	UpdateName(id int, nombre string) (*Product, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]*Product, error) {
	return s.repo.GetAll()
}

func (s *service) UpdateName(id int, nombre string) (*Product, error) {
	return s.repo.UpdateName(id, nombre)
}

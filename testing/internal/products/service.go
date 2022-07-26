package products

type Service interface {
	GetAll() ([]*Product, error)
	UpdateName(id int, nombre string) (*Product, error)
	Update(id int, nombre string, stock int, precio float64) (*Product, error)
	Delete(id int) error
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

func (s *service) Update(id int, nombre string, stock int, precio float64) (*Product, error) {
	return s.repo.Update(id, nombre, stock, precio)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}

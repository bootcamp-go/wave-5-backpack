package products

type Service interface {
	GetAll() ([]Product, error)
	Store(name, productType string, count int, price float64) (Product, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]Product, error) {
	return s.repo.GetAll()
}

func (s *service) Store(name, productType string, count int, price float64) (Product, error) {
	lastID, err := s.repo.LastID()
	if err != nil {
		return Product{}, err
	}

	lastID++
	return s.repo.Store(lastID, name, productType, count, price)
}

func (s *service) Update(id int, name, productType string, count int, price float64) (Product, error) {
	return s.repo.Update(id, name, productType, count, price)
}

func (s *service) UpdateName(id int, name string) (Product, error) {
	return s.repo.UpdateName(id, name)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}

package products

type Service interface {
	GetAll() ([]Product, error)
	Store(nombre, tipo string, cantidad int, precio float64) (Product, error)
	Update(id int, nombre, tipo string, cantidad int, precio float64) (Product, error)
	Patch(id int, nombre string, precio float64) (Product, error)
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

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(nombre, tipo string, cantidad int, precio float64) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, nombre, tipo, cantidad, precio)
	if err != nil {
		return Product{}, err
	}

	return producto, nil
}

func (s *service) Update(id int, nombre, tipo string, cantidad int, precio float64) (Product, error) {
	return s.repository.Update(id, nombre, tipo, cantidad, precio)
}

func (s *service) Patch(id int, nombre string, precio float64) (Product, error) {
	return s.repository.Patch(id, nombre, precio)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

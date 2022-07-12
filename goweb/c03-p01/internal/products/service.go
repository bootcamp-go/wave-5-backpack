package products

type Service interface {
	GetAll() ([]Producto, error)
	Store(nombre string, cantidad int, precio float64) (Producto, error)
	Update(id int, nombre string, cantidad int, precio float64) (Producto, error)
	UpdateName(id int, nombre string) (Producto, error)
	Delete(id int) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]Producto, error) {
	ps, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) Store(nombre string, cantidad int, precio float64) (Producto, error) {
	lastID, err := s.repo.LastID()
	if err != nil {
		return Producto{}, err
	}

	lastID++
	producto, err := s.repo.Store(lastID, nombre, cantidad, precio)
	if err != nil {
		return Producto{}, err
	}

	return producto, nil
}

func (s *service) Update(id int, nombre string, cantidad int, precio float64) (Producto, error) {
	return s.repo.Update(id, nombre, cantidad, precio)
}

func (s *service) UpdateName(id int, nombre string) (Producto, error) {
	return s.repo.UpdateName(id, nombre)
}

func (s *service) UpdateNamePrice(id int, nombre string, precio float64) (Producto, error) {
	return s.repo.UpdateNamePrice(id, nombre, precio)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}

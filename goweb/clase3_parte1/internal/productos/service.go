package productos

type Service interface {
	Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Productos, error)
	GetAll() ([]Productos, error)
	Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Productos, error)
	UpdatePrecio(id int, precio float64) (Productos, error)
	Delete(id int) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Productos, error) {
	return s.repo.Update(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
}

func (s *service) Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Productos, error) {
	lastID, err := s.repo.LastID()
	if err != nil {
		return Productos{}, err
	}
	lastID++
	return s.repo.Store(lastID, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
}

func (s *service) GetAll() ([]Productos, error) {
	return s.repo.GetAll()
}

func (s *service) UpdatePrecio(id int, precio float64) (Productos, error) {
	return s.repo.UpdatePrecio(id, precio)
}

func (s *service) Delete(id int) error {
	return s.repo.Delete(id)
}

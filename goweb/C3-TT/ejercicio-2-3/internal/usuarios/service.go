package usuarios

type Service interface {
	GetAll() ([]*Usuario, error)
	Registrar(nombre, apellido, email string, edad, altura int, activo bool, fecha_creacion string) (*Usuario, error)
	Modificar(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha_creacion string) (*Usuario, error)
	Eliminar(id int) error
	ModificarAE(id int, apellido string, edad int) (*Usuario, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]*Usuario, error) {
	usuarios, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return usuarios, nil
}

func (s *service) Registrar(nombre, apellido, email string, edad, altura int, activo bool, fecha_creacion string) (*Usuario, error) {
	usuarios, err := s.repository.Registrar(nombre, apellido, email, edad, altura, activo, fecha_creacion)
	if err != nil {
		return nil, err
	}

	return usuarios, nil
}

func (s *service) Modificar(id int, nombre, apellido, email string, edad, altura int, activo bool, fecha_creacion string) (*Usuario, error) {
	return s.repository.Modificar(id, nombre, apellido, email, edad, altura, activo, fecha_creacion)
}

func (s *service) Eliminar(id int) error {
	return s.repository.Eliminar(id)
}

func (s *service) ModificarAE(id int, apellido string, edad int) (*Usuario, error) {
	return s.repository.ModificarAE(id, apellido, edad)
}

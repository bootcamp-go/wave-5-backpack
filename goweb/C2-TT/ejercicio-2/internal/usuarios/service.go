package usuarios

type Service interface {
	GetAll() ([]*Usuario, error)
	Registrar(nombre, apellido, email string, edad, altura int, activo bool, fecha_creacion string) (*Usuario, error)
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

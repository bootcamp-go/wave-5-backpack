package usuarios

import (
	"context"
	"time"

	"github.com/del_rio/web-server/internal/domain"
)

type servicio struct {
	repo Repository
}

type Servicio interface {
	GetAll() ([]domain.Usuario, error)
	Save(domain.Usuario) (domain.Usuario, error)
	UpdateUsuario(ctx context.Context, Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo *bool) (domain.Usuario, error)
	UpdateAtributos(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo *bool) (domain.Usuario, error)
	DeleteUsuario(id int) error
}

func (s *servicio) GetAll() ([]domain.Usuario, error) {
	listUsuar, err := s.repo.GetAll()
	if err != nil {
		return []domain.Usuario{}, err
	}
	return listUsuar, nil
}
func (s *servicio) Save(usuario domain.Usuario) (domain.Usuario, error) {
	lastId, _ := s.repo.LastId()
	lastId++
	usuario.Fecha_creacion = time.Now().String()
	nuevoUsuario, err := s.repo.Store(usuario)

	if err != nil {
		return domain.Usuario{}, err
	}
	return nuevoUsuario, nil
}

func (s *servicio) UpdateUsuario(ctx context.Context, Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo *bool) (domain.Usuario, error) {
	nullUsuario := domain.Usuario{}
	usuarioChanging, err := s.repo.GetById(Id)
	if err != nil {
		return nullUsuario, err
	}
	validationAttributeInt := []int{Id, Edad, Altura}
	validationAttributeStr := []string{Nombre, Apellido, Email, Fecha_creacion}
	validationAttributeBool := []*bool{Activo}
	attributeInt := []*int{&usuarioChanging.Id, &usuarioChanging.Edad, &usuarioChanging.Altura}
	attributeStr := []*string{&usuarioChanging.Nombre, &usuarioChanging.Apellido, &usuarioChanging.Email, &usuarioChanging.Fecha_creacion}
	attributeBool := []*bool{&usuarioChanging.Activo}
	for i := 0; i < len(attributeInt); i++ {
		if validationAttributeInt[i] != 0 {
			*attributeInt[i] = validationAttributeInt[i]
		}
	}
	for i := 0; i < len(attributeStr); i++ {
		if validationAttributeStr[i] != "" {
			*attributeStr[i] = validationAttributeStr[i]
		}
	}
	for i := 0; i < len(attributeBool); i++ {
		if validationAttributeBool[i] != nil {
			*attributeBool[i] = (*validationAttributeBool[i])
		}
	}

	return s.repo.UpdateUsuario(ctx, usuarioChanging)
}

func (s *servicio) UpdateAtributos(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo *bool) (domain.Usuario, error) {
	return s.repo.UpdateAtributos(
		Nombre, Apellido, Email, Fecha_creacion, Id, Edad, Altura, Activo,
	)
}
func (s *servicio) DeleteUsuario(id int) error {
	return s.repo.DeleteUsuario(id)
}
func NewService(r Repository) Servicio {
	return &servicio{
		repo: r,
	}
}

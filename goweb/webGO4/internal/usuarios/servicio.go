package usuarios

import (
	"time"

	"github.com/del_rio/web-server/internal/domain"
)

type servicio struct {
	repo Repository
}

type Servicio interface {
	GetAll() ([]domain.Usuario, error)
	Save(Nombre, Apellido, Email string, Edad, Altura int, Activo bool) (domain.Usuario, error)
	// UpdateUsuario(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo bool) (domain.Usuario, error)
	// UpdateAtributos(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo *bool) (domain.Usuario, error)
	// DeleteUsuario(id int) error
}

func (s *servicio) GetAll() ([]domain.Usuario, error) {
	listUsuar, err := s.repo.GetAll()
	if err != nil {
		return []domain.Usuario{}, err
	}
	return listUsuar, nil
}
func (s *servicio) Save(Nombre, Apellido, Email string, Edad, Altura int, Activo bool) (domain.Usuario, error) {
	lastId, _ := s.repo.LastId()
	lastId++
	nuevoUsuario, err := s.repo.Save(
		Nombre,
		Apellido,
		Email,
		time.Now().String(),
		lastId,
		Edad,
		Altura,
		Activo,
	)

	if err != nil {
		return domain.Usuario{}, err
	}
	return nuevoUsuario, nil
}

// func (s *servicio) UpdateUsuario(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo bool) (domain.Usuario, error) {
// 	return s.repo.UpdateUsuario(
// 		Nombre,
// 		Apellido,
// 		Email,
// 		Fecha_creacion,
// 		Id,
// 		Edad,
// 		Altura,
// 		Activo,
// 	)
// }
// func (s *servicio) UpdateAtributos(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int, Activo *bool) (domain.Usuario, error) {
// 	return s.repo.UpdateAtributos(
// 		Nombre, Apellido, Email, Fecha_creacion, Id, Edad, Altura, Activo,
// 	)
// }
// func (s *servicio) DeleteUsuario(id int) error {
// 	return s.repo.DeleteUsuario(id)
// }
func NewService(r Repository) Servicio {
	return &servicio{
		repo: r,
	}
}

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
	Save(Nombre, Apellido, Email string, Edad, Altura int) (domain.Usuario, error)
}

func (s *servicio) GetAll() ([]domain.Usuario, error) {
	listUsuar, err := s.repo.GetAll()
	if err != nil {
		return []domain.Usuario{}, err
	}
	return listUsuar, nil
}
func (s *servicio) Save(Nombre, Apellido, Email string, Edad, Altura int) (domain.Usuario, error) {
	lastId, _ := s.repo.LastId()
	nuevoUsuario, err := s.repo.Save(
		Nombre,
		Apellido,
		Email,
		time.Now().String(),
		lastId,
		Edad,
		Altura,
	)
	if err != nil {
		return domain.Usuario{}, err
	}
	return nuevoUsuario, nil
}
func NewService(r Repository) Servicio {
	return &servicio{
		repo: r,
	}
}

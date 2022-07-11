package usuarios

import (
	"fmt"

	"github.com/del_rio/web-server/internal/domain"
)

var lastId int
var listUsuarios []domain.Usuario

type Repository interface {
	GetAll() ([]domain.Usuario, error)
	Save(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int) (domain.Usuario, error)
	LastId() (int, error)
}

type repository struct{}

func (r *repository) GetAll() ([]domain.Usuario, error) {
	if len(listUsuarios) == 0 {
		return listUsuarios, fmt.Errorf("lista vacia perdone usted")
	}
	return listUsuarios, nil
}

func (r *repository) Save(Nombre, Apellido, Email, Fecha_creacion string, Id, Edad, Altura int) (domain.Usuario, error) {
	if Nombre == "" || Apellido == "" || Email == "" {
		usuarioNulo := domain.Usuario{}
		return usuarioNulo, fmt.Errorf("campo invalido o campos invalidos")
	}
	usuario := domain.Usuario{
		Id:             Id,
		Nombre:         Nombre,
		Apellido:       Apellido,
		Email:          Email,
		Edad:           Edad,
		Altura:         Altura,
		Fecha_creacion: Fecha_creacion,
	}
	listUsuarios = append(listUsuarios, usuario)
	return usuario, nil
}
func (r *repository) LastId() (int, error) {
	//aqui podria pasar algo
	return lastId, nil
}
func NewRepository() Repository {
	return &repository{}
}

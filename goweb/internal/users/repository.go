package users

import (
	"time"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
)

var users []domain.ModelUser
var lastId int

type Repository interface {
	GetAll() ([]domain.ModelUser, error)
	Store(id int, nombre string, apellido string, email string, edad int, altura float64) (domain.ModelUser, error)
	LastId() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]domain.ModelUser, error) {
	return users, nil
}

func (*repository) LastId() (int, error) {
	if len(users) > 0 {
		lastId = 0
		for _, u := range users {
			if u.Id > lastId {
				lastId = u.Id
			}
		}
		return lastId, nil
	} else {
		return 0, nil
	}
}

func (r *repository) Store(id int, nombre string, apellido string, email string, edad int, altura float64) (domain.ModelUser, error) {
	u := domain.ModelUser{Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: false, FechaCreacion: time.Now()}
	users = append(users, u)
	return u, nil
}

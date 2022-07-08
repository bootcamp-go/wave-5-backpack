/*Repositorio, debe tener el acceso a la variable guardada en memoria.
OK Se debe crear el archivo repository.go
OK Se deben crear las variables globales donde guardar las entidades
OK Se debe generar la interface Repository con todos sus métodos
OK Se debe generar la estructura repository
OK Se debe generar una función que devuelva el Repositorio
OK Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..)
*/
package usuarios

import "github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"

var us []domain.Usuarios
var lastID int

type Repository interface {
	GetAll() ([]domain.Usuarios, error)
	Guardar(id int, nombre string, apellido string, email string, edad int, altura float64, actico bool, fecha string) (domain.Usuarios, error)
	lastId() (int, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]domain.Usuarios, error) {
	return us, nil
}
func (r *repository) Guardar(id int, nombre string, apellido string, email string, edad int, altura float64, actico bool, fecha string) (domain.Usuarios, error) {
	u := domain.Usuarios{Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: actico, FechaCreacion: fecha}
	us = append(us, u)
	lastID = u.Id
	return u, nil
}
func (r *repository) lastId() (int, error) {
	return lastID, nil
}

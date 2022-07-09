package repositorio

import (
	"fmt"

	"github.com/rodrigoeshard/goweb/Practica2.2/pkg/store"

	"github.com/rodrigoeshard/goweb/Practica2.2/internal/domain"
)

/*
type User struct {
	Id           int     `json:"-"`
	FirstName    string  `json:"firstName" binding:"required"`
	LastName     string  `json:"lastName" binding:"required"`
	Email        string  `json:"email" binding:"required"`
	Age          int     `json:"age" binding:"required"`
	Height       float64 `json:"height" binding:"required"`
	Active       bool    `json:"active" binding:"required"`
	CreationDate string  `json:"creationDate" binding:"required"`
}
*/
type Repository interface {
	GetAll() ([]domain.User, error)
	Store(id int, firstName string, lastName string, email string, age int, height float64, activo bool, createdAt string) (domain.User, error)
	LastID() (int, error)
	Update(id int, firstName string, lastName string, email string, age int, height float64, activo bool, createdAt string) (domain.User, error)
	UpdateLastNameAge(id int, lastName string, age int) (domain.User, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

/*func newRepository() Repository{
	return &repository{}
}*/

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

var ps []domain.User
var lastID int

func (r *repository) GetAll() ([]domain.User, error) {
	var ps []domain.User
	r.db.Read(&ps)
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	var ps []domain.User
	if err := r.db.Read(&ps); err != nil {
		return 0, nil
	}
	return ps[len(ps)-1].Id, nil
}

func (r *repository) Store(id int, firstName string, lastName string, email string, age int, height float64, activo bool, createdAt string) (domain.User, error) {
	var ps []domain.User
	r.db.Read(&ps)
	p := domain.User{Id: id, FirstName: firstName, LastName: lastName, Email: email, Age: age, Height: height, Activo: activo, CreatedAt: createdAt}
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return domain.User{}, err
	}
	return p, nil
}

func (r *repository) Update(id int, firstName string, lastName string, email string, age int, height float64, activo bool, createdAt string) (domain.User, error) {
	u := domain.User{Id: id, FirstName: firstName, LastName: lastName, Email: email, Age: age, Height: height, Activo: activo, CreatedAt: createdAt}
	var update bool = false
	for k, v := range ps {
		if u.Id == v.Id {
			ps[k] = u
			update = true
		}
	}
	if update != true {
		return domain.User{}, fmt.Errorf("No se pudo actualizar el usuario")
	}

	return u, nil
}

func (r *repository) UpdateLastNameAge(id int, lastName string, age int) (domain.User, error) {
	u := domain.User{}
	var update bool = false
	for k, v := range ps {
		if v.Id == id {
			ps[k].Age = age
			ps[k].LastName = lastName
			u = ps[k]
			update = true
		}
	}
	if update != true {
		return domain.User{}, fmt.Errorf("No se pudo actualizar el usuario")
	}

	return u, nil
}
func (r *repository) Delete(id int) error {
	fmt.Println(id)
	var delete bool = false
	var idDel = 0
	for k, v := range ps {
		if v.Id == id {
			idDel = k
			delete = true
		}

	}
	if delete != true {
		return fmt.Errorf("No se pudo eliminar el usuario")
	}
	ps = append(ps[:idDel], ps[idDel+1:]...)

	return nil
}

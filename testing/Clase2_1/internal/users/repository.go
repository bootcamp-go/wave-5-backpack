package users

import (
	"clase2_2/internal/domain"
	"clase2_2/pkg/storage"
	"errors"
)

type Repository interface {
	LastId() (int, error)
	GetAll() ([]domain.User, error)
	AddUser(name, lastName, mail, createDate string, year, id int, tall float64, enable bool) (domain.User, error)
	UpdateUser(name, lastName, mail, createDate string, year, id int, tall float64, enable bool) (domain.User, error)
	UpdateUserName(name string, id int) (domain.User, error)
	Delete(id int) error
}

var lastId int

type repository struct {
	db storage.Store
}

func NewRepository(db storage.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Read(&users)
	if err != nil {
		return nil, errors.New("Ocurrió un error al intentar leer el archivo")
	}
	return users, nil
}

func (r *repository) AddUser(name, lastName, mail, createDate string, year, id int, tall float64, enable bool) (domain.User, error) {
	var users []domain.User
	err := r.db.Read(&users)
	if err != nil {
		return domain.User{}, errors.New("Ocurrió un error al intentar leer el archivo")
	}
	user := domain.User{
		Id:         id,
		Name:       name,
		LastName:   lastName,
		Mail:       mail,
		CreateDate: createDate,
		Years:      year,
		Tall:       tall,
		Enable:     enable,
	}
	users = append(users, user)
	err = r.db.Write(users)
	if err != nil {
		return domain.User{}, errors.New("Ocurrió un error al intentar escribir el archivo")
	}
	return user, nil
}

func (r *repository) LastId() (int, error) {
	var users []domain.User
	err := r.db.Read(&users)
	if err != nil {
		return 0, errors.New("Ocurrió un error al intentar escribir el archivo")
	}
	if len(users) == 0 {
		return 0, nil
	}

	return users[len(users)-1].Id, nil
}

//clase 3_1
func (r *repository) UpdateUser(name, lastName, mail, createDate string, year, id int, tall float64, enable bool) (domain.User, error) {
	var users []domain.User
	err := r.db.Read(&users)
	if err != nil {
		return domain.User{}, errors.New("Ocurrió un error al intentar leer el archivo")
	}
	user := domain.User{
		Id:         id,
		Name:       name,
		LastName:   lastName,
		Mail:       mail,
		CreateDate: createDate,
		Years:      year,
		Tall:       tall,
		Enable:     enable,
	}
	updated := false
	for i, u := range users {
		if u.Id == id {
			users[i] = user
			updated = true
		}
	}
	if !updated {
		return domain.User{}, errors.New("no se a encontrado el id")
	}
	err = r.db.Write(users)
	if err != nil {
		return domain.User{}, errors.New("Ocurrió un error al intentar escribir el archivo")
	}
	return user, nil
}
func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	var users []domain.User
	err := r.db.Read(&users)
	if err != nil {
		return errors.New("Ocurrió un error al intentar leer el archivo")
	}
	for i := range users {
		if users[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return errors.New("no se a encontrado el id")
	}
	users = append(users[:index], users[index+1:]...)
	err = r.db.Write(users)
	if err != nil {
		return errors.New("Ocurrió un error al intentar escribir el archivo")
	}
	return nil
}
func (r *repository) UpdateUserName(name string, id int) (domain.User, error) {
	var users []domain.User
	err := r.db.Read(&users)
	if err != nil {
		return domain.User{}, errors.New("Ocurrió un error al intentar leer el archivo")
	}
	updated := false
	user := domain.User{}
	for i, u := range users {
		if u.Id == id {
			users[i].Name = name
			u.Name = name
			user = u
			updated = true
		}
	}
	if !updated {
		return domain.User{}, errors.New("no se a encontrado el id")
	}
	err = r.db.Write(users)
	if err != nil {
		return domain.User{}, errors.New("Ocurrió un error al intentar escribir el archivo")
	}
	return user, nil
}

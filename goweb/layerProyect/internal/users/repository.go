package users

import "errors"

type Repository interface {
	LastId() (int, error)
	GetAll() ([]*User, error)
	AddUser(name, lastName, mail, createDate string, year, id int, tall float64, enable bool) (*User, error)
	UpdateUser(name, lastName, mail, createDate string, year, id int, tall float64, enable bool) (*User, error)
	UpdateUserName(name string, id int) (*User, error)
	Delete(id int) error
}

type User struct {
	Id         int
	Name       string
	LastName   string
	Mail       string
	Years      int
	Tall       float64
	Enable     bool
	CreateDate string
}

var users []*User
var lastId int

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]*User, error) {
	return users, nil
}

func (r *repository) AddUser(name, lastName, mail, createDate string, year, id int, tall float64, enable bool) (*User, error) {
	user := &User{
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
	lastId = user.Id
	return user, nil
}

func (r *repository) LastId() (int, error) {
	return lastId, nil
}

//clase 3_1
func (r *repository) UpdateUser(name, lastName, mail, createDate string, year, id int, tall float64, enable bool) (*User, error) {
	user := &User{
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
		return &User{}, errors.New("no se a encontrado el id")
	}
	return user, nil
}
func (r *repository) Delete(id int) error {
	deleted := false
	var index int
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
	return nil
}
func (r *repository) UpdateUserName(name string, id int) (*User, error) {
	updated := false
	user := &User{}
	for i, u := range users {
		if u.Id == id {
			user = u
			users[i].Name = name
			updated = true
		}
	}
	if !updated {
		return &User{}, errors.New("no se a encontrado el id")
	}
	return user, nil
}

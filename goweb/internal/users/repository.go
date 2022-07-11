package users

import (
	"fmt"
	"github.com/bootcamp-go/wave-5-backpack/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/pkg/store"
)

type Repository interface {
	GetAll() ([]domain.User, error)
	LastId() (int, error)
	GetById(id int) (domain.User, error)
	StoreUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error)
	UpdateUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error)
}

const (
	UserNotFound = "user id %d not found"
	FailReading  = "fail to read database"
	FailWriting  = "fail to write database: %w"
)

type repository struct {
	db store.Store
}

func NewRepositoy(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

// var allUsers []domain.User
// var lastId int

func (r *repository) GetAll() ([]domain.User, error) {
	var allUsers []domain.User
	if err := r.db.Read(&allUsers); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return allUsers, nil
}
func (r *repository) GetById(id int) (domain.User, error) {

	allUsers, err := r.GetAll()
	if err != nil {
		return domain.User{}, err
	}
	for _, u := range allUsers {
		if u.ID == id {
			return u, nil
		}
	}

	return domain.User{}, fmt.Errorf(UserNotFound, id)
}

func (r *repository) LastId() (int, error) {
	allUsers, err := r.GetAll()
	if err != nil {
		return 0, err
	}
	if len(allUsers) == 0 {
		return 0, nil
	}

	return allUsers[len(allUsers)-1].ID, nil
}

func (r *repository) StoreUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error) {
	allUsers, err := r.GetAll()
	if err != nil {
		return domain.User{}, err
	}
	newUser := domain.User{
		ID:         id,
		Name:       name,
		Lastname:   lastname,
		Email:      email,
		Age:        age,
		Height:     height,
		Active:     active,
		DoCreation: doCreation,
	}
	allUsers = append(allUsers, newUser)

	if err := r.db.Write(allUsers); err != nil {
		return domain.User{}, fmt.Errorf(FailWriting, err)
	}
	return newUser, nil
}

func (r *repository) UpdateUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error) {
	allUsers, err := r.GetAll()
	if err != nil {
		return domain.User{}, err
	}

	newUser := domain.User{
		ID:         id,
		Name:       name,
		Lastname:   lastname,
		Email:      email,
		Age:        age,
		Height:     height,
		Active:     active,
		DoCreation: doCreation,
	}
	for index := range allUsers {
		if allUsers[index].ID == id {
			allUsers[index] = newUser
			if err := r.db.Write(allUsers); err != nil {
				return domain.User{}, fmt.Errorf(FailWriting, err)
			}
			return newUser, nil
		}
	}
	return domain.User{}, fmt.Errorf(UserNotFound, id)

}

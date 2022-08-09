package users

import (
	"context"
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/pkg/store"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.User, error)
	LastId() (int, error)
	GetById(ctx context.Context,id int) (domain.User, error)
	StoreUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error)
	UpdateUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error)
	DeleteUser(id int) error
	GetByName(ctx context.Context,name string) ([]domain.User, error)

	UpdateLastnameAndAge(id int, lastname string, age int)(*domain.User, error)
}

const (
	UserNotFound = "user id %d not found"
	FailReading  = "fail to read database"
	FailWriting  = "fail to write database"
)

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

// var allUsers []domain.User
// var lastId int

func (r *repository) GetAll(ctx context.Context) ([]domain.User, error) {
	var allUsers []domain.User
	if err := r.db.Read(&allUsers); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return allUsers, nil
}
func (r *repository) GetById(ctx context.Context, id int) (domain.User, error) {

	allUsers, err := r.GetAll(context.TODO())
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
func (r *repository) DeleteUser(id int) error {
	allUsers, err := r.GetAll(context.TODO())
	if err != nil {
		return err
	}
	for i := range allUsers {
		if allUsers[i].ID == id {
			allUsers = append(allUsers[:i], allUsers[i+1:]...)
			break
		}
	}

	if err := r.db.Write(allUsers); err != nil {
		return fmt.Errorf(FailWriting)
	}

	return nil
}

func (r *repository) LastId() (int, error) {
	allUsers, err := r.GetAll(context.TODO())
	if err != nil {
		return 0, err
	}
	if len(allUsers) == 0 {
		return 0, nil
	}

	return allUsers[len(allUsers)-1].ID, nil
}

func (r *repository) StoreUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error) {
	allUsers, err := r.GetAll(context.TODO())
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
		return domain.User{}, fmt.Errorf(FailWriting)
	}
	return newUser, nil
}

func (r *repository) UpdateUser(id int, name, lastname, email string, age int, height float32, active bool, doCreation string) (domain.User, error) {
	allUsers, err := r.GetAll(context.TODO())
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
				return domain.User{}, fmt.Errorf(FailWriting)
			}
			return newUser, nil
		}
	}
	return domain.User{}, fmt.Errorf(UserNotFound, id)

}

func (r *repository) UpdateLastnameAndAge(id int, lastname string, age int) (*domain.User, error) {
	allUsers, err := r.GetAll(context.TODO())
	if err != nil {
		return nil, err
	}


	for index := range allUsers {
		if allUsers[index].ID == id {
			allUsers[index].Lastname = lastname
			allUsers[index].Age = age

			if err := r.db.Write(allUsers); err != nil {
				return nil, fmt.Errorf(FailWriting)
			}
			return &allUsers[index], nil
		}
	}
	return nil, fmt.Errorf(UserNotFound, id)

}

func (r *repository) GetByName(ctx context.Context, name string) ([]domain.User, error) {
	var users []domain.User
	return users, nil
}
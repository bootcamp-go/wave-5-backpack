package users

import (
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/store"
)

const (
	UserNotFound = "user %d not found"
	FailReading  = "cant read database"
	FailWriting  = "cant write database, error: %w"
)

type Repository interface {
	GetAll() ([]domain.Users, error)
	GetByName(name string) ([]domain.Users, error)
	Store(id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error)
	LastID() (int, error)
	Update(id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error)
	UpdateLastNameAndAge(id, age int, lastName string) (domain.Users, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Users, error) {
	var users []domain.Users
	if err := r.db.Read(&users); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return users, nil
}

func (r *repository) GetByName(name string) ([]domain.Users, error) {
	var users []domain.Users
	return users, nil
}

func (r *repository) LastID() (int, error) {
	var users []domain.Users
	if err := r.db.Read(&users); err != nil {
		return 0, fmt.Errorf(FailReading)
	}
	if len(users) == 0 {
		return 0, nil
	}

	return users[len(users)-1].Id, nil
}

func (r *repository) Store(id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	var users []domain.Users

	if err := r.db.Read(&users); err != nil {
		return domain.Users{}, fmt.Errorf(FailReading)
	}

	user := domain.Users{
		Id: id, Name: name, LastName: lastName, Email: email, Age: age, Height: height, Active: active, CreationDate: creationDate}
	users = append(users, user)

	if err := r.db.Write(users); err != nil {
		return domain.Users{}, fmt.Errorf(FailWriting, err)
	}

	return user, nil
}

func (r *repository) Update(id, age int, name, lastName, email, creationDate string, height float64, active bool) (domain.Users, error) {
	var users []domain.Users

	if err := r.db.Read(&users); err != nil {
		return domain.Users{}, fmt.Errorf(FailReading)
	}

	for i := range users {
		user := &users[i]
		if user.Id == id {
			user.Name = name
			user.LastName = lastName
			user.Email = email
			user.CreationDate = creationDate
			user.Active = active
			user.Age = age
			user.Height = height

			if err := r.db.Write(users); err != nil {
				return domain.Users{}, fmt.Errorf(FailWriting, err)
			}
			return *user, nil
		}
	}
	return domain.Users{}, fmt.Errorf(UserNotFound, id)
}

func (r *repository) UpdateLastNameAndAge(id, age int, lastName string) (domain.Users, error) {
	var users []domain.Users

	if err := r.db.Read(&users); err != nil {
		return domain.Users{}, fmt.Errorf(FailReading)
	}

	for i := range users {
		user := &users[i]
		if user.Id == id {
			user.LastName = lastName
			user.Age = age

			if err := r.db.Write(users); err != nil {
				return domain.Users{}, fmt.Errorf(FailWriting, err)
			}
			return *user, nil
		}
	}
	return domain.Users{}, fmt.Errorf(UserNotFound, id)
}

func (r *repository) Delete(id int) error {
	var users []domain.Users

	if err := r.db.Read(&users); err != nil {
		return fmt.Errorf(FailReading)
	}

	for i := range users {
		user := users[i]
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)

			if err := r.db.Write(users); err != nil {
				return fmt.Errorf(FailWriting, err)
			}
			return nil
		}
	}
	return fmt.Errorf(UserNotFound, id)
}

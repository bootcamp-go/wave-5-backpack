package users

import (
	"proyectoFinal/internal/domain"
	"proyectoFinal/pkg/store"
	"strconv"
	"time"
)

type Repository interface {
	GetAll() ([]domain.User, error)
	LastId() (int, error)
	Store(Id int, Name string, LastName string, Email string, Age int, Height float64, Active bool, Date time.Time) (domain.User, error)
	Update(Id int, Name string, LastName string, Email string, Age int, Height float64, Active bool) (domain.User, error)
	Delete(Id int) error
	GetById(Id int) (domain.User, error)
}

var users []domain.User = make([]domain.User, 0)

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db,
	}
}

func (r *repository) GetAll() ([]domain.User, error) {
	return users, nil
}

func (r *repository) LastId() (int, error) {
	users, err := r.GetAll()
	if err != nil {
		return 0, err
	}

	return users[len(users)-1].Id, nil
}

func (r *repository) Store(Id int,
	Name string,
	LastName string,
	Email string,
	Age int,
	Height float64,
	Active bool,
	Date time.Time) (domain.User, error) {

	users, err := r.GetAll()
	if err != nil {
		return domain.User{}, err
	}

	user := domain.User{
		Id:       Id,
		Name:     Name,
		LastName: LastName,
		Email:    Email,
		Age:      Age,
		Height:   Height,
		Active:   Active,
		Date:     Date,
	}

	users = append(users, user)
	if err := r.db.Write(users); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *repository) Update(Id int, Name string, LastName string, Email string, Age int, Height float64, Active bool) (domain.User, error) {
	users, err := r.GetAll()

	if err != nil {
		return domain.User{}, err
	}

	userToUpdate := domain.User{}

	update := false

	for _, user := range users {

		if user.Id == Id {
			user.Name = Name
			user.LastName = LastName
			user.Email = Email
			user.Age = Age
			user.Height = Height
			user.Active = Active

			userToUpdate = user
			update = true
		}
	}

	if !update {
		return domain.User{}, &NotFound{searchValue: strconv.Itoa(Id), fileName: "Id"}
	}

	if err := r.db.Write(users); err != nil {
		return domain.User{}, err
	}

	return userToUpdate, nil
}

func (r *repository) Delete(Id int) error {
	users, err := r.GetAll()

	if err != nil {
		return err
	}
	delete := false

	for i, user := range users {
		if user.Id == Id {
			users = append(users[:i], users[i+1])
			delete = true
			break
		}
	}

	if !delete {
		return &NotFound{searchValue: strconv.Itoa(Id), fileName: "Id"}
	}

	if err := r.db.Write(users); err != nil {
		return err
	}

	return nil
}

func (r *repository) GetById(Id int) (domain.User, error) {
	users, err := r.GetAll()
	if err != nil {
		return domain.User{}, err
	}

	for _, user := range users {
		if user.Id == Id {
			return user, nil
		}
	}

	return domain.User{}, &NotFound{searchValue: strconv.Itoa(Id), fileName: "Id"}
}

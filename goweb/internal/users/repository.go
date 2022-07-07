package users

import (
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/file"
)

type Repository interface {
	GetById(Id int) (domain.User, error)
	GetAll() ([]domain.User, error)
	Store(Id, Age int, FirstName, LastName, Email, CreatedAt string, Height float64, Active bool) (domain.User, error)
	LastID() (int, error)
}

type repository struct{}

func (r *repository) LastID() (int, error) {
	db, err := file.ReadJSONFile("users.json")
	if err != nil {
		return 0, err
	}
	return db.LastId, nil
}

func (r *repository) GetById(Id int) (domain.User, error) {
	db, err := file.ReadJSONFile("users.json")
	if err != nil {
		return domain.User{}, err
	}

	for _, user := range db.Users {
		if Id == user.Id {
			return user, nil
		}
	}

	return domain.User{}, nil
}

func (r *repository) GetAll() ([]domain.User, error) {
	db, err := file.ReadJSONFile("users.json")

	if err != nil {
		return nil, err
	}

	return db.Users, nil
}

func (r *repository) Store(Id, Age int, FirstName, LastName, Email, CreatedAt string, Height float64, Active bool) (domain.User, error) {

	db, err := file.ReadJSONFile("users.json")
	if err != nil {
		return domain.User{}, err
	}

	user := domain.User{
		Id:        Id,
		Age:       Age,
		FirstName: FirstName,
		LastName:  LastName,
		Email:     Email,
		CreatedAt: CreatedAt,
		Height:    Height,
		Active:    Active,
	}

	db.Users = append(db.Users, user)
	db.LastId++

	err = file.WriteJSONFile("users.json", db)

	return user, err
}

func NewRepository() Repository {
	return &repository{}
}

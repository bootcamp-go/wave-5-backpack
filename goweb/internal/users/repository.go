package users

import (
	"errors"

	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/file"
)

type Repository interface {
	GetById(Id int) (domain.User, error)
	GetAll(filters map[string]interface{}) ([]domain.User, error)
	Store(Id, Age int, FirstName, LastName, Email, CreatedAt string, Height float64, Active bool) (domain.User, error)
	Update(Id, Age int, FirstName, LastName, Email, CreatedAt string, Height float64, Active bool) (domain.User, error)
	UpdateAgeLastName(Id, Age int, LastName string) (domain.User, error)
	Delete(Id int) error
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

func (r *repository) GetAll(filters map[string]interface{}) ([]domain.User, error) {
	db, err := file.ReadJSONFile("users.json")

	if err != nil {
		return nil, err
	}

	users, err := filterUsers(filters, db.Users)

	return *users, err
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

func (r *repository) Update(Id, Age int, FirstName, LastName, Email, CreatedAt string, Height float64, Active bool) (domain.User, error) {
	db, err := file.ReadJSONFile("users.json")
	if err != nil {
		return domain.User{}, err
	}

	u := domain.User{
		Id:        Id,
		Age:       Age,
		FirstName: FirstName,
		LastName:  LastName,
		Email:     Email,
		CreatedAt: CreatedAt,
		Height:    Height,
		Active:    Active,
	}

	updated := false
	for i, user := range db.Users {
		if Id == user.Id {
			db.Users[i] = u
			updated = true
			break
		}
	}

	if !updated {
		return domain.User{}, errors.New("usuario no encontrado")
	}

	err = file.WriteJSONFile("users.json", db)
	return u, err
}

func (r *repository) UpdateAgeLastName(Id, Age int, LastName string) (domain.User, error) {
	db, err := file.ReadJSONFile("users.json")
	if err != nil {
		return domain.User{}, err
	}
	u := domain.User{}

	updated := false
	for i, user := range db.Users {
		if Id == user.Id {
			if Age != 0 {
				db.Users[i].Age = Age
			}
			if LastName != "" {
				db.Users[i].LastName = LastName
			}
			u = db.Users[i]
			updated = true
			break
		}
	}

	if !updated {
		return u, errors.New("usuario no encontrado")
	}

	err = file.WriteJSONFile("users.json", db)
	return u, err
}

func (r *repository) Delete(Id int) error {
	db, err := file.ReadJSONFile("users.json")
	if err != nil {
		return err
	}

	updated := false
	for i, user := range db.Users {
		if Id == user.Id {
			db.Users = append(db.Users[:i], db.Users[i+1:]...)
			updated = true
			break
		}
	}
	if !updated {
		return errors.New("usuario no encontrado")
	}

	err = file.WriteJSONFile("users.json", db)
	return err
}

func filterUsers(filters map[string]interface{}, users []domain.User) (*[]domain.User, error) {
	resultUsers := []domain.User{}

	for _, user := range users {
		if Id, ok := filters["Id"]; ok && Id != user.Id {
			continue
		}
		if Age, ok := filters["Age"]; ok && Age != user.Age {
			continue
		}
		if FirstName, ok := filters["FirstName"]; ok && FirstName != user.FirstName {
			continue
		}
		if LastName, ok := filters["LastName"]; ok && LastName != user.LastName {
			continue
		}
		if Email, ok := filters["Email"]; ok && Email != user.Email {
			continue
		}
		if CreatedAt, ok := filters["CreatedAt"]; ok && CreatedAt != user.CreatedAt {
			continue
		}
		if Height, ok := filters["Height"]; ok && Height != user.Height {
			continue
		}
		if Active, ok := filters["Active"]; ok && Active != user.Active {
			continue
		}

		resultUsers = append(resultUsers, user)
	}

	return &resultUsers, nil
}

func NewRepository() Repository {
	return &repository{}
}

package products

import (
	"errors"
	"fmt"
	"goweb/go-web-II/internal/domain"
	"goweb/go-web-II/pkg/store"
)

type Repository interface {
	GetAll() (*[]domain.User, error)
	Store(age int, name, surname, email, created string, active bool) (domain.User, error)
	Update(id, age int, name, surname, email, created string, active bool) (domain.User, error)
	Delete(id int) error
	LastId() (int, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) LastId() (int, error) {
	var us []domain.User
	if err := r.db.Read(&us); err != nil {
		return 0, err
	}
	if len(us) > 0 {
		return us[len(us)-1].Id, nil
	}
	return 0, nil
}

func (r *repository) GetById(Id int) (domain.User, error) {
	var us []domain.User
	if err := r.db.Read(&us); err != nil {
		return domain.User{}, err
	}

	for _, user := range us {
		if Id == user.Id {
			return user, nil
		}
	}

	return domain.User{}, errors.New("usuario no encontrado")
}

func (r *repository) GetAll() (*[]domain.User, error) {
	var us *[]domain.User
	if err := r.db.Read(&us); err != nil {
		return nil, err
	}

	return us, nil
}

func (r *repository) Store(age int, name, surname, email, created string, active bool) (domain.User, error) {

	var us []domain.User
	if err := r.db.Read(&us); err != nil {
		return domain.User{}, err
	}
	id, _ := r.LastId()

	user := domain.User{
		Id:      id,
		Age:     age,
		Name:    name,
		Surname: surname,
		Email:   email,
		Created: created,
		Active:  active,
	}

	us = append(us, user)

	err := r.db.Write(us)

	return user, err
}

func (r *repository) Update(id, age int, name, surname, email, created string, active bool) (domain.User, error) {
	var us []domain.User
	if err := r.db.Read(&us); err != nil {
		return domain.User{}, err
	}

	u := domain.User{
		Id:      id,
		Age:     age,
		Name:    name,
		Surname: surname,
		Email:   email,
		Created: created,
		Active:  active,
	}

	updated := false
	for i, user := range us {
		if id == user.Id {
			us[i] = u
			updated = true
			break
		}
	}

	if !updated {
		return domain.User{}, errors.New("usuario no encontrado")
	}

	err := r.db.Write(us)
	return u, err
}

func (r *repository) UpdateAgeLastName(Id, Age int, LastName string) (domain.User, error) {
	var us []domain.User
	if err := r.db.Read(&us); err != nil {
		return domain.User{}, err
	}
	u := domain.User{}

	updated := false
	for i, user := range us {
		if Id == user.Id {
			if Age != 0 {
				us[i].Age = Age
			}
			if LastName != "" {
				us[i].Surname = LastName
			}
			u = us[i]
			updated = true
			break
		}
	}

	if !updated {
		return u, errors.New("usuario no encontrado")
	}

	err := r.db.Write(us)
	return u, err
}

func (r *repository) Delete(Id int) error {
	var us []domain.User
	if err := r.db.Read(&us); err != nil {
		return err
	}

	updated := false
	for i, user := range us {
		if Id == user.Id {
			us = append(us[:i], us[i+1:]...)
			updated = true
			break
		}
	}
	if !updated {
		return fmt.Errorf("el usuario %v, no ha sido encontrado", Id)
	}

	err := r.db.Write(us)
	return err
}

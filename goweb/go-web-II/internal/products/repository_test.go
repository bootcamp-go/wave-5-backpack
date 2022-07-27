package products

import (
	"goweb/go-web-II/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

func (s *StubStore) GetAll() (*[]domain.User, error) {
	return &[]domain.User{
		{
			Id:      0,
			Name:    "Nahuel",
			Surname: "Monserrat",
			Email:   "nahuelmonserrat@gmail.com",
			Age:     20,
			Active:  true,
			Created: "20/20/2021",
		},
		{
			Id:      0,
			Name:    "Nahuel",
			Surname: "Dominguez",
			Email:   "nahuelDominguez@gmail.com",
			Age:     22,
			Active:  false,
			Created: "09/10/2010",
		},
	}, nil
}

func (s *StubStore) LastId() (int, error) {
	return 2, nil
}

func (s *StubStore) Store(age int, name, surname, email, created string, active bool) (domain.User, error) {
	return domain.User{
		Id:      0,
		Name:    "Nahuel",
		Surname: "Monserrat",
		Email:   "nahuelmonserrat@gmail.com",
		Age:     20,
		Active:  true,
		Created: "20/20/2021",
	}, nil
}

func (s *StubStore) Update(id, age int, name, surname, email, created string, active bool) (domain.User, error) {
	return domain.User{
		Id:      0,
		Name:    "Nahuel",
		Surname: "Monserrat",
		Email:   "nahuelmonserrat@gmail.com",
		Age:     20,
		Active:  true,
		Created: "20/20/2021",
	}, nil
}

func (s *StubStore) Delete(id int) error {
	return nil
}

func TestGetAll(t *testing.T) {
	// Arrange
	service := NewService(&StubStore{})
	us := &[]domain.User{
		{
			Id:      0,
			Name:    "Nahuel",
			Surname: "Monserrat",
			Email:   "nahuelmonserrat@gmail.com",
			Age:     20,
			Active:  true,
			Created: "20/20/2021",
		},
		{
			Id:      0,
			Name:    "Nahuel",
			Surname: "Dominguez",
			Email:   "nahuelDominguez@gmail.com",
			Age:     22,
			Active:  false,
			Created: "09/10/2010",
		},
	}

	// Act
	users, err := service.GetAll()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, us, users)
}

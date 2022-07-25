package products

import (
	"goweb/go-web-II/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStore struct {
	Read bool
}

func (s *MockStore) GetAll() ([]*domain.User, error) {
	return []*domain.User{
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

func (s *MockStore) LastId() (int, error) {
	return 2, nil
}

func (s *MockStore) Store(id, age int, name, surname, email, created string, active bool) (*domain.User, error) {
	return &domain.User{
		Id:      0,
		Name:    "Nahuel",
		Surname: "Monserrat",
		Email:   "nahuelmonserrat@gmail.com",
		Age:     20,
		Active:  true,
		Created: "20/20/2021",
	}, nil
}

func (s *MockStore) Update(id, age int, name, surname, email, created string, active bool) (*domain.User, error) {
	s.Read = true
	return &domain.User{
		Id:      id,
		Name:    name,
		Surname: surname,
		Email:   email,
		Age:     age,
		Active:  active,
		Created: created,
	}, nil
}

func (s *MockStore) Delete(id int) error {
	return nil
}

func TestUpdate(t *testing.T) {
	//Arrange
	mock := MockStore{}
	service := NewService(&mock)
	beforeUpdate := &domain.User{
		Id:      0,
		Name:    "Nahuel",
		Surname: "Monserrat",
		Email:   "nahuelmonserrat@gmail.com",
		Age:     20,
		Active:  true,
		Created: "20/20/2021",
	}
	userExpected := &domain.User{
		Id:      0,
		Name:    "Cristiano",
		Surname: "Ronaldo",
		Email:   "nahuelmonserrat@gmail.com",
		Age:     20,
		Active:  true,
		Created: "20/20/2021",
	}

	// Act
	user, err := service.Update(beforeUpdate.Id, beforeUpdate.Age, "Cristiano", "Ronaldo", beforeUpdate.Email, beforeUpdate.Created, beforeUpdate.Active)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, userExpected, user)
	assert.True(t, mock.Read)
}

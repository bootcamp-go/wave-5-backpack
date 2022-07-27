package products

import (
	"fmt"
	"goweb/go-web-II/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	DataMock      []domain.User
	errWrite      string
	errRead       string
	ReadWasCalled bool
}

func (s *MockStorage) GetAll() (*[]domain.User, error) {
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

func (s *MockStorage) LastId() (int, error) {
	return 2, nil
}

func (s *MockStorage) Store(age int, name, surname, email, created string, active bool) (domain.User, error) {
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

func (s *MockStorage) Update(id, age int, name, surname, email, created string, active bool) (domain.User, error) {
	s.ReadWasCalled = true
	return domain.User{
		Id:      id,
		Name:    name,
		Surname: surname,
		Email:   email,
		Age:     age,
		Active:  active,
		Created: created,
	}, nil
}

func (s *MockStorage) Delete(id int) error {
	return nil
}

func (ms *MockStorage) Write(data interface{}) error {
	if ms.errWrite != "" {
		return fmt.Errorf(ms.errWrite)
	}
	a := data.([]domain.User)
	ms.DataMock = append(ms.DataMock, a...)
	return nil
}

func (ms *MockStorage) Read(data interface{}) error {
	if ms.errRead != "" {
		return fmt.Errorf(ms.errRead)
	}
	user := data.(*[]domain.User)
	*user = ms.DataMock
	ms.ReadWasCalled = true
	return nil
}

func TestUpdate(t *testing.T) {
	//Arrange
	mock := MockStorage{}
	service := NewService(&mock)
	beforeUpdate := domain.User{
		Id:      0,
		Name:    "Nahuel",
		Surname: "Monserrat",
		Email:   "nahuelmonserrat@gmail.com",
		Age:     20,
		Active:  true,
		Created: "20/20/2021",
	}
	userExpected := domain.User{
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
	assert.True(t, mock.ReadWasCalled)
}

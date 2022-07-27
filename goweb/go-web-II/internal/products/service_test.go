package products

import (
	"fmt"
	"goweb/go-web-II/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationTestUpdate(t *testing.T) {
	//arrange
	usersTest := []domain.User{
		{
			Id:      1,
			Name:    "Nahuel",
			Surname: "Monserrat",
			Email:   "nahuelmonserrat@gmail.com",
			Age:     20,
			Active:  true,
			Created: "20/20/2021",
		},
	}

	userUpdated := domain.User{
		Id:      1,
		Name:    "Ramon",
		Surname: "Perez",
		Email:   "aramon@gmail.com",
		Age:     20,
		Active:  false,
		Created: "24/11/2001",
	}
	db := MockStorage{
		DataMock:      usersTest,
		ReadWasCalled: false,
	}
	//act
	repo := NewRepository(&db)
	service := NewService(repo)
	result, err := service.Update(1, 20, "Ramon", "Perez", "aramon@gmail.com", "24/11/2001", false)
	//assert
	assert.Nil(t, err)
	assert.Equal(t, userUpdated, result)
	assert.True(t, db.ReadWasCalled)
}

func TestIntegrationDelete(t *testing.T) {
	//arrange
	userNotFound := fmt.Errorf("el usuario %v, no ha sido encontrado", 50)
	usersTest := []domain.User{
		{
			Id:      1,
			Name:    "Nahuel",
			Surname: "Monserrat",
			Email:   "nahuelmonserrat@gmail.com",
			Age:     20,
			Active:  true,
			Created: "20/20/2021",
		},
	}
	db := MockStorage{
		DataMock:      usersTest,
		ReadWasCalled: false,
	}
	//act
	repo := NewRepository(&db)
	service := NewService(repo)
	errUserNotFound := service.Delete(50)
	err2 := service.Delete(1)
	//assert
	assert.Equal(t, errUserNotFound, userNotFound)
	assert.Nil(t, err2)
}

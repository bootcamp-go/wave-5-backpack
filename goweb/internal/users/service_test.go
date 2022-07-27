package users

import (
	"goweb/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceUpdate(t *testing.T) {
	// arrange
	usersDatabase := []domain.User{
		{
			ID:           1,
			Name:         "Cristian",
			LastName:     "Ladino",
			Email:        "test@gmail.com",
			Age:          23,
			Height:       1.68,
			Active:       true,
			CreationDate: "21-07-2022",
		},
		{
			ID:           2,
			Name:         "Camilo",
			LastName:     "Pinzon",
			Email:        "test1@gmail.com",
			Age:          30,
			Height:       1.78,
			Active:       false,
			CreationDate: "01-09-2021",
		},
	}

	updateUser := domain.User{
		ID:           1,
		Name:         "Alexis",
		LastName:     "Dueñas",
		Email:        "testUpdate@gmail.com",
		Age:          24,
		Height:       1.69,
		Active:       false,
		CreationDate: "20-07-2022",
	}

	mockStorage := MockStore{
		Data: usersDatabase,
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	userUpdate, err := service.Update(updateUser.ID, updateUser.Name, updateUser.LastName, updateUser.Email, updateUser.Age, updateUser.Height, updateUser.Active, updateUser.CreationDate)

	assert.Nil(t, err)
	assert.Equal(t, updateUser.ID, userUpdate.ID)
	assert.Equal(t, true, mockStorage.ReadCalled)

}

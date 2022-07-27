package users

import (
	"errors"
	"goweb/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationUpdate(t *testing.T) {
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
		Data:      usersDatabase,
		errDelete: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	userUpdate, err := service.Update(updateUser.ID, updateUser.Name, updateUser.LastName, updateUser.Email, updateUser.Age, updateUser.Height, updateUser.Active, updateUser.CreationDate)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, updateUser.ID, userUpdate.ID)
	assert.Equal(t, true, mockStorage.ReadCalled)

}

func TestServiceIntegrationDelete(t *testing.T) {
	// arrange
	id := 1
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
	mockStorage := MockStore{
		Data:      usersDatabase,
		errDelete: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	err := service.Delete(id)

	// assert
	assert.Nil(t, err)

}

func TestServiceIntegrationDeleteFail(t *testing.T) {
	// arrange
	id := 0
	expectedError := errors.New("usuario no encontrado")
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
	mockStorage := MockStore{
		Data:      usersDatabase,
		errDelete: "usuario no encontrado",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	err := service.Delete(id)

	// assert
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, expectedError.Error())
}

// aumento coverage

func TestServiceIntegrationGetAll(t *testing.T) {
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
	mockStorage := MockStore{
		Data:      usersDatabase,
		errDelete: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	users, err := service.GetAll()

	// assert
	assert.Nil(t, err)
	assert.Equal(t, usersDatabase, users)

}

func TestServiceIntegrationNewUser(t *testing.T) {
	// arrange

	newUser := domain.User{
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
		Data:      nil,
		errDelete: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	createdUser, err := service.NewUser(newUser.Name, newUser.LastName, newUser.Email, newUser.Age, newUser.Height, newUser.Active, newUser.CreationDate)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, newUser, createdUser)
	assert.Equal(t, true, mockStorage.ReadCalled)
}

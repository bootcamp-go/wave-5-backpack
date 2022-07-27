package users

import (
	"goweb/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

func (fs *StubStore) Read(data interface{}) error {
	a := data.(*[]domain.User)
	p1 := domain.User{
		ID:           1,
		Name:         "Cristian",
		LastName:     "Ladino",
		Email:        "test@gmail.com",
		Age:          23,
		Height:       1.68,
		Active:       true,
		CreationDate: "21-07-2022",
	}

	p2 := domain.User{
		ID:           2,
		Name:         "Camilo",
		LastName:     "Pinzon",
		Email:        "test1@gmail.com",
		Age:          30,
		Height:       1.78,
		Active:       false,
		CreationDate: "01-09-2021",
	}

	*a = append(*a, p1)
	*a = append(*a, p2)
	return nil
}

func (fs *StubStore) Write(data interface{}) error {
	return nil
}

func (fs *StubStore) Ping() error {
	return nil
}

func TestGetAll(t *testing.T) {
	// arrage
	expected := []domain.User{
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

	// act

	myStubUser := &StubStore{}
	repo := NewRepository(myStubUser)
	users, err := repo.GetAll()

	// assert
	assert.Nil(t, err)
	assert.Equal(t, expected, users)
}

// Ejercicio 2 - Test Unitario UpdateName()

type MockStore struct {
	ReadCalled bool
	Data       []domain.User
	errDelete  string
}

func (ms *MockStore) Read(data interface{}) error {
	ms.ReadCalled = true
	user := data.(*[]domain.User)
	*user = ms.Data
	return nil
}

func (ms *MockStore) Write(data interface{}) error {
	return nil
}

func (ms *MockStore) Ping() error {
	return nil
}

func TestUpdateName(t *testing.T) {
	// arrage
	id, name := 1, "Update After"
	users := []domain.User{{ID: 1, Name: "Update Before", LastName: "Ladino", Email: "test@gmail.com", Age: 23, Height: 1.68, Active: true, CreationDate: "21-07-2022"}}

	// act
	mock := MockStore{Data: users}
	repo := NewRepository(&mock)

	userUpdate, err := repo.UpdateName(id, name)

	// assert
	assert.Nil(t, err)

	assert.Equal(t, id, userUpdate.ID)
	assert.Equal(t, name, userUpdate.Name)
	assert.Equal(t, true, mock.ReadCalled)

}

// Test Integracion
func TestUpdate(t *testing.T) {
	// arrage
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
	// act
	mock := MockStore{Data: usersDatabase}
	repo := NewRepository(&mock)

	userUpdate, err := repo.Update(updateUser.ID, updateUser.Name, updateUser.LastName, updateUser.Email, updateUser.Age, updateUser.Height, updateUser.Active, updateUser.CreationDate)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, updateUser, userUpdate)
	assert.Equal(t, updateUser.ID, userUpdate.ID)
	assert.Equal(t, true, mock.ReadCalled)

}

func TestDelete(t *testing.T) {
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
	mock := MockStore{Data: usersDatabase}
	repo := NewRepository(&mock)

	err := repo.Delete(id)

	assert.Nil(t, err)
}

// subiendo coverage
func TestNewUser(t *testing.T) {
	// arrage
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

	newUser := domain.User{
		ID:           3,
		Name:         "Alexis",
		LastName:     "Dueñas",
		Email:        "testUpdate@gmail.com",
		Age:          24,
		Height:       1.69,
		Active:       false,
		CreationDate: "20-07-2022",
	}

	// act
	mock := MockStore{Data: usersDatabase}
	repo := NewRepository(&mock)

	createdUser, err := repo.NewUser(newUser.ID, newUser.Name, newUser.LastName, newUser.Email, newUser.Age, newUser.Height, newUser.Active, newUser.CreationDate)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, newUser, createdUser)
	assert.Equal(t, newUser.ID, createdUser.ID)
}

func TestLastId(t *testing.T) {
	// arrage
	expectedId := 2
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

	// act
	mock := MockStore{Data: usersDatabase}
	repo := NewRepository(&mock)

	lastId, err := repo.LastID()

	// assert

	assert.Nil(t, err)
	assert.Equal(t, expectedId, lastId)
}

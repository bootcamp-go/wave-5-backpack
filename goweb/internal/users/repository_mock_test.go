package users

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockStore struct {
	readWasCalled bool
}

func (ms *MockStore) Write(data interface{}) error {
	return nil
}

func (ms *MockStore) Ping() error {
	return nil
}

func (ms *MockStore) Read(data interface{}) error {
	BeforeUpd := data.(*[]domain.Users)
	*BeforeUpd = []domain.Users{
		{Id: 1, Name: "Juan", LastName: "Perez", Height: 1.82, CreationDate: "1992"},
		{Id: 2, Name: "Simon", LastName: "Fernandez", Height: 1.65, CreationDate: "1232"},
	}
	ms.readWasCalled = true
	return nil

}

func TestUpdateNAL(t *testing.T) {
	myMockStore := MockStore{}
	repo := NewRepository(&myMockStore)
	expected := domain.Users{Id: 1, Name: "Juan", LastName: "Nuevo", Height: 1.82, Age: 15, CreationDate: "1992"}

	user, err := repo.UpdateLastNameAndAge(1, 15, "Nuevo")
	assert.True(t, myMockStore.readWasCalled)
	assert.Equal(t, user, expected)
	assert.Nil(t, err)
}

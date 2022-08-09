package users

import (
	"context"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

func (ms *StubStore) Write(data interface{}) error {
	return nil
}

func (ms *StubStore) Ping() error {
	return nil
}

func (ss *StubStore) Read(data interface{}) error {
	a := data.(*[]domain.Users)
	*a = []domain.Users{
		{Id: 1, Name: "Juan", LastName: "Perez", Height: 1.82, CreationDate: "1992"},
		{Id: 2, Name: "Simon", LastName: "Fernandez", Height: 1.65, CreationDate: "1232"},
	}
	return nil
}

func TestGetAllRepo(t *testing.T) {
	myStubStore := StubStore{}
	repo := NewRepository(&myStubStore)
	expected := []domain.Users{
		{Id: 1, Name: "Juan", LastName: "Perez", Height: 1.82, CreationDate: "1992"},
		{Id: 2, Name: "Simon", LastName: "Fernandez", Height: 1.65, CreationDate: "1232"},
	}
	user, err := repo.GetAll(context.TODO())

	assert.Equal(t, user, expected)
	assert.Nil(t, err)
}

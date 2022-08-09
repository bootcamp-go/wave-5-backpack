package users

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	_ "github.com/go-sql-driver/mysql"
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

func TestUpdate(t *testing.T) {
	myMockStore := MockStore{}
	repo := NewRepository(&myMockStore)
	expected := domain.Users{Id: 1, Name: "Juan", LastName: "Nuevo", Height: 1.82, Age: 15, CreationDate: "1992"}

	user, err := repo.UpdateLastNameAndAge(context.TODO(), 1, 15, "Nuevo")
	assert.True(t, myMockStore.readWasCalled)
	assert.Equal(t, user, expected)
	assert.Nil(t, err)
}

func TestGetAllWithContextTO(t *testing.T) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/users_db")
	if err != nil {
		t.Log(err)
	}
	repo := NewRepositoryDB(db)
	time := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), time)
	defer cancel()
	_, err = repo.GetAll(ctx)
	if err != nil {
		t.Errorf("el error deberia ser nulo %v", err)
	}
}

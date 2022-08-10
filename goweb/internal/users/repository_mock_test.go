package users

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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

func TestRepositoryStoreAndGetOneTXDB(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/users_db")
	db, err := sql.Open("txdb", uuid.New().String())
	assert.NoError(t, err)
	repo := NewRepositoryDB(db)
	u, err := repo.Store(context.TODO(), 0, 10, "abc", "abc", "a@a.com", "2020", 1.78, true)
	assert.NoError(t, err)
	assert.NotZero(t, u)
	u, err2 := repo.GetOne(context.TODO(), u.Id)
	assert.NoError(t, err2)
}

func TestDeleteGetAllGetOne(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")
	db, err := sql.Open("txdb", uuid.New().String())
	assert.NoError(t, err)
	var exist bool
	repo := NewRepositoryDB(db)
	er2 := repo.Delete(context.TODO(), 9)
	listUsers, er := repo.GetAll(context.TODO())
	for i := 0; i < len(listUsers); i++ {
		if listUsers[i].Id == 9 {
			exist = true
		}
	}
	assert.False(t, exist)
	assert.Nil(t, er)
	assert.Nil(t, er2)
}

func TestSaveErrorMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO storage.users(nombre,apellido,email,edad,altura,activo,fechaCreacion) VALUES(?,?,?,?,?,?,?)"))
	mock.ExpectExec("INSERT INTO storage.users").WillReturnError(fmt.Errorf("Ocurrio un error al ejecutar la BBDD"))
	repo := NewRepositoryDB(db)
	user := domain.Users{
		Id:           12,
		Name:         "Timon",
		LastName:     "ABC",
		Email:        "a@a.com",
		Age:          22,
		Height:       1.8,
		Active:       true,
		CreationDate: "2020",
	}
	userO, er2 := repo.Store(context.TODO(), user.Id, user.Age, user.Name, user.LastName, user.Email, user.CreationDate, user.Height, user.Active)
	assert.Error(t, er2)
	assert.Zero(t, userO)
}

func TestSaveMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO users(name, last_name, email, age, height, active, creation_date) VALUES( ?, ?, ?, ?, ?, ?, ? )"))
	mock.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(12, 1))
	repo := NewRepositoryDB(db)
	user := domain.Users{
		Id:           12,
		Name:         "Timon",
		LastName:     "ABC",
		Email:        "a@a.com",
		Age:          22,
		Height:       1.8,
		Active:       true,
		CreationDate: "2020",
	}
	userO, er2 := repo.Store(context.TODO(), user.Id, user.Age, user.Name, user.LastName, user.Email, user.CreationDate, user.Height, user.Active)
	assert.NoError(t, er2)
	assert.NotZero(t, userO)
	assert.Equal(t, user.Id, userO.Id)
}

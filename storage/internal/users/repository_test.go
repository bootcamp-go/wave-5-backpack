package users

import (
	"context"
	"database/sql"
	"goweb/db"
	"goweb/internal/domain"
	"testing"

	"github.com/google/uuid"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	db := db.Connection()

	repo := NewRepository(db)

	users, err := repo.GetAll(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))
}

func TestStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	mock.ExpectPrepare("INSERT INTO users")
	mock.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(1, 1))

	repo := NewRepository(db)

	usr := domain.User{
		Id:            1,
		Nombre:        "Dala",
		Apellido:      "Algo",
		Email:         "algo@gmail.com",
		Edad:          20,
		Altura:        1.65,
		Activo:        true,
		FechaCreacion: "2022-08-10",
	}

	u, err := repo.Store(context.TODO(), usr)

	assert.NoError(t, err)
	assert.Equal(t, 1, u)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestStoreTXDB(t *testing.T) {
	txdb.Register("txdb", "mysql", db.DataSource) //conectar db
	db, err := sql.Open("txdb", uuid.New().String())
	assert.NoError(t, err)

	defer db.Close()

	user := domain.User{
		Id:            1,
		Nombre:        "Dala",
		Apellido:      "Algo",
		Email:         "algo@gmail.com",
		Edad:          20,
		Altura:        1.65,
		Activo:        true,
		FechaCreacion: "2022-08-10 00:00:00",
	}

	repo := NewRepository(db)
	usr, err := repo.Store(context.TODO(), user)

	user.Id = usr
	assert.NoError(t, err)

	u, err := repo.GetById(context.TODO(), usr)

	assert.NoError(t, err)
	assert.NotZero(t, usr)
	assert.Equal(t, user, u)
}

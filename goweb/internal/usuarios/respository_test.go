package usuarios

import (
	"context"
	"database/sql"
	"errors"
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

type StubStore struct {
	errRead string
}

func (ss *StubStore) Validate() error {
	return nil
}

func (ss *StubStore) Write(data interface{}) error {
	return nil
}

func (ss *StubStore) Read(data interface{}) error {
	if ss.errRead != "" {
		return errors.New("error al leer la bd")
	}
	a := data.(*[]domain.Usuarios) //ACA ESTOY RECIBIENDO DESDE REPOSITORY UN PUNTERO DE LISTA DE USUARIOS
	*a = []domain.Usuarios{        //ACA LLENO ESOS VALORES DEL PUNTERO, por eso lo desreferencio
		{Id: 1, Nombre: "Yvo", Apellido: "Pintos", Altura: 3, FechaCreacion: "1992"},
		{Id: 2, Nombre: "Pedro", Apellido: "Juan", Altura: 3, FechaCreacion: "1232"},
	}
	return nil

}

func TestGetAllRepo(t *testing.T) {
	myStubStore := StubStore{}
	repo := NewRepository(&myStubStore) //Probando el repository, yo le paso datos dummy a lo que quiero probar
	expected := []domain.Usuarios{
		{Id: 1, Nombre: "Yvo", Apellido: "Pintos", Altura: 3, FechaCreacion: "1992"},
		{Id: 2, Nombre: "Pedro", Apellido: "Juan", Altura: 3, FechaCreacion: "1232"},
	}

	user, err := repo.GetAll(context.TODO())

	assert.Equal(t, user, expected)
	assert.Nil(t, err)
}

func TestGetAllWithContextTO(t *testing.T) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/storage")
	if err != nil {
		t.Log(err)
	}
	repo := NewRepositoryBD(db)
	time := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), time)
	defer cancel()
	_, err = repo.GetAll(ctx)
	if err != nil {
		t.Errorf("el error deberia ser nulo %v", err)
	}
}

func TestSaveMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()
	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO storage.users(nombre,apellido,email,edad,altura,activo,fechaCreacion) VALUES(?,?,?,?,?,?,?)"))
	mock.ExpectExec("INSERT INTO storage.users").WillReturnResult(sqlmock.NewResult(12, 1))

	repo := NewRepositoryBD(db)
	user := domain.Usuarios{
		Id:            12,
		Nombre:        "TayronaT",
		Apellido:      "Fante",
		Email:         "titi",
		Edad:          30,
		Altura:        20,
		Activo:        false,
		FechaCreacion: "2020",
	}
	userO, er2 := repo.Guardar(user.Id, user.Nombre, user.Apellido, user.Email, user.Edad, user.Altura, user.Activo, user.FechaCreacion)

	assert.NoError(t, er2)
	assert.NotZero(t, userO)
	assert.Equal(t, user.Id, userO.Id)
}

func TestSaveErrorMock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()
	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO storage.users(nombre,apellido,email,edad,altura,activo,fechaCreacion) VALUES(?,?,?,?,?,?,?)"))
	mock.ExpectExec("INSERT INTO storage.users").WillReturnError(fmt.Errorf("Ocurrio un error al ejecutar la BBDD"))

	repo := NewRepositoryBD(db)
	user := domain.Usuarios{
		Id:            12,
		Nombre:        "TayronaT",
		Apellido:      "Fante",
		Email:         "titi",
		Edad:          30,
		Altura:        20,
		Activo:        false,
		FechaCreacion: "2020",
	}
	userO, er2 := repo.Guardar(user.Id, user.Nombre, user.Apellido, user.Email, user.Edad, user.Altura, user.Activo, user.FechaCreacion)

	//assert.Error(t, fmt.Errorf("Ocurrio un error al ejecutar la BBDD"))
	assert.Error(t, er2)
	assert.Zero(t, userO)
}

func TestSaveAndGetWithContext(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")
	db, err := sql.Open("txdb", uuid.New().String())
	assert.NoError(t, err)

	repo := NewRepositoryBD(db)
	user := []domain.Usuarios{{
		Id:            26,
		Nombre:        "TayronaT",
		Apellido:      "Fante",
		Email:         "titi",
		Edad:          30,
		Altura:        20,
		Activo:        false,
		FechaCreacion: "2020",
	},
	}
	userO, er2 := repo.Guardar(24, "TayronaT", "Fante", "titi", 30, 20, false, "2020")
	users, er := repo.GetByName("TayronaT")
	assert.NoError(t, er)
	assert.NoError(t, er2)
	fmt.Println(userO)
	assert.NotZero(t, userO)
	assert.NotEmpty(t, users)
	assert.Equal(t, user, users)
}

func TestUpdateBD(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")
	db, err := sql.Open("txdb", uuid.New().String())
	assert.NoError(t, err)

	repo := NewRepositoryBD(db)
	user := domain.Usuarios{
		Id:            15,
		Nombre:        "TayronaT",
		Apellido:      "Fante",
		Email:         "titi",
		Edad:          30,
		Altura:        20,
		Activo:        false,
		FechaCreacion: "2020",
	}
	userO, er2 := repo.Update(context.TODO(), user.Id, user.Nombre, user.Apellido, user.Email, user.Edad, user.Altura, user.Activo, user.FechaCreacion)

	assert.NoError(t, er2)
	assert.NotZero(t, userO)
	assert.Equal(t, user, userO)
}

func TestDeleteGetAllGetOne(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")
	db, err := sql.Open("txdb", uuid.New().String())
	assert.NoError(t, err)
	var exist bool

	repo := NewRepositoryBD(db)
	er2 := repo.Delete(9)

	listUsers, er := repo.GetAll(context.TODO())
	for i := 0; i < len(listUsers); i++ {
		if listUsers[i].Id == 9 {
			exist = true
		}
	}
	userOne, er3 := repo.GetById(context.TODO(), 9)
	assert.False(t, exist)
	assert.Zero(t, userOne)
	assert.Nil(t, er)
	assert.Nil(t, er2)
	assert.Nil(t, er3)
}

func TestGetOneInexist(t *testing.T) {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")
	db, err := sql.Open("txdb", uuid.New().String())
	assert.NoError(t, err)

	repo := NewRepositoryBD(db)
	user, er := repo.GetByName("XXX")
	assert.Zero(t, user)
	assert.Nil(t, er)
}

func TestGetAllRepoErrRead(t *testing.T) {
	myStubStore := StubStore{
		errRead: "error",
	}
	repo := NewRepository(&myStubStore) //Probando el repository, yo le paso datos dummy a lo que quiero probar
	expected := "error al leer la bd"

	user, err := repo.GetAll(context.TODO())

	assert.EqualError(t, err, expected)
	assert.Nil(t, user)
}

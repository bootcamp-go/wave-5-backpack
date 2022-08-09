package usuarios

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	_ "github.com/go-sql-driver/mysql"
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

package users

import (
	"testing"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

func (fs *StubStore) Read(data interface{}) error {
	fecha, _ := time.Parse("2006-01-02", "2022-07-25")
	user := data.(*[]domain.ModelUser)
	*user = []domain.ModelUser{
		{Id: 1, Nombre: "Juan", Apellido: "Perez", Email: "juan.perez@gmail.com", Edad: 22, Altura: 1.60, Activo: true, FechaCreacion: fecha, Borrado: false},
		{Id: 2, Nombre: "Norma", Apellido: "Carrasco", Email: "norma.carrasco@gmail.com", Edad: 28, Altura: 1.56, Activo: false, FechaCreacion: fecha, Borrado: false},
	}
	return nil
}

func (fs *StubStore) Write(data interface{}) error {
	return nil
}

func (fs *StubStore) Ping() error {
	return nil
}

// Ejercicio 1
func TestGet(t *testing.T) {
	// Arrange
	fecha, _ := time.Parse("2006-01-02", "2022-07-25")
	store := StubStore{}
	repo := NewRepository(&store)
	esperado := []domain.ModelUser{
		{Id: 1, Nombre: "Juan", Apellido: "Perez", Email: "juan.perez@gmail.com", Edad: 22, Altura: 1.60, Activo: true, FechaCreacion: fecha, Borrado: false},
		{Id: 2, Nombre: "Norma", Apellido: "Carrasco", Email: "norma.carrasco@gmail.com", Edad: 28, Altura: 1.56, Activo: false, FechaCreacion: fecha, Borrado: false},
	}

	// Act
	users, err := repo.GetAll()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, esperado, users)
}

// Ejercicio 2
type MockStorage struct {
	ReadInvoked      bool
	BeforeUpdateUser []domain.ModelUser
}

func (fs *MockStorage) Read(data interface{}) error {
	fs.ReadInvoked = true
	user := data.(*[]domain.ModelUser)
	*user = fs.BeforeUpdateUser
	return nil
}

func (fs *MockStorage) Write(data interface{}) error {
	return nil
}

func (fs *MockStorage) Ping() error {
	return nil
}

func TestUpdateApellidoEdad(t *testing.T) {
	// Arrange
	id, apellido, edad := 2, "Paez", 22
	fecha, _ := time.Parse("2006-01-02", "2022-07-25")
	users := []domain.ModelUser{
		{Id: 1, Nombre: "Juan", Apellido: "Perez", Email: "juan.perez@gmail.com", Edad: 22, Altura: 1.60, Activo: true, FechaCreacion: fecha, Borrado: false},
		{Id: 2, Nombre: "Norma", Apellido: "Carrasco", Email: "norma.carrasco@gmail.com", Edad: 28, Altura: 1.56, Activo: false, FechaCreacion: fecha, Borrado: false},
	}
	mock := MockStorage{BeforeUpdateUser: users}
	repo := NewRepository(&mock)

	// Act
	updated, err := repo.UpdateApellidoEdad(id, apellido, edad)

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, id, updated.Id)
	assert.Equal(t, apellido, updated.Apellido)
	assert.True(t, true, mock.ReadInvoked)
}

func TestUpdate(t *testing.T) {
	// Arrange
	fecha, _ := time.Parse("2006-01-02", "2022-07-25")
	users := []domain.ModelUser{
		{Id: 1, Nombre: "Juan", Apellido: "Perez", Email: "juan.perez@gmail.com", Edad: 22, Altura: 1.60, Activo: true, FechaCreacion: fecha, Borrado: false},
		{Id: 2, Nombre: "Norma", Apellido: "Carrasco", Email: "norma.carrasco@gmail.com", Edad: 28, Altura: 1.56, Activo: false, FechaCreacion: fecha, Borrado: false},
	}
	user := domain.ModelUser{Id: 2, Nombre: "Jose", Apellido: "Juan", Email: "jose.juan@gmail.com", Edad: 22, Altura: 1.60, Activo: false, FechaCreacion: fecha, Borrado: false}
	mock := MockStorage{BeforeUpdateUser: users}
	repo := NewRepository(&mock)

	// Act
	updated, err := repo.Update(user.Id, user.Nombre, user.Apellido, user.Email, user.Edad, user.Altura)

	// Assert
	assert.Nil(t, err)

	assert.Equal(t, user.Id, updated.Id)
	assert.Equal(t, user, updated)
	assert.True(t, true, mock.ReadInvoked)
}

func TestUpdateFail(t *testing.T) {
	// Arrange
	fecha, _ := time.Parse("2006-01-02", "2022-07-25")
	users := []domain.ModelUser{
		{Id: 1, Nombre: "Juan", Apellido: "Perez", Email: "juan.perez@gmail.com", Edad: 22, Altura: 1.60, Activo: true, FechaCreacion: fecha, Borrado: false},
		{Id: 2, Nombre: "Norma", Apellido: "Carrasco", Email: "norma.carrasco@gmail.com", Edad: 28, Altura: 1.56, Activo: false, FechaCreacion: fecha, Borrado: false},
	}
	user := domain.ModelUser{Id: 3, Nombre: "Jose", Apellido: "Juan", Email: "jose.juan@gmail.com", Edad: 22, Altura: 1.60, Activo: false, FechaCreacion: fecha, Borrado: false}
	mock := MockStorage{BeforeUpdateUser: users}
	repo := NewRepository(&mock)

	// Act
	updated, err := repo.Update(user.Id, user.Nombre, user.Apellido, user.Email, user.Edad, user.Altura)

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, updated.Id, 0)
	assert.True(t, true, mock.ReadInvoked)
}

func TestGetById(t *testing.T) {
	// Arrange
	id := 2
	fecha, _ := time.Parse("2006-01-02", "2022-07-25")
	users := []domain.ModelUser{
		{Id: 1, Nombre: "Juan", Apellido: "Perez", Email: "juan.perez@gmail.com", Edad: 22, Altura: 1.60, Activo: true, FechaCreacion: fecha, Borrado: false},
		{Id: 2, Nombre: "Norma", Apellido: "Carrasco", Email: "norma.carrasco@gmail.com", Edad: 28, Altura: 1.56, Activo: false, FechaCreacion: fecha, Borrado: false},
	}
	mock := MockStorage{BeforeUpdateUser: users}
	repo := NewRepository(&mock)

	// Act
	user, err := repo.GetById(id)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, user.Id, id)
	assert.True(t, true, mock.ReadInvoked)
}

func TestGetByIdFail(t *testing.T) {
	// Arrange
	id := 3
	fecha, _ := time.Parse("2006-01-02", "2022-07-25")
	users := []domain.ModelUser{
		{Id: 1, Nombre: "Juan", Apellido: "Perez", Email: "juan.perez@gmail.com", Edad: 22, Altura: 1.60, Activo: true, FechaCreacion: fecha, Borrado: false},
		{Id: 2, Nombre: "Norma", Apellido: "Carrasco", Email: "norma.carrasco@gmail.com", Edad: 28, Altura: 1.56, Activo: false, FechaCreacion: fecha, Borrado: false},
	}
	mock := MockStorage{BeforeUpdateUser: users}
	repo := NewRepository(&mock)

	// Act
	user, err := repo.GetById(id)

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, user.Id, 0)
	assert.True(t, true, mock.ReadInvoked)
}

func TestSearch(t *testing.T) {
	// Arrange
	nombre, paterno, email, edad := "Juan", "Perez", "juan.perez@gmail.com", "22"
	fecha, _ := time.Parse("2006-01-02", "2022-07-25")
	users := []domain.ModelUser{
		{Id: 1, Nombre: "Juan", Apellido: "Perez", Email: "juan.perez@gmail.com", Edad: 22, Altura: 1.60, Activo: true, FechaCreacion: fecha, Borrado: false},
		{Id: 2, Nombre: "Norma", Apellido: "Carrasco", Email: "norma.carrasco@gmail.com", Edad: 28, Altura: 1.56, Activo: false, FechaCreacion: fecha, Borrado: false},
	}
	mock := MockStorage{BeforeUpdateUser: users}
	repo := NewRepository(&mock)

	// Act
	user, err := repo.SearchUser(nombre, paterno, email, edad, "", "", "")

	// Assert
	assert.Nil(t, err)
	assert.NotEmpty(t, user)
	assert.True(t, true, mock.ReadInvoked)
}

func TestSearchFail(t *testing.T) {
	// Arrange
	nombre := "Pedro"
	fecha, _ := time.Parse("2006-01-02", "2022-07-25")
	users := []domain.ModelUser{
		{Id: 1, Nombre: "Juan", Apellido: "Perez", Email: "juan.perez@gmail.com", Edad: 22, Altura: 1.60, Activo: true, FechaCreacion: fecha, Borrado: false},
		{Id: 2, Nombre: "Norma", Apellido: "Carrasco", Email: "norma.carrasco@gmail.com", Edad: 28, Altura: 1.56, Activo: false, FechaCreacion: fecha, Borrado: false},
	}
	mock := MockStorage{BeforeUpdateUser: users}
	repo := NewRepository(&mock)

	// Act
	user, err := repo.SearchUser(nombre, "", "", "", "", "", "")

	// Assert
	assert.Nil(t, err)
	assert.Empty(t, user)
	assert.True(t, true, mock.ReadInvoked)
}

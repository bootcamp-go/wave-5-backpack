package usuarios

import (
	"errors"
	"fmt"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type StubDB struct {
	errRead string
	errLast string
	errSave string
}

func (sdb *StubDB) GetAll() ([]domain.Usuarios, error) {
	if sdb.errRead != "" {
		return nil, errors.New("fail to getAll")
	}
	sliceUsers := []domain.Usuarios{{Id: 1, Nombre: "Yvo", Apellido: "Pintos", Email: "yvo", Edad: 30, Altura: 3, Activo: true, FechaCreacion: "1992"}, {Id: 2, Nombre: "Mat", Apellido: "Fant", Email: "mat", Edad: 33, Altura: 3, Activo: true, FechaCreacion: "1990"}}
	return sliceUsers, nil
}

func (sdb *StubDB) GetById(id int) (domain.Usuarios, error) {
	sliceUsers := []domain.Usuarios{{Id: 1, Nombre: "Yvo", Apellido: "Pintos", Email: "yvo", Edad: 30, Altura: 3, Activo: true, FechaCreacion: "1992"}, {Id: 2, Nombre: "Mat", Apellido: "Fant", Email: "mat", Edad: 33, Altura: 3, Activo: true, FechaCreacion: "1990"}}
	for i := 0; i < len(sliceUsers); i++ {
		if sliceUsers[i].Id == id {
			return sliceUsers[i], nil
		}
	}
	return domain.Usuarios{}, errors.New("error al encontrar el user")
}

func (sdb *StubDB) UpdateNameAndLastName(id int, name string, last string) (domain.Usuarios, error) {
	return domain.Usuarios{}, nil
}

func (sdb *StubDB) Update(id int, nombre, apellido, email string, edad int, altura float64, activo bool, fecha string) (domain.Usuarios, error) {
	return domain.Usuarios{}, nil
}

func (sdb *StubDB) Delete(id int) error {
	return nil
}

func (sdb *StubDB) Guardar(id int, nombre string, apellido string, email string, edad int, altura float64, actico bool, fecha string) (domain.Usuarios, error) {
	if sdb.errLast != "" {
		return domain.Usuarios{}, errors.New("error in last")
	}
	if sdb.errSave != "" {
		return domain.Usuarios{}, errors.New("error saving")
	}
	user := domain.Usuarios{
		Id:            id,
		Nombre:        nombre,
		Apellido:      apellido,
		Email:         email,
		Edad:          edad,
		Altura:        altura,
		Activo:        actico,
		FechaCreacion: fecha,
	}
	return user, nil
}
func (sdb *StubDB) LastId() (int, error) {
	if sdb.errLast != "" {
		return 0, errors.New("error in last id")
	}
	sliceUsers := []domain.Usuarios{{Id: 1, Nombre: "Yvo", Apellido: "Pintos", Email: "yvo", Edad: 30, Altura: 3, Activo: true, FechaCreacion: "1992"}, {Id: 2, Nombre: "Mat", Apellido: "Fant", Email: "mat", Edad: 33, Altura: 3, Activo: true, FechaCreacion: "1990"}}
	lastId := len(sliceUsers)
	return lastId, nil
}

func TestGetAll(t *testing.T) {
	myStubDB := StubDB{}
	servi := NewService(&myStubDB) //Probando el service, yo le paso un mock de repo a lo q quiero probar

	sliceUsers := []domain.Usuarios{{Id: 1, Nombre: "Yvo", Apellido: "Pintos", Email: "yvo", Edad: 30, Altura: 3, Activo: true, FechaCreacion: "1992"}, {Id: 2, Nombre: "Mat", Apellido: "Fant", Email: "mat", Edad: 33, Altura: 3, Activo: true, FechaCreacion: "1990"}}

	resultado, err := servi.GetAll()

	assert.Equal(t, sliceUsers, resultado)
	assert.Nil(t, err)
}

func TestGetAllErrRead(t *testing.T) {
	myStubDB := StubDB{
		errRead: "fail to getAll",
	}
	servi := NewService(&myStubDB)
	_, err := servi.GetAll()

	assert.EqualError(t, err, myStubDB.errRead)
}

func TestGuardar(t *testing.T) {
	myStubDB := StubDB{}
	servi := NewService(&myStubDB)
	result, err := servi.Guardar("new", "new", "new", 20, 20, true, "2020")

	expected := domain.Usuarios{Id: 3, Nombre: "new", Apellido: "new", Email: "new", Edad: 20, Altura: 20, Activo: true, FechaCreacion: "2020"}
	assert.Equal(t, result, expected)
	assert.Nil(t, err)
}

func TestGuardarErrorLastID(t *testing.T) {
	writeLast := fmt.Errorf("error in last id")
	expectedError := fmt.Errorf("error al obtener usuario con id: %w", writeLast)
	myStubDB := StubDB{
		errLast: "error in last id",
	}
	servi := NewService(&myStubDB)
	result, err := servi.Guardar("new", "new", "new", 20, 20, true, "2020")

	assert.Equal(t, err, expectedError)
	assert.Equal(t, domain.Usuarios{}, result)
}
func TestGuardarError(t *testing.T) {
	expectedError := fmt.Errorf("error saving")
	myStubDB := StubDB{
		errSave: "error saving",
	}
	servi := NewService(&myStubDB)
	result, err := servi.Guardar("new", "new", "new", 20, 20, true, "2020")

	assert.Equal(t, err, expectedError)
	assert.Equal(t, domain.Usuarios{}, result)
}

func TestGetById(t *testing.T) {
	myStubDB := StubDB{}

	user := domain.Usuarios{Id: 1, Nombre: "Yvo", Apellido: "Pintos", Email: "yvo", Edad: 30, Altura: 3, Activo: true, FechaCreacion: "1992"}
	servi := NewService(&myStubDB)
	result, err := servi.GetById(1)

	assert.Equal(t, result, user)
	assert.Nil(t, err)
}
func TestGetByIdNotFound(t *testing.T) {
	expectedError := "error al encontrar el user"
	myStubDB := StubDB{}
	servi := NewService(&myStubDB)
	result, err := servi.GetById(12)

	assert.Equal(t, result, domain.Usuarios{})
	assert.EqualError(t, err, expectedError)
}

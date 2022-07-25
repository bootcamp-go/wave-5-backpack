package usuarios

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type StubDB struct{}

func (sdb *StubDB) GetAll() ([]domain.Usuarios, error) {
	sliceUsers := []domain.Usuarios{{Id: 1, Nombre: "Yvo", Apellido: "Pintos", Email: "yvo", Edad: 30, Altura: 3, Activo: true, FechaCreacion: "1992"}, {Id: 2, Nombre: "Mat", Apellido: "Fant", Email: "mat", Edad: 33, Altura: 3, Activo: true, FechaCreacion: "1990"}}
	return sliceUsers, nil
}

func (sdb *StubDB) GetById(id int) (domain.Usuarios, error) {
	return domain.Usuarios{}, nil
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
	return domain.Usuarios{}, nil
}
func (sdb *StubDB) LastId() (int, error) {
	return 0, nil
}

func TestGetAll(t *testing.T) {
	myStubDB := StubDB{}
	servi := NewService(&myStubDB) //Probando el service, yo le paso datos truchos a lo q quiero probar

	sliceUsers := []domain.Usuarios{{Id: 1, Nombre: "Yvo", Apellido: "Pintos", Email: "yvo", Edad: 30, Altura: 3, Activo: true, FechaCreacion: "1992"}, {Id: 2, Nombre: "Mat", Apellido: "Fant", Email: "mat", Edad: 33, Altura: 3, Activo: true, FechaCreacion: "1990"}}

	resultado, err := servi.GetAll()

	assert.Equal(t, sliceUsers, resultado)
	assert.Nil(t, err)
}

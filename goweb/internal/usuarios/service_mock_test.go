package usuarios

import (
	"context"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockDB struct {
	GetWasCalled bool
}

func (mdb *MockDB) GetAll(ctx context.Context) ([]domain.Usuarios, error) {
	sliceUsers := []domain.Usuarios{{Id: 1, Nombre: "Yvo", Apellido: "Pintos", Email: "yvo", Edad: 30, Altura: 3, Activo: true, FechaCreacion: "1992"}, {Id: 2, Nombre: "Mat", Apellido: "Fant", Email: "mat", Edad: 33, Altura: 3, Activo: true, FechaCreacion: "1990"}}
	return sliceUsers, nil
}

func (mdb *MockDB) GetById(id int) (domain.Usuarios, error) {
	return domain.Usuarios{}, nil
}

func (mdb *MockDB) UpdateNameAndLastName(id int, name string, last string) (domain.Usuarios, error) {
	BeforeUpdate, _ := mdb.GetAll(context.TODO())
	mdb.GetWasCalled = true
	var user domain.Usuarios
	for i := 0; i < len(BeforeUpdate)-1; i++ {
		if BeforeUpdate[i].Id == id {
			user = BeforeUpdate[i]
			user.Nombre = name
			user.Apellido = last
		}
	}

	return user, nil
}

func (mdb *MockDB) Update(ctx context.Context, id int, nombre, apellido, email string, edad int, altura float64, activo bool, fecha string) (domain.Usuarios, error) {
	return domain.Usuarios{}, nil
}

func (mdb *MockDB) Delete(id int) error {
	return nil
}

func (mdb *MockDB) Guardar(id int, nombre string, apellido string, email string, edad int, altura float64, actico bool, fecha string) (domain.Usuarios, error) {
	return domain.Usuarios{}, nil
}
func (mdb *MockDB) LastId() (int, error) {
	return 0, nil
}
func (mdb *MockDB) GetByName(name string) ([]domain.Usuarios, error) {
	return nil, nil
}
func (mdb *MockDB) Store(userD domain.Usuarios) (domain.Usuarios, error) {
	return domain.Usuarios{}, nil
}

func TestUpdate(t *testing.T) {
	myMock := MockDB{}
	servi := NewService(&myMock) //Estoy probnando el service, porq uso la implementacion real de service.

	user1 := domain.Usuarios{Id: 1, Nombre: "Perla", Apellido: "Oceano", Email: "yvo", Edad: 30, Altura: 3, Activo: true, FechaCreacion: "1992"}

	resultado, _ := servi.UpdateNameAndLastName(1, "Perla", "Oceano")

	assert.Equal(t, user1, resultado)
	assert.True(t, myMock.GetWasCalled)
}

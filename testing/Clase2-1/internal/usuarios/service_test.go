package usuarios

import (
	"Clase2-1/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubDB struct{}

/*GetAll() ([]domain.Usuario, error)
Store(id int, nombre, apellido, email string, edad, altura int, activo bool) (domain.Usuario, error)
LastID() (int, error)
Update(id int, nombre, apellido, email string, edad, altura int, activo bool) (domain.Usuario, error)
UpdateSurnameAndAge(id int, surname string, age int) (domain.Usuario, error)
Delete(id int) error*/

func (d StubDB) GetAll() ([]domain.Usuario, error) {
	var users []domain.Usuario
	usuario1 := domain.Usuario{
		Id:              1,
		Nombre:          "prueba1",
		Apellido:        "prueba1apellido",
		Email:           "prueba1Email",
		Edad:            25,
		Altura:          180,
		Activo:          true,
		FechaDeCreacion: "29/10/2004",
	}
	usuario2 := domain.Usuario{
		Id:              4,
		Nombre:          "prueba2",
		Apellido:        "prueba2apellido",
		Email:           "prueba2Email",
		Edad:            25,
		Altura:          190,
		Activo:          false,
		FechaDeCreacion: "05/10/2004",
	}
	users = append(users, usuario1)
	users = append(users, usuario2)

	return users, nil
}
func (d StubDB) Store(id int, nombre, apellido, email string, edad, altura int, activo bool) (domain.Usuario, error) {
	return domain.Usuario{}, nil
}
func (d StubDB) LastID() (int, error) {
	return 0, nil
}
func (d StubDB) Update(id int, nombre, apellido, email string, edad, altura int, activo bool) (domain.Usuario, error) {
	return domain.Usuario{}, nil
}
func (d StubDB) UpdateSurnameAndAge(id int, surname string, age int) (domain.Usuario, error) {
	return domain.Usuario{}, nil
}
func (d StubDB) Delete(id int) error {
	return nil
}

func TestGetAll(t *testing.T) {
	//arrange

	myStubDB := StubDB{}
	servicio := NewService(myStubDB)

	usuario1 := domain.Usuario{
		Id:              1,
		Nombre:          "prueba1",
		Apellido:        "prueba1apellido",
		Email:           "prueba1Email",
		Edad:            25,
		Altura:          180,
		Activo:          true,
		FechaDeCreacion: "29/10/2004",
	}
	usuario2 := domain.Usuario{
		Id:              4,
		Nombre:          "prueba2",
		Apellido:        "prueba2apellido",
		Email:           "prueba2Email",
		Edad:            25,
		Altura:          190,
		Activo:          false,
		FechaDeCreacion: "05/10/2004",
	}

	var usuariosEsperados []domain.Usuario
	usuariosEsperados = append(usuariosEsperados, usuario1)
	usuariosEsperados = append(usuariosEsperados, usuario2)

	//act
	resultado, err := servicio.GetAll()

	//assert
	assert.Nil(t, err)
	assert.Equal(t, usuariosEsperados, resultado)
}

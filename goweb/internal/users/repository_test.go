package users

import (
	"github.com/bootcamp-go/wave-5-backpack/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

type StubStore struct{

}

func (ss *StubStore) Read(data interface{}) error {
	userDB := data.(*[]domain.User)
	*userDB = []domain.User{{
			ID:         1,
			Name:       "name1",
			Lastname:   "lastname1",
			Email:      "1@mail.com",
			Age:        22,
			Height:     1.55,
			Active:     true,
			DoCreation: "02-03-2020",
		},{
			ID:         2,
			Name:       "name2",
			Lastname:   "lastname2",
			Email:      "2@mail.com",
			Age:        23,
			Height:     1.56,
			Active:     true,
			DoCreation: "02-03-2021",
		},
	}
		
	return nil
}

func (ss *StubStore) Write(data interface{}) error {
	return nil
}
func (ss *StubStore) Ping() error {
	return nil
}

/* Ejercicio 1 - Test Unitario GetAll()
Generar un Stub del Store cuya función “Read” retorne dos productos con las especificaciones que deseen.
Comprobar que GetAll() retorne la información exactamente igual a la esperada. Para esto:
1. Dentro de la carpeta /internal/products, crear un archivo repository_test.go con el test diseñado. */

func TestGetAllUsers(t *testing.T) {
	stub := StubStore{}

	repo := NewRepository(&stub)
	expectedResponse := []domain.User{
		{
			ID:         1,
			Name:       "name1",
			Lastname:   "lastname1",
			Email:      "1@mail.com",
			Age:        22,
			Height:     1.55,
			Active:     true,
			DoCreation: "02-03-2020",
		}, {
			ID:         2,
			Name:       "name2",
			Lastname:   "lastname2",
			Email:      "2@mail.com",
			Age:        23,
			Height:     1.56,
			Active:     true,
			DoCreation: "02-03-2021",
		},
	}
	users, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, users)
}

/* Ejercicio 2 - Test Unitario UpdateName()
Diseñar Test de UpdateName, donde se valide que la respuesta retornada sea correcta para la actualización del nombre de un producto específico.
Y además se compruebe que efectivamente se usa el método “Read” del Storage para buscar el producto. Para esto:
1. Crear un mock de Storage, dicho mock debe contener en su data un producto específico cuyo nombre puede ser “Before Update”.
2. El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado.
Puede ser a través de un boolean como se observó en la clase.
3. Para dar el test como OK debe validarse que al invocar el método del Repository UpdateName,
con el id del producto mockeado y con el nuevo nombre “After Update”, efectivamente haga la actualización.
También debe validarse que el método Read haya sido ejecutado durante el test. */

type MockStore struct {
	ReadOk bool
}
func (fs *MockStore) Ping() error {
	return nil
}

func (fs *MockStore) Write(data interface{}) error {
	return nil
}

func (fs *MockStore) Read(data interface{}) error {
	InfoUsers := data.(*[]domain.User)
	*InfoUsers = []domain.User{
		{
			ID:         1,
			Name:       "name1",
			Lastname:   "lastname1",
			Email:      "1@mail.com",
			Age:        22,
			Height:     1.55,
			Active:     true,
			DoCreation: "02-03-2020",
		}, {
			ID:         2,
			Name:       "name2",
			Lastname:   "lastname2",
			Email:      "2@mail.com",
			Age:        23,
			Height:     1.56,
			Active:     true,
			DoCreation: "02-03-2021",
		},
	}
	fs.ReadOk = true

	return nil
}


/* Test with Mock */
func TestUpdateNameUser(t *testing.T) {
	id, name := 1, "After Update"
	
	mock := MockStore{}

	r := NewRepository(&mock)
	response, err := r.UpdateUser(id, name, "","",22, 1.88,true,"")
	assert.Nil(t, err)

	assert.True(t, true, mock.ReadOk)
	assert.Equal(t, id, response.ID)
	assert.Equal(t, name, response.Name)
}


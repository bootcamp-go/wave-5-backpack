/* EJERCICIO 2
Diseñar Test de UpdateName, donde se valide que la respuesta retornada sea correcta para la actualización del nombre de un producto/usuario/transacción específico. Y además se compruebe que efectivamente se usa el método “Read” del Storage para buscar el producto. Para esto:
1. Crear un mock de Storage, dicho mock debe contener en su data un producto/usuario/transacción específico cuyo nombre puede ser “Before Update”.
2. El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado. Puede ser a través de un boolean como se observó en la clase.
3. Para dar el test como OK debe validarse que al invocar el método del Repository UpdateName, con el id del producto/usuario/transacción mockeado y con el nuevo nombre “After Update”, efectivamente haga la actualización. También debe validarse que el método Read haya sido ejecutado durante el test.
*/

package users

import (
	"fmt"
	"goweb/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1) creo lo que sería la pkg/store de mi test
type MockStore struct {
	ReadWasCalled bool
	dataMock []domain.User
	errRead string
	errWrite string
}

// así quedaría para el test
func (m *MockStore) Read(data interface{}) error{
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	BeforeUpdate := data.(*[]domain.User)//ACA ESTOY RECIBIENDO DESDE REPOSITORY UN PUNTERO DE LISTA DE USUARIOS
	*BeforeUpdate = m.dataMock
	m.ReadWasCalled = true
	return nil
}

// van los otros métodos
func (m *MockStore) Write(data interface{}) error{
	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}
	a := data.([]domain.User)// acá lo que defino es que la interfaz vacía "data" es de tipo "ista de Productos"
	m.dataMock = append(m.dataMock, a...) // agrego la data recibida a la info mockeada que ya tenía
	return nil
}

func (m *MockStore) Ping() error{
	return nil
}

// ahora va el testeo
func TestGetAllUsers(t *testing.T){
	dataBase := []domain.User{
		{Id: 1, Name: "nombre1", LastName: "apellido1", Email: "mail1@mail.com", Age: 22, Height:1.83, Active: true, CreatedAt: "25/07/2022"},
		{Id: 2, Name: "nombre2", LastName: "apellido2", Email: "mail2@mail.com", Age: 23, Height:1.60, Active: true, CreatedAt: "25/07/2022"},
	}

	mock := MockStore{
		dataMock: dataBase,
	}
	
    repo := NewRepository(&mock)
	resultadoEsperado := dataBase

	a, err := repo.GetAllUsers()

	assert.Nil(t,err)
	assert.Equal(t, resultadoEsperado,a)

}



func TestUpdateTotal(t *testing.T){
	dataBase := []domain.User{
		{Id: 1, Name: "nombre1", LastName: "apellido1", Email: "mail1@mail.com", Age: 22, Height:1.83, Active: true, CreatedAt: "25/07/2022"},
		{Id: 2, Name: "nombre2", LastName: "apellido2", Email: "mail2@mail.com", Age: 23, Height:1.60, Active: true, CreatedAt: "25/07/2022"},
	}

	mock := MockStore{
		dataMock: dataBase,
	}

    repo := NewRepository(&mock)

	afterUpdate, err := repo.UpdateTotal(1, "nuevoNombre", "nuevoApellido1", "mail1@mail.com", 22, 1.83, true, "25/07/2022")

	assert.Equal(t, mock.dataMock[0],afterUpdate)
	assert.True(t, mock.ReadWasCalled)
	assert.Nil(t,err)
	
}


func TestDelete(t *testing.T){
	// arrange
	dataBase := []domain.User{
		{Id: 1, Name: "nombre1", LastName: "apellido1", Email: "mail1@mail.com", Age: 22, Height:1.83, Active: true, CreatedAt: "25/07/2022"},
		{Id: 2, Name: "nombre2", LastName: "apellido2", Email: "mail2@mail.com", Age: 23, Height:1.60, Active: true, CreatedAt: "25/07/2022"},
	}

	mockStorage := MockStore{
		dataMock: dataBase,
		errRead: "",
		errWrite: "",
		ReadWasCalled: false,
	}
	
	// act
    repo := NewRepository(&mockStorage)
	err := repo.Delete(1)

	//assert
	assert.True(t, mockStorage.ReadWasCalled)
	assert.Nil(t, err)
	assert.NotEqual(t, mockStorage.dataMock[0].Id, 1)

}


/* EJERCICIO 2
Diseñar Test de UpdateName, donde se valide que la respuesta retornada sea correcta para la actualización del nombre de un producto/usuario/transacción específico. Y además se compruebe que efectivamente se usa el método “Read” del Storage para buscar el producto. Para esto:
1. Crear un mock de Storage, dicho mock debe contener en su data un producto/usuario/transacción específico cuyo nombre puede ser “Before Update”.
2. El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado. Puede ser a través de un boolean como se observó en la clase. 
3. Para dar el test como OK debe validarse que al invocar el método del Repository UpdateName, con el id del producto/usuario/transacción mockeado y con el nuevo nombre “After Update”, efectivamente haga la actualización. También debe validarse que el método Read haya sido ejecutado durante el test. 
 */

package users

import (
	"goweb/internal/domain"
	"testing"
	"github.com/stretchr/testify/assert"
)

// 1) creo lo que sería la pkg/store de mi test
type MockStore struct {
	ReadWasCalled bool
}

// así quedaría para el test
func (m *MockStore) Read(data interface{}) error{
	
	BeforeUpdate := data.(*[]domain.User)//ACA ESTOY RECIBIENDO DESDE REPOSITORY UN PUNTERO DE LISTA DE USUARIOS
	*BeforeUpdate = []domain.User{//ACA LLENO ESOS VALORES DEL PUNTERO, por eso lo desreferencio
		{Id: 1, Name: "nombre1", LastName: "apellido1", Email: "mail1@mail.com", Age: 22, Height:1.83, Active: true, CreatedAt: "25/07/2022"},
		{Id: 2, Name: "nombre2", LastName: "apellido2", Email: "mail2@mail.com", Age: 23, Height:1.60, Active: true, CreatedAt: "25/07/2022"},
	}

	m.ReadWasCalled = true
	return nil
}

// van los otros métodos
func (m *MockStore) Write(data interface{}) error{
	return nil
}
func (m *MockStore) Ping() error{
	return nil
}

// ahora va el testeo
func TestUpdateTotal(t *testing.T){
	mock := MockStore{}
    repo := NewRepository(&mock) //Probando el repository, yo le paso datos dummy a lo que quiero probar
	resultadoEsperado := domain.User{Id: 1, Name: "nuevoNombre", LastName: "nuevoApellido1", Email: "mail1@mail.com", Age: 22, Height:1.83, Active: true, CreatedAt: "25/07/2022"}

	afterUpdate, err := repo.UpdateTotal(1, "nuevoNombre", "nuevoApellido1", "mail1@mail.com", 22, 1.83, true, "25/07/2022")

	assert.Equal(t, resultadoEsperado,afterUpdate)
	assert.True(t, mock.ReadWasCalled)
	assert.Nil(t,err)
	
}



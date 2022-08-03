package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

func (fs *StubStore) Read(data interface{}) error {

	a := data.(*[]*Product)
	p1 := &Product{
		Id:     1,
		Nombre: "Tomate",
		Stock:  1,
		Precio: 10.0,
	}

	p2 := &Product{
		Id:     1,
		Nombre: "Tomate",
		Stock:  1,
		Precio: 10.0,
	}

	*a = append(*a, p1)
	*a = append(*a, p2)

	return nil
}

func (fs *StubStore) Write(data interface{}) error {
	return nil
}

/*
Ejercicio 1 - Test Unitario GetAll()
	Generar un Stub del Store cuya función “Read” retorne dos productos con las especificaciones que deseen.
	Comprobar que GetAll() retorne la información exactamente igual a la esperada. Para esto:
1. Dentro de la carpeta /internal/products, crear un archivo repository_test.go con el test diseñado.
*/

func TestGetAll(t *testing.T) {
	stub := &StubStore{}
	repo := NewRepository(stub)
	expected := []*Product{
		{
			Id:     1,
			Nombre: "Tomate",
			Stock:  1,
			Precio: 10.0,
		},
		{
			Id:     1,
			Nombre: "Tomate",
			Stock:  1,
			Precio: 10.0,
		},
	}

	a, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, expected, a)
}

// /*
// Ejercicio 2 - Test Unitario UpdateName()
// Diseñar Test de UpdateName, donde se valide que la respuesta retornada sea correcta para la actualización del nombre de un producto específico.
// Y además se compruebe que efectivamente se usa el método “Read” del Storage para buscar el producto. Para esto:
//     1. Crear un mock de Storage, dicho mock debe contener en su data un producto específico cuyo nombre puede ser “Before Update”.
//     2. El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado.
// 	Puede ser a través de un boolean como se observó en la clase.
//     3. Para dar el test como OK debe validarse que al invocar el método del Repository UpdateName,
// 	con el id del producto mockeado y con el nuevo nombre “After Update”, efectivamente haga la actualización.
// 	También debe validarse que el método Read haya sido ejecutado durante el test.
// */

// type MockStore struct {
// 	ReadInvoked bool
// 	Data        []*Product
// }

// func (fs *MockStore) Read(data interface{}) error {
// 	fs.ReadInvoked = true
// 	a := data.(*[]*Product)
// 	*a = fs.Data
// 	return nil
// }

// func (fs *MockStore) Write(data interface{}) error {
// 	return nil
// }

// /* Test with Mock */
// func TestUpdateName(t *testing.T) {
// 	id, nombre := 1, "Update After"
// 	products := []*Product{{Id: 1, Nombre: "Update Before", Stock: 1, Precio: 12}}

// 	mock := MockStore{Data: products}

// 	r := NewRepository(&mock)
// 	productUpdated, err := r.UpdateName(id, nombre)
// 	assert.Nil(t, err)

// 	assert.Equal(t, id, productUpdated.Id)
// 	assert.Equal(t, nombre, productUpdated.Nombre)
// 	assert.True(t, true, mock.ReadInvoked)
// }

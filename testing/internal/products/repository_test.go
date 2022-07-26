package products

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/bootcamp-go/go-testing/pkg/store"
	"github.com/stretchr/testify/assert"
)

/*
Ejercicio 2 - Test Unitario UpdateName()
Diseñar Test de UpdateName, donde se valide que la respuesta retornada sea correcta para la actualización del nombre de un producto específico. Y además se compruebe que efectivamente se usa el método “Read” del Storage para buscar el producto. Para esto:
    1. Crear un mock de Storage, dicho mock debe contener en su data un producto específico cuyo nombre puede ser “Before Update”.
    2. El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado. Puede ser a través de un boolean como se observó en la clase.
    3. Para dar el test como OK debe validarse que al invocar el método del Repository UpdateName, con el id del producto mockeado y con el nuevo nombre “After Update”, efectivamente haga la actualización. También debe validarse que el método Read haya sido ejecutado durante el test.
*/

/* Test with Mock */
func TestUpdateName(t *testing.T) {
	id, nombre := 1, "Update After"
	products := []*Product{{Id: 1, Nombre: "Update Before", Stock: 1, Precio: 12}}

	data, _ := json.Marshal(products)
	mock := store.Mock{Data: data}

	stubStore := store.FileStore{
		FileName: "",
		Mock:     &mock,
	}

	r := NewRepository(&stubStore)
	productUpdated, err := r.UpdateName(id, nombre)
	assert.Nil(t, err)

	assert.True(t, mock.ReadInvoked)
	assert.Equal(t, id, productUpdated.Id)
	assert.Equal(t, nombre, productUpdated.Nombre)
}

/*
Ejercicio 1 - Test Unitario GetAll()
Generar un Stub del Store cuya función “Read” retorne dos productos con las especificaciones que deseen. Comprobar que GetAll() retorne la información exactamente igual a la esperada. Para esto:
    1. Dentro de la carpeta /internal/products, crear un archivo repository_test.go con el test diseñado.
*/

/* Test with Stub */
func TestGetAll(t *testing.T) {
	products := []*Product{
		{
			Id:     1,
			Nombre: "Pepsi",
			Stock:  12,
			Precio: 100,
		},
		{
			Id:     2,
			Nombre: "Coca Cola",
			Stock:  1,
			Precio: 20,
		},
	}

	data, _ := json.Marshal(products)
	dbMock := store.Mock{
		Data:  data,
		Error: nil,
	}

	stubStore := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	r := NewRepository(&stubStore)
	products, err := r.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, products, products)
}

func TestGetAllError(t *testing.T) {
	errorExpected := errors.New("error for GetAll")
	dbMock := store.Mock{ // No le pasamos información
		Error: errorExpected, // Deberia fallar
	}

	stubStore := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	r := NewRepository(&stubStore)
	products, err := r.GetAll()

	assert.Equal(t, errorExpected, err)
	assert.Nil(t, products)
}

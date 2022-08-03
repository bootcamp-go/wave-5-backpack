// package products

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

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
// 	// Si la interfaz 'data' es del tipo slice de punteros de productos -> Asigna en "a" la direccion de memoria del slice de punteros de productos
// 	// Se desreferencia para obtener el tipo de datos slice de punteros de productos porque recibe la direccion de memoria y no el valor
// 	a := data.(*[]*Product)
// 	// asigna al valor de esa direccion el valor de fs.Data
// 	*a = fs.Data
// 	return nil
// }

// func (fs *MockStore) Write(data interface{}) error {
// 	return nil
// }

// func (fs *MockStore) Ping() error {
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

// 	assert.Equal(t, id, productUpdated.ID)
// 	assert.Equal(t, nombre, productUpdated.Name)
// 	assert.True(t, true, mock.ReadInvoked)
// }

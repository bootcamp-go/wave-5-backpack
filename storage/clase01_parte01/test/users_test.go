/* 
Ejercicio 1 - Functional Testing Update() 
Se requiere probar la funcionalidad de “actualización de producto”, pudiendo reutilizar las funciones creadas en la clase. Para lograrlo realizar los siguientes pasos:
Dentro de la carpeta /test, crear un archivo products_test.go.
Levantar el Servidor y definir la ruta para este test.
Crear Request y Response apropiados.
Solicitar al servidor que atienda al Request.
Validar Response.

 Ejercicio 2 - Functional Testing Delete() 
Se solicita probar la funcionalidad de “eliminar producto”, pudiendo reutilizar las funciones creadas en la clase. Para lograrlo realizar los siguientes pasos:
Dentro de la carpeta /test, crear un archivo products_test.go.
Levantar el Servidor y definir la ruta para este test.
Crear Request y Response apropiados.
Solicitar al servidor que atienda al Request.
Validar Response.
 */


package test

import (
	"bytes"
	"goweb/cmd/server/handler"
	"goweb/internal/users"
	"goweb/pkg/store"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

// creo el servidor
func createServer() *gin.Engine {

	_ = os.Setenv("TOKEN", "123456")
	db := store.NewStore("users_test.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	u := handler.NewUser(service)
	r := gin.Default()
	pr := r.Group("/users")
	pr.PUT("/:id", u.UpdateTotal())
	pr.DELETE("/:id", u.Delete())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestUpdateTotal(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodPut, "/users/3",
		`{"name": "Lionel Andres",	"lastname": "Messi","email": "lioGOAT@mail.com","age": 35,"height": 1.7,"active": true,"createdat": "24/06/05"}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestDelete(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodDelete, "/users/3","")

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

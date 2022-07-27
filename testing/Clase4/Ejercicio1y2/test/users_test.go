package test

import (
	"bytes"
	"clase2_2/cmd/server/handler"
	"clase2_2/internal/users"
	"clase2_2/pkg/storage"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	// "encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := storage.NewStore("./usuarios.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	handler := handler.NewUser(service)
	r := gin.Default()

	us := r.Group("users")
	//ejercicio1
	us.PUT("/:id", handler.UpdateUser())
	us.DELETE("/:id", handler.Delete())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

//Ejercicio1
func TestUpdateUser(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer()
	// crear Request del tipo GET y Response para obtener el resultado
	rJson := `{ "name": "Luis","last_name": "tapia","mail": "mail@mail", "years": 25,"tall": 1.68,"enable": true,"create_date": "18/07/1997"}`
	req, rr := createRequestTest(http.MethodPut, "/users/1", rJson)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

//Ejercicio2
func TestDeleteUser(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer()
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodDelete, "/users/1", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

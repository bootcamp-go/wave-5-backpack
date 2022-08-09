package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/anesquivel/wave-5-backpack/storage/arquitectura_ejercicio/cmd/server/handler"
	"github.com/anesquivel/wave-5-backpack/storage/arquitectura_ejercicio/db"
	"github.com/anesquivel/wave-5-backpack/storage/arquitectura_ejercicio/internal/usuarios"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	//db := store.NewStore("usuarios.json")
	db := db.StorageDB
	repo := usuarios.NewRepository(db)
	service := usuarios.NewService(repo)
	u := handler.NewUsuario(service)
	r := gin.Default()

	usrsGroup := r.Group("/usuarios")
	usrsGroup.GET("/", u.GetAll())
	usrsGroup.POST("/", u.Store())
	usrsGroup.PUT("/:id", u.Update())
	usrsGroup.DELETE("/:id", u.Delete())

	return r

}

func createRequestTest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func Test_GetAllUsers_OK(t *testing.T) {
	// crear el server y definir las rutas
	r := createServer()

	// Crear requerst del tipo GET y Response para obtener el result

	req, res := createRequestTest(http.MethodGet, "/usuarios/", "")

	// indicar al server que responda la solicitud
	var objRes []handler.RequestUser

	r.ServeHTTP(res, req)

	// assert
	assert.Equal(t, 200, res.Code)
	err := json.Unmarshal(res.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.True(t, len(objRes) > 0)
}

func Test_StoreUser_OK(t *testing.T) {
	// crear el server y definir las rutas
	r := createServer()

	// Crear requerst del tipo GET y Response para obtener el result

	nwUsr := `{
		"nombre": "testing",
		"apellido": "funcional",
		"edad": 32,
		"email": "testfun@gsjk.com",
		"altura": 1.53
	}`
	req, res := createRequestTest(http.MethodPost, "/usuarios/", nwUsr)

	// indicar al server que responda la solicitud
	r.ServeHTTP(res, req)

	// assert
	var objRes handler.RequestUser

	assert.Equal(t, 200, res.Code)
	err := json.Unmarshal(res.Body.Bytes(), &objRes)
	assert.Nil(t, err)
}

func Test_UpdateUser_OK(t *testing.T) {
	// crear el server y definir las rutas
	r := createServer()

	// Crear requerst del tipo GET y Response para obtener el result

	nwUsr := `{
        "id": 15,
        "nombre": "test",
        "apellido": "TESTING",
        "edad": 28,
        "fechaCreacion": "27 JUL 2022",
        "altura": 0,
        "email": "fdsfds@gmail.com",
        "activo": true
    }`
	req, res := createRequestTest(http.MethodPut, "/usuarios/15", nwUsr)

	// indicar al server que responda la solicitud
	r.ServeHTTP(res, req)

	// assert
	var objRes handler.RequestUser

	assert.Equal(t, 200, res.Code)
	err := json.Unmarshal(res.Body.Bytes(), &objRes)
	assert.Nil(t, err)
}

func Test_DeleteUser_OK(t *testing.T) {
	// crear el server y definir las rutas
	r := createServer()

	// Crear requerst del tipo GET y Response para obtener el result
	req, res := createRequestTest(http.MethodDelete, "/usuarios/15", "")

	// indicar al server que responda la solicitud
	r.ServeHTTP(res, req)

	// assert
	var objRes handler.RequestUser

	assert.Equal(t, 200, res.Code)
	err := json.Unmarshal(res.Body.Bytes(), &objRes)
	assert.Nil(t, err)
}

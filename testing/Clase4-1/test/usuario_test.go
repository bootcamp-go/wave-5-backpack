package test

import (
	"Clase4-1/cmd/server/handler"
	"Clase4-1/internal/usuarios"
	"Clase4-1/pkg/store"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(pathDB string) *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")

	db := store.NewStore(pathDB)
	repo := usuarios.NewRepository(db)
	service := usuarios.NewService(repo)
	user := handler.NewUser(service)

	r := gin.Default()

	pr := r.Group("/usuarios")
	pr.POST("/", user.Store())
	pr.GET("/", user.GetAll())
	pr.PATCH("/:id", user.Update())
	pr.DELETE("/:id", user.Delete())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

/*
func TestGetAllUsers(t *testing.T) {

	type usuario struct {
		Nombre          string `json:"nombre"`
		Apellido        string `json:"apellido"`
		Email           string `json:"email"`
		Edad            int    `json:"edad"`
		Altura          int    `json:"altura"`
		Activo          bool   `json:"activo"`
		FechaDeCreacion string `json:"fecha_de_creacion"`
	}

	// crear el Server y definir las Rutas
	r := createServer("usuarios.json")
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/usuarios/", "")

	var objRes []usuario

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	assert.True(t, len(objRes) > 0)
}
*/
func TestSaveUser(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer("usuarios.json")
	// crear Request del tipo POST y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPost, "/usuarios/", `{
        "nombre": "prueba8",
    "apellido": "BeforeUpdate", 
    "email": "prueba1Email", 
    "edad": 25, 
    "Altura": 180, 
    "activo": true, 
    "fecha_de_creacion": "29/10/2004"
    }`)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestDeleteUser(t *testing.T) {
	src := "usuarios.json"
	dest := "users_delete.json"
	bytesRead, err := ioutil.ReadFile(src)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(dest, bytesRead, 0644)

	if err != nil {
		log.Fatal(err)
	}

	// crear el Server y definir las Rutas
	r := createServer(dest)
	// crear Request del tipo POST y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodDelete, "/usuarios/1", "")

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestUpdateUser(t *testing.T) {
	r := createServer("users_update.json")

	req, rr := createRequestTest(http.MethodPatch, "/usuarios/3",
		`{"nombre": "Jose",
		"apellido": "AfterUpdate", 
		"email": "prueba1Email", 
		"edad": 25, 
		"Altura": 180, 
		"activo": true, 
		"fecha_de_creacion": "29/10/2004"}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

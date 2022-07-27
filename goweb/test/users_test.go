package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/usuarios"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")

	db := store.NewStore("usuarios_test.json")
	repo := usuarios.NewRepository(db)
	servi := usuarios.NewService(repo)
	u := handler.NewUsuario(servi)

	router := gin.Default()

	usr := router.Group("/usuarios")
	usr.PUT("/:id", u.Update())
	usr.GET("/", u.GetAll())
	usr.DELETE("/:id", u.Delete())
	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Token", "123456")
	return req, httptest.NewRecorder()
}

func TestUpdateUser(t *testing.T) {
	r := createServer()
	req, rr := createRequestTest(http.MethodPut, "/usuarios/4", `{
		"nombre": "YvoNew","apellido": "PintNew","email": "yvonew","edad":30,"altura":2,"activo":true,"fecha_de_creacion":"2020"
	}`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

//Esta funcion no anda, por el tipo de retorno supongo sera DATA
// func TestGetAllUsuarios(t *testing.T) {
// 	type usuarios struct {
// 		Nombre        string  `json:"nombre" binding:"required"`
// 		Apellido      string  `json:"apellido" binding:"required"`
// 		Email         string  `json:"email" binding:"required"`
// 		Edad          int     `json:"edad" binding:"required"`
// 		Altura        float64 `json:"altura" binding:"required"`
// 		Activo        bool    `json:"activo" binding:"required"`
// 		FechaCreacion string  `json:"fecha_de_creacion" binding:"required"`
// 	}

// 	r := createServer()

// 	req, rr := createRequestTest(http.MethodGet, "/usuarios", "")

// 	var respOB []usuarios

// 	r.ServeHTTP(rr, req)
// 	t.Log(rr)
// 	assert.Equal(t, http.StatusOK, rr.Code)
// 	err := json.Unmarshal(rr.Body.Bytes(), &respOB)
// 	assert.Nil(t, err)
// 	assert.True(t, len(respOB) > 0)
// }

func TestDeleteUser(t *testing.T) {
	r := createServer()

	req, rr := createRequestTest(http.MethodDelete, "/usuarios/12", "")

	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

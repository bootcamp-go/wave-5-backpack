package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/del_rio/web-server/cmd/server/controlador"
	"github.com/del_rio/web-server/internal/domain"
	"github.com/del_rio/web-server/internal/usuarios"
	"github.com/del_rio/web-server/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	fileStore := store.NewStore("usuarios.json")
	repo := usuarios.NewRepository(fileStore)
	service := usuarios.NewService(repo)
	controlador := controlador.NewControlador(service)
	router := gin.Default()

	pr := router.Group("/usuarios")
	pr.PATCH("/:id", controlador.ActualizarAtribUsuario())
	pr.DELETE("/:id", controlador.BorrarUsuario())
	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}
func TestPatchUsuario(t *testing.T) {
	router := createServer()
	req, responseRecorder := createRequestTest(http.MethodPatch, "/usuarios/2", `{
		"nombre": "test.9876",
		"apellido": "del guerrero",
		"email": "dg@gmail.com",
		"edad": 13,
		"altura": 170,
		"activo": false
	}`)
	var objRes domain.Usuario

	router.ServeHTTP(responseRecorder, req)
	assert.Equal(t, 200, responseRecorder.Code)
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &objRes)
	assert.Nil(t, err)
}
func TestDelete(t *testing.T) {
	router := createServer()
	req, responseRecorder := createRequestTest(http.MethodDelete, "/usuarios/1", "")

	router.ServeHTTP(responseRecorder, req)
	assert.Equal(t, 200, responseRecorder.Code)
}

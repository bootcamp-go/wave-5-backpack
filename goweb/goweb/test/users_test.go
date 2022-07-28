package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/goweb/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/users"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.NewStore("./usuarios.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	u := handler.NewUser(service)
	r := gin.Default()

	pr := r.Group("/users")
	pr.PUT("/:id", u.Update())
	pr.DELETE("/:id", u.Delete())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func Test_Update_OK(t *testing.T) {
	type usuario struct {
		Id            int       `json:"id"`
		Nombre        string    `json:"nombre"`
		Apellido      string    `json:"apellido"`
		Email         string    `json:"email"`
		Edad          int       `json:"edad"`
		Altura        float64   `json:"altura"`
		Activo        bool      `json:"activo"`
		FechaCreacion time.Time `json:"fecha_creacion"`
		Borrado       bool      `json:"borrado"`
	}

	type response struct {
		Code  string      `json:"code"`
		Data  interface{} `json:"data,omitempty"`
		Error interface{} `json:"error,omitempty"`
	}

	// Crear el Server y definir las Rutas
	r := createServer()

	// Crear Request del tipo PUT y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPut, "/users/2", `{
        "nombre": "Arleth","apellido": "Alvarez","email": "arleth.alvarez@gmail.com","edad": 29,"altura": 1.55,"activo": true
    }`)

	var objRes response

	// Indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	if err != nil {
		objUsr := objRes.Data.(usuario)
		assert.Equal(t, "Arleth", objUsr.Nombre)
	}
}

func Test_Delete_OK(t *testing.T) {
	type response struct {
		Code  string      `json:"code"`
		Data  interface{} `json:"data,omitempty"`
		Error interface{} `json:"error,omitempty"`
	}

	// Crear el Server y definir las Rutas
	r := createServer()

	// Crear Request del tipo PUT y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodDelete, "/users/5", "")

	var objRes response

	// Indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)

	if err != nil {
		dataResponse := objRes.Data.(string)
		assert.Equal(t, "el usuario 5 ha sido eliminado", dataResponse)
	}
}

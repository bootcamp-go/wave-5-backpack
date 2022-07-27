package test

import (
	"C3ejercicioTT/cmd/server/handler"
	users "C3ejercicioTT/internal/users"
	"C3ejercicioTT/pkg/store"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(pathDB string) *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")

	db := store.New(store.FileType, pathDB)
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	p := handler.NewUser(service)

	r := gin.Default()

	pr := r.Group("usuarios")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PATCH("/:id", p.UpdateLastAge())
	pr.DELETE("/:id", p.Delete())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestGetAllUsers(t *testing.T) {

	type usuario struct {
		Id       int       `json:"id"`
		Nombre   string    `json:"nombre"`
		Apellido string    `json:"apellido"`
		Email    string    `json:"email"`
		Edad     int       `json:"edad"`
		Altura   float64   `json:"altura"`
		Activo   bool      `json:"activo"`
		Fecha    time.Time `json:"fecha"`
	}

	// crear el Server y definir las Rutas
	r := createServer("users.json")
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

func TestSaveUser(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer("users.json")
	// crear Request del tipo POST y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPost, "/usuarios/", `{
        "nombre": "Angela","email": "angela@gmail.es","edad": 34,"altura": 1.60,
    }`)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestUpdateLatAgeUser(t *testing.T) {
	r := createServer("user_update_las_age.json")

	req, rr := createRequestTest(http.MethodPatch, "/usuarios/6",
		`{"apellido": "Pinzón","edad":31}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestDeleteUser(t *testing.T) {
	r := createServer()
	req, rr := createRequestTest(http.MethodDelete, "usuarios/6", "")
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rrCode)
}

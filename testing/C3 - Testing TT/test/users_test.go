package test

import (
	"C3ejercicioTT/cmd/server/handler"
	"C3ejercicioTT/internal/domain"
	users "C3ejercicioTT/internal/users"
	"C3ejercicioTT/pkg/store"
	"C3ejercicioTT/pkg/web"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(pathDB string) *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")

	/* if _, err := os.ReadFile(pathDB); err != nil {
		panic("no me lee el archivo!! ")
	} */
	db := store.New(store.FileType, pathDB)
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	p := handler.NewUser(service)

	gin.SetMode(gin.ReleaseMode) // optional.
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

/*
	Esta roto tu archivo users.json por eso no leia el la data []domains.User y db.Read siempre retornaba el error.
*/
func TestGetAllUsers(t *testing.T) {

	/* type usuario struct { // no es necesario si utilizamos la struct web.Response
		Id       int       `json:"id"`
		Nombre   string    `json:"nombre"`
		Apellido string    `json:"apellido"`
		Email    string    `json:"email"`
		Edad     int       `json:"edad"`
		Altura   float64   `json:"altura"`
		Activo   bool      `json:"activo"`
		Fecha    time.Time `json:"fecha"`
	} */

	// crear el Server y definir las Rutas
	r := createServer("users.json")
	// crear Request del tipo GET y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodGet, "/usuarios/", "")

	//var objRes []usuario
	var objRes web.Response // Usamos la misma struct de respuesta que cuando realizamos un request normal

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	assert.Nil(t, err)
	//assert.True(t, len(objRes) > 0)
	assert.True(t, reflect.ValueOf(objRes.Data).Len() > 0)
}

// Luz este te devolvia 400 porque no se le pasaba el apellido que es requerido por lo que siempre retornaba 400 bad request
func TestSaveUser(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer("users.json")

	user := domain.Usuarios{
		Nombre:   "Angela",
		Apellido: "algun apellido",
		Email:    "angela@gmail.es",
		Edad:     34,
		Altura:   1.60,
		Activo:   false,
		Fecha:    time.Now(),
	}

	dataJson, err := json.Marshal(user)
	assert.Nil(t, err)

	// crear Request del tipo POST y Response para obtener el resultado
	/* req, rr := createRequestTest(http.MethodPost, "/usuarios/", `{
	    "nombre": "Angela","email": "angela@gmail.es","edad": 34,"altura": 1.60,
	}`) */
	body := string(dataJson)
	req, rr := createRequestTest(http.MethodPost, "/usuarios/", body)
	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

func TestUpdateLatAgeUser(t *testing.T) {
	r := createServer("user_update_las_age.json") // este archivo estaba roto en el .json
	req, rr := createRequestTest(http.MethodPatch, "/usuarios/6",
		`{"apellido": "Pizón","edad":25}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

// Cuando eliminas un recurso se rompe el .json
func TestDeleteUser(t *testing.T) {
	r := createServer("users.json")
	req, rr := createRequestTest(http.MethodDelete, "/usuarios/11", "")

	r.ServeHTTP(rr, req)
	t.Log(rr)
	assert.Equal(t, 200, rr.Code)
}

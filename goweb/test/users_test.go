package test

import (
	"bytes"
	"encoding/json"
	"goweb/cmd/server/handler"
	"goweb/internal/users"
	"goweb/pkg/dataStore"
	"goweb/pkg/web"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(pathDB string) *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")

	db := dataStore.NewStore(dataStore.FileType, pathDB)
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	u := handler.NewUsers(service)

	r := gin.Default()

	us := r.Group("/users")
	//pr.POST("/", p.NewUser())
	us.GET("/", u.GetAll())
	us.PUT("/:id", u.Update())
	us.DELETE("/:id", u.Delete())
	//pr.PATCH("/:id", p.UpdateName())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestGetAllUsers(t *testing.T) {

	// crear el Server y definir las Rutas
	r := createServer("users.json")
	// crear Request del tipo GET y Response para obtener el resultado
	req, resp := createRequestTest(http.MethodGet, "/users/", "")

	var userList web.Response

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	//log.Fatal(resp.Body.Bytes())
	err := json.Unmarshal(resp.Body.Bytes(), &userList)
	assert.Nil(t, err)
	assert.True(t, reflect.ValueOf(userList.Data).Len() > 0)
}

func TestUpdateUser(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer("users.json")
	// crear Request del tipo GET y Response para obtener el resultado
	req, resp := createRequestTest(http.MethodPut, "/users/1", `{
        															"name": "Alexis",
        															"lastname": "Esteban",
        															"email": "testUpdate@gmail.com",
        															"age": 30,
        															"height": 1.70,
        															"active": false,
        															"creation-date": "10-10-2010"
    															}`)

	r.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

func TestDeleteUser(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer("users.json")
	// crear Request del tipo GET y Response para obtener el resultado
	req, resp := createRequestTest(http.MethodDelete, "/users/4", "")

	r.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/internal/users"
	"github.com/bootcamp-go/wave-5-backpack/pkg/store"
	"github.com/bootcamp-go/wave-5-backpack/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(pathDB string) *gin.Engine {
	_ = os.Setenv("TOKEN", "lalala")

	db := store.NewStore(pathDB)
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	p := handler.NewUser(service)

	r := gin.Default()

	us := r.Group("/users")
	us.POST("/", p.StoreUser())
	us.GET("/", p.GetAll())
	us.PUT("/:id", p.UpdateUser())
	us.DELETE("/:id", p.DeleteUser())
	us.PATCH("/:id", p.UpdateLastnameAndAge())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "lalala")

	return req, httptest.NewRecorder()
}



func TestGetAllUsuarios(t *testing.T) {

	r := createServer("users.json")

	req, rr := createRequestTest(http.MethodGet, "/users/", "")

	var resp web.Response

	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.True(t, reflect.ValueOf(resp.Data).Len() > 0)
	//valor := reflect.ValueOf(resp.Data)
	//fmt.Println(valor)
}


func TestDeleteUser(t *testing.T) {
	r := createServer("users.json")
	req, rr := createRequestTest(http.MethodDelete, "/users/3", "")
	fmt.Println(rr)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

func TestUpdateUser(t *testing.T) {
	r := createServer("users.json")
	req, rr := createRequestTest(http.MethodPut, "/users/4", `{"name":"eimi", "lastname": "gal", "email":"a@mail","age":22,"height":1.22,"active":true,"doCreation":"9jun89"}`)
	r.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

func TestSaveUser(t *testing.T) {
	// crear el Server y definir las Rutas
	r := createServer("users.json")
	// crear Request del tipo POST y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPost, "/users/", `{
        "name":"eimita", "lastname": "new", "email":"new@mail","age":22,"height":1.22,"active":true,"doCreation":"9jun89"
    }`)

	// indicar al servidor que pueda atender la solicitud
	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestUpdateLastnameAge(t *testing.T) {
	r := createServer("users.json")

	req, rr := createRequestTest(http.MethodPatch, "/users/1",
		`{"lastname": "cambiado", "age":44}`)

	r.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

package test

import (
	"C4-Testing/cmd/handler"
	"C4-Testing/internal/repositorio"
	"C4-Testing/internal/servicio"
	"C4-Testing/pkg/store"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(pathDB string) *gin.Engine {
	err := os.Setenv("TOKEN", "123GO")
	if err != nil {
		fmt.Println(err)
	}
	db := store.NewStore(".json", pathDB)
	fmt.Println("my_debug", db)
	repo := repositorio.NewRepository(db)
	fmt.Println("repo", repo)

	service := servicio.NewService(repo)
	fmt.Println("service", service)

	p := handler.NewUser(service)
	fmt.Println("handler", p)

	r := gin.Default()

	pr := r.Group("/users")
	pr.POST("/", p.CreateUser())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.UpdateUser())
	pr.DELETE("/:id", p.Delete())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123GO")

	return req, httptest.NewRecorder()
}

func TestCreateUser(t *testing.T) {
	//Crear el server y definir las rutas
	r := createServer("users.json")
	//Crear Request del tipo POST y Response para obtener el resultado
	req, rr := createRequestTest(http.MethodPost, "/users/", `{
		"firstName": "RokoTest",
		"lastName": "Testing",
		"email": "roko@dogmail.com",
		"age": 1,
		"height": 1.7,
		"activo": true,
		"createdAt": "07/07/22"
	}`)

	//indicar al servidor que atienda la solicitud
	r.ServeHTTP(rr, req)
	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestUpdateUser(t *testing.T) {
	//Crear el server y definimos las rutas
	r := createServer("users.json")
	//Crear Request del PUT y Response para obtener el Resultado
	req, rr := createRequestTest(http.MethodPut, "/users/7", `{
		"id": 6,
		"firstName": "RokoUpdateTest",
		"lastName": "Testing",
		"email": "roko@dogmail.com",
		"age": 1,
		"height": 1.7,
		"activo": False,
		"createdAt": "07/08/22"
		}`)

	//El server atiende la solicitud
	r.ServeHTTP(rr, req)
	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetAll(t *testing.T) {
	type User struct {
		Id        int     `json:"id"`
		FirstName string  `json:"firstName" binding:"required"`
		LastName  string  `json:"lastName" binding:"required"`
		Email     string  `json:"email" binding:"required"`
		Age       int     `json:"age" binding:"required"`
		Height    float64 `json:"height" binding:"required"`
		Activo    bool    `json:"activo" binding:"required"`
		CreatedAt string  `json:"createdAt" binding:"required"`
	}

	// Crear el server y definir las rutas
	r := createServer("users.json")

	//Crear Request del POST y Response para obtener el Resultado
	req, rr := createRequestTest(http.MethodGet, "/users/", "")

	//var objRes []User

	//El server atiende la solicitud
	r.ServeHTTP(rr, req)
	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestDeleteUser(t *testing.T) {
	//Crear el server y definimos las rutas
	r := createServer("users.json")
	//Crear Request del PUT y Response para obtener el Resultado
	req, rr := createRequestTest(http.MethodDelete, "/users/8", "")

	//El server atiende la solicitud
	r.ServeHTTP(rr, req)
	resp := rr.Result()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	assert.Equal(t, http.StatusOK, rr.Code)
}

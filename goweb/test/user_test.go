package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/cmd/server/middleware"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/users"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/pkg/store"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func createServer(db store.Store) *gin.Engine {
	_ = godotenv.Load()
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	u := handler.NewUser(service)

	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.GET("/", u.GetAll)
		userRouter.GET("/:Id", u.GetById)

		userRouter.Use(middleware.Authorization)
		userRouter.POST("/", u.Store)
		userRouter.PUT("/:Id", u.Update)
		userRouter.PATCH("/:Id", u.UpdateAgeLastName)
		userRouter.DELETE("/:Id", u.Delete)
	}

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", os.Getenv("TOKEN"))

	return req, httptest.NewRecorder()
}

func getTestDB() *store.MockStorage {
	return &store.MockStorage{
		DataMock: []domain.User{
			{
				Id:        9,
				Age:       19,
				FirstName: "Agustin",
				LastName:  "Flores",
				Email:     "Agustin.Flores@lalala.com",
				CreatedAt: "08/05/1999",
				Height:    1.67,
				Active:    true,
			},
		},
	}
}

func Test_UpdateUser_Ok(t *testing.T) {
	db := getTestDB()
	r := createServer(db)

	updatedUser := domain.User{
		Id:        9,
		Age:       20,
		FirstName: "Agustina",
		LastName:  "Flores",
		Email:     "Agustin.Flores@lalala.com",
		CreatedAt: "08/05/1999",
		Height:    1.67,
		Active:    true,
	}
	updatedUserBytes, _ := json.Marshal(updatedUser)
	objRes := web.Response{}

	// Test Con ID existente
	req, rr := createRequestTest(http.MethodPut, "/users/9",
		string(updatedUserBytes),
	)
	r.ServeHTTP(rr, req)

	{
		// Test codigo 200 de respuesta
		assert.Equal(t, 200, rr.Code)

		// Test cuerpo de respuesta válido
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)

		// Test respuesta con usuario actualizado
		user := domain.User{}
		jsonUser, _ := json.Marshal(objRes.Data)
		_ = json.Unmarshal(jsonUser, &user)
		assert.Equal(t, updatedUser, user)
	}

	// Test Con ID inexistente
	req, rr = createRequestTest(http.MethodPut, "/users/10",
		string(updatedUserBytes),
	)
	r.ServeHTTP(rr, req)

	{
		// Test codigo 404 de respuesta
		assert.Equal(t, 404, rr.Code)

		// Test cuerpo de respuesta válido
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)

		// Test mensaje de error válido
		assert.Equal(t, objRes.Error, "usuario no encontrado")
	}
}

func Test_DeleteUser_Ok(t *testing.T) {
	db := getTestDB()
	r := createServer(db)

	objRes := web.Response{}

	// Test Con ID existente
	req, rr := createRequestTest(http.MethodDelete, "/users/9", "")
	r.ServeHTTP(rr, req)

	{
		// Test codigo 200 de respuesta
		assert.Equal(t, 204, rr.Code)

		// Test cuerpo de respuesta vacio
		assert.Equal(t, len(rr.Body.Bytes()), 0)
	}

	// Test Con ID inexistente (borrado en test anterior)
	req, rr = createRequestTest(http.MethodDelete, "/users/9", "")
	r.ServeHTTP(rr, req)

	{
		// Test codigo 404 de respuesta
		assert.Equal(t, 404, rr.Code)

		// Test cuerpo de respuesta válido
		err := json.Unmarshal(rr.Body.Bytes(), &objRes)
		assert.Nil(t, err)

		// Test mensaje de error válido
		assert.Equal(t, objRes.Error, "usuario no encontrado")
	}
}

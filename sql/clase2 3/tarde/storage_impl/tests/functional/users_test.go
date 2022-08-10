package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bootcamp-go/storage/cmd/server/handler"
	cnn "github.com/bootcamp-go/storage/db"
	"github.com/bootcamp-go/storage/internal/domains"
	"github.com/bootcamp-go/storage/internal/users"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var serverUser = createServerUsers()

func createServerUsers() *gin.Engine {
	err := godotenv.Load("./../../.env")
	if err != nil {
		panic("can't connect to database")
	}

	db := cnn.InitDynamo()
	repo := users.NewDynamoRepository(db, "Users")
	serv := users.NewService(repo)

	p := handler.NewUser(serv)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	pr := r.Group("/api/v1/users")
	pr.GET("/", p.GetOne())
	pr.POST("/", p.Store())

	return r
}

func TestStoreUser_Ok(t *testing.T) {
	new := domains.User{
		Firstname:  "Digital",
		Lastname:   "House",
		Username:   "DH",
		Password:   "ADF23LKAS%3434",
		MacAddress: "ER:23:D3:SD:SD",
		Email:      "dh@digitalhouse.com",
		Website:    "digitalhouse.com",
		IP:         "127.0.0.1",
		Image:      "test.png",
	}

	user, err := json.Marshal(new)
	require.Nil(t, err)

	req, rr := createRequest(http.MethodPost, "/api/v1/users/", string(user))
	serverUser.ServeHTTP(rr, req)

	// assert code
	assert.Equal(t, http.StatusCreated, rr.Code)

	// struct for assertion
	u := struct{ Data domains.User }{}
	err = json.Unmarshal(rr.Body.Bytes(), &u)
	require.Nil(t, err)

	new.Id = u.Data.Id
	assert.Equal(t, new, u.Data)
}

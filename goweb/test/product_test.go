package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goweb/cmd/server/handler"
	"goweb/internal/domain"
	"goweb/internal/products"
	"goweb/pkg/store"
	"goweb/pkg/web"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
    _ = os.Setenv("TOKEN", "123456")
    db := store.NewStore("./../productos.json")//store.New(store.FileType, "./products.json")
    repo := products.NewRepository(db)
    service := products.NewService(repo)
    p := handler.NewProduct(service)
    r := gin.Default()

    pr := r.Group("/products")
    pr.POST("/", p.Create())
    pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())
    return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
    req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("token", "123456")

    return req, httptest.NewRecorder()
}

func TestGetProductOK(t *testing.T) {
    // crear el Server y definir las Rutas
    r := createServer()
    // crear Request del tipo GET y Response para obtener el resultado
    req, rr := createRequestTest(http.MethodGet, "/products/", "")

    // indicar al servidor que pueda atender la solicitud
    r.ServeHTTP(rr, req)

    assert.Equal(t, 200, rr.Code)
    objRes := []domain.Product{}
    fmt.Println(rr.Body)
    err := json.Unmarshal(rr.Body.Bytes(), &objRes)
    assert.Nil(t, err)
   	assert.True(t, len(objRes) > 0)
}

func TestSaveProductOK(t *testing.T) {
    // crear el Server y definir las Rutas
    r := createServer()
    // crear Request del tipo POST y Response para obtener el resultado
    req, rr := createRequestTest(http.MethodPost, "/products/", `{
        "name": "Tester","color": "prueba","price": 1020,"stock": 10, "code": "AAA333", "publisher": true
    }`)

    // indicar al servidor que pueda atender la solicitud
    r.ServeHTTP(rr, req)
    

    assert.Equal(t, 201, rr.Code)
    objRes := web.Response{}
    err := json.Unmarshal(rr.Body.Bytes(), &objRes)
    assert.Nil(t, err)
    assert.True(t, reflect.ValueOf(objRes.Data).Len() > 0)
}

func TestUpdateOk(t *testing.T) {
 // crear el Server y definir las Rutas
 r := createServer()
 // crear Request del tipo PUT y Response para obtener el resultado
 req, rr := createRequestTest(http.MethodPut, "/products/16", `{
	 "name": "Tester Update","color": "prueba update","price": 1021,"stock": 11, "code": "BBB333", "publisher": false
 }`)

 // indicar al servidor que pueda atender la solicitud
 r.ServeHTTP(rr, req)

 assert.Equal(t, 200, rr.Code)
 objRes := web.Response{}
 err := json.Unmarshal(rr.Body.Bytes(), &objRes)
 assert.Nil(t, err)
 assert.True(t, reflect.ValueOf(objRes.Data).Len() > 0)

}

func TestDeleteOk(t *testing.T) {
 // crear el Server y definir las Rutas
 r := createServer()
 // crear Request del tipo PUT y Response para obtener el resultado
 req, rr := createRequestTest(http.MethodDelete, "/products/16", "")

 // indicar al servidor que pueda atender la solicitud
 r.ServeHTTP(rr, req)

 assert.Equal(t, 204, rr.Code)
}


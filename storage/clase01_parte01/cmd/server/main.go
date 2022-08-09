package main

import (
	"goweb/cmd/server/handler"
	"goweb/internal/users"
	"log"
	"github.com/gin-gonic/gin"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// @title Franco Visintini API
// @version 1.0
// @description This API Handle MELI Users. You can get info about the users of the platform, store new users in the database, edit existing users and more!
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	dataSource := "root:rootpass@tcp(localhost:3306)/storage"
    // Open inicia un pool de conexiones. Sólo abrir una vez
    var err error
    StorageDB, err := sql.Open("mysql", dataSource)
    if err != nil {
        log.Fatal(err)
    }
    if err = StorageDB.Ping(); err != nil {
		log.Fatal(err)
	}
    log.Println("database Configured")

	repo := users.NewRepository(StorageDB)
	service := users.NewService(repo)
	u := handler.NewUser(service)

	router := gin.Default()

	userGroup := router.Group("/users")
	userGroup.GET("/", u.GetAllUsers())
	userGroup.GET("/:id", u.GetUserById())
	userGroup.GET("/name/:name", u.GetUserByName())
	userGroup.POST("/", u.StoreUser())
	userGroup.PUT("/:id", u.UpdateTotal())
	userGroup.PATCH("/:id", u.UpdatePartial())
	userGroup.DELETE("/:id", u.Delete())

	if err := router.Run(); err != nil {
		panic(err)
	}

}

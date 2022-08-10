package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/bootcamp-go/wave-5-backpack/goweb/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/goweb/docs"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/usuarios"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/store"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swaggo/swag/example/basic/docs"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Usuarios.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	error := godotenv.Load()
	if error != nil {
		log.Fatal("error al intentar cargar el .env")
	}

	db := store.NewStore("usuarios.json")
	err1 := db.Validate()
	if err1 != nil {
		log.Fatal("error al intentar abrir el json")
	}

	//UTILIZANDO JSON
	// repo := usuarios.NewRepository(db)
	// servi := usuarios.NewService(repo)
	// u := handler.NewUsuario(servi)

	dataSource := "root@tcp(localhost:3306)/storage"
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	//UTILIZANDO BBDD RELACIONAL
	bdRepo := usuarios.NewRepositoryBD(StorageDB)
	servi := usuarios.NewService(bdRepo)
	u := handler.NewUsuario(servi)

	router := gin.Default()

	// router.GET("/usuarios", GetAll)
	// router.GET("/usuarios/filtroNombre", FilterByName)
	// router.GET("/usuarios/filtroApellido", FilterByLastName)
	// router.GET("/usuarios/filtroCorreo", filterByEmail)
	// router.GET("/usuarios/filtroEdad", filterByEdad)

	//SWAGGER
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//

	us := router.Group("/usuarios")
	us.Use(tokenAuthMiddleware())
	//Aca ponemos los corchetes para dar una tabulacion
	{
		///////////////BBDD/////////////////
		us.GET("/name/:nombre", u.GetByName())
		///////////////BBDD/////////////////
		us.PUT("/:id", u.Update())
		us.GET("/:id", u.GetById())
		us.POST("/", u.Guardar())
		us.GET("/", u.GetAll())
		us.DELETE("/:id", u.Delete())
		us.PATCH("/:id", u.UpdateNameAndLastName())
	}
	ele := router.Run()
	if ele != nil {
		panic("Router error")
	}
}

func tokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal("no se encontro el token en la variable de entorno")
	}
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "falta token en cabecera"))
			return
		}
		if token != requiredToken {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "token incorrecto"))
			return
		}
		ctx.Next()
	}
}

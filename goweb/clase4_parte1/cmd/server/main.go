package main
// Se importan e inyectan el repositorio, servicio y handler
import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"goweb/clase4_parte1/cmd/server/handler"
	"goweb/clase4_parte1/internal/users"
	"goweb/clase4_parte1/pkg/store"
	"log"
)

// Se implementa el router para los diferentes endpoints
func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error al cargar archivo .env")
	}

	db := store.NewStore("../../users.json")
	if err := db.Ping(); err != nil {
		log.Fatal("Error al intentar cargar archivo json")
	}

	repo := users.NewRepository(db)
	service := users.NewService(repo)
	u := handler.NewUser(service)

	r := gin.Default()
	ur := r.Group("/users")
	ur.POST("/", u.Store())
	ur.GET("/", u.GetAll())
	ur.PUT("/:id", u.Update())
	ur.PATCH("/:id", u.UpdateLastNameAndAge())
	ur.DELETE("/:id", u.Delete())
	r.Run()
}
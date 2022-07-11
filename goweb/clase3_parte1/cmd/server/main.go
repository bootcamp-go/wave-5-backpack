package main
// Se importan e inyectan el repositorio, servicio y handler
import (
	"github.com/gin-gonic/gin"
	"goweb/clase3_parte1/cmd/server/handler"
	"goweb/clase3_parte1/internal/users"
)

// Se implementa el router para los diferentes endpoints
func main() {
	repo := users.NewRepository()
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
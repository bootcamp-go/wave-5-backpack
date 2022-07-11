package main
// Se debe importar e inyectar el repositorio, servicio y handler
import (
	"github.com/gin-gonic/gin"
	"goweb/clase2_parte2/cmd/server/handler"
	"goweb/clase2_parte2/internal/users"
)

// Se debe implementar el router para los diferentes endpoints
func main() {
	repo := users.NewRepository()
	service := users.NewService(repo)
	u := handler.NewUser(service)

	r := gin.Default()
	ur := r.Group("/users")
	ur.POST("/", u.Store())
	ur.GET("/", u.GetAll())
	r.Run()
}
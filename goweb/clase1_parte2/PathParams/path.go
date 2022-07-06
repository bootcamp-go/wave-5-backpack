package main

import "github.com/gin-gonic/gin"

var empleados = map[string]string{
	"644": "Empleado A",
	"755": "Empleado B",
	"777": "Empleado C",
}

//Este handler se encargará de responder a /.
func PaginaPrincipal(ctx *gin.Context) {
	ctx.String(200, "¡Bienvenido a la Empresa Gophers!")
}

//Este handler verificará si la id que pasa el cliente existe en nuestra base de datos.
func BuscarEmpleado(ctx *gin.Context) {
	empleado, ok := empleados[ctx.Param("id")]
	if ok {
		ctx.String(200, "Información del empleado %s, nombre: %s", ctx.Param("id"), empleado)
	} else {
		ctx.String(404, "Información del empleado ¡No existe!")
	}
}

func main() {
	server := gin.Default()
	server.GET("/", PaginaPrincipal)
	server.GET("/empleados/:id", BuscarEmpleado)
	server.Run(":8080")
}

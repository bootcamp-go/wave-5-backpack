package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type usuarios struct {
	//el guion lo ignora
	Id            int     `json:"-"`
	Nombre        string  `json:"nombre" binding:"required"`
	Apellido      string  `json:"apellido" binding:"required"`
	Email         string  `json:"email" binding:"required"`
	Edad          int     `json:"edad" binding:"required"`
	Altura        float64 `json:"altura" binding:"required"`
	Activo        bool    `json:"activo" binding:"required"`
	FechaCreacion string  `json:"fecha_de_creacion" binding:"required"`
}

var users []usuarios = []usuarios{
	{Id: 1, Nombre: "Yvo", Apellido: "P", Email: "a@a.com", Edad: 21, Altura: 1.82, Activo: true, FechaCreacion: "2022-02-20"},
	{Id: 4, Nombre: "Mati", Apellido: "F", Email: "b@b.com", Edad: 30, Altura: 1.76, Activo: false, FechaCreacion: "2022-02-25"},
	{},
}
var user2 []usuarios = []usuarios{}
var lastID int

func GetAll(ctx *gin.Context) {
	ctx.JSON(200, users)
}

func GetById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	for i := 0; i < len(users); i++ {
		if id == users[i].Id {
			fmt.Println("LO ENCONTRE!")
			ctx.JSON(200, users[i])
			return
		} else {
			ctx.JSON(400, nil)
		}
	}
	fmt.Println("NO LO ENCONTRE!")
}

func FilterByName(ctx *gin.Context) {
	name := ctx.Query("name")
	result := []usuarios{}
	for i := 0; i < len(users); i++ {
		if strings.HasPrefix(users[i].Nombre, name) {
			result = append(result, users[i])
		}
	}
	ctx.JSON(200, result)
}

// func filterByEmail(ctx *gin.Context) {
// 	email := ctx.Query("email")
// 	result := []usuarios{}
// 	for i := 0; i < len(users); i++ {
// 		if strings.HasPrefix(users[i].Email, email) {
// 			result = append(result, users[i])
// 		}
// 	}
// 	ctx.JSON(200, result)
// }

// //Este no estaria andando
// func filterByEdad(ctx *gin.Context) {
// 	//EL .QUERY ARMA LA URL BASURITA
// 	edad := ctx.Query("edad")
// 	result := []usuarios{}
// 	campEdad, _ := strconv.Atoi(edad)
// 	for i := 0; i < len(users); i++ {
// 		if users[i].Edad == campEdad {
// 			result = append(result, users[i])
// 		}
// 	}
// 	ctx.JSON(200, result)
// }

func FilterByLastName(ctx *gin.Context) {
	lastName := ctx.Query("apellido")
	result := []usuarios{}
	for i := 0; i < len(users); i++ {
		if strings.HasPrefix(users[i].Apellido, lastName) {
			result = append(result, users[i])
		}
	}
	ctx.JSON(200, result)
}

func Guardar() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var user usuarios
		token := ctx.GetHeader("token")
		if token != "12345678" {
			ctx.JSON(401, gin.H{
				"error": "no tiene permisos para realizar la peticion solicitada",
			})
			return
		}

		if err := ctx.ShouldBindJSON(&user); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				result := ""
				for i, field := range ve {
					if i != len(ve)-1 {
						result += fmt.Sprintf("el campo %s es requerido y ", field.Tag())
					} else {
						result += fmt.Sprintf("el campo %s es requerido", field.Tag())
					}
				}
				ctx.JSON(404, result)
			}
		}
		lastID++
		user.Id = lastID
		user2 = append(user2, user)
		ctx.JSON(200, user2)

	}
}

func main() {
	// para leer json desde archivo
	/*
		byteArr, err := os.ReadFile("./usuarios.json")
		if err != nil {
			log.Fatal("No se pudo abrir el .json")
		}
		var usuarios []usuarios
		err2 := json.Unmarshal(byteArr, &usuarios)
		if err2 != nil {
			fmt.Println(err2.Error())
			log.Fatal("No representa a usuarios")
		}
	*/
	// dummy
	/*
		router.GET("hola", func(ctx *gin.Context) {
			ctx.JSON(200, "Hola Yvo!")
		})
	*/
	router := gin.Default()
	// router.GET("usuarios/:id", GetById)
	// router.GET("/usuarios", GetAll)
	// router.GET("/usuarios/filtroNombre", FilterByName)
	// router.GET("/usuarios/filtroApellido", FilterByLastName)
	// router.GET("/usuarios/filtroCorreo", filterByEmail)
	// router.GET("/usuarios/filtroEdad", filterByEdad)
	us := router.Group("/usuarios")
	us.POST("/", Guardar())
	router.Run()
}

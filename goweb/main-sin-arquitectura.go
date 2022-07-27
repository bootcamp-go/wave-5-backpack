package main

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id           int     `json:"-"`
	Name         string  `json:"name" binding:"required"`
	LastName     string  `json:"lastname" binding:"required"`
	Email        string  `json:"email" binding:"required"`
	Age          int     `json:"age" binding:"required"`
	Height       float64 `json:"height" binding:"required"`
	Active       bool    `json:"active" binding:"required"`
	CreationDate string  `json:"creation-date" `
}

var lastID int
var usersList []User

//Este handler se encargará de responder a /.
func HomePage(ctx *gin.Context) {
	ctx.String(200, "¡Bienvenido a la Empresa Gophers!")
}

// GET

//EJERCICIO 2 M
func GetName(ctx *gin.Context) {
	name := "Cristian Ladino"
	ctx.JSON(200, gin.H{"message": "Hola " + name})
}

//EJERCICIO 3 M
func GetUserList() ([]User, error) {
	jsonData, err := ioutil.ReadFile("users.json")
	if err != nil {
		return []User{}, err
	}
	var userList []User
	err = json.Unmarshal(jsonData, &userList)
	if err != nil {
		return []User{}, err
	}
	return userList, nil
}

/*func GetAllUsers(ctx *gin.Context) {
	users, err := GetUserList()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{"message": &users})
}*/

// EJERCICIO 1 T
func FilterList(users *[]User, params url.Values) error {
	res := []User{}
	name := params.Get("name")
	lastname := params.Get("lastname")
	email := params.Get("email")
	age := params.Get("age")
	height := params.Get("height")
	active := params.Get("active")
	creationDate := params.Get("creation-date")
	for _, user := range *users {
		valid := true
		if name != "" {
			if user.Name != name {
				valid = false
			}
		}
		if lastname != "" && valid {
			if user.LastName != lastname {
				valid = false
			}
		}
		if email != "" && valid {
			if user.Email != email {
				valid = false
			}
		}
		if age != "" && valid {
			compareage, err := strconv.Atoi(age)
			if err != nil {
				return err
			}
			if user.Age != compareage {
				valid = false
			}
		}
		if height != "" && valid {
			compareage, err := strconv.ParseFloat(height, 64)
			if err != nil {
				return err
			}
			if user.Height != compareage {
				valid = false
			}
		}
		if active != "" && valid {
			compareactive, err := strconv.ParseBool(active)
			if err != nil {
				return err
			}
			if user.Active != compareactive {
				valid = false
			}
		}
		if creationDate != "" && valid {
			if user.CreationDate != creationDate {
				valid = false
			}
		}

		if valid {
			res = append(res, user)
		}
	}
	*users = res
	return nil
}

func GetAll(ctx *gin.Context) {
	users, err := GetUserList()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	params := ctx.Request.URL.Query()
	if len(params) > 0 {
		err = FilterList(&users, params)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
	}
	ctx.JSON(200, users)
}

// EJERCICIO 2 T
func GetUserById(ctx *gin.Context) {
	var usersFilterById []User
	users, err := GetUserList()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	for _, user := range users {
		strId := strconv.Itoa(user.Id)
		if ctx.Param("id") == strId {
			usersFilterById = append(usersFilterById, user)
		}
	}
	if len(usersFilterById) > 0 {
		ctx.JSON(200, gin.H{"message": &usersFilterById})
	} else {
		ctx.String(404, "Información del empleado ¡No existe!")
	}
}

// POST

// EJERCICIO 1, 2, 3 M
func NewEntity() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req User
		// Validar token
		token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		// Si el token fue valido, avanzo
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		lastID++
		req.Id = lastID

		usersList = append(usersList, req)

		c.JSON(200, req)
	}
}

func main() {

	// Crea un router con gin
	router := gin.Default()

	// EJERCICIO 1 T
	//Cada vez que llamamos a GET y le pasamos una ruta, definimos un nuevo endpoint.
	router.GET("/", HomePage)
	router.GET("/hello-world", GetName)
	router.GET("users/filter", GetAll)
	router.GET("/users/:id", GetUserById)
	rt := router.Group("/users")
	rt.POST("/", NewEntity())

	if err := router.Run(); err != nil {
		panic(err)
	}

}

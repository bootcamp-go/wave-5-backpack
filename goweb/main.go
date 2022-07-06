package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	LastName     string  `json:"lastname"`
	Email        string  `json:"email"`
	Age          int     `json:"age"`
	Height       float64 `json:"height"`
	Active       bool    `json:"active"`
	CreationDate string  `json:"creation-date"`
}

//Este handler se encargará de responder a /.
func HomePage(ctx *gin.Context) {
	ctx.String(200, "¡Bienvenido a la Empresa Gophers!")
}

//EJERCICIO 2 M
func GetName(ctx *gin.Context) {
	name := "Cristian Ladino"
	ctx.JSON(200, gin.H{"message": "Hola " + name})
}

//EJERCICIO 3 M
func GetAllUsers(ctx *gin.Context) {
	var users []User
	jsonData, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(jsonData), &users); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(200, gin.H{"message": &users})
}

// EJERCICIO 1 T
func GetUserByName(ctx *gin.Context) {
	var users []User
	var usersFilter []User
	jsonData, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(jsonData), &users); err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		if ctx.Param("name") == user.Name {
			usersFilter = append(usersFilter, user)
		}
	}
	if len(usersFilter) > 0 {
		ctx.JSON(200, gin.H{"message": &usersFilter})
	} else {
		ctx.String(404, "Información del empleado ¡No existe!")
	}
}

func GetUserByLastName(ctx *gin.Context) {
	var users []User
	var usersFilter []User
	jsonData, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(jsonData), &users); err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		if ctx.Param("lastname") == user.LastName {
			usersFilter = append(usersFilter, user)
		}
	}
	if len(usersFilter) > 0 {
		ctx.JSON(200, gin.H{"message": &usersFilter})
	} else {
		ctx.String(404, "Información del empleado ¡No existe!")
	}
}

func GetUserByEmail(ctx *gin.Context) {
	var users []User
	var usersFilter []User
	jsonData, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(jsonData), &users); err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		if ctx.Param("email") == user.Email {
			usersFilter = append(usersFilter, user)
		}
	}
	if len(usersFilter) > 0 {
		ctx.JSON(200, gin.H{"message": &usersFilter})
	} else {
		ctx.String(404, "Información del empleado ¡No existe!")
	}
}

func GetUserByAge(ctx *gin.Context) {
	var users []User
	var usersFilter []User
	jsonData, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(jsonData), &users); err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		AgeStr := strconv.Itoa(user.Age)
		if ctx.Param("age") == AgeStr {
			usersFilter = append(usersFilter, user)
		}
	}
	if len(usersFilter) > 0 {
		ctx.JSON(200, gin.H{"message": &usersFilter})
	} else {
		ctx.String(404, "Información del empleado ¡No existe!")
	}
}

func GetUserByCreationDate(ctx *gin.Context) {
	var users []User
	var usersFilter []User
	jsonData, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(jsonData), &users); err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		if ctx.Param("creation-date") == user.CreationDate {
			usersFilter = append(usersFilter, user)
		}
	}
	if len(usersFilter) > 0 {
		ctx.JSON(200, gin.H{"message": &usersFilter})
	} else {
		ctx.String(404, "Información del empleado ¡No existe!")
	}
}

func GetUserByActive(ctx *gin.Context) {
	var users []User
	var usersFilter []User
	jsonData, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(jsonData), &users); err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		ActiveStr := strconv.FormatBool(user.Active)
		if ctx.Param("active") == ActiveStr {
			usersFilter = append(usersFilter, user)
		}
	}
	if len(usersFilter) > 0 {
		ctx.JSON(200, gin.H{"message": &usersFilter})
	} else {
		ctx.String(404, "Información del empleado ¡No existe!")
	}
}

func GetUserByHeight(ctx *gin.Context) {
	var users []User
	var usersFilter []User
	jsonData, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(jsonData), &users); err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		HeightStr := strconv.FormatFloat(user.Height, 'f', 0, 64)
		if ctx.Param("height") == HeightStr {
			usersFilter = append(usersFilter, user)
		}
	}
	if len(usersFilter) > 0 {
		ctx.JSON(200, gin.H{"message": &usersFilter})
	} else {
		ctx.String(404, "Información del empleado ¡No existe!")
	}
}

// EJERCICIO 2 T
func GetUserById(ctx *gin.Context) {
	var users []User
	var usersFilter []User
	jsonData, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(jsonData), &users); err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		if ctx.Param("id") == user.Id {
			usersFilter = append(usersFilter, user)
		}
	}
	if len(usersFilter) > 0 {
		ctx.JSON(200, gin.H{"message": &usersFilter})
	} else {
		ctx.String(404, "Información del empleado ¡No existe!")
	}
}

func main() {

	// Crea un router con gin
	router := gin.Default()

	// EJERCICIO 1 T
	//Cada vez que llamamos a GET y le pasamos una ruta, definimos un nuevo endpoint.
	param := "active"
	switch param {
	case "id":
		router.GET("/users/:id", GetUserById)
	case "name":
		router.GET("/users/:name", GetUserByName)
	case "lastname":
		router.GET("/users/:lastname", GetUserByLastName)
	case "email":
		router.GET("/users/:email", GetUserByEmail)
	case "age":
		router.GET("/users/:age", GetUserByAge)
	case "height":
		router.GET("/users/:height", GetUserByHeight)
	case "active":
		router.GET("/users/:active", GetUserByActive)
	case "creation-date":
		router.GET("/users/:creation-date", GetUserByCreationDate)
	default:
		router.GET("/users/:id", GetUserById)
	}
	router.GET("/", HomePage)
	router.GET("/hello-world", GetName)

	router.Run(":8080")

}

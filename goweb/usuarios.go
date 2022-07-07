package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Id          int     `json:"id"`
	Names       string  `json:"nombre"`
	LastName    string  `json:"apellido"`
	Age         int     `json:"edad"`
	DateCreated string  `json:"fechaCreacion"`
	Estatura    float64 `json:"altura"`
	Email       string  `json:"email"`
	IsActivo    bool    `json:"activo"`
}

type UserResult struct {
	usuario  Usuario
	posicion int
}

type Usuarios struct {
	Users []Usuario
}

func leerArchivo() (Usuarios, error) {
	var usersArr Usuarios
	jsonFile, errOpenFile := os.Open("usuarios.json")
	if errOpenFile != nil {
		return Usuarios{}, errOpenFile
	}
	fmt.Println("····· Successfully Opened users.json ✅")
	defer jsonFile.Close()

	byteValue, eReadingJsonFile := ioutil.ReadAll(jsonFile)

	if eReadingJsonFile != nil {
		return Usuarios{}, eReadingJsonFile
	}

	eUnmarshal := json.Unmarshal(byteValue, &usersArr)

	if eUnmarshal != nil {
		return Usuarios{}, eUnmarshal
	}

	return usersArr, nil

}

func formatToJSON(users ...Usuario) ([]byte, error) {
	jsonData, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func getAllUsuarios(ctx *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, "Ocurrió un error...")
		}
	}()
	users := gettingDataFromFile()
	usersJSON, errUsersJSON := formatToJSON(users.Users...)

	if errUsersJSON != nil {
		panic(errUsersJSON)
	}

	ctx.JSON(http.StatusOK, string(usersJSON))

}

func getUsuariosByFilter(ctx *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, "Ocurrió un error...")
		}
	}()

	category := ctx.Query("categoria")
	value := ctx.Query("valor")

	users := gettingDataFromFile()

	usersFiltered := []Usuario{}

	for _, user := range users.Users {
		v := reflect.ValueOf(user)
		tipoUser := v.Type()

		for i := 0; i < v.NumField(); i++ {
			actualField := strings.ToUpper(fmt.Sprintf("%s", tipoUser.Field(i).Name))
			actualValue := strings.ToUpper(fmt.Sprintf("%v", v.Field(i).Interface()))

			if (actualField == strings.ToUpper(category)) && actualValue == strings.ToUpper(value) {
				usersFiltered = append(usersFiltered, user)
			}
		}

	}

	usersJSON, errUsersJSON := formatToJSON(usersFiltered...)

	if errUsersJSON != nil {
		panic(errUsersJSON)
	}

	ctx.JSON(http.StatusOK, string(usersJSON))
}

func getUserById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	users := gettingDataFromFile()

	var userGet UserResult
	for pos, user := range users.Users {
		if user.Id == id {
			userGet.usuario = user
			userGet.posicion = pos
		}
	}

	if (UserResult{} == userGet) {

		ctx.JSON(http.StatusNotFound, gin.H{
			"Error": errors.New("No se encontró el usuario").Error(),
		})
	}

	userJSON, errUsers := formatToJSON(userGet.usuario)
	posJSON, errPos := json.Marshal(userGet.posicion)

	if errUsers != nil || errPos != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": errors.New("Hubo un problema.").Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"usuario":  string(userJSON),
		"posicion": string(posJSON),
	})
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hola, Andy!",
		})
	})
	usuariosRouting := router.Group("/usuarios")
	{
		usuariosRouting.GET("/", getAllUsuarios)
		usuariosRouting.GET("/filtrar", getUsuariosByFilter)
		usuariosRouting.GET("/:id", getUserById)
	}

	router.Run(":8000")

	users := gettingDataFromFile()

	fmt.Printf("····· Hay %d users registrados por ahora😀\n", len(users.Users))
}

func gettingDataFromFile() Usuarios {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	users, errReading := leerArchivo()
	if errReading != nil {
		panic(errReading)
	}

	return users
}

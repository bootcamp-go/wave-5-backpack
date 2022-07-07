package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	Id int
	Name string
	LastName string
	Email string
	Age int
	Height float32
	Active bool
	CreatedAt string
}

// trayendo datos de users.json
func GetAllUsers(c *gin.Context){

//TRAIGO DATOS DEL JSON----------------------------------------------------------------------------------------------
	// Open our jsonFile
	jsonFile, err := os.Open("users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// paso el JSON file a []bytes
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// En vez de crear una var tipo struct, uso una interface vacía que se adapte a cualquier estructura
	var totalUsuarios []*User
	// decodifico el archivo []byte al formato que tenía el JSON
	json.Unmarshal([]byte(byteValue), &totalUsuarios)
//----------------------------------------------------------------------------------------------------------------------

// FILTROS---------------------------------------------------------------------------------------------------------------
	var filters User
	if c.ShouldBindQuery(&filters) == nil { // Setea las variables obtenidas de c.Query("nombredelavariable")
		log.Println(filters.Name)
		log.Println(filters.LastName)
		log.Println(filters.Email)
		log.Println(filters.Age)
		log.Println(filters.Height)
		log.Println(filters.Active)
		log.Println(filters.CreatedAt)
	}

	var filtrado []*User
	for _, user := range totalUsuarios { // filtrado por todos los campos. De esta forma necesito todos los campos para filtrar
		if filters.Name == user.Name && filters.LastName == user.LastName && filters.Email == user.Email && filters.Age == user.Age && filters.Height == user.Height && filters.Active == user.Active && filters.CreatedAt == user.CreatedAt {
			filtrado = append(filtrado, user)
		}
	}
// ---------------------------------------------------------------------------------------------------------------------------
	c.JSON(http.StatusOK, filtrado)

}

func GetUserById(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))

		// Open our jsonFile
		jsonFile, err := os.Open("users.json")
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
		}
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()
	
		// paso el JSON file a []bytes
		byteValue, _ := ioutil.ReadAll(jsonFile)
	
		// En vez de crear una var tipo struct, uso una interface vacía que se adapte a cualquier estructura
		var totalUsuarios []*User
	
		// decodifico el archivo []byte al formato que tenía el JSON
		json.Unmarshal([]byte(byteValue), &totalUsuarios)

	find := false
	var usuarioEncontrado User
	for _,u := range totalUsuarios{
		if u.Id == id {
			find = true
			usuarioEncontrado = *u
		}
	}
	if !find{
		c.JSON(http.StatusNotFound, "No se encontró ningún usuario con ese Id")
	} else {
		c.JSON(http.StatusOK, usuarioEncontrado)
	}


}


func main(){
	
	router := gin.Default()

	router.GET("/inicio", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hola Franco!",
		})
	})

	router.GET("/users", GetAllUsers)
	router.GET("/users/:id", GetUserById) 

	router.Run(":8080")

}


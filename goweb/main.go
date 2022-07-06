package main

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	//"log"
	"os"
	"fmt"
	"io/ioutil"
)

type Usuario struct {
	Id int
	Name string
	LastName string
	Email string
	Age int
	Height float32
	Active bool
	CreatedAt string
}

var (
	u1 = Usuario{
		Id: 1,
		Name: "Jorge",
		LastName: "Gonzalez",
		Email: "jorge@mail.com",
		Age: 30,
		Height: 2.02,
		Active: true,
		CreatedAt: "21/06/2",
	}

	u2 = Usuario{
		Id: 2,
        Name: "Romina",
        LastName: "Gutierrez",
        Email: "romi@mail.com",
        Age: 29,
        Height: 1.64,
        Active: false,
        CreatedAt: "03/10/20",
	}

)

var miSlice = []Usuario{u1,u2}




func GetAll1(c *gin.Context){
	c.JSON(200, miSlice)
}

func GetAll2(c *gin.Context){
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
	var result map[string]interface{}

	// decodifico el archivo []byte al formato que tenía el JSON
	json.Unmarshal([]byte(byteValue), &result)

	// lo devuelvo
	c.JSON(200, result["users"])

}


func main(){
	
	router := gin.Default()

	router.GET("/inicio", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hola Franco!",
		})
	})

	router.GET("/users1", GetAll1)
	router.GET("/users2", GetAll2)

	router.Run(":8080")

}


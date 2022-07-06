package main

import "github.com/gin-gonic/gin"

func main() {
	// Crea un router con gin
	router := gin.Default()

	// Captura la solicitud GET “/hello-world”
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(401, gin.H{"message": "Hola Mundo!"})
	})

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run(":8080")
}

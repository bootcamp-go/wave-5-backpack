package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
) //github.com/usuario/web-server
type usuario struct {
	Id     int    `json:"-"`
	Nombre string `json:"nickname" binding:"required"`
	Edad   int    `json:"edad" binding:"required"`
}

var id int
var listUsur []usuario

func main() {
	fmt.Println("hoal")
	router := gin.Default()

	router.POST("/usuario", func(ctx *gin.Context) {
		token := "123456"
		if token == ctx.GetHeader("token") {
			id += 1
			fmt.Println(id)
			var usuar usuario
			if err := ctx.ShouldBindJSON(&usuar); err != nil {
				fmt.Println("mira")

				var bodyBytes []byte

				fmt.Println(ctx)
				fmt.Println(ioutil.ReadAll((*ctx.Request).Body))
				bodyBytes, _ = ioutil.ReadAll(ctx.Request.Body)

				fmt.Println(bodyBytes)
				// Restore the io.ReadCloser to its original state
				ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
				fmt.Println(ioutil.NopCloser(bytes.NewBuffer(bodyBytes)))
				fmt.Println(ctx.Request.Body)
				ctx.JSON(404, gin.H{
					"error": err.Error(),
				})
			} else {
				usuar.Id = id
				listUsur = append(listUsur, usuar)
				fmt.Println(listUsur)
				ctx.JSON(200, listUsur)
			}
		} else {
			ctx.JSON(401, "no tiene permisos para realizar la petición solicitada")
		}

	})
	router.Run()
}

package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.POST("transactions", PostTransactions())
	router.Run(":3000")
}

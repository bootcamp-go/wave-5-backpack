package main

import (
	"fmt"
	"time"
)

func main() {
	statusCode := 404
	if statusCode >= 400 {
		err := fmt.Errorf("momento del error: %v", time.Now())
		fmt.Println("error ocurrido: ", err)
		return
	}
	fmt.Println("la peticion reuslto con el codigo:", statusCode)
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	statusCode := 404
	if statusCode >= 400 {
		fmt.Println(errors.New("La petición ha fallado."))
		return
	}
	fmt.Println("El programa finalizó correctamente.")
}

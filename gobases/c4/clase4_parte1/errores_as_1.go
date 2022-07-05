package main

import (
	"errors"
	"fmt"
)

type myError struct {
	msg string
	x   string
}

type myError2 struct {
	msg string
	x   string
}

func (e *myError) Error() string {
	return fmt.Sprintf("ha ocurrido un error: %s, %s", e.msg, e.x)
}

func (e *myError2) Error() string {
	return fmt.Sprintf("ha ocurrido un error: %s, %s", e.msg, e.x)
}

func main() {
	// Errores de tipo iguales
	err1 := &myError{"nuevo error", "404"}
	err2 := &myError{"otro error", "400"}

	isMyError := errors.As(err1, &err2)

	fmt.Println("Son del mismo tipo? :", isMyError)

	// Errores de tipo distintos
	err3 := &myError{"nuevo error", "404"}
	err4 := &myError2{"otro error", "400"}

	isMyError = errors.As(err3, &err4)

	fmt.Println("Son del mismo tipo? :", isMyError)
}

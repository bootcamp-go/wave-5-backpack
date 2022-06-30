package main

import (
	"errors"
	"fmt"
)

type Student struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (s Student) detalle(DNI int) (string, error) {
	var result string

	if DNI == s.DNI {
		result = s.Nombre
	} else {
		return "", errors.New("DNI del estudiante es incorrecto")
	}

	return result, nil
}

func main() {
	a1 := Student{"Camilo", "Carreño", 123, "12/10/2022"}
	res, err := a1.detalle(13)

	if err != nil {
		fmt.Printf("Ocurrio un error: %v \n", err)
	} else {
		fmt.Printf("El nombre del estudiante consultado es: %v \n", res)
	}

}

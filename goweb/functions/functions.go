package functions

import (
	"errors"
	"math/rand"

	"github.com/go-playground/validator/v10"
)

func IdRandom(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}

func ValidateErrors(campo string, v validator.FieldError) string {
	switch v.Tag() {
	case "required":
		return "El campo " + campo + " es requerido"
	case "email":
		return "Direccion de correo electronico invalida"
	}
	return "Error desconoodido..."
}

func ValidateToken(tokenAuth string) error {
	if tokenAuth != "62c5b68a0cc23a33375c85f8" {
		return errors.New("no tiene permisos para realizar la petición solicitada")
	}
	return nil
}

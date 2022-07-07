package middleware

import (
	"errors"
)

func validate(token string) (string, error) {

	if token != "123456" {
		return "", errors.New("Algo salio mal")
	}

	return "valido", nil

}

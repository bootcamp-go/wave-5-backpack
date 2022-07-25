package handler

import (
	"os"
)

func validarToken(token string) bool {
	if token != os.Getenv("TOKEN") {
		return true
	}
	return false
}

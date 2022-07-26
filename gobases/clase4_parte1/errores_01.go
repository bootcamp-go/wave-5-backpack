package main

import (
	"fmt"
)

type statusError struct {
	status int
	msg    string
}

//hacemos que nuestro struct implemente el método Error()
func (e *statusError) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}

func api(token int) (int, error) {
	secretToken := 123
	if token != secretToken {
		return 401, &statusError{
			status: 401,
			msg:    "token incorrecto",
		}
	}
	return 200, nil
}

func main() {
	status, err := api(123)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Status %d, Funciona!", status)
}

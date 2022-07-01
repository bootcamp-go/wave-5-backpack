package main

import (
	"fmt"
)

type statusPer struct {
	estado int
	msg    string
}

//Aca hago que mi estructura IMPLEMENTE el metodo ERROR() de la interfaz
func (e *statusPer) Error() string {
	return fmt.Sprintf("%d,%v", e.estado, e.msg)
}

func miErrorCustom(salario int) (int, error) {
	if salario < 150000 {
		return 1, &statusPer{
			estado: 1,
			msg:    "error: el salario ingresado no alcanza el mínimo imponible",
		}
	} else if salario >= 150000 {
		return 0, &statusPer{
			estado: 1,
			msg:    "Debe pagar impuestos",
		}
	}
	return 0, nil
}

func main() {
	var salary int = 232321321
	estado, err := miErrorCustom(salary)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("todo excelente!", estado)
}

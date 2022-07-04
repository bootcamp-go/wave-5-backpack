package main

import (
	"errors"
	"fmt"
)

type salary struct {
	Monto   int
	Mensaje string
}

func (s *salary) Error() string {
	return s.Mensaje
}

func controlError(salarys int, err error) {
	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
	} else {
		fmt.Printf("Debe pagar impuestos. Su salario ($%d) supera los $150000\n", salarys)
	}
}

func impuestoSalaryError(salarys int) (int, error) {
	if salarys < 150000 {
		return salarys, &salary{
			Monto:   salarys,
			Mensaje: "error: el salario ingresado no alcanza el mínimo imponible"}
	} else {
		return salarys, nil
	}
}

func impuestoSalaryErrorNew(salarys int) (int, error) {
	if salarys < 150000 {
		return salarys, errors.New("error: el salario ingresado no alcanza el mínimo imponible")
	} else {
		return salarys, nil
	}
}

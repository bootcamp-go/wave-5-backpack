package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	empleado := empleado{Nombre: "Francisco", Apellido: "Monay", HsTrabajadas: 100, MesesTrabajados: 6}

	salario, err := salarioTotal(empleado.HsTrabajadas, 2000)

	if err != nil {
		fmt.Println(err)
	} else {
		empleado.Salario = salario
	}

	aguinaldo, err2 := calcularAguinaldo(empleado.MesesTrabajados, empleado.Salario)

	if err != nil {
		err = fmt.Errorf("err2: %w", err2)
		fmt.Println(errors.Unwrap(err))
		os.Exit(1)
	} else {
		empleado.Aguinaldo = aguinaldo
	}

	empleado.detalle()

}

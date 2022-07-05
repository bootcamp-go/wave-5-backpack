package main

import (
	"errors"
	"fmt"
)

const (
	SUMA           = "+"
	RESTA          = "-"
	MULTIPLICACION = "*"
	DIVISION       = "/"
)

func operacionSuma(val1, val2 float64) (float64, error) {
	return val1 + val2, nil
}

func operacionResta(val1, val2 float64) (float64, error) {
	return val1 - val2, nil
}

func operacionMultiplicacion(val1, val2 float64) (float64, error) {
	return val1 * val2, nil
}

func operacionDivision(val1, val2 float64) (float64, error) {
	var err error
	if val2 == 0 {
		err = errors.New("cant divide by 0")
	}
	return val1 / val2, err
}

type operacion func(float64, float64) (float64, error)

func maestroOperaciones(tipoOperacion operacion, params ...float64) (float64, error) {
	var res float64 = params[0]
	var err error
	for i := 1; i < len(params); i++ {
		res, err = tipoOperacion(res, params[i])
	}
	return res, err
}

func operar(operacionString string, params ...float64) (float64, error) {
	var tipoOperacion operacion
	switch operacionString {
	case SUMA:
		tipoOperacion = operacionSuma
	case RESTA:
		tipoOperacion = operacionResta
	case MULTIPLICACION:
		tipoOperacion = operacionMultiplicacion
	case DIVISION:
		tipoOperacion = operacionDivision
	}

	return maestroOperaciones(tipoOperacion, params...)
}

func main() {
	res, err := operar("/", 4, 2)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}
	fmt.Println(res)
}

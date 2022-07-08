package main

import (
	"fmt"
	"errors"
)

const (
	CATA = "cata"
	CATB = "catb"
	CATC = "catc"
)

func calcHoras(minutos int) float32 {
	return float32(minutos) / 60
}

func calcSalario(bono float32, salario float32, tiempo int) float32 {
	salarioMensual := salario * calcHoras(tiempo)
	return salarioMensual + (salarioMensual * bono)
}

func orquestradorDeOperaciones(cat string, minutosTrabajados int) (float32,error) {
	switch cat {
		case CATA:
			return calcSalario(0.5, 3000.0, minutosTrabajados), nil
		case CATB:
			return calcSalario(0.2, 1500.0, minutosTrabajados), nil
		case CATC:
			return calcSalario(0.0, 1000.0, minutosTrabajados), nil
		default:
			return 0, errors.New("No se reconoce la categoría")
	}
}

func main() {
	cataValue, errcata := orquestradorDeOperaciones(CATA, 1000)
	catbValue, errcatb := orquestradorDeOperaciones(CATB, 1000)
	catcValue, errcatc := orquestradorDeOperaciones(CATC, 1000)

	if errcata == nil {
		fmt.Printf("Salario categoría A: %.2f\n", cataValue)
	} else {
		fmt.Println(errcata)
	}

	if errcatb == nil {
		fmt.Printf("Salario categoría B: %.2f\n", catbValue)
	} else {
		fmt.Println(errcatb)
	}

	if errcatc == nil {
		fmt.Printf("Salario categoría C: %.2f\n", catcValue)
	} else {
		fmt.Println(errcatc)
	}
}
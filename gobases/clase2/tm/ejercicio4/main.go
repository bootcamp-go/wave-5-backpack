package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

// Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los alumnos de un curso,
// requiriendo calcular los valores mínimo, máximo y promedio de sus calificaciones.

// Se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio)
// y que devuelva otra función ( y un mensaje en caso que el cálculo no esté definido) que se le puede pasar una cantidad N de enteros
// y devuelva el cálculo que se indicó en la función anterior

func main() {
	minOp, err := operacion("minimo")
	if err != nil {
		log.Fatal(err)
	}

	min := minOp(3,6,10,-3)
	fmt.Printf("Min: %v\n", min)

	maxOp, err := operacion("maximo")
	if err != nil {
		log.Fatal(err)
	}

	max := maxOp(-5,3,256,123)
	fmt.Printf("Max: %v\n", max)

	promOp, err := operacion("promedio")
	if err != nil {
		log.Fatal(err)
	}

	prom := promOp(3,-4,9,10)
	fmt.Printf("Promedio: %v\n", prom)
}

const (
  min = "minimo"
  max = "maximo"
  prom = "promedio"
)

func operacion(operacion string) (func(n ...int) float64, error){
  switch operacion {
  case min:
    return opMin, nil
  case max:
    return opMax, nil
  case prom:
    return opProm, nil
  }

  return nil, errors.New("operación no conocida")
}

func opMin(nums ...int) float64 {
  min := math.MaxInt

  for _, v := range nums {
    if  v < min {
      min = v
    }
  }

  return float64(min)
}

func opMax(nums ...int) float64 {
  max := math.MinInt

  for _, v := range nums {
    if v > max {
      max = v
    }
  }

  return float64(max)
}

func opProm(nums ...int) float64 {
  total := float64(len(nums))
  sum := 0

  for _, v := range nums {
    sum += v
  }
  promedio := float64(sum)/total

  return promedio
}

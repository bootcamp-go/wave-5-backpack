package main

import (
	"errors"
	"fmt"
	"log"
)

// Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
// Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

// perro necesitan 10 kg de alimento
// gato 5 kg
// Hamster 250 gramos.
// Tarántula 150 gramos.

// Se solicita:
// Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado
// y que retorne una función y un mensaje (en caso que no exista el animal)

// Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.


func main() {
  perroFunc, err := Animal("perro")
  if err != nil {
    log.Println(err)
  }
  fmt.Printf("Necesito %v kg de comida para 5 perros\n", perroFunc(5))

  gatoFunc, err := Animal("gato")
  if err != nil {
    log.Println(err)
  }
  fmt.Printf("Necesito %v kg de comida para 3 gatos\n", gatoFunc(3)) 
  
  hamsterFunc, err := Animal("hamster")
  if err != nil {
    log.Println(err)
  }
  fmt.Printf("Necesito %v gr de comida para 7 hamsters\n", hamsterFunc(7))

  tarantulaFunc, err := Animal("tarantula")
  if err != nil {
    log.Println(err)
  }
  fmt.Printf("Necesito %v gr de comida para 2 tarantulas\n", tarantulaFunc(2))
}

const (
  perro = "perro"
  gato = "gato"
  hamster = "hamster"
  tarantula = "tarantula"
)

func Animal(animal string) (func(n uint) uint, error) {
  switch animal {
  case perro:
    return perroFunc, nil
  case gato:
    return gatoFunc, nil
  case hamster:
    return hamsterFunc, nil
  case tarantula:
    return tarantulaFunc, nil
  }

  return nil, errors.New("Animal no conocido")
}

// Retorna la cantidad de alimento necesario en kg
func perroFunc(n uint) uint {
  return n * 10
}

// Retorna la cantidad de alimento necesario en kg
func gatoFunc(n uint) uint {
  return n * 5 
}

// Retorna la cantidad de alimento necesario en gr
func hamsterFunc(n uint) uint {
  return n * 250
}

// Retorna la cantidad de alimento necesario en gr
func tarantulaFunc(n uint) uint {
  return n * 150
}

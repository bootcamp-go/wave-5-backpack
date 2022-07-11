package main

import (
	"math/rand"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/gobases/clase3/tt/ejercicio4/ordenamiento"
)

func main() {
	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	ordenamiento.Ordenar(variable1)
	ordenamiento.Ordenar(variable2)
	ordenamiento.Ordenar(variable3)
}

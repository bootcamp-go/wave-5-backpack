package main

import "fmt"

const (
	Minimum = "minimum"
	Vverage = "average"
	Maximum = "maximum"
)

func minimum(valores ...float64) float64 {
	minimum := valores[0]
	for _, value := range valores {
		if value < minimum {
			minimum = value
		}
	}
	return minimum

}

func operation(operacion string, minimun func(valores ...int)) float64 {
	switch operacion {
	case Minimum:
		fmt.Printf("hh")
		return minimum()
	}
	return 0

}

func main() {
	valores := operation(Minimum, 1, 2, 3)

	fmt.Println(valores)
}

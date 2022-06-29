package main

import(
	"fmt"
	"errors"
)

func main(){
	const(
		Minimum = "minimum"
		Average = "average"
		Maximum = "maximum"
	)

	minFunc, _ := operation(Minimum)
	averageFunc, _ := operation(Average)
	maxFunc, _ := operation(Maximum)

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("Estadísticas")
	fmt.Printf("valor mínimo:%.2f", minValue)
	fmt.Printf("valor promedio:%.2f", averageValue)
	fmt.Printf("valor máximo:%.2f", maxValue)

}

func minFunction(valores ...float64) float64{
	valueMin := 0.0
	for _, valor := range valores{
		if valueMin <= valor {
			valueMin = valor
		}
	}
	return valueMin
}

func averageFunction(valores ...float64) float64{
	var resultado float64 = 0.0
	for _, valor := range valores{
		resultado += valor
	}
	return (resultado/float64(len(valores)))
}

func maxFunction(valores ...float64) float64{
	valueMax := 800000000.0
	for _, valor := range valores{
		if valueMax <= valor {
			valueMax = valor
		}
	}
	return valueMax
}

func operation(calculo string) (func(valores ...float64) float64, error){
	switch calculo {
		case "minimum":
			return minFunction,nil
		case "average":
			return averageFunction,nil
		case "maximum":
			return maxFunction,nil
		return nil, errors.New("La operación no es válida")
	}
}

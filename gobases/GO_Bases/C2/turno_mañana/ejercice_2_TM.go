package main

import "fmt"

func averag(values ...int) (float64, string) {
	var count int = 0
	for _, value := range values {
		if value <= 0 {
			return 0, fmt.Sprint("ERROR : Una de las notas es negativa o cero")
		}
		count += value
	}
	return float64(count) / float64(len(values)), ""
}

func main() {
	avg, err := averag(-4, 5, 2, 1)
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Printf("El promedio de las notas es %.2f \n", avg)
	}
}

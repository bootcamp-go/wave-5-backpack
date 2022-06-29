package main

import (
	"errors"
	"fmt"
)

const (
	DOG     = "dog"
	CAT     = "cat"
	HAMSTER = "hamster"
	SPIDER  = "spider"
)

const (
	KG    = "kg"
	GRAMS = "grams"
)

func calcAnimalFood(requiredByAnimal float32, metric string) func(currentFood float32) float32 {
	requieredKg := metricToKg(requiredByAnimal, metric)
	return func(currentFood float32) float32 {
		if requieredKg > currentFood {
			return requieredKg - currentFood
		}
		return 0
	}
}

func metricToKg(metricValue float32, metric string) float32 {
	switch metric {
	case KG:
		return metricValue
	case GRAMS:
		return metricValue / 1000
	}
	return 0
}

func animal(typeAnimal string) (func(currentFood float32) float32, error) {
	switch typeAnimal {
	case DOG:
		return calcAnimalFood(10, KG), nil
	case SPIDER:
		return calcAnimalFood(150, GRAMS), nil
	case CAT:
		return calcAnimalFood(5, KG), nil
	case HAMSTER:
		return calcAnimalFood(250, GRAMS), nil
	}

	return nil, errors.New("Error Animal no encontrado: " + typeAnimal)
}

func getFoodRequiredByStorage(storage map[string]float32) float32 {
	requiredAmount := float32(0)
	for ani, currentQty := range storage {
		calcAnimal, errAnimal := animal(ani)
		if errAnimal != nil {
			fmt.Println(errAnimal)
		} else {
			requiredAmount += calcAnimal(currentQty)
		}
	}
	return requiredAmount
}

func main() {
	actualStorage := map[string]float32{
		CAT:     2,
		HAMSTER: metricToKg(200, GRAMS),
		SPIDER:  metricToKg(100, GRAMS),
		DOG:     5,
	}
	fmt.Println("Cantidad total a comprar", getFoodRequiredByStorage(actualStorage), "KG")
}

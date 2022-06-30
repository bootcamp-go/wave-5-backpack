package main

import ("fmt"
		"errors")

func main(){

	const(
		cat 	= "cat"
		dog 	= "dog"
		spider 	= "spider"
		hamster	= "hamster"
	)

	quantityFood,err:= Animal(cat)

	

	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Printf("la cantidad de comida es %v \n", quantityFood(5.0))
	}

}


func Animal(animal string) (func(num float64) float64,error) {

	switch animal {
	case "dog":
		return dogFood,nil
	case "cat":
		return catFood,nil
	case "spider":
		return spiderFood,nil
	case "hamster":
		return hamsterFood,nil
	}
	return 0, errors.New("animal incorrecto , no encontrado")

}


func dogFood(num float64) float64{
	return num*10
}
func catFood(num float64) float64{
	return num*5
}
func spiderFood(num float64) float64{
	return num*0.15
}
func hamsterFood(num float64) float64{
	return num*0.25
}

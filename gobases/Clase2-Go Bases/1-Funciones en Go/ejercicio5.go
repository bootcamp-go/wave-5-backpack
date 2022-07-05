package main

import "fmt"

const (
	dog = "dog"
	cat = "cat"
	hamster ="hamster"
	tarantula="tarantula"
 )
 
 

func main(){
	var amount float64
	//Calculo Perro
	animalDog, msg := Animal(dog)
	if(msg!=""){
		fmt.Printf("%v\n",msg)
	}else{
		amount+= animalDog(5)
		fmt.Printf("Cantidad acumulada %v \n",amount) 
	}

	//Calculo Gato
	animalCat, msg := Animal(cat)
	if(msg!=""){
		fmt.Printf("%v\n",msg)
	}else{
		amount+= animalCat(8)
		fmt.Printf("Cantidad acumulada %v \n",amount)   
	}

}

func Animal(nombreAnimal string) (func(int) float64,string){
	switch nombreAnimal {
	case cat:
		return comidaGato,""
	case dog:
		return comidaPerro,""
	case hamster:
		return comidaHamster,""
	case tarantula:
		return comidaTarantula,""
	default:
		return nil,"No se reconoció el animal "+nombreAnimal
	}
	
} 

func comidaPerro(valor int) float64{
	
	resultado:=float64(valor)*10
	return resultado 
}

func comidaGato(valor int) float64{
	
	resultado:=float64(valor)*5
	return resultado 
}
func comidaHamster(valor int) float64{
	
	resultado:=float64(valor)*0.25
	return resultado 
}
func comidaTarantula(valor int) float64{
	
	resultado:=float64(valor)*0.15
	return resultado  
}





/*Ejercicio 5 - Calcular cantidad de alimento

Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan haber muchos más animales que refugiar.

perro necesitan 10 kg de alimento
gato 5 kg
Hamster 250 gramos.
Tarántula 150 gramos.

Se solicita:
Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y un mensaje (en caso que no exista el animal)
Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.

*/
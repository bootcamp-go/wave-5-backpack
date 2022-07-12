package main

import (
	"fmt"
	"errors"
)

const (
	perro = "perro"
	gato = "gato"
	hamster = "hamster"
	tarantula = "tarantula"
)

func main(){
	calculaAlimento,e:= Animal(perro)

	if e != nil {
		fmt.Printf("%v\n",e)
	}else{
		alimento := calculaAlimento(2)
		fmt.Printf("%s\n",alimento)
	}
	
}
func Animal(ops string) (func(cantida int)string,error){
	switch ops {
		case perro:
			return CalcPerro,nil
		case gato:
			return CalcGato,nil
		case hamster:
			return CalcHamster,nil
		case tarantula:
			return CalcTarantula,nil
		}
		return nil,errors.New("Opcion invalida")
}
func CalcPerro (cantida int)string {
	alimentoNum := cantida * 10
	return fmt.Sprintf("Perros: %d Alimento: %d kg\n", cantida,alimentoNum)
}
func CalcGato (cantida int)string {
	alimentoNum := cantida * 5
	return fmt.Sprintf("Gatos: %d Alimento: %d kg\n", cantida,alimentoNum)
}
func CalcHamster (cantida int)string {
	alimentoNum := cantida * 250
	return fmt.Sprintf("Hamnster: %d Alimento: %d g\n", cantida,alimentoNum)
}
func CalcTarantula (cantida int)string {
	alimentoNum := cantida * 150
	return fmt.Sprintf("Tarantula: %d Alimento: %d g\n", cantida,alimentoNum)
}

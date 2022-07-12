package main

import (
	"fmt"
	"errors"
)

const (
	valMin = "valMin"
	valMax = "valMax"
	prom = "prom"
)

func main(){
	ctrl,e:= controlador(valMin)

	if e != nil {
		fmt.Printf("%v\n",e)
	}else{
		resultado := ctrl(1,2,3,4,5)
		fmt.Printf("%d\n",resultado)
	}
	
}
func controlador(ops string) (func(values ...int)int,error){
	switch ops {
		case valMin:
			return CalcMin,nil
		case valMax:
			return CalcMax,nil
		case prom:
			return CalcProm,nil
		}
		return nil,errors.New("Opcion invalida")
}

func CalcMin(values ...int) int {
	fmt.Printf("Minimo\n")
	var min int 
	for key,val := range values {
		if key == 0{
			fmt.Printf("vaor: %d\n",val)
			min = val
		}
		if val < min {
			min = val
		}
	}
	return min
}
func CalcMax(values ...int) int{
	fmt.Printf("Maximo\n")
	max := values[0]
	for key,val := range values {
		if key == 0{
			fmt.Printf("vaor: %d\n",val)
			max = val 
		}
		if val > max {
			max = val
		}
	}
	return max
}
func CalcProm(values ...int) int{
	fmt.Printf("Promedio\n")
	var prom,sum int 
	for _,val := range values {
		sum += val
	}
	prom = sum/len(values)
	return prom
}
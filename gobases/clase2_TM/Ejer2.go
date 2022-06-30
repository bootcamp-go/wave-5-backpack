package main

import("fmt"
		"errors")
func main(){

	var qualifications =[]float64{1.2,3.0,-10.0,5.0}
	var value float64
	var err error
	value , err = Average(qualifications...)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Printf("el promedio del estudiantes es %v \n",value)

	}
}


func Average(values ...float64) (float64,error){
	var qualification float64

	for _,v := range values{
		qualification+=v
	}
	qualification/=float64(len(values))

	if qualification<0{
		return 0,errors.New("Calificaciones negativas")
	}else {
		return qualification,nil
	}
}

// func inputQualifications(){
// 	slice 

// 	if (value==)
// 	return 
// }
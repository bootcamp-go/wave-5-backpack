package main

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main(){
	var qualification =[]float64{1.2,3.3,4.2,3.2,5.0}
	oper := operation(minimum)
	m := oper(qualification...)

	fmt.Printf("Operacion %v resultado %v \n",minimum,m)

}


func  operation(statisticOperator string) (func(qualification ...float64) float64){
	switch statisticOperator {
	case "minimum":
		return minimumFunc
	case "average":
		return averageFunc
	case "maximum":
		return maximumFunc
	}
	return nil
}



func averageFunc(qualification ...float64) float64{
	var result float64

	for _,v := range qualification{
		result+=v
	}

	return result/float64(len(qualification))
}

func minimumFunc(qualification ...float64) float64{
	var min float64= qualification[0]
	for _,v := range qualification{
		if min>=v{
			min=v
		}
	}
	return min
}

func maximumFunc(qualification ...float64) float64{
	var max float64= qualification[0]
	for _,v := range qualification{
		if max<=v{
			max=v
		}
	}

	return max
}
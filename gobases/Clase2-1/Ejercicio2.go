package main

import "fmt"

func main(){
	var prom float64 = CaluloPromedio(17,25,36,85)
	fmt.Printf("promedio: %.2f\n",prom)
}
func CaluloPromedio(califs ...float64) float64{
	var prom, sum float64 = 0,0
	for _,valor := range califs{
		sum += valor
	}
	prom = sum/float64(len(califs))
	return prom
}
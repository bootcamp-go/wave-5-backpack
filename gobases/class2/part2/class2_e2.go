package main

import (
	"fmt"
)	

type Matrix struct{
	Values 		[]float64
	Height 		int
	Width 		int
	Quadratic 	bool
	MaxValue 	float64
}

func maxValue(values ...float64)float64 {
	result := 0.0
	for _, value := range values{
		if result == 0{
			result = value
		}else if result < value{
			result = value
		}
	}
	
	return result
}

func (matrix *Matrix) Set(height int, width int, quadratic bool, values ...float64) {
	matrix.Values = 	values
	matrix.Height =  	height
	matrix.Width =  	width
	matrix.Quadratic = 	quadratic
	matrix.MaxValue = 	maxValue(values...)
}

func (matrix Matrix)Print()  {

	flag := 0
	for j:= 0; j < matrix.Height; j++ {
		for i := 0; i < matrix.Width; i++ {
			fmt.Printf("%.2f\t", matrix.Values[flag])
			flag++
		}
		fmt.Println()
	}
}

func main()  {
	matrix := Matrix{}
	matrix.Set(2, 2, true, 1, 2, 3, 4)
	matrix.Print()
}
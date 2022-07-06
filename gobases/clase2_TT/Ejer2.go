package main

import "fmt"

type Matrix struct {
	matriz     [][]float64 //[m][n]
	m          int
	n          int
	cuadratica bool
	maximum    float64
}

func main() {
	m := Matrix{}
	m.set(2, 2, 12.2, 12, 11, 11)
	m.print()

}

// referenced method
func (mz *Matrix) set(x, y int, params ...float64) {
	if x == y {
		mz.cuadratica = true
	}
	mz.m = x
	mz.n = y
	var count int               // walk throught params
	var max float64 = params[0] // init max value with first found
	for i := 0; i < mz.n; i++ {
		mz.matriz = append(mz.matriz, []float64{})
		for j := 0; j < mz.m; j++ {
			mz.matriz[i] = append(mz.matriz[i], params[count])
			if params[count] > max {
				max = params[count]
			}
			count++
		}
	}
}

func (mz *Matrix) print() {

	fmt.Printf("Dimension %d x %d \n", mz.m, mz.n)
	for j := 0; j < len(mz.matriz); j++ { //walk throught
		for i := 0; i < len(mz.matriz[0]); i++ {
			if i == 0 {
				fmt.Print("[")
			}
			fmt.Printf("%f  ", mz.matriz[j][i])
			if i == len(mz.matriz)-1 {
				fmt.Print("]")
			}
		}
		fmt.Println("")
	}

}

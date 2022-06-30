package main

import "fmt"

func impuestoDeSueldo(monto float64) float64 {

	if monto > 50000 {
		return monto * 0.17
	}
	if monto > 150000 {
		monto1 := monto * 0.17
		monto2 := (monto - monto1) * 0.10

		return monto2
	}

	return monto * 0.10

}

func main() {
	monto := impuestoDeSueldo(151000)

	fmt.Println("impuesto:", monto)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	var matriz [3][4]int
	valores := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	var iniPos int = -1
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			iniPos += 1
			matriz[i][j] = valorPos(valores, iniPos)
		}
	}

	Print(matriz)
}

func valorPos(valores []int, i int) int {
	return valores[i]
}

func Print(matriz [3][4]int) {
	for i := 0; i < len(matriz); i++ {
		fmt.Printf("\n\n")
		for j := 0; j < len(matriz[j]); j++ {
			time.Sleep(time.Second * 2)
			fmt.Printf("\t%d ", matriz[i][j])
		}
		fmt.Printf("\n\n")
	}
}

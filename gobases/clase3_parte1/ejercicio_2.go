package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calcularTotal() float64 {
	var total float64

	return total
}

func main() {

	data, err := os.ReadFile("./informacion_producto.csv")
	slit := strings.Split(string(data), "\n")

	fmt.Println(len(slit))
	var total float64
	for _, value := range slit {
		line := strings.Split(value, ";")
		precio, _ := strconv.ParseFloat(line[1], 64)
		cantidad, _ := strconv.Atoi(line[2])
		total += precio * float64(cantidad)
		fmt.Printf("%s\t%12s%12s\n", line[0], line[1], line[2])
	}
	fmt.Printf("%s\t%12.2f\t%12s\n", "Total", total, " ")
	/* 	for i := 0; i < len(slit); i++ {
		fmt.Printf("%s\t", slit[i])
	} */

	if err != nil {
		fmt.Println(err)
	}

}

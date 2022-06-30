package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	byteArray, err := os.ReadFile("../result.csv")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("ID\tPrecio\t\tCantidad\n")
	text := string(byteArray)
	lines := strings.Split(text, "\n")
	var total float64
	for _, line := range lines {
		values := strings.Split(line, ",")
		if len(values) == 3 {
			precio, _ := strconv.ParseFloat(values[1], 32)
			cantidad, _ := strconv.ParseFloat(values[2], 32)
			total += precio * cantidad
			fmt.Printf("%s\t%f\t%f\n", values[0], precio, cantidad)
		} else {
			continue
		}
	}
	fmt.Printf("Total\t%f\n", total)
}

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./gobases/3/tm/ejercicio_1.csv")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s %10s %5s\n", "ID", "Precio", "Cantidad")
	var total float64

	for _, col := range rows[1:] {
		fmt.Printf("%s %11s %8s\n", col[0], col[1], col[2])

		precio, err := strconv.ParseFloat(col[1], 64)
		if err != nil {
			fmt.Println(err)
		}
		total += precio
	}

	fmt.Printf("%13s\n", fmt.Sprint(total))
}

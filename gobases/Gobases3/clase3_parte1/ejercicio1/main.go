package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	data := [][]string{
		{"ID", "Precio", "Cantidad"},
		{"111223", "30012.00", "1"},
		{"444321", "1000000.00", "4"},
		{"434321", "50.50", "1"},
		{"0", "4030062.50", "0"},
	}

	file, err := os.Create("example.csv")

	if err != nil {
		log.Fatal(err)
	}

	filecsv := csv.NewWriter(file)

	for _, row := range data {
		_ = filecsv.Write(row)
	}

	filecsv.Flush()
	file.Close()

}

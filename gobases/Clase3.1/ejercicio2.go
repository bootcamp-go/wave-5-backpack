package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, _ := os.ReadFile("./products.csv")
	f := csv.NewReader(bytes.NewBuffer(data))
	f.Comma = ';'
	line, err := f.Read()
	if err != nil {
		os.Exit(0)
	}
	fmt.Printf("----------------------------------------\n")
	fmt.Printf("| %-10v | %10v | %10v |\n", line[0], line[1], line[2])
	fmt.Printf("|------------|------------|------------|\n")

	sum := 0
	for {
		line, err := f.Read()
		if err != nil {
			break
		}
		fmt.Printf("| %-10v | %10v | %10v |\n", line[0], line[1], line[2])
		num, _ := strconv.Atoi(line[1])
		sum += num
	}
	fmt.Printf("----------------------------------------\n")
	fmt.Printf("  %10v : %10v \n", "Total", sum)
}

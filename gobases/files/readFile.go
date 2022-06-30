package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// open file
	f, err := os.Open("personal.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(" %-10s", strings.Split(rec[0], ";")[0])
		fmt.Printf(" %10s", strings.Split(rec[0], ";")[1])
		fmt.Printf(" %20s \n", strings.Split(rec[0], ";")[2])

	}
}

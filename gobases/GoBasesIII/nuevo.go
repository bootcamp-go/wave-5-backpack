package main

import "fmt"

import (
	"log"
	"os"

	"github.com/tobgu/qframe"
)

func main() {

	file, err := os.Open("example.csv")

	if err != nil {
		log.Fatal(err)
	}

	read := qframe.ReadCSV(file)

}

// usando la libreria qframehttps://github.com/tobgu/qframe
package main

import (
	"fmt"
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
	fmt.Println("----------")
	fmt.Println(read)
	fmt.Println("----------")
}

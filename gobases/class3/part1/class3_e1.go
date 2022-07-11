package main

import (
	"fmt"
	"os"
)

func checkFile(fileName string) bool {
	files, _ := os.ReadDir(".")
	exists := false
	for _, file := range files {
		if file.Name() == fileName {
			exists = true
			break
		}
	}

	return exists
}

func main() {

	if !checkFile("Products.cvs") {
		myfile, _ := os.Create("Products.cvs")
		myfile.Close()
	}

	myfile, err := os.OpenFile("Products.cvs", os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		fmt.Println(err)
	}

	d1 := []byte("\"1\";\"200.00\";\"2\"\n")

	_, err = myfile.WriteString(string(d1))

	if err != nil {
		fmt.Println(err)
	}

	myfile.Close()
	file, _ := os.ReadFile("./Products.cvs")

	fmt.Println(string(file))

}

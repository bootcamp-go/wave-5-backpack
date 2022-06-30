package main

import (
	"io/ioutil"
	"log"
)

func main() {
	b := []byte("id;precio;cantidad\n1;10;5\n2;5;10\n")
	err := ioutil.WriteFile("personal.csv", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {

	res, err := random(33, 10)

	if err != nil {
		panic("Hubo un error con el legajo")
	}
	fmt.Println(res)

	str := "customers.txt"
	read(str, res)

}

func random(max, min int) (int, error) {
	num := rand.Intn(max-min) + min
	return num, nil
}

func read(str string, lega int) {

	file, err := os.Open(str)

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		panic("error: el archivo indicado no fue encontrado o está dañado")
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")

		fmt.Println(items[0])

	}

}

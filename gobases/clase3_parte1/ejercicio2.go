package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := "./myFile.csv"
	showFile(fileName)
}

func showFile(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err == nil {
		printFile(data)
		return nil
	} else {
		return errors.New("Error al leer el archivo")
	}
}

func printFile(data []byte) {
	var sep byte = 10
	splitLines := strings.Split(string(data), string(sep))
	tabLines := ""
	total := 0

	for _, line := range splitLines {
		splitData := strings.Split(line, ",")

		precio, _ := strconv.Atoi(splitData[1])
		cantidad, _ := strconv.Atoi(splitData[2])
		total += precio * cantidad

		tabData := fmt.Sprintf("%s%10s%10s", splitData[0], splitData[1], splitData[2])
		tabLines = fmt.Sprintf("%s%s\n", tabLines, tabData)
	}
	tabData := fmt.Sprintf("%s%10d%10s", " ", total, " ")
	tabLines = fmt.Sprintf("%s%s\n", tabLines, tabData)
	fmt.Print(tabLines)
}

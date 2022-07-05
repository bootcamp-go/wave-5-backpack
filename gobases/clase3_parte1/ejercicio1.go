package main

import (
	"fmt"
	"os"
)

type products struct {
	id       int
	nombre   string
	precio   float64
	cantidad int
}

var path = "prueba.csv"

func main() {
	p1 := products{id: 10, nombre: "Papaya", precio: 1500.0, cantidad: 5}
	p2 := products{id: 1, nombre: "Mango", precio: 2500.0, cantidad: 3}
	p3 := products{id: 12, nombre: "Maracuya", precio: 3000.0, cantidad: 6}
	//convertbyte(p1)
	os.Create(path)
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if existeError(err) {
		return
	}
	defer file.Close()
	file.WriteString(string(convertbyte(p1)))
	file.WriteString(string(convertbyte(p2)))
	file.WriteString(string(convertbyte(p3)))
}

func printconst(p products) string {
	text := fmt.Sprintf("%d , %s , %.2f , %d\n", p.id, p.nombre, p.precio, p.cantidad)
	return text
}

func convertbyte(p products) []byte {
	d := []byte(printconst(p))
	return d
}

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

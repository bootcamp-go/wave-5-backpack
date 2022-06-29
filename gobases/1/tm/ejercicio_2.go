package main

import "fmt"

var (
	t float64
	p float64
	h string
)

func main() {
	t = 17.2
	p = 1.01
	h = "50%"

	fmt.Printf("Temperatura: %f, Presion: %f, Humedad: %s \n", t, p, h)
}
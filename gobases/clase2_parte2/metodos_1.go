package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	radio float64
}

func (c Circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c Circulo) perim() float64 {
	return 2 * math.Pi * c.radio
}

func (c *Circulo) setRadio(r float64) {
	c.radio = r
}

func main() {
	c := Circulo{radio: 5}
	fmt.Println(c.area())
	fmt.Println(c.perim())
	c.setRadio(10)
	fmt.Println(c.area())
	fmt.Println(c.perim())
}

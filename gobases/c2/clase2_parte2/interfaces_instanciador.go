package main

import (
	"fmt"
	"math"
)

const (
	recType    = "RECT"
	circleType = "CIRC"
)

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

type rectangle struct {
	width, height float64
}

func (r *rectangle) area() float64 {
	return r.width * r.height
}
func (r *rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func newGeometry(geoType string, values ...float64) geometry {
	switch geoType {
	case recType:
		return &rectangle{width: values[0], height: values[1]}
	case circleType:
		return &circle{radius: values[0]}
	}

	return nil
}

func main() {
	c := newGeometry(circleType, 3)
	fmt.Println("Circulo")
	fmt.Printf("area: %.3f \n", c.area())
	fmt.Printf("perimetro: %.3f\n", c.perim())
	r := newGeometry(recType, 2.0, 5)
	fmt.Println("Rectangulo")
	fmt.Printf("area: %.3f \n", r.area())
	fmt.Printf("perimetro: %.3f\n", r.perim())
}

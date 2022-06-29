package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

func details(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
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

func newCircle(values float64) geometry {
	return &circle{radius: values}
}

func newRectangle(h, w float64) geometry {
	return &rectangle{height: h, width: w}
}

func main() {
	c := newCircle(3)
	fmt.Println("Circulo")
	fmt.Printf("area: %.3f \n", c.area())
	fmt.Printf("perimetro: %.3f\n", c.perim())
	r := newRectangle(2, 5)
	fmt.Println("Rectangulo")
	fmt.Printf("area: %.3f \n", r.area())
	fmt.Printf("perimetro: %.3f\n", r.perim())
}

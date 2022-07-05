package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}
type polygon struct {
	numSides   int
	sideLength float64
}

func (p polygon) area() float64 {
	alpha := 2 * math.Pi / float64(p.numSides)
	apothem := (p.sideLength / 2) / (math.Tan(alpha / 2))
	return p.sideLength * apothem * float64(p.numSides) / 2
}

func (p polygon) perim() float64 {
	return p.sideLength * float64(p.numSides)
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 { // declarea method contained on interface geometry
	return r.height * r.width
}

func (r rectangle) perim() float64 {
	return 2*r.height + 2*r.width
}

const (
	rectType   = "RECT"
	circleType = "CIRClE"
)

func newGeometry(geoType string, values ...float64) geometry {
	switch geoType {
	case rectType:
		return rectangle{width: values[0], height: values[1]}
	case circleType:
		return circle{radius: values[0]}
	}
	return nil
}

func details(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

type geometry interface {
	area() float64 //the type of value returned by function should be specificated
	perim() float64
}

func main() {
	c := newGeometry(rectType, 13, 13)
	r := rectangle{10, 10}
	p := polygon{6, 10.0}
	details(c)
	details(r)
	details(p)

}

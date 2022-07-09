package main

import (
	"fmt"
	"math"
)

const (
	circType = "CIRCLE"
	rectType = "RECT"
)

type rect struct {
	width, height float64
}

type circ struct {
	radio float64
}

type geometry interface {
	area() float64
	perim() float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circ) area() float64 {
	return math.Pi * (c.radio * c.radio)
}

func (c circ) perim() float64 {
	return math.Pi * c.radio * 2
}

func details(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func newCircle(values float64) geometry {
	return &circ{radio: values}
}

func newGeometry(types string, values ...float64) geometry {
	switch types {
	case rectType:
		return rect{width: values[0], height: values[1]}
	case circType:
		return circ{radio: values[0]}
	default:
		return nil
	}
}

type List struct {
	Data []interface{}
}

func main() {
	r := rect{width: 4, height: 3}
	c := circ{radio: 5}
	details(r)
	details(c)
	nc := newCircle(3)
	fmt.Println(nc.area())
	fmt.Println(nc.perim())
	fmt.Println("NEW")
	g := newGeometry(rectType, 3, 8)
	fmt.Println(g.area())
	fmt.Println(g.perim())
	fmt.Println("CIRC")
	q := newGeometry(circType, 5, 9)
	fmt.Println(q.area())
	fmt.Println(q.perim())
	fmt.Println("DATA INTERFACE")
	i := List{}
	i.Data = append(i.Data, 1)
	i.Data = append(i.Data, 9, 7, 5)
	fmt.Println(i.Data)
}

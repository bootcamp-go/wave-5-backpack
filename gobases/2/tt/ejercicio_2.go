package main

import (
	"errors"
	"fmt"
)

type Matrix struct {
	values    []float64
	height    int
	width     int
	quad      bool
	max_value float64
}

func main() {
	v := []float64{1.25, 7.5, 2.15, 4.8, 5.25, 6.75}
	h := 2
	w := 3

	matrix := Matrix{}
	err := matrix.set(h, w, v...)
	if err != nil {
		fmt.Println(err)
	}
	matrix.print()
}

func (m *Matrix) set(h int, w int, values ...float64) error {
	if len(values) != h*w {
		return errors.New("La cantidad de valores no coincide con la dimension especificada")
	}

	m.height = h
	m.width = w
	if h == w {
		m.quad = true
	} else {
		m.quad = false
	}

	max := values[0]
	for _, v := range values {
		m.values = append(m.values, v)

		if max < v {
			max = v
		}
	}
	m.max_value = max

	return nil
}

func (m *Matrix) print() {
	for i, v := range m.values {
		fmt.Printf("%.2f ", v)
		if i%m.width == m.width-1 {
			fmt.Println("")
		}
	}

	if m.quad {
		fmt.Println("La matriz es cuadratica")
	} else {
		fmt.Println("La matriz no es cuadratica")
	}

	fmt.Printf("El valor máximo es %.2f\n", m.max_value)
}

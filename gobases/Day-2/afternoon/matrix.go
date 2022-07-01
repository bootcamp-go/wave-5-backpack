package main

import "fmt"

type Matrix struct {
	Height  int
	Width   int
	Matriz  [][]float64
	maxNmbr float64
}

func (m *Matrix) Set(floats ...float64) {
	max := floats[0]
	nmbrToLvt := 2
	for _, float := range floats {
		if float > max {
			max = float
		}
	}
	m.maxNmbr = max
	for nmbrToLvt*nmbrToLvt <= len(floats) {
		if nmbrToLvt*nmbrToLvt == len(floats) {
			m.Height = nmbrToLvt
			m.Width = nmbrToLvt
			for i := 0; i < len(floats); i += nmbrToLvt {
				cut := floats[i : i+nmbrToLvt]
				m.Matriz = append(m.Matriz, cut)
			}
		}
		nmbrToLvt++
	}
	if nmbrToLvt*nmbrToLvt > len(floats) && len(m.Matriz) == 0 {
		for i := 0; i < len(floats); i += nmbrToLvt {
			if !(i+nmbrToLvt > len(floats)) {
				cut := floats[i : i+nmbrToLvt]
				m.Matriz = append(m.Matriz, cut)
			} else {
				cut := floats[i:]
				m.Matriz = append(m.Matriz, cut)
			}
		}
		m.Height = len(m.Matriz)
		m.Width = len(m.Matriz[len(m.Matriz)-1])
	}
}
func (m Matrix) Print() {
	fmt.Println("o The matrix: ")
	if m.Width == m.Height {
		fmt.Println("Is a cuadratic Matrix")
	} else {
		fmt.Println("Is'nt a cuadratic Matrix")
	}
	if m.Width%2 != 0 {
		fmt.Printf("The width is an iregular: %d \n", m.Width)
	}
	fmt.Printf("The height is: %d \n", m.Width)
	for _, slice := range m.Matriz {
		fmt.Println(slice)
	}
	fmt.Printf("The max number is: %2.f \n", m.maxNmbr)
}

func main() {
	matrix1 := Matrix{}
	matrix1.Set(3.0, 2.0, 4.0, 5.0, 3.0, 2.0, 4.0)
	matrix1.Print()
}

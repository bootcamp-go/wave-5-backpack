package main

import "fmt"

type Matrix struct{
	Matriz [][]float64
	Alto int
	Ancho int
	Cuadratica bool
	Max int
}
func main (){
	m := Matrix{}
	m.Set(1,16,3,4)
	m.Set(2,4,3,6)

	m.Print()

}
func (m *Matrix)Set(valores ...float64){

	m.Matriz = append(m.Matriz,valores)

	fils := m.Matriz
	for i:=0;i<len(fils);i++{
		col := m.Matriz[i]
		for j:=0;j<len(col);j++{
			v := col[j]
			if(int(v)>m.Max){
				m.Max = int(v)
			}
			m.Alto = len(fils)
			m.Ancho = len(col)
			if(m.Alto==m.Ancho){
				m.Cuadratica = true
			}
		}
	}
}

func (m Matrix) Print(){
	for _,fils := range m.Matriz{
		fmt.Printf("\n\n")
		for _,col := range fils {
			fmt.Printf("%v\t",col)
		}
	}
	fmt.Printf("\n\n")
	fmt.Printf("Alto:%d\n",m.Alto)
	fmt.Printf("Ancho:%d\n",m.Ancho)
	fmt.Printf("Es cuadrativa:%t\n",m.Cuadratica)
	fmt.Printf("Valor maximo:%d\n",m.Max)
}

package main

func main(){

}

const (
	C string = "c"
	B string = "b"
	A string = "a"
)

func calcularSalario(horas int, categoria string)int{
	switch categoria {
		case C: 
			return 1000 * horas
		case B:
			return 1500 * horas 
		case A: 
			return 3000 * horas
	}
	return 0
}
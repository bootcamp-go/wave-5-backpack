package main

func main() {
	//1
	alumno := Alumno{
		Nombre:       "Francisco",
		Apellido:     "Monay",
		DNI:          "35123321",
		FechaIngreso: "21/06/2022",
	}

	alumno.Detalle()

	//2
	matriz := Matriz{
		Alto:       2,
		Ancho:      2,
		Cuadratica: true,
	}

	matriz.setMatriz(1, 2, 3, 4, 5, 6, 7, 8)
	matriz.printMatriz()
}

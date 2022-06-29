package main

const (
	tarantulas = "tarantulas"
	hamsters   = "hamsters"
	perros     = "perros"
	gatos      = "gatos"
)

func main() {

}

func tarantula(str string, x float64) (float64, string) {

	var cantidad float64
	if str == tarantulas {
		cantidad = x * 150 / 1000
	}
	return cantidad, "El animal no corresponde"
}

func hamster(str string, x float64) (float64, string) {

	var cantidad float64
	if str == hamsters {
		cantidad = x * 250 / 1000
	}
	return cantidad, ""
}

func perro(str string, x float64) (float64, string) {

	var cantidad float64
	if str == perros {
		cantidad = x * 10
	}
	return cantidad, "El aninal no corresponde"
}

func gato(str string, x float64) (float64, string) {

	var cantidad float64
	if str == gatos {
		cantidad = x * 5
	}
	return cantidad, ""
}

func Animal(pet, mensaje string) func(str string, x float64) (float64, string) {

	switch pet {
	case tarantulas:
		return tarantula
	case hamsters:
		return hamster
	case perros:
		return perro
	case gatos:
		return gato
	}
	return nil
}

package main

import (
	"errors"
	"fmt"
	"unicode"
)

func main() {
	var totalOfSignatures int = 0

	fmt.Println("------Ingrese el total de materias que se calificaron")
	fmt.Scanf("%v", &totalOfSignatures)

	notes, err := takingNotes()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("### --->> Notas ingresadas: ", notes)
		var avg = avg(notes, totalOfSignatures)
		fmt.Println("---El promedio del alumno es", avg)
	}

}

func takingNotes() ([]float32, error) {
	var notes []float32
	var asignature int = 1
	var inputNote float32 = 1
	for inputNote != 0 {

		fmt.Println("---Ingrese la calificación de la materia", asignature, "o cero para salir")
		fmt.Scanf("%v", &inputNote)
		if inputNote == 0 {
			break
		} else if !unicode.IsNumber(rune(inputNote)) == true || inputNote < 0 {
			return nil, errors.New("###### -> Ingresó una calificación inválida, intente de nuevo")
		} else if inputNote > 0 {
			notes = append(notes, inputNote)
			asignature++
		}

	}

	return notes, nil
}

func avg(notes []float32, totalOfAsignatures int) (avg float32) {
	var totalOfGrades float32 = addGrades(notes)
	avg = totalOfGrades / float32(totalOfAsignatures)

	return
}

func addGrades(notes []float32) float32 {
	var totalOfGrades float32 = 0
	for _, grade := range notes {
		totalOfGrades += grade
	}
	return totalOfGrades
}

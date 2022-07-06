package utils

import (
	"fmt"
	"strings"
)

// Función para dar formato a moneda
func FormatearMoneda(m float64) string {
	// Formateamos la cantidad a string
	money := fmt.Sprintf("%.2f", m)
	// Separamos la cantidad de su decimal
	moneyElements := strings.Split(money, ".")
	// Invertimos la cantidad
	moneyInverted := ""
	for _, v := range moneyElements[0] {
		moneyInverted = string(v) + moneyInverted
	}
	// Reinvertimos la cantidad y agregamos las comas
	moneyValid := ""
	for i, v := range moneyInverted {
		if (i+1)%3 == 0 && (i+1) != len(moneyInverted) {
			moneyValid = "," + string(v) + moneyValid
		} else {
			moneyValid = string(v) + moneyValid
		}
	}
	// Regresamos el resultado
	return "$" + moneyValid + "." + moneyElements[1]
}

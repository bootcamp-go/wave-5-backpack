package utils

import (
	"fmt"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/gobases/hackaton-go-bases/internal/service"
)

func ImprimirPasaje(pasaje service.Ticket) {
	p := []service.Ticket{}
	p = append(p, pasaje)
	ImprimirPasajes(p)
}

func ImprimirPasajes(pasajes []service.Ticket) {
	tabla := ""

	// Cantidad de letras de cada columna del encabezado
	idTitle, namesTitle, emailTitle, destinationTitle, dateTitle, priceTitle := "ID", "Nombre", "Email", "Destino", "Hora", "Precio"
	idSize, namesSize, emailSize, destinationSize, dateSize, priceSize := len(idTitle), len(namesTitle), len(emailTitle), len(destinationTitle), len(dateTitle), len(priceTitle)

	// Se calcula la longitud de cada columna
	for _, p := range pasajes {
		if len(strconv.Itoa(p.Id)) > idSize {
			idSize = len(strconv.Itoa(p.Id))
		}
		if len(p.Names) > namesSize {
			namesSize = len(p.Names)
		}
		if len(p.Email) > emailSize {
			emailSize = len(p.Email)
		}
		if len(p.Destination) > destinationSize {
			destinationSize = len(p.Destination)
		}
		if len(p.Date) > dateSize {
			dateSize = len(p.Date)
		}
		formatMoneda := FormatearMoneda(float64(p.Price))
		if len(formatMoneda) > priceSize {
			//priceSize = len(strconv.Itoa(p.Price))
			priceSize = len(formatMoneda)
		}
	}

	// Se agregan 1 espacios a cada columna
	idSize += 1
	namesSize += 1
	emailSize += 1
	destinationSize += 1
	dateSize += 1
	priceSize += 1
	total := idSize + namesSize + emailSize + destinationSize + dateSize + priceSize + 13

	// Se imprime el separador principal
	for i := 0; i < total; i++ {
		tabla += "-"
	}
	tabla += "\n"

	// Se imprimen los titulos
	tabla += "| " + idTitle
	for i := 0; i < (idSize - len(idTitle)); i++ {
		tabla += " "
	}
	tabla += "| " + namesTitle
	for i := 0; i < (namesSize - len(namesTitle)); i++ {
		tabla += " "
	}
	tabla += "| " + emailTitle
	for i := 0; i < (emailSize - len(emailTitle)); i++ {
		tabla += " "
	}
	tabla += "| " + destinationTitle
	for i := 0; i < (destinationSize - len(destinationTitle)); i++ {
		tabla += " "
	}
	tabla += "| " + dateTitle
	for i := 0; i < (dateSize - len(dateTitle)); i++ {
		tabla += " "
	}
	tabla += "| " + priceTitle
	for i := 0; i < (priceSize - len(priceTitle)); i++ {
		tabla += " "
	}
	tabla += "|\n"

	// Se imprime el separador
	for i := 0; i < total; i++ {
		tabla += "-"
	}
	tabla += "\n"

	// Se imprime la información de los tickets
	for _, t := range pasajes {
		tabla += "| " + strconv.Itoa(t.Id)
		for i := 0; i < (idSize - len(strconv.Itoa(t.Id))); i++ {
			tabla += " "
		}
		tabla += "| " + t.Names
		for i := 0; i < (namesSize - len(t.Names)); i++ {
			tabla += " "
		}
		tabla += "| " + t.Email
		for i := 0; i < (emailSize - len(t.Email)); i++ {
			tabla += " "
		}
		tabla += "| " + t.Destination
		for i := 0; i < (destinationSize - len(t.Destination)); i++ {
			tabla += " "
		}
		tabla += "| " + t.Date
		for i := 0; i < (dateSize - len(t.Date)); i++ {
			tabla += " "
		}
		priceFormat := FormatearMoneda(float64(t.Price))
		tabla += "|"
		for i := 0; i < (priceSize - len(priceFormat)); i++ {
			tabla += " "
		}
		tabla += priceFormat + " |\n"

		// Se imprime el separador
		for i := 0; i < total; i++ {
			tabla += "-"
		}
		tabla += "\n"
	}

	fmt.Print(tabla)
}

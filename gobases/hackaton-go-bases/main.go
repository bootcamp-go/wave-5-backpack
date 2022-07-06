package main

import (
	"fmt"
	"strconv"

	"github.com/bootcamp-go/wave-5-backpack/gobases/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/wave-5-backpack/gobases/hackaton-go-bases/internal/service"
	"github.com/bootcamp-go/wave-5-backpack/gobases/hackaton-go-bases/internal/utils"
)

func main() {
	fmt.Println("Sistema de reservas de pasajes")

	// Leemos la información del archivo
	var ticketsFile file.File
	ticketsFile.SetPath("./tickets.csv")
	dataPasajes, err := ticketsFile.Read()

	// Verificamos que se haya podido abrir el archivo
	if err != nil {
		panic(err)
	} else {
		// Mostramos menú
		fmt.Print("\n  1. Crear registro\n  2. Leer registro\n  3. Actualizar registro\n  4. Borrar registro\n  5. Mostrar tickets\n  6. Mostrar menú\n  7. Salir\n")

		// Funcion para obtener tickets del archivo csv
		bookings := service.NewBookings(dataPasajes)

		option := ""
		for option != "7" {
			fmt.Print("\nSeleccione una opción del menú: ")
			fmt.Scanf("%s", &option)

			// Se verifica que la opción sea la correcta
			optionInt, err := strconv.Atoi(option)
			if err != nil {
				fmt.Println("Opcion no valida")
			} else if optionInt < 8 && optionInt > 0 {
				switch optionInt {
				case 1: // Crear registro
					valid := true
					nameReg, emailReg, destinationReg, dateReg, priceReg := "", "", "", "", ""
					price := 0

					// Se validan los datos ingresados
					for valid {
						fmt.Print("Ingrese el Nombre: ")
						fmt.Scanf("%s", &nameReg)
						fmt.Print("Ingrese el Email: ")
						fmt.Scanf("%s", &emailReg)
						fmt.Print("Ingrese el Destino: ")
						fmt.Scanf("%s", &destinationReg)
						fmt.Print("Ingrese la Hora: ")
						fmt.Scanf("%s", &dateReg)
						fmt.Print("Ingrese el Precio: ")
						fmt.Scanf("%s", &priceReg)

						if nameReg == "" || emailReg == "" || destinationReg == "" || dateReg == "" || priceReg == "" {
							fmt.Print("Todos los campos son obligatorios!\n")
						} else {
							priceInt, err := strconv.Atoi(priceReg)
							if err != nil {
								fmt.Print("\nEl precio es incorrecto")
							} else {
								price = priceInt
								valid = false
							}
						}
					}

					ticket := service.Ticket{Names: nameReg, Email: emailReg, Destination: destinationReg, Date: dateReg, Price: price}
					created, err := bookings.Create(ticket)
					if err != nil {
						fmt.Print("No se pudo crear el ticket")
					} else {
						err := ticketsFile.Write(bookings.GetTickets(0))
						if err != nil {
							fmt.Print("No se pudo guardar la información")
						} else {
							fmt.Print("Se creo el siguiente ticket:\n")
							utils.ImprimirPasaje(created)
						}
					}
					option = ""
				case 2: // Imprimir información del ticket
					idString := ""
					valid := true
					for valid {
						fmt.Print("Ingrese el id del ticket (CERO para regresar al menu): ")
						fmt.Scanf("%s", &idString)

						// Se verifica que el id sea en formato número
						idInt, err := strconv.Atoi(idString)
						if err != nil {
							fmt.Print("ID incorrecto!")
						} else {
							if idInt == 0 {
								valid = false
							} else {
								// Buscamos el ID en el archivo
								ticket, err := bookings.Read(idInt)
								if err != nil {
									fmt.Print(err, ", ID: ", idInt, "\n")
								} else {
									utils.ImprimirPasaje(ticket)
								}
							}
						}
					}
					option = ""
				case 3: // Actualizar ticket
					idString := ""
					valid := true
					for valid {
						fmt.Print("Ingrese el id del ticket a actualizar (CERO para regresar al menu): ")
						fmt.Scanf("%s", &idString)

						// Se verifica que el id sea en formato número
						idInt, err := strconv.Atoi(idString)
						if err != nil {
							fmt.Print("ID incorrecto!")
						} else {
							if idInt == 0 {
								valid = false
							} else {
								// Buscamos el ID en el archivo
								ticket, err := bookings.Read(idInt)
								if err != nil {
									fmt.Print(err, ", ID: ", idInt, "\n")
								} else {
									valid = false
									utils.ImprimirPasaje(ticket)
									emailReg, destinationReg, dateReg, priceReg := "", "", "", ""
									fmt.Print("Ingrese el Email (ENTER para no modificar): ")
									fmt.Scanf("%s", &emailReg)
									fmt.Print("Ingrese el Destino (ENTER para no modificar): ")
									fmt.Scanf("%s", &destinationReg)
									fmt.Print("Ingrese la Hora (ENTER para no modificar): ")
									fmt.Scanf("%s", &dateReg)
									fmt.Print("Ingrese el Precio (ENTER para no modificar): ")
									fmt.Scanf("%s", &priceReg)

									if emailReg != "" {
										ticket.Email = emailReg
									}
									if destinationReg != "" {
										ticket.Destination = destinationReg
									}
									if dateReg != "" {
										ticket.Date = dateReg
									}

									if priceReg != "" {
										priceInt, err := strconv.Atoi(priceReg)
										if err != nil {
											fmt.Print("El precio es incorrecto, no se modificará!\n")
										} else if priceInt < 0 {
											fmt.Print("El precio no debe ser negativo, no se modificará!")
										} else {
											ticket.Price = priceInt
										}
									}

									updated, err := bookings.Update(ticket.Id, ticket)
									if err != nil {
										fmt.Print(err)
									} else {
										err := ticketsFile.Write(bookings.GetTickets(0))
										if err != nil {
											fmt.Print("No se pudo guardar la información")
										} else {
											fmt.Print("\nSe actualzó siguiente ticket:\n")
											utils.ImprimirPasaje(updated)
										}
									}
								}
							}
						}
					}
					option = ""
				case 4: // Borrar ticket
					idString := ""
					valid := true
					for valid {
						fmt.Print("Ingrese el id del ticket a borrar (CERO para regresar al menu): ")
						fmt.Scanf("%s", &idString)

						// Se verifica que el id sea en formato número
						idInt, err := strconv.Atoi(idString)
						if err != nil {
							fmt.Print("ID incorrecto!")
						} else {
							if idInt == 0 {
								valid = false
							} else {
								// Buscamos el ID en el archivo
								ticket, err := bookings.Read(idInt)
								if err != nil {
									fmt.Print(err, ", ID: ", idInt, "\n")
								} else {
									valid = false
									_, err := bookings.Delete(idInt)
									if err != nil {
										fmt.Print(err)
									} else {
										err := ticketsFile.Write(bookings.GetTickets(0))
										if err != nil {
											fmt.Println(err)
										} else {
											fmt.Print("\nSe borró siguiente ticket:\n")
											utils.ImprimirPasaje(ticket)
										}
									}
								}
							}
						}
					}
					option = ""
				case 5:
					total := ""
					valid := true
					for valid {
						fmt.Print("Ingrese el total de tickets a mostrar: ")
						fmt.Scanf("%s", &total)
						totalInt, err := strconv.Atoi(total)
						if err != nil {
							fmt.Print("Cantidad incorrecta!")
						} else {
							valid = false
							utils.ImprimirPasajes(bookings.GetTickets(totalInt))
						}
					}
					option = ""
				case 6:
					fmt.Print("\n  1. Crear registro\n  2. Leer registro\n  3. Actualizar registro\n  4. Borrar registro\n  5. Mostrar tickets\n  6. Mostrar menú\n  7. Salir\n")
					option = ""
				}
			} else {
				fmt.Println("No existe esa opción")
				option = ""
			}
		}
	}

	fmt.Print("\nFin de la ejecución\n\n")
}

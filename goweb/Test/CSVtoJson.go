package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Ticket struct {
	Id      string
	Name    string
	Email   string
	Country string
	Time    string
	Price   float64
}

func main() {
	csvFile, err := os.Open("./tickets.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var ticket Ticket
	var tickets []Ticket

	for _, each := range csvData {
		ticket.Id = each[0]
		ticket.Name = each[1]
		ticket.Email = each[2]
		ticket.Country = each[3]
		ticket.Time = each[4]
		ticket.Price, _ = strconv.ParseFloat(each[5], 64)

		tickets = append(tickets, ticket)
	}

	// Convert to JSON
	jsonData, err := json.Marshal(tickets)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))

	jsonFile, err := os.Create("./data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	jsonFile.Close()
}

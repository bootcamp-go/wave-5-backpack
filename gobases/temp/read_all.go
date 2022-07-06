package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// type User struct {
// 	firstName  string
// 	lastName   string
// 	occupation string
// }
type User struct {
	Id                              int
	Names, Email, Destination, Date string
	Price                           int
}

func main() {

	records, err := readData("users.csv")

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {

		id, _ := strconv.Atoi(record[0])
		price, _ := strconv.Atoi(record[5])
		user := User{
			Id:          id,
			Names:       record[1],
			Email:       record[2],
			Destination: record[3],
			Date:        record[4],
			Price:       price,
		}

		fmt.Printf("%d %s %s %s %s %d\n", user.Id, user.Names,
			user.Email, user.Destination, user.Date, user.Price)
	}
}

func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

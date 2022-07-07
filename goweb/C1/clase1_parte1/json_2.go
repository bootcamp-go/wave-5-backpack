package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type product struct {
		Name      string `json:"name"`
		Price     int
		Published bool
	}

	jsonData := []byte(`{"name": "MacBook Air", "Price": 900, "Published": true, "Date":"2020"}`)

	var p product

	if err := json.Unmarshal(jsonData, &p); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("p: %+v\n", p)
}

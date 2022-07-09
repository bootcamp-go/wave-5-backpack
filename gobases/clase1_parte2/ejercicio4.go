package main

import "fmt"

func main()  {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println(employees)
	count := 0
	if employees["Benjamin"] > 21{
		count ++
	}
	if employees["Nahuel"] > 21{
		count ++
	}
	if employees["Brenda"] > 21{
		count ++
	}
	if employees["Dario"] > 21{
		count ++
	}
	if employees["Pedro"] > 21{
		count ++
	}
	fmt.Println(count)

	employees["Camilo"] = 30
	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println(employees)

}

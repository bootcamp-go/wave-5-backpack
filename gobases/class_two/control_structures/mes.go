package main

import "fmt"

func main() {
	month := 1

	months := []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "Octuber", "November", "December"}

	if month < len(months) {
		fmt.Println("El mes seleccionado es: ", months[month])
	} else {
		fmt.Println("That month does not exists! are you ok?")
	}

	var monthsMap = map[int]string{
		0:  "January",
		1:  "February",
		2:  "March",
		3:  "April",
		4:  "May",
		5:  "June",
		6:  "July",
		7:  "August",
		8:  "September",
		9:  "October",
		10: "November",
		11: "December",
	}
	fmt.Println("mes con map", monthsMap[month])
	/*switch month {
	case 0:
		fmt.Println("January")
	case 1:
		fmt.Println("February")

	case 2:
		fmt.Println("March")

	case 3:
		fmt.Println("April")

	case 4:
		fmt.Println("May")

	case 5:
		fmt.Println("June")

	case 6:
		fmt.Println("July")

	case 7:
		fmt.Println("August")

	case 8:
		fmt.Println("September")

	case 9:
		fmt.Println("Octuber")

	case 10:
		fmt.Println("November")

	case 11:
		fmt.Println("December")

	default:
		fmt.Println("That month does not exists! are you ok?")

	}

	*/
}

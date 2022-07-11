package main

import "fmt"
import "strings"
import "bufio"
import "os"
import "strconv"

func salaryTax(salary float64) float64{
	if salary > 150000{
		// 27%
		return salary * 0.27
	}else if salary > 50000{
		// 17%
		return salary * 0.17
	}
	return 0
}

func main(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("--------Salary----------")

	fmt.Print("-> ")

	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	salary, err := strconv.ParseFloat(text, 2)
	if err != nil {
		fmt.Println(err)
	}else {
		tax := salaryTax(salary)
		fmt.Println("Tax: ", tax)
	}

}
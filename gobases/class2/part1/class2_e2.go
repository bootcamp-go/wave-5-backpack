package main

import "fmt"
import "errors"

func average(values ...int)(int, error){
	var result = 0
	ran := 0
	for _, value := range values {
		ran++
		if value < 0{
			return 0, errors.New("Negative value")
		}
		result += value 
	}

	if (ran == 0){
		return result, nil
	} else{	
		return result/ ran, nil
	}
}


func main(){
	result, err := average(1, 2, 5, 8, 9, -1)
	if err !=  nil {
		fmt.Println(err)
	}else{
		fmt.Println("Result: ", result);
	}
	
}
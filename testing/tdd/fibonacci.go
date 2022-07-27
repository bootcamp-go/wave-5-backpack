package tdd

import "fmt"

func fibonacci(x int) int {
	fmt.Println("VALIDA X : ", x)
	t1, t2 := 0, 0
	for i := 0; i <= x; i++ {
		if i == 0 {
			fmt.Printf("t1: %d, t2:%d\n", t1, t2)
			continue
		}
		if i == 1 {
			t2 = 1
			fmt.Printf("t1: %d, t2:%d\n", t1, t2)
			continue
		}
		temp := t1 + t2
		t1 = t2
		t2 = temp
		fmt.Printf("t1: %d, t2:%d\n", t1, t2)
	}
	fmt.Println("RETORNA : ", t2)
	return t2
}

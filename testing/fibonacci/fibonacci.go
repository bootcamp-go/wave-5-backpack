package fibonacci

func Fibonacci(n int) int64 {
	var a, b int64 = 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

//func Fibonacci(n int) int64 {
//	if n == 0 {
//		return 0
//	}
//	if n == 1 || n == 2 {
//		return 1
//	}
//	var a, b int64 = 1, 1
//	for i := 3; i <= n; i++ {
//		a, b = b, a+b
//	}
//	return b
//}

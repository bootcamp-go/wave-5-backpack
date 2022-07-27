package clase4testingprt1tdd

const LIM = 41

var facts [LIM]uint64

func FactorialMemoization(n uint64) (res uint64) {
	if facts[n] != 0 {
		res = facts[n]
		return res
	}

	if n > 0 {
		res = n * FactorialMemoization(n-1)
		return res
	}

	return 1
}

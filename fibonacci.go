package piscine

func Fibonacci(index int) int {
	index = 0

	if index < 0 {
		return -1
	} else {
		return index + Fibonacci(index+1)
	}

}

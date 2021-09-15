package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index < 2 {
		return index
	}
	return Fibonacci(index-2) + Fibonacci(index-1)
}

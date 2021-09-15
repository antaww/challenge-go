package piscine

func Fibonacci(index int) int {
	if index < 2 {
		return index
	}
	if index < 0 {
		return -1
	}
	return Fibonacci(index-2) + Fibonacci(index-1)
}

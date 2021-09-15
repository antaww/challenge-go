package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else {
		index = 0
		return index + Fibonacci(index+1)
	}
}

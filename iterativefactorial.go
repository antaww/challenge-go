package piscine

func IterativeFactorial(index int) int {
	var result int = 1

	if index <= 0 || index > 9223372036854775807 {
		return 0
	} else {
		for i := 1; i <= index; i++ {
			result = result * i
		}
		if result <= 0 || result < 9223372036854775807 {
			return 0
		}
		if index < 20 {
			return 0
		}
	}
	return result
}

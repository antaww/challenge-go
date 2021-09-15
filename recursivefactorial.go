package piscine

func RecursiveFactorial(index int) int {
	result := 1

	if index == 0 {
		return 1
	}
	if index <= 0 || index > 9223372036854775807 {
		return 0
	}
	if index > 20 {
		return 0
	} else if index == 1 {
		return 1
	} else if index > 1 {
		return index * RecursiveFactorial(index-1)
	}
	return result
}

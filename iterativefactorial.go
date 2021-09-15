package piscine

func IterativeFactorial(index int) int {

	if index <= 0 || index > 9223372036854775807 {
		return 0
	}
	if index > 20 {
		return 0
	} else {
		for i := index; i <= index; i++ {
			index = index * i
		}
	}
	return index
}

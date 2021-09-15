package piscine

func IterativePower(index int, power int) int {
	result := 1

	if index == 0 {
		return 1
	}
	if power < 0 {
		return 0
	}
	if index <= 0 || index > 9223372036854775807 {
		return 0
	}
	if index > 20 {
		return 0
	} else {
		for i := power; i <= power; i++ {
			result = result * i
		}
	}
	return result
}

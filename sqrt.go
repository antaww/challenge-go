package piscine

func Sqrt(nb int, power int) int {
	result := nb

	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	} else {
		for i := 1; i <= power-1; i++ {
			result = nb / result
		}
	}
	return result
}

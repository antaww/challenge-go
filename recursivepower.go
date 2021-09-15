package piscine

func RecursivePower(nb int, power int) int {
	result := nb

	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	} else if nb > 1 {
		return result * RecursivePower(nb, power-1)
	}
	return result
}

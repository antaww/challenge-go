package piscine

func FindNextPrime(nb int) int {
	if nb < 0 {
		return 0
	} else {
		return FindNextPrime(nb)
	}
}

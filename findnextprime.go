package piscine

func PrimeNumber(nb int) bool {
	prime := true
	for i := 2; i <= nb-1; i++ {
		if nb%i == 0 {
			prime = false
		}
	}
	return prime
}

func FindNextPrime(nb int) int {
	if PrimeNumber(nb) == true {
		return nb
	} else {
		return FindNextPrime(nb + 1)
	}
}

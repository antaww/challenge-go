package piscine

func Any(f func(string) bool, a []string) bool {
	for i := range a {
		if f(a[i]) == true {
			return true
		}
	}
	return false
}

package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := range a {
		if a[i-1] <= a[i] {
			return true
		}
	}
	return false
}

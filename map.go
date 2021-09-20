package piscine

func Map(f func(int) bool, a []int) []bool {
	result := []bool{}
	for i := range a {
		f(a[i])
	}
	return result
}

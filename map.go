package piscine

func Map(f func(int) bool, a []int) []bool {
	for i := range a {
		f(a[i])
	}
}

package piscine

func Map(f func(int) bool, a []int) []bool {
	result := []bool{}
	for i := range a {
		result = append(result, f(i))
	}
	return result
}

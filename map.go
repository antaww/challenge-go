package piscine

func Map(f func(int) bool, a []int) []bool {
	result := []bool{}
	for i := range a {
		if f(a[i]) == true {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
		return result
	}
}

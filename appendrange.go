package piscine

func AppendRange(min, max int) []int {
	var tableau []int
	if min >= max {
		return nil
	}
	if max >= min {
		for i := min; i < max; i++ {
			tableau = append(tableau, i)
		}
	}
	return tableau
}

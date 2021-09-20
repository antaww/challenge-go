package piscine

func MakeRange(min, max int) []int {
	tableau := make([]int, max-min)
	if min >= max {
		return nil
	}
	if max >= min {
		for i := range tableau {
			tableau[i] += min + i
		}
	}
	return tableau
}

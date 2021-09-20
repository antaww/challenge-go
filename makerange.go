package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	tableau := make([]int, max-min)
	for i := range tableau {
		tableau[i] += min + i
	}
	return tableau
}

package piscine

func MakeRange(min, max int) []int {
	tableau := make([]int, max-min)
	if min >= max {
		return nil
	}
	for i := range tableau {
		tableau[i] = min + 1
	}
	return tableau
}

package piscine

func MakeRange(min, max int) []int {
	var tableau []int
	if min >= max {
		return nil
	}
	if max >= min {
		for i := min; i < max; i++ {
			tableau = make([]int, i)
		}
	}
	return tableau
}

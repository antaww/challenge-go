package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	decroissant := true
	for i := range a {
		if f(a[i], a[i+1]) > 0 { // i < i+1
			if i == 0 {
				decroissant = true
			} else if i != 0 && decroissant == false {
				return false
			}
		} else if f(a[i], a[i+1]) < 0 { // i > i+1
			if i == 0 {
				decroissant = false
			} else if i != 0 && decroissant == true {
				return false
			}
		}
	}
	return true
}

package piscine

func Swap(a *int, b *int) {
	afix := *a
	bfix := *b
	*a = bfix
	*b = afix
}

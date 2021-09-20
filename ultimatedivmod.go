package piscine

func UltimateDivMod(a *int, b *int) {
	div := *a / *b
	*b = *a % *b
	*a = div
}

package piscine

func UltimateDivMod(a *int, b *int) {
	div := a / b
	*a = div
	*b = div %
}

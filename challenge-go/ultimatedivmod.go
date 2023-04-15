package piscine

func UltimateDivMod(a *int, b *int) {
	t := *a
	*a = *a / *b
	*b = t % *b
}
